package http

import (
	"bytes"
	"encoding/json"
	ps "fase-4-hf-voucher/external/strings"
	"fase-4-hf-voucher/internal/core/application"
	"fase-4-hf-voucher/internal/core/domain/entity"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type VoucherHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type voucherHandler struct {
	app application.Application
}

func NewHandler(app application.Application) VoucherHandler {
	return voucherHandler{app: app}
}

func (h voucherHandler) Handler(rw http.ResponseWriter, req *http.Request) {

	var routesVoucher = map[string]http.HandlerFunc{
		"get hermes_foods/voucher/{id}": h.getByID,
		"post hermes_foods/voucher":     h.save,
	}

	handler, err := router(req.Method, req.URL.Path, routesVoucher)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h voucherHandler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func getID(handlerName, url string) string {
	index := strings.Index(url, handlerName+"/")

	if index == -1 {
		return ""
	}

	id := strings.ReplaceAll(url[index:], handlerName+"/", "")

	return id
}

func (h *voucherHandler) getByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	v, err := h.app.GetVoucherByID(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(v)))
}

func (h *voucherHandler) save(rw http.ResponseWriter, req *http.Request) {
	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}
	percentage, err := strconv.ParseInt(reqVoucher.Percentage, 10, 64)
	if err != nil { }
	voucher := entity.Voucher{
		Code:       reqVoucher.Code,
		Percentage: percentage,
	}

	if len(reqVoucher.ExpiresAt) > 0 {
		voucher.ExpiresAt.Value = new(time.Time)
		if err := voucher.ExpiresAt.SetTimeFromString(reqVoucher.ExpiresAt); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
			return
		}
	}

	reqVoucher.ExpiresAt = voucher.ExpiresAt.Format()

	v, err := h.app.SaveVoucher(reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}

func (h *voucherHandler) UpdateVoucherByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("voucher", req.URL.Path)

	var buff bytes.Buffer

	var reqVoucher dto.RequestVoucher

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqVoucher); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	// voucher := entity.Voucher{
	// 	Code:       reqVoucher.Code,
	// 	Percentage: reqVoucher.Percentage,
	// }

	// func (app *application.Application) UpdateVoucherByID(id string, reqVoucher dto.RequestVoucher) (entity.Voucher, error) {
	// 	// Implementation of the UpdateVoucherByID method
	// }

	v, err := h.app.UpdateVoucherByID(id, reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to update voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(v)))
}

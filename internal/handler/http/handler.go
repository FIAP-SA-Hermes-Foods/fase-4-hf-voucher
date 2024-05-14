package http

import (
	"bytes"
	"encoding/json"
	ps "fase-4-hf-voucher/external/strings"
	"fase-4-hf-voucher/internal/core/application"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"strings"
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
		"post hermes_foods/voucher":      h.save,
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

func (h voucherHandler) getByCPF(rw http.ResponseWriter, req *http.Request) {
	id := getId(req.URL.Path)

	c, err := h.app.GetVoucherByID(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get voucher by ID: %v"} `, err)
		return
	}

	if c == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "voucher not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(c)))
}

func (h voucherHandler) save(rw http.ResponseWriter, req *http.Request) {
	var (
		buff      bytes.Buffer
		reqVoucher dto.RequestVoucher
	)

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

	c, err := h.app.SaveVoucher(reqVoucher)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save voucher: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(c)))
}

func getId(url string) string {
	indexId := strings.Index(url, "voucher/")

	if indexId == -1 {
		return ""
	}

	return strings.ReplaceAll(url[indexId:], "voucher/", "")
}

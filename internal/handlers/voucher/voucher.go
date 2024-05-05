package voucher

import (
	"errors"
	"net/http"
	"time"

	"github.com/PauloLucas94/fase-4-hf-voucher/internal/controllers/voucher"
	EntityVoucher "github.com/PauloLucas94/fase-4-hf-voucher/internal/entities/voucher"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/handlers"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/repository/adapter"
	Rules "github.com/PauloLucas94/fase-4-hf-voucher/internal/rules"
	RulesVoucher "github.com/PauloLucas94/fase-4-hf-voucher/internal/rules/voucher"
	HttpStatus "github.com/PauloLucas94/fase-4-hf-voucher/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Handler struct {
	handlers.Interface

	Controller voucher.Interface
	Rules      Rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: voucher.NewController(repository),
		Rules:      RulesVoucher.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	response, err := h.Controller.ListOne(ID)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	voucherBody, err := h.getBodyAndValidate(r, uuid.Nil)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	ID, err := h.Controller.Create(voucherBody)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	voucherBody, err := h.getBodyAndValidate(r, ID)
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(ID, voucherBody); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	if err := h.Controller.Remove(ID); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityVoucher.Voucher, error) {
	voucherBody := &EntityVoucher.Voucher{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, voucherBody)
	if err != nil {
		return &EntityVoucher.Voucher{}, errors.New("body is required")
	}

	voucherParsed, err := EntityVoucher.InterfaceToModel(body)
	if err != nil {
		return &EntityVoucher.Voucher{}, errors.New("error on convert body to model")
	}

	setDefaultValues(voucherParsed, ID)

	return voucherParsed, h.Rules.Validate(voucherParsed)
}

func setDefaultValues(voucher *EntityVoucher.Voucher, ID uuid.UUID) {
	voucher.UpdatedAt = time.Now().AddDate(0, 1, 0)
	if ID == uuid.Nil {
		voucher.ID = uuid.New()
		voucher.CreatedAt = time.Now()
	} else {
		voucher.ID = ID
	}
}

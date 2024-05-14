package rpc

import (
	"context"
	cp "fase-4-hf-voucher/client_proto"
	"fase-4-hf-voucher/internal/core/application"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedVoucherServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) Create(ctx context.Context, req *cp.CreateRequest) (*cp.CreateResponse, error) {

	input := dto.RequestVoucher{
		Code:       req.Code,
		Percentage: req.Percentage,
	}

	c, err := h.app.SaveVoucher(input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.CreateResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: c.Percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil

}

func (h *handlerGRPC) GetById(ctx context.Context, req *cp.GetByIDRequest) (*cp.GetByIDResponse, error) {
	//Update to GetUuid
	c, err := h.app.GetVoucherByID(req.GetUuid())

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.GetByIDResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: c.Percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil
}

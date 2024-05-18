package rpc

import (
	"context"
	"fase-4-hf-voucher/internal/core/application"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	cp "fase-4-hf-voucher/voucher_proto"
	"fmt"
	"strconv"
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
		Percentage: strconv.FormatInt(req.Percentage, 10),
		ExpiresAt:  req.ExpiresAt,
	}


	c, err := h.app.SaveVoucher(input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}


	percentage, err := strconv.ParseInt(c.Percentage, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse percentage: %v", err)
	}
	out := &cp.CreateResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil

}

func (h *handlerGRPC) GetById(ctx context.Context, req *cp.GetByIDRequest) (*cp.GetByIDResponse, error) {

	c, err := h.app.GetVoucherByID(req.Uuid)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	percentage, err := strconv.ParseInt(c.Percentage, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse percentage: %v", err)
	}

	out := &cp.GetByIDResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil
}

func (h *handlerGRPC) GetByID(ctx context.Context, req *cp.GetByIDRequest) (*cp.GetByIDResponse, error) {
	c, err := h.app.GetVoucherByID(req.Uuid)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	percentage, err := strconv.ParseInt(c.Percentage, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse percentage: %v", err)
	}

	out := &cp.GetByIDResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil
}

func (h *handlerGRPC) UpdateByID(ctx context.Context, req *cp.UpdateByIDRequest) (*cp.UpdateByIDResponse, error) {

	input := dto.RequestVoucher{
		Code:       req.Code,
		Percentage: strconv.FormatInt(req.Percentage, 10),
		ExpiresAt:  req.ExpiresAt,
	}

	c, err := h.app.UpdateVoucherByID(req.Uuid, input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	percentage, err := strconv.ParseInt(c.Percentage, 10, 64)

	if err != nil {
	}

	out := &cp.UpdateByIDResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil
}

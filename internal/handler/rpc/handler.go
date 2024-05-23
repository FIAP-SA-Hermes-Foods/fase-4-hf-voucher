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

func (h *handlerGRPC) CreateVoucher(ctx context.Context, req *cp.CreateVoucherRequest) (*cp.CreateVoucherResponse, error) {

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
	out := &cp.CreateVoucherResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil

}

func (h *handlerGRPC) GetVoucherByID(ctx context.Context, req *cp.GetVoucherByIDRequest) (*cp.GetVoucherByIDResponse, error) {

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

	out := &cp.GetVoucherByIDResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil
}

func (h *handlerGRPC) UpdateVoucherByID(ctx context.Context, req *cp.UpdateVoucherByIDRequest) (*cp.UpdateVoucherByIDResponse, error) {

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

	out := &cp.UpdateVoucherByIDResponse{
		Uuid:       c.UUID,
		Code:       c.Code,
		Percentage: percentage,
		CreatedAt:  c.CreatedAt,
		ExpiresAt:  c.ExpiresAt,
	}

	return out, nil
}

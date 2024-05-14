package dto

import (
	"fase-4-hf-voucher/internal/core/domain/entity"
)

type VoucherDB struct {
	UUID       string `json:"uuid,omitempty"`
	Code       string `json:"code,omitempty"`
	Percentage string `json:"percentage,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	ExpiresAt  string `json:"expires_at,omitempty"`
}

type (
	RequestVoucher struct {
		UUID       string `json:"uuid,omitempty"`
		Code       string `json:"code,omitempty"`
		Percentage string `json:"percentage,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		ExpiresAt  string `json:"expires_at,omitempty"`
	}

	OutputVoucher struct {
		UUID       string `json:"uuid,omitempty"`
		Code       string `json:"code,omitempty"`
		Percentage string `json:"percentage,omitempty"`
		CreatedAt  string `json:"created_at,omitempty"`
		ExpiresAt  string `json:"expires_at,omitempty"`
	}
)

func (r RequestVoucher) Voucher() entity.Voucher {
	return entity.Voucher{
		Code: r.Code,
		// r.uuid: vo.uuid{
		// 	Value: r.uuid,
		// },
		Percentage: r.Percentage,
	}
}

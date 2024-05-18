package repository

import (
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
)

type VoucherRepository interface {
	GetVoucherByID(id string) (*dto.VoucherDB, error)
	SaveVoucher(voucher dto.VoucherDB) (*dto.VoucherDB, error)
	UpdateVoucherByID(id string, voucher dto.VoucherDB) (*dto.VoucherDB, error)
}

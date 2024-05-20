package useCase

import "fase-4-hf-voucher/internal/core/domain/entity/dto"

type VoucherUseCase interface {
	SaveVoucher(reqVoucher dto.RequestVoucher) error
	GetVoucherByID(id string) error
	UpdateVoucherByID(id string, reqVoucher dto.RequestVoucher) error
}

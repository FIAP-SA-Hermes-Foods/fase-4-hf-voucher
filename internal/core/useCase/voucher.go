package useCase

import (
	"errors"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/internal/core/domain/useCase"
)

var _ useCase.VoucherUseCase = (*voucherUseCase)(nil)

type voucherUseCase struct {
}

func NewVoucherUseCase() voucherUseCase {
	return voucherUseCase{}
}

func (c voucherUseCase) SaveVoucher(reqVoucher dto.RequestVoucher) error {
	voucher := reqVoucher.Voucher()

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return err
	}

	if len(voucher.Code) == 0 {
		return errors.New("the voucher code is null or not valid")
	}

	if voucher.Percentage < 0 || voucher.Percentage > 101 {
		return errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return nil
}

func (c voucherUseCase) GetVoucherByID(id string) error {
	if len(id) == 0 {
		return errors.New("the id is not valid for consult")
	}

	return nil
}

func (c voucherUseCase) UpdateVoucherByID(id string, reqVoucher dto.RequestVoucher) error {
	if len(id) == 0 {
		return errors.New("the id is not valid for consult")
	}

	voucher := reqVoucher.Voucher()

	if err := voucher.ExpiresAt.Validate(); err != nil {
		return err
	}

	if len(voucher.Code) == 0 {
		return errors.New("the voucher code is null or not valid")
	}

	if voucher.Percentage < 0 || voucher.Percentage > 101 {
		return errors.New("the porcentage is not valid try a number between 0 and 100")
	}

	return nil
}

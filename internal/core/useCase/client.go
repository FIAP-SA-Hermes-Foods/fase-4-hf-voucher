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

	if err := voucher.ID.Validate(); err != nil {
		return err
	}

	return nil
}

func (c voucherUseCase) GetVoucherByID(id string) error {
	if len(id) == 0 {
		return errors.New("the id is not valid for consult")
	}

	return nil
}

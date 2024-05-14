package mocks

import (
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"strings"
)

type MockVoucherUseCase struct {
	WantOutNull string
	WantErr     error
}

func (m MockVoucherUseCase) GetVoucherByID(voucher string) error {
	if m.WantErr != nil && strings.EqualFold("errGetVoucherByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockVoucherUseCase) SaveVoucher(reqVoucher dto.RequestVoucher) error {
	if m.WantErr != nil && strings.EqualFold("errSaveVoucher", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}

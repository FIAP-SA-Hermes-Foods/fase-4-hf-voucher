package mocks

import (
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"strings"
)

type MockVoucherRepository struct {
	WantOut     *dto.VoucherDB
	WantOutNull string
	WantErr     error
}

func (m MockVoucherRepository) GetVoucherByID(id string) (*dto.VoucherDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetVoucherByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetVoucherByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockVoucherRepository) SaveVoucher(voucher dto.VoucherDB) (*dto.VoucherDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveVoucher", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	if strings.EqualFold(m.WantOutNull, "nilSaveVoucher") {
		return nil, nil
	}

	return m.WantOut, nil
}

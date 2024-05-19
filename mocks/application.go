package mocks

import (
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"strings"
)

type MockApplication struct {
	WantOut     *dto.OutputVoucher
	WantErr     error
	WantOutNull string
}

// UpdateVoucherByID implements application.Application.
func (m MockApplication) UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	panic("unimplemented")
}

func (m MockApplication) GetVoucherByID(id string) (*dto.OutputVoucher, error) {
	if m.WantErr != nil && strings.EqualFold("errGetVoucherByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetVoucherByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) SaveVoucher(reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveVoucher", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveVoucher") {
		return nil, nil
	}
	return m.WantOut, nil
}

// GetVoucherByID(id string) (*dto.OutputVoucher, error)
// SaveVoucher(reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error)
// UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)

// Repository Callers
type MockApplicationRepostoryCallers struct {
	WantOut *dto.VoucherDB
	WantErr error
}

func (m MockApplicationRepostoryCallers) GetVoucherByIDRepository(id string) (*dto.VoucherDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetVoucherByIDRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) SaveVoucherRepository(voucher dto.VoucherDB) (*dto.VoucherDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveVoucherRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

// UseCase callers
type MockApplicationUseCaseCallers struct {
	WantErr error
}

func (m MockApplicationUseCaseCallers) GetVoucherByIDUseCase(id string) error {
	if m.WantErr != nil && strings.EqualFold("errGetVoucherByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) SaveVoucherUseCase(id dto.RequestVoucher) error {
	if m.WantErr != nil && strings.EqualFold("errSaveVoucherUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

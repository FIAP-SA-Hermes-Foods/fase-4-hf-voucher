package application

import (
	"errors"
	l "fase-4-hf-voucher/external/logger"
	ps "fase-4-hf-voucher/external/strings"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/internal/core/domain/repository"
	"fase-4-hf-voucher/internal/core/domain/useCase"

	"github.com/google/uuid"
)

type Application interface {
	GetVoucherByID(id string) (*dto.OutputVoucher, error)
	SaveVoucher(reqVoucher dto.RequestVoucher) (*dto.OutputVoucher, error)
	UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error)
}

type application struct {
	voucherRepo repository.VoucherRepository
	voucherUC   useCase.VoucherUseCase
}

func NewApplication(voucherRepo repository.VoucherRepository, voucherUC useCase.VoucherUseCase) Application {
	return application{voucherRepo: voucherRepo, voucherUC: voucherUC}
}

func (app application) GetVoucherByID(id string) (*dto.OutputVoucher, error) {
	l.Infof("GetVoucherByIDApp: ", " | ", id)
	if err := app.GetVoucherByIDUseCase(id); err != nil {
		l.Errorf("GetVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	cOutDb, err := app.GetVoucherByIDRepository(id)

	if err != nil {
		l.Errorf("GetVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("GetVoucherByIDApp output: ", " | ", cOutDb)
		return nil, nil
	}

	out := &dto.OutputVoucher{
		UUID:       cOutDb.UUID,
		Code:       cOutDb.Code,
		Percentage: cOutDb.Percentage,
		CreatedAt:  cOutDb.CreatedAt,
		ExpiresAt:  cOutDb.ExpiresAt,
	}

	l.Infof("GetVoucherByIDApp output: ", " | ", ps.MarshalString(out))
	return out, err
}

func (app application) SaveVoucher(voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	l.Infof("SaveVoucherApp: ", " | ", ps.MarshalString(voucher))
	voucherWithId, err := app.GetVoucherByID(voucher.UUID)

	if err != nil {
		l.Errorf("SaveVoucherApp error: ", " | ", err)
		return nil, err
	}

	if voucherWithId != nil {
		l.Errorf("SaveVoucherApp error: ", " | ", "is not possible to save voucher because this id is already in use")
		return nil, errors.New("is not possible to save voucher because this id is already in use")
	}

	if err := app.SaveVoucherUseCase(voucher); err != nil {
		l.Errorf("SaveVoucherApp error: ", " | ", err)
		return nil, err
	}

	voucherDB := dto.VoucherDB{
		UUID:       uuid.NewString(),
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		CreatedAt:  voucher.CreatedAt,
		ExpiresAt:  voucher.ExpiresAt,
	}

	cOutDb, err := app.SaveVoucherRepository(voucherDB)

	if err != nil {
		l.Errorf("SaveVoucherApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("SaveVoucherApp output: ", " | ", nil)
		return nil, nil
	}

	out := &dto.OutputVoucher{
		UUID:       cOutDb.UUID,
		Code:       cOutDb.Code,
		Percentage: cOutDb.Percentage,
		CreatedAt:  cOutDb.CreatedAt,
		ExpiresAt:  cOutDb.ExpiresAt,
	}

	l.Infof("SaveVoucherApp output: ", " | ", ps.MarshalString(out))

	return out, nil
}

func (app application) UpdateVoucherByID(id string, voucher dto.RequestVoucher) (*dto.OutputVoucher, error) {
	l.Infof("UpdateVoucherByIDApp: ", " | ", id, " | ", ps.MarshalString(voucher))

	err := app.UpdateVoucherByIDUseCase(id, voucher)

	if err != nil {
		l.Errorf("UpdateVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	voucherDB := dto.VoucherDB{
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		CreatedAt:  voucher.CreatedAt,
		ExpiresAt:  voucher.ExpiresAt,
	}

	rVoucher, err := app.UpdateVoucherByIDRepository(id, voucherDB)

	if err != nil {
		l.Errorf("UpdateVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	vOut := dto.OutputVoucher{
		UUID:       rVoucher.UUID,
		Code:       rVoucher.Code,
		Percentage: rVoucher.Percentage,
		CreatedAt:  rVoucher.CreatedAt,
		ExpiresAt:  rVoucher.ExpiresAt,
	}

	l.Infof("UpdateVoucherByIDApp output: ", " | ", ps.MarshalString(vOut))
	return &vOut, nil
}

package application

import (
	l "fase-4-hf-voucher/external/logger"
	ps "fase-4-hf-voucher/external/strings"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/internal/core/domain/repository"
	"fase-4-hf-voucher/internal/core/domain/useCase"
	"time"

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

	if err := app.SaveVoucherUseCase(voucher); err != nil {
		l.Errorf("SaveVoucherApp error: ", " | ", err)
		return nil, err
	}

	createdAtFmt := time.Now().Format(`02-01-2006 15:04:05`)

	voucherDB := dto.VoucherDB{
		UUID:       uuid.NewString(),
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		CreatedAt:  createdAtFmt,
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

	foundVoucher, err := app.GetVoucherByID(id)

	if err != nil {
		l.Errorf("UpdateVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	if foundVoucher == nil {
		return nil, nil
	}

	if err := app.UpdateVoucherByIDUseCase(id, voucher); err != nil {
		l.Errorf("UpdateVoucherByIDApp error: ", " | ", err)
		return nil, err
	}

	var (
		code       = foundVoucher.Code
		percentage = foundVoucher.Percentage
		createdAt  = foundVoucher.CreatedAt
		expiresAt  = foundVoucher.ExpiresAt
	)

	if len(voucher.Code) != 0 {
		code = voucher.Code
	}

	if len(voucher.Percentage) != 0 {
		percentage = voucher.Percentage
	}

	if len(voucher.CreatedAt) != 0 {
		createdAt = voucher.CreatedAt
	}

	if len(voucher.Code) != 0 {
		expiresAt = voucher.ExpiresAt
	}

	voucherDB := dto.VoucherDB{
		Code:       code,
		Percentage: percentage,
		CreatedAt:  createdAt,
		ExpiresAt:  expiresAt,
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

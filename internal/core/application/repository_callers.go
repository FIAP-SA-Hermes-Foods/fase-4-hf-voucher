package application

import "fase-4-hf-voucher/internal/core/domain/entity/dto"

func (app application) GetVoucherByIDRepository(id string) (*dto.VoucherDB, error) {
	return app.voucherRepo.GetVoucherByID(id)
}

func (app application) SaveVoucherRepository(voucher dto.VoucherDB) (*dto.VoucherDB, error) {
	return app.voucherRepo.SaveVoucher(voucher)
}

func (app application) UpdateVoucherByIDRepository(id string, voucher dto.VoucherDB) (*dto.VoucherDB, error) {
	return app.voucherRepo.UpdateVoucherByID(id, voucher)
}



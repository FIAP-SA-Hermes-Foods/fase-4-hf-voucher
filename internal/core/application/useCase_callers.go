package application

import "fase-4-hf-voucher/internal/core/domain/entity/dto"

func (app application) GetVoucherByIDUseCase(id string) error {
	return app.voucherUC.GetVoucherByID(id)
}

func (app application) SaveVoucherUseCase(voucher dto.RequestVoucher) error {
	return app.voucherUC.SaveVoucher(voucher)
}

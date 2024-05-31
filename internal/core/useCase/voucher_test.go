package useCase

import (
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"log"
	"testing"
)

// go test -v -failfast -run ^Test_GetVoucherByID$
func Test_GetVoucherByID(t *testing.T) {
	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: "1000000000",
			},
			wantErr: false,
		},
		{
			name: "not_valid_code",
			args: args{
				id: "",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewVoucherUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.GetVoucherByID(tc.args.id)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_SaveVoucher$
func Test_SaveVoucher(t *testing.T) {

	type args struct {
		reqVoucher dto.RequestVoucher
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				reqVoucher: dto.RequestVoucher{
					UUID:       "",
					Code:       "abc123",
					Percentage: "50",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
			},
			wantErr: false,
		},
		{
			name: "not_valid_code",
			args: args{
				reqVoucher: dto.RequestVoucher{
					UUID:       "",
					Code:       "",
					Percentage: "",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewVoucherUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.SaveVoucher(tc.args.reqVoucher)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}

}

// go test -v -failfast -run ^Test_UpdateVoucher$
func Test_UpdateVoucher(t *testing.T) {

	type args struct {
		id         string
		reqVoucher dto.RequestVoucher
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: "1",
				reqVoucher: dto.RequestVoucher{
					UUID:       "",
					Code:       "abc123",
					Percentage: "50",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_id",
			args: args{
				reqVoucher: dto.RequestVoucher{
					UUID:       "",
					Code:       "abc123",
					Percentage: "-",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid_code",
			args: args{
				reqVoucher: dto.RequestVoucher{
					UUID:       "",
					Code:       "",
					Percentage: "-",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "not_valid_code",
			args: args{
				reqVoucher: dto.RequestVoucher{
					UUID:       "",
					Code:       "",
					Percentage: "",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewVoucherUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.UpdateVoucherByID(tc.args.id, tc.args.reqVoucher)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}

}

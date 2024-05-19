package application

import (
	"errors"
	ps "fase-4-hf-voucher/external/strings"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/mocks"
	"log"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_Get$
func Test_GetVoucherByID(t *testing.T) {
	type args struct {
		id string
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockVoucherRepository
		mockUseCase    mocks.MockVoucherUseCase
		wantOut        dto.OutputVoucher
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				id: "10000000",
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: &dto.VoucherDB{
					UUID:       "001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "2001-01-01 15:30:00",
					ExpiresAt:  "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputVoucher{
				UUID:       "001",
				Code:       "MYDISCOUNTCODE10",
				Percentage: "10",
				CreatedAt:  "2001-01-01 15:30:00",
				ExpiresAt:  "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "error_repository",
			args: args{
				id: "10000000",
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: nil,
				WantErr: errors.New("errGetVoucherByID"),
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut:       dto.OutputVoucher{},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				id: "10000000",
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: errors.New("errGetVoucherByID"),
			},
			wantOut:       dto.OutputVoucher{},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {

			out, err := app.GetVoucherByID(tc.args.id)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveVoucher$
func Test_SaveVoucher(t *testing.T) {
	type args struct {
		req dto.RequestVoucher
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockVoucherRepository
		mockUseCase    mocks.MockVoucherUseCase
		wantOut        dto.OutputVoucher
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				req: dto.RequestVoucher{
					UUID:       "001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "2001-01-01 15:30:00",
					ExpiresAt:  "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOutNull: "nilGetVoucherByID",
				WantOut: &dto.VoucherDB{
					UUID:       "001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "2001-01-01 15:30:00",
					ExpiresAt:  "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputVoucher{
				UUID:       "001",
				Code:       "MYDISCOUNTCODE10",
				Percentage: "10",
				CreatedAt:  "2001-01-01 15:30:00",
				ExpiresAt:  "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "success_null",
			args: args{
				req: dto.RequestVoucher{
					UUID:       "001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "2001-01-01 15:30:00",
					ExpiresAt:  "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputVoucher{
				UUID:      "001",
				Code:      "MYDISCOUNTCODE10",
				Percentage:       "10",
				CreatedAt: "2001-01-01 15:30:00",
				ExpiresAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},

		{
			name: "error_user_exists",
			args: args{
				req: dto.RequestVoucher{
					UUID:      "001",
					Code:      "MYDISCOUNTCODE10",
					Percentage:       "10",
					CreatedAt: "2001-01-01 15:30:00",
					ExpiresAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: &dto.VoucherDB{
					UUID:      "001",
					Code:      "MYDISCOUNTCODE10",
					Percentage:       "10",
					CreatedAt: "2001-01-01 15:30:00",
					ExpiresAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputVoucher{
				UUID:      "001",
				Code:      "MYDISCOUNTCODE10",
				Percentage:       "10",
				CreatedAt: "2001-01-01 15:30:00",
				ExpiresAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				req: dto.RequestVoucher{
					UUID:      "001",
					Code:      "MYDISCOUNTCODE10",
					Percentage:       "10",
					CreatedAt: "2001-01-01 15:30:00",
					ExpiresAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: errors.New("errSaveVoucher"),
			},
			wantOut: dto.OutputVoucher{
				UUID:      "001",
				Code:      "MYDISCOUNTCODE10",
				Percentage:       "10",
				CreatedAt: "2001-01-01 15:30:00",
				ExpiresAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_repository",
			args: args{
				req: dto.RequestVoucher{
					UUID:      "001",
					Code:      "MYDISCOUNTCODE10",
					Percentage:       "10",
					CreatedAt: "2001-01-01 15:30:00",
					ExpiresAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: nil,
				WantErr: errors.New("errSaveVoucher"),
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputVoucher{
				UUID:      "001",
				Code:      "MYDISCOUNTCODE10",
				Percentage:       "10",
				CreatedAt: "2001-01-01 15:30:00",
				ExpiresAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_getVoucherByID",
			args: args{
				req: dto.RequestVoucher{
					UUID:      "001",
					Code:      "MYDISCOUNTCODE10",
					Percentage:       "10",
					CreatedAt: "2001-01-01 15:30:00",
					ExpiresAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockVoucherRepository{
				WantOut: nil,
				WantErr: errors.New("errGetVoucherByID"),
			},
			mockUseCase: mocks.MockVoucherUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputVoucher{
				UUID:      "001",
				Code:      "MYDISCOUNTCODE10",
				Percentage:       "10",
				CreatedAt: "2001-01-01 15:30:00",
				ExpiresAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {
			out, err := app.SaveVoucher(tc.args.req)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

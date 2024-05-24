package http

import (
	"errors"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/mocks"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_Handler$
func Test_Handler(t *testing.T) {

	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name            string
		args            args
		mockApplication mocks.MockApplication
		wantOut         string
		isWantedErr     bool
	}{
		{
			name: "success_getByID",
			args: args{
				method: "GET",
				url:    "hermes_foods/voucher/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputVoucher{
					UUID:       "",
					Code:       "",
					Percentage: "",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     "{}",
			isWantedErr: false,
		},
		{
			name: "voucher_null_getByID",
			args: args{
				method: "GET",
				url:    "hermes_foods/voucher/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "voucher not found"}`,
			isWantedErr: false,
		},
		{
			name: "error_getByID",
			args: args{
				method: "GET",
				url:    "hermes_foods/voucher/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     errors.New("errGetVoucherByID"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to get voucher by ID: errGetVoucherByID"}`,
			isWantedErr: false,
		},
		{
			name: "success_save",
			args: args{
				method: "POST",
				url:    "hermes_foods/voucher",
				body:   strings.NewReader(`{"code":"MYDISCOUNTCODE10", "percentage":"10", "created_at": "2021-10-10", "expires_at": "2021-10-20"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputVoucher{
					UUID:       "0001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "2021-10-10",
					ExpiresAt:  "2021-10-20",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"uuid":"0001","code":"MYDISCOUNTCODE10","percentage":"10","created_at":"2021-10-10","expires_at":"2021-10-20"}`,
			isWantedErr: false,
		},
		{
			name: "error_save_unmarshal",
			args: args{
				method: "POST",
				url:    "hermes_foods/voucher/",
				body:   strings.NewReader(`<=>`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputVoucher{
					UUID:       "0001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to Unmarshal: invalid character '<' looking for beginning of value"}`,
			isWantedErr: true,
		},
		{
			name: "error_save",
			args: args{
				method: "POST",
				url:    "hermes_foods/voucher",
				body:   strings.NewReader(`{"code":"MYDISCOUNTCODE10", "percentage":"10", "created_at": "2021-10-10", "expires_at": "2021-10-20"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputVoucher{
					UUID:       "0001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
				WantErr:     errors.New("errSaveVoucher"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to save voucher: errSaveVoucher"}`,
			isWantedErr: false,
		},
		{
			name: "error_route_not_found",
			args: args{
				method: "PATCH",
				url:    "/hermes_foods/voucher",
				body:   strings.NewReader(`{"code":"MYDISCOUNTCODE10", "percentage":"10", "created_at": "2021-10-10", "expires_at": "2021-10-20"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputVoucher{
					UUID:       "0001",
					Code:       "MYDISCOUNTCODE10",
					Percentage: "10",
					CreatedAt:  "",
					ExpiresAt:  "",
				},
				WantErr:     errors.New("errSaveVoucher"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "route PATCH /hermes_foods/voucher not found"}`,
			isWantedErr: false,
		},
	}

	for _, tc := range tests {
		h := NewHandler(tc.mockApplication)
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, "/", tc.args.body)
			req.URL.Path = tc.args.url
			rw := httptest.NewRecorder()

			h.Handler(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}

		})
	}
}

// go test -v -count=1 -failfast -run ^Test_HealthCheck$
func Test_HealthCheck(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name            string
		args            args
		wantOut         string
		mockApplication mocks.MockApplication
		isWantedErr     bool
	}{
		{
			name: "success",
			args: args{
				method: "GET",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"status": "OK"}`,
			isWantedErr: false,
		},
		{
			name: "error_method_not_allowed",
			args: args{
				method: "POST",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"error": "method not allowed"}`,
			isWantedErr: true,
		},
	}

	for _, tc := range tests {
		h := NewHandler(tc.mockApplication)
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, tc.args.url, tc.args.body)
			rw := httptest.NewRecorder()

			h.HealthCheck(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}
		})
	}
}

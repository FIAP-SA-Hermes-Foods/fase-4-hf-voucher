package bdd

import (
	"context"
	"fase-4-hf-voucher/external/db/dynamo"
	l "fase-4-hf-voucher/external/logger"
	ps "fase-4-hf-voucher/external/strings"
	repositories "fase-4-hf-voucher/internal/adapters/driven/repositories/nosql"
	"fase-4-hf-voucher/internal/core/application"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/internal/core/useCase"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/marcos-dev88/genv"
)

// go test -v -count=1 -failfast -run ^Test_GetVoucherByID$
func Test_GetVoucherByID(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST GetVoucherByID <====")

	type Input struct {
		ID string `json:"id"`
	}

	type Output struct {
		Output *dto.OutputVoucher `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid and existing ID",
			name:     "success_valid_id",
			input: Input{
				ID: "1",
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputVoucher{
					UUID:       "1",
					Code:       "promo10",
					Percentage: "10",
					CreatedAt:  "19-05-2024 23:13:29",
					ExpiresAt:  "19-12-2024 23:13:29",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewVoucherRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewVoucherUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			voucher, err := app.GetVoucherByID(tc.input.ID)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if voucher.Code != tc.expectedOutput.Output.Code {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Code, voucher.Code)
				}

				if voucher.Percentage != tc.expectedOutput.Output.Percentage {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Percentage, voucher.Percentage)
				}

				if voucher.ExpiresAt != tc.expectedOutput.Output.ExpiresAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.ExpiresAt, voucher.ExpiresAt)
				}

				if voucher.CreatedAt != tc.expectedOutput.Output.CreatedAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.CreatedAt, voucher.CreatedAt)
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test GetVoucherByID <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveVoucher$
func Test_SaveVoucher(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST SaveVoucher <====")

	type Input struct {
		Input *dto.RequestVoucher `json:"input"`
	}

	type Output struct {
		Output *dto.OutputVoucher `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				Input: &dto.RequestVoucher{
					Code:       "promo15",
					Percentage: "15",
					ExpiresAt:  "20-05-2024 13:45:05",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputVoucher{
					Code:       "promo15",
					Percentage: "15",
					ExpiresAt:  "20-05-2024 13:45:05",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewVoucherRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewVoucherUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			voucher, err := app.SaveVoucher(*tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if voucher.Code != tc.expectedOutput.Output.Code {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Code, voucher.Code)
				}

				if voucher.Percentage != tc.expectedOutput.Output.Percentage {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Percentage, voucher.Percentage)
				}

				if voucher.ExpiresAt != tc.expectedOutput.Output.ExpiresAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.ExpiresAt, voucher.ExpiresAt)
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test SaveVoucher <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateVoucherByID$
func Test_UpdateVoucherByID(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST UpdateVoucherByID <====")

	type Input struct {
		UUID  string              `json:"uuid"`
		Input *dto.RequestVoucher `json:"input"`
	}

	type Output struct {
		Output *dto.OutputVoucher `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				UUID: "fb8e8118-bd7f-4f66-b358-acb9586d1708",
				Input: &dto.RequestVoucher{
					Code:       "promo50",
					Percentage: "50",
					ExpiresAt:  "20-05-2024 13:45:05",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputVoucher{
					Code:       "promo50",
					Percentage: "50",
					ExpiresAt:  "20-05-2024 13:45:05",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewVoucherRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewVoucherUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			voucher, err := app.UpdateVoucherByID(tc.input.UUID, *tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if voucher.Code != tc.expectedOutput.Output.Code {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Code, voucher.Code)
				}

				if voucher.Percentage != tc.expectedOutput.Output.Percentage {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Percentage, voucher.Percentage)
				}

				if voucher.ExpiresAt != tc.expectedOutput.Output.ExpiresAt {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.ExpiresAt, voucher.ExpiresAt)
				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test UpdateVoucherByID <====")
	}
}

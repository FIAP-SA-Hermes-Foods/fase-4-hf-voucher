package entity

import (
	vo "fase-4-hf-voucher/internal/core/domain/entity/valueObject"
)

type Voucher struct {
	ID        	int64        		`json:"id,omitempty"`
	Code      	string       		`json:"code,omitempty"`
	Percentage  vo.Percentage       `json:"percentage,omitempty"`
	CreatedAt 	vo.CreatedAt 		`json:"createdAt,omitempty"`
	ExpiresAt   vo.ExpiresAt      	`json:"expiresAt,omitempty"`
}

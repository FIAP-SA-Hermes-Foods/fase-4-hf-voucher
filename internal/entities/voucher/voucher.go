package voucher

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/PauloLucas94/fase-4-hf-voucher/internal/entities"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type Voucher struct {
	entities.Base
	Code       string  `json:"code"`
	Percentage float64 `json:"percentage"` // Atualizado para float64
}

func InterfaceToModel(data interface{}) (instance *Voucher, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Voucher) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Voucher) TableName() string {
	return "voucher"
}

func (p *Voucher) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Voucher) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":        p.ID.String(),
		"code":       p.Code,
		"percentage": p.Percentage,
		"createdAt":  p.CreatedAt.Format(entities.GetTimeFormat()),
		"expiresAt":  p.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}

func ParseDynamoAtributeToStruct(response map[string]*dynamodb.AttributeValue) (p Voucher, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return p, errors.New("item not found")
	}

	for key, value := range response {
		if key == "_id" {
			p.ID, err = uuid.Parse(*value.S)
			if p.ID == uuid.Nil {
				err = errors.New("item not found")
			}
		}
		if key == "code" {
			p.Code = *value.S
		}

		if key == "percentage" {
			percentage, err := strconv.ParseFloat(*value.N, 64)
			if err != nil {
				return p, err
			}
			p.Percentage = percentage
		}

		if key == "createdAt" {
			p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "expiresAt" {
			p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}

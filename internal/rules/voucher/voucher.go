package voucher

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/PauloLucas94/fase-4-hf-voucher/internal/entities"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/entities/voucher"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model)
}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	return r.createTable(connection)
}

func (r *Rules) createTable(connection *dynamodb.DynamoDB) error {
	table := &voucher.Voucher{}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(table.TableName()),
	}
	response, err := connection.CreateTable(input)
	if err != nil && strings.Contains(err.Error(), "Table already exists") {
		return nil
	}
	if response != nil && strings.Contains(response.GoString(), "TableStatus: \"CREATING\"") {
		time.Sleep(3 * time.Second)
		err = r.createTable(connection)
		if err != nil {
			return err
		}
	}
	return err
}

func (r *Rules) GetMock() interface{} {
	return voucher.Voucher{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Code: uuid.New().String(),
	}
}

func (r *Rules) Validate(model interface{}) error {
	voucherModel, err := voucher.InterfaceToModel(model)
	if err != nil {
		return err
	}

	return Validation.ValidateStruct(voucherModel,
		Validation.Field(&voucherModel.ID, Validation.Required, is.UUIDv4),
		Validation.Field(&voucherModel.Code, Validation.Required, Validation.Length(3, 50)),
	)
}

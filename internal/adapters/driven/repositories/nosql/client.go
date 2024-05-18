package reponosql

import (
	"fase-4-hf-voucher/internal/core/db"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/internal/core/domain/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var _ repository.VoucherRepository = (*voucherDB)(nil)

type voucherDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

// GetVoucherByID implements repository.VoucherRepository.

func NewVoucherRepository(database db.NoSQLDatabase, tableName string) *voucherDB {
	return &voucherDB{Database: database, tableName: tableName}
}

func (c *voucherDB) GetVoucherByID(id string) (*dto.VoucherDB, error) {
	filter := "id = :value"
	attrSearch := map[string]*dynamodb.AttributeValue{
		":value": {
			S: aws.String(id),
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(c.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := c.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var userList = make([]dto.VoucherDB, 0)
	for _, item := range result.Items {
		var c dto.VoucherDB
		if err := dynamodbattribute.UnmarshalMap(item, &c); err != nil {
			return nil, err
		}
		userList = append(userList, c)
	}

	if len(userList) > 0 {
		return &userList[0], nil
	}

	return nil, nil
}

func (c *voucherDB) SaveVoucher(voucher dto.VoucherDB) (*dto.VoucherDB, error) {

	putItem := map[string]*dynamodb.AttributeValue{
		"uuid": {
			S: aws.String(voucher.UUID),
		},
		"code": {
			S: aws.String(voucher.Code),
		},
		"percentage": {
			S: aws.String(voucher.Percentage),
		},
		"created_at": {
			S: aws.String(voucher.CreatedAt),
		},
		"expires_at": {
			S: aws.String(voucher.ExpiresAt),
		},
	}

	inputPutItem := &dynamodb.PutItemInput{
		Item:      putItem,
		TableName: aws.String(c.tableName),
	}

	putOut, err := c.Database.PutItem(inputPutItem)

	if err != nil {
		return nil, err
	}

	var out *dto.VoucherDB

	if err := dynamodbattribute.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func (c *voucherDB) UpdateVoucherByID(id string, voucher dto.VoucherDB) (*dto.VoucherDB, error) {
	updateItem := map[string]*dynamodb.AttributeValue{
		":code": {
			S: aws.String(voucher.Code),
		},
		":percentage": {
			S: aws.String(voucher.Percentage),
		},
		":created_at": {
			S: aws.String(voucher.CreatedAt),
		},
		":expires_at": {
			S: aws.String(voucher.ExpiresAt),
		},
	}

	inputUpdateItem := &dynamodb.UpdateItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(id),
			},
		},
		UpdateExpression: aws.String("set code = :code, percentage = :percentage, created_at = :created_at, expires_at = :expires_at"),
		ExpressionAttributeValues: updateItem,
		ReturnValues:              aws.String("ALL_NEW"),
	}

	updateOut, err := c.Database.UpdateItem(inputUpdateItem)
	if err != nil {
		return nil, err
	}

	var out *dto.VoucherDB

	if err := dynamodbattribute.UnmarshalMap(updateOut.Attributes, &out); err != nil {
		return nil, err
	}

	return out, nil
}

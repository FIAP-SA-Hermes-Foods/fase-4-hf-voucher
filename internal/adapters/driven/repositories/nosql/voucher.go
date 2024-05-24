package reponosql

import (
	"fase-4-hf-voucher/internal/core/db"
	"fase-4-hf-voucher/internal/core/domain/entity/dto"
	"fase-4-hf-voucher/internal/core/domain/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var _ repository.VoucherRepository = (*voucherDB)(nil)

type voucherDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

func NewVoucherRepository(database db.NoSQLDatabase, tableName string) *voucherDB {
	return &voucherDB{Database: database, tableName: tableName}
}

func (c *voucherDB) GetVoucherByID(id string) (*dto.VoucherDB, error) {
	partitionKeyName := "uuid"
	partitionKeyValue := id

	input := &dynamodb.QueryInput{
		TableName:              aws.String(c.tableName),
		KeyConditionExpression: aws.String("#pk = :value"),
		ExpressionAttributeNames: map[string]string{
			"#pk": partitionKeyName,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: partitionKeyValue},
		},
	}

	result, err := c.Database.Query(input)
	if err != nil {
		return nil, err
	}

	var userList = make([]dto.VoucherDB, 0)
	for _, item := range result.Items {
		var c dto.VoucherDB
		if err := attributevalue.UnmarshalMap(item, &c); err != nil {
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

	putItem := map[string]types.AttributeValue{
		"uuid": &types.AttributeValueMemberS{
			Value: voucher.UUID,
		},
		"code": &types.AttributeValueMemberS{
			Value: voucher.Code,
		},
		"percentage": &types.AttributeValueMemberS{
			Value: voucher.Percentage,
		},
		"createdAt": &types.AttributeValueMemberS{
			Value: voucher.CreatedAt,
		},
		"expiresAt": &types.AttributeValueMemberS{
			Value: voucher.ExpiresAt,
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

	if err := attributevalue.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	out.UUID = voucher.UUID
	out.Code = voucher.Code
	out.Percentage = voucher.Percentage
	out.CreatedAt = voucher.CreatedAt
	out.ExpiresAt = voucher.ExpiresAt

	return out, nil
}

func (c *voucherDB) UpdateVoucherByID(id string, voucher dto.VoucherDB) (*dto.VoucherDB, error) {
	update := expression.Set(expression.Name("code"), expression.Value(voucher.Code))
	update.Set(expression.Name("percentage"), expression.Value(voucher.Percentage))
	update.Set(expression.Name("createdAt"), expression.Value(voucher.CreatedAt))
	update.Set(expression.Name("expiresAt"), expression.Value(voucher.ExpiresAt))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	inputUpdateItem := &dynamodb.UpdateItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{
				Value: id,
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	updateOut, err := c.Database.UpdateItem(inputUpdateItem)
	if err != nil {
		return nil, err
	}

	var out *dto.VoucherDB

	if err := attributevalue.UnmarshalMap(updateOut.Attributes, &out); err != nil {
		return nil, err
	}

	out.UUID = id
	out.Code = voucher.Code
	out.Percentage = voucher.Percentage
	out.CreatedAt = voucher.CreatedAt
	out.ExpiresAt = voucher.ExpiresAt

	return out, nil
}

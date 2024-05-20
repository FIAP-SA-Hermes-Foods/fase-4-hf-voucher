package db

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type NoSQLDatabase interface {
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

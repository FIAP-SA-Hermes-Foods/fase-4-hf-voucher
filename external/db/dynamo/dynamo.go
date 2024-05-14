package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoDB struct {
	session *session.Session
	voucher *dynamodb.DynamoDB
}

func NewDynamoDB(session *session.Session) *dynamoDB {
	return &dynamoDB{session: session}
}

func (d *dynamoDB) voucherDynamo() {
	d.voucher = dynamodb.New(d.session)
}

func (d *dynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if d.voucher == nil {
		d.voucherDynamo()
	}
	return d.voucher.Scan(input)
}

func (d *dynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if d.voucher == nil {
		d.voucherDynamo()
	}
	return d.voucher.PutItem(input)
}

package mocks

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Mock DB NoSQL
type MockDbNoSQL struct {
	WantResultScan    *dynamodb.ScanOutput
	WantResultPutItem *dynamodb.PutItemOutput
	WantErr           error
}

// UpdateItem implements db.NoSQLDatabase.
func (m *MockDbNoSQL) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	panic("unimplemented")
}

func (m MockDbNoSQL) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errScan", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResultScan, nil

}
func (m MockDbNoSQL) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if m.WantErr != nil && strings.EqualFold("errPutItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantResultPutItem, nil
}

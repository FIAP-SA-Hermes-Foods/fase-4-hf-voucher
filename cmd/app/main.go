package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PauloLucas94/fase-4-hf-voucher/config"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/repository/adapter"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/repository/instance"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/routes"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/rules"
	RulesVoucher "github.com/PauloLucas94/fase-4-hf-voucher/internal/rules/voucher"
	"github.com/PauloLucas94/fase-4-hf-voucher/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()

	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("Waiting service starting.... ", nil)

	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migrate: ", err)
		}
	}
	logger.PANIC("", checkTables(connection))

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("Service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error

	callMigrateAndAppendError(&errors, connection, &RulesVoucher.Rules{})

	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables not found: ", nil)
		}
		for _, tableName := range response.TableNames {
			logger.INFO("Table found: ", *tableName)
		}
	}
	return err
}

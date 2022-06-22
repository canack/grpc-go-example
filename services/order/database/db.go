package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"os"
)

type dbConfigs struct {
	tableName, databaseUrl string
}

var dbConfig = dbConfigs{
	// Table name for order service's database
	tableName: "order",

	// Caution!!
	// Addresses should change in docker-compose too.
	// Otherwise, they will not be able to communicate with each other
	databaseUrl: "http://ddblocal:8000",
}

var svc *dynamodb.Client

func DBStart() {

	if env := os.Getenv("DATABASE_URL"); env != "" {
		dbConfig.databaseUrl = env
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	svc = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolver = dynamodb.EndpointResolverFromURL(dbConfig.databaseUrl)
	})
}

//func DBStartAWS() {
//	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
//		o.Region = "us-east-1"
//		return nil
//	})
//	if err != nil {
//		panic(err)
//	}
//	svc = dynamodb.NewFromConfig(cfg)
//}

// Creates table on dynamoDB local instance
func createOrderTable() error {
	_, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("OrderUUID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("CustomerUUID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("OrderUUID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{{
			IndexName: aws.String("customerIndex"),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("CustomerUUID"),
					KeyType:       types.KeyTypeHash,
				},
			},
			Projection: &types.Projection{ProjectionType: types.ProjectionTypeAll},
		}},
		TableName:   aws.String(dbConfig.tableName),
		BillingMode: types.BillingModePayPerRequest,
	})
	return err
}

func handleTableCreateErrors(err error) {
	if err != nil {
		var resourceAlreadyAvailable *types.ResourceInUseException
		var resourceUnreachable *http.ResponseError

		if errors.As(err, &resourceAlreadyAvailable) {
			return
		} else if errors.As(err, &resourceUnreachable) {
			log.Fatalln("cannot connect to database")
		} else {
			fmt.Println("an unexpected error has occurred")
			log.Fatalln(err)
		}
	}
}

func CreateTables() {
	orderErr := createOrderTable()
	handleTableCreateErrors(orderErr)
}

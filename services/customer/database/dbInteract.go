// This file responsible for connections between database for triggered functions
package database

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"
)

func (c Customer) CreateCustomer() (string, error) {

	av, _ := attributevalue.MarshalMap(c)
	_, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(dbConfig.tableName),
		Item:      av,
	})

	if err != nil {
		return "", err
	}
	return c.CustomerUUID, nil
}

func (c Customer) GetCustomer() ([]byte, error) {
	var tmp []Customer

	if c.CustomerUUID == "" {
		out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
			TableName: aws.String(dbConfig.tableName),
		})
		if err != nil {
			return []byte{}, err
		}

		attributevalue.UnmarshalListOfMaps(out.Items, &tmp)
		CustomerBytes, err := json.Marshal(tmp)
		if err != nil {
			return []byte{}, err
		}
		return CustomerBytes, nil

	} else {
		out, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:              aws.String(dbConfig.tableName),
			KeyConditionExpression: aws.String("CustomerUUID = :CustomerUUID"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":CustomerUUID": &types.AttributeValueMemberS{Value: c.CustomerUUID},
			},
		})

		if err != nil {
			return []byte{}, err
		}

		if out.Count > 0 {
			attributevalue.UnmarshalListOfMaps(out.Items, &tmp)
			CustomerBytes, err := json.Marshal(tmp)
			if err != nil {
				return []byte{}, err
			}
			return CustomerBytes, nil
		}

		return []byte{}, nil

	}
}

func (c Customer) UpdateCustomer() (bool, error) {

	CustomerBytes, err := c.GetCustomer()
	if err != nil {
		return false, err
	}
	var oldCustomer []Customer

	err = json.Unmarshal(CustomerBytes, &oldCustomer)
	if err != nil {
		return false, err
	}

	c.CreatedAt = oldCustomer[0].CreatedAt
	c.UpdatedAt = time.Now()
	c.CustomerUUID = oldCustomer[0].CustomerUUID

	if len(oldCustomer) == 0 {
		return false, nil
	}

	deleteStatus, err := c.DeleteCustomer()
	if err != nil {
		return false, err
	}
	if !deleteStatus {
		return false, nil
	}

	createStatus, err := c.CreateCustomer()
	if err != nil {
		return false, err
	}

	if len(createStatus) > 1 {
		return true, nil
	}

	return false, nil

}

func (c Customer) DeleteCustomer() (bool, error) {
	isAvailable, err := c.CheckCustomer()
	if err != nil {
		return false, err
	}
	if !isAvailable {
		return false, nil
	}

	_, err = svc.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(dbConfig.tableName),
		Key: map[string]types.AttributeValue{
			"CustomerUUID": &types.AttributeValueMemberS{Value: c.CustomerUUID},
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c Customer) CheckCustomer() (bool, error) {
	out, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(dbConfig.tableName),
		KeyConditionExpression: aws.String("CustomerUUID = :CustomerUUID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":CustomerUUID": &types.AttributeValueMemberS{Value: c.CustomerUUID},
		},
	})

	if err != nil {
		return false, err
	}

	var tmp []Customer
	attributevalue.UnmarshalListOfMaps(out.Items, &tmp)

	if len(tmp) > 0 {
		return true, nil
	}

	return false, nil
}

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

func (o Order) CreateOrder() (string, error) {

	av, _ := attributevalue.MarshalMap(o)
	_, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(dbConfig.tableName),
		Item:      av,
	})

	if err != nil {
		return "", err
	}
	return o.OrderUUID, nil
}

func (o Order) GetOrder() ([]byte, error) {
	var tmp []Order

	if o.OrderUUID == "" {
		out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
			TableName: aws.String(dbConfig.tableName),
		})
		if err != nil {
			return []byte{}, err
		}

		attributevalue.UnmarshalListOfMaps(out.Items, &tmp)
		orderBytes, err := json.Marshal(tmp)
		if err != nil {
			return []byte{}, err
		}
		return orderBytes, nil

	} else {
		out, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
			TableName:              aws.String(dbConfig.tableName),
			KeyConditionExpression: aws.String("OrderUUID = :OrderUUID"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":OrderUUID": &types.AttributeValueMemberS{Value: o.OrderUUID},
			},
		})

		if err != nil {
			return []byte{}, err
		}

		if out.Count > 0 {
			attributevalue.UnmarshalListOfMaps(out.Items, &tmp)
			orderBytes, err := json.Marshal(tmp)
			if err != nil {
				return []byte{}, err
			}
			return orderBytes, nil
		}

		return []byte{}, nil

	}
}

func (o Order) UpdateOrder() (bool, error) {

	orderBytes, err := o.GetOrder()
	if err != nil {
		return false, err
	}
	var oldOrder []Order

	err = json.Unmarshal(orderBytes, &oldOrder)
	if err != nil {
		return false, err
	}

	o.CreatedAt = oldOrder[0].CreatedAt
	o.CustomerUUID = oldOrder[0].CustomerUUID
	o.UpdatedAt = time.Now()
	o.OrderUUID = oldOrder[0].OrderUUID

	if len(oldOrder) == 0 {
		return false, nil
	}

	deleteStatus, err := o.DeleteOrder()
	if err != nil {
		return false, err
	}
	if !deleteStatus {
		return false, nil
	}

	createStatus, err := o.CreateOrder()
	if err != nil {
		return false, err
	}

	if len(createStatus) > 1 {
		return true, nil
	}

	return false, nil

}

func (o Order) DeleteOrder() (bool, error) {
	isAvailable, err := o.CheckOrder()
	if err != nil {
		return false, err
	}
	if !isAvailable {
		return false, nil
	}

	_, err = svc.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(dbConfig.tableName),
		Key: map[string]types.AttributeValue{
			"OrderUUID": &types.AttributeValueMemberS{Value: o.OrderUUID},
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (o Order) ChangeStatus() (bool, error) {

	var oldOrder []Order
	oldOrderBytes, err := o.GetOrder()
	if err != nil {
		return false, err
	}
	json.Unmarshal(oldOrderBytes, &oldOrder)

	if len(oldOrder) == 0 {
		return false, nil
	}

	oldOrder[0].Status = o.Status

	deleteStatus, err := o.DeleteOrder()
	if err != nil {
		return false, err
	}
	if !deleteStatus {
		return false, nil
	}

	oldOrder[0].UpdatedAt = time.Now()

	createStatus, err := oldOrder[0].CreateOrder()
	if err != nil {
		return false, err
	}

	if len(createStatus) > 1 {
		return true, nil
	}

	return false, nil
}

func (o Order) CheckOrder() (bool, error) {
	out, err := svc.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(dbConfig.tableName),
		KeyConditionExpression: aws.String("OrderUUID = :OrderUUID"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":OrderUUID": &types.AttributeValueMemberS{Value: o.OrderUUID},
		},
	})

	if err != nil {
		return false, err
	}

	var tmp []Order
	attributevalue.UnmarshalListOfMaps(out.Items, &tmp)

	if len(tmp) > 0 {
		return true, nil
	}

	return false, nil
}

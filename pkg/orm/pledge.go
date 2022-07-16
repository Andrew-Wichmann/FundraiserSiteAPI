// Package orm contains GoLangs binding for accessing and modifying DynamoDB documents.
package orm

import (
	"context"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Pledge struct {
	Email   string  `valid:"required,email"`
	PlanID  string  `valid:"required,plan"`
	Maximum float32 `valid:"required"`
	Rate    float32 `valid:"required"`
}

var client *dynamodb.Client

const pledgeTable = "pledge"

func init() {
	// TODO: query Dynamo for plan
	govalidator.TagMap["plan"] = govalidator.Validator(func(str string) bool {
		return true
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("personal"))
	if err != nil {
		panic(err)
	}
	client = dynamodb.NewFromConfig(cfg)
}

func (pledge Pledge) Save() error {
	valid, err := govalidator.ValidateStruct(pledge)
	if !valid {
		return err
	}
	item := map[string]types.AttributeValue{
		"email":   &types.AttributeValueMemberS{Value: pledge.Email},
		"plan_id": &types.AttributeValueMemberS{Value: pledge.PlanID},
		"rate":    &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", pledge.Rate)},
		"maximum": &types.AttributeValueMemberS{Value: fmt.Sprintf("%f", pledge.Maximum)},
	}
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{TableName: aws.String(pledgeTable), Item: item})
	return err
}

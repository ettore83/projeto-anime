package wire

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	cfg "github.com/ettore83/projeto-anime/cmd/api/internal/config"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewDb(ctx context.Context, cfg *cfg.Config) (*dynamodb.Client, error) {
	cf, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: cfg.DynamoDB.Host}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	svc := dynamodb.NewFromConfig(cf)

	err = createTables(ctx, cfg, svc)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return svc, nil
}

func createTables(ctx context.Context, cfg *cfg.Config, c *dynamodb.Client) error {
	// c.DeleteTable(ctx, &dynamodb.DeleteTableInput{TableName: aws.String(cfg.DynamoDB.Tables.Anime)})
	resp, err := c.ListTables(ctx, &dynamodb.ListTablesInput{
		Limit: aws.Int32(1),
	})

	if len(resp.TableNames) > 0 {
		return nil
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(cfg.DynamoDB.Tables.Anime),
		BillingMode: types.BillingModePayPerRequest,
	}

	_, err = c.CreateTable(ctx, input)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

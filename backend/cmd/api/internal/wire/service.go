package wire

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/ettore83/projeto-anime/cmd/api/internal/config"
	"github.com/ettore83/projeto-anime/internal/service"
)

func NewService(db *dynamodb.Client, cfg *config.Config) *service.Service {
	return &service.Service{DB: db, TableName: cfg.DynamoDB.Tables.Anime}
}

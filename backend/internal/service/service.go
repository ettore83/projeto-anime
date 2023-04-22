package service

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/ettore83/projeto-anime/anime"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service struct {
	DB        *dynamodb.Client
	TableName string
}

func NewService(db *dynamodb.Client, tableName string) *Service {
	return &Service{DB: db, TableName: tableName}
}

func (svc Service) Insert(ctx context.Context, anime *anime.Anime) error {
	now := time.Now()
	anime.ID = uuid.NewString()
	anime.CreateAt = now.Format("2006-01-02")
	anime.UpdateAt = now.Format("2006-01-02")

	item, err := attributevalue.MarshalMap(anime)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = svc.DB.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(svc.TableName), Item: item,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (svc Service) List(ctx context.Context) ([]anime.Anime, error) {
	var animes []anime.Anime

	response, err := svc.DB.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(svc.TableName),
	})

	if err != nil {
		return animes, errors.WithStack(err)
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &animes)
	if err != nil {
		return animes, errors.WithStack(err)
	}

	return animes, nil
}

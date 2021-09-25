package models

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"github.com/a180024/golang-api/dto"
)

type userRepository struct {
	db *dynamodb.DynamoDB
}

type UserRepository interface {
	Save(userDto dto.UserDto) error
	FindOneByID(id string) (*dto.UserResponseDto, error)
	FindOneByUserName(id string) (*dto.UserResponseDto, error)
}

func NewUserRepository(db *dynamodb.DynamoDB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (userRepository *userRepository) Save(userDto dto.UserDto) error {
	db := userRepository.db
	id := uuid.NewString()
	user := dto.UserResponseDto{
		ID:        id,
		UserName:  userDto.UserName,
		Password:  userDto.Password,
		Email:     userDto.Email,
		CreatedAt: time.Now().UnixNano(),
		UpdatedAt: time.Now().UnixNano(),
	}
	item, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}
	params := &dynamodb.PutItemInput{
		Item:                item,
		TableName:           aws.String("User"),
		ConditionExpression: aws.String("attribute_not_exists(username) AND attribute_not_exists(email)"),
	}
	if _, err := db.PutItem(params); err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepository) FindOneByID(id string) (*dto.UserResponseDto, error) {
	db := userRepository.db
	params := &dynamodb.GetItemInput{
		TableName: aws.String("User"),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(id),
			},
		},
		ConsistentRead: aws.Bool(true),
	}
	resp, err := db.GetItem(params)
	if err != nil {
		return nil, err
	}
	fmt.Println("Resp", resp)
	var user *dto.UserResponseDto
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
		return nil, err
	}
	return user, err
}

/* This method requires a DynamoDB GSI on the username field */
func (userRepository *userRepository) FindOneByUserName(username string) (*dto.UserResponseDto, error) {
	db := userRepository.db
	params := &dynamodb.QueryInput{
		TableName:              aws.String("User"),
		IndexName:              aws.String("username-index"),
		KeyConditionExpression: aws.String("username = :UserName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":UserName": {
				S: aws.String(username),
			},
		},
	}
	resp, err := db.Query(params)
	if err != nil {
		return nil, err
	}
	fmt.Println("Resp", resp)
	var users []dto.UserResponseDto
	if err := dynamodbattribute.UnmarshalListOfMaps(resp.Items, &users); err != nil {
		return nil, err
	} else if len(users) > 0 {
		return &users[0], nil
	}
	return nil, nil
}

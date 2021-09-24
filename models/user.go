package models

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"github.com/a180024/nft_api/dto"
)

type User struct {
	ID        string `json:"user_id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at;omitempty"`
}

type userRepository struct {
	db *dynamodb.DynamoDB
}

type UserRepository interface {
	Save(userDto dto.UserDto) error
	FindOneByID(id string) error
}

func NewUserRepository(db *dynamodb.DynamoDB) *userRepository {
	return &userRepository{
		db: db,
	}
}

/* Methods */
func (userRepository *userRepository) Save(userDto dto.UserDto) error {
	db := userRepository.db
	id := uuid.NewString()
	user := User{
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

func (userRepository *userRepository) FindOneByID(id string) error {
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
		return err
	}
	var user *User
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
		return err
	}
	return nil
}

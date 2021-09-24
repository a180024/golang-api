package models

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"github.com/a180024/nft_api/dto/users"
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
	Save(userDto users.UserDto) (*User, error)
	FindOneByID(id string) (*User, error)
}

func NewUserRepository(db *dynamodb.DynamoDB) *userRepository {
	return &userRepository{
		db: db,
	}
}

/* Methods */
func (userRepository *userRepository) Save(userDto users.UserDto) (*User, error) {
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
		log.Println(err)
		return nil, err
	}
	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("User"),
	}
	if _, err := db.PutItem(params); err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindOneByID(id string) (*User, error) {
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
		log.Println(err)
		return nil, err
	}
	var user *User
	if err := dynamodbattribute.UnmarshalMap(resp.Item, &user); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

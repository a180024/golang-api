# Go API Template

Golang Gin API with JWT authentication, DynamoDB support and Swagger documentation.

TODO:
Tests & Github Actions

## Installation

```
$ make deps
```

## Set up

- Fill in env variables in the yaml files under /config. dbEndpointURL should be localhost for running DynamoDB locally or get the region endpoints [https://docs.aws.amazon.com/general/latest/gr/ddb.html](https://docs.aws.amazon.com/general/latest/gr/ddb.html) for running on cloud.

Then, 
```
$ export AWS_ACCESS_KEY_ID=$(aws configure get aws_access_key_id)                        │
$ export AWS_SECRET_ACCESS_KEY=$(aws configure get aws_secret_access_key)
```

If running DynamoDB locally in development,

`$ docker compose up`

Create DynamoDB tables,

`$ ./create-tables.sh <dbEndpointURL>`

## Usage

```
$ make run
```

## Swagger

Go to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Credits

[https://github.com/vsouza/go-gin-boilerplate](https://github.com/vsouza/go-gin-boilerplate)

#!/bin/sh
set -e

endpoint_url=$1

echo "Creating User Table"

aws dynamodb create-table \
    --table-name User\
    --attribute-definitions AttributeName=user_id,AttributeType=S \
    --key-schema AttributeName=user_id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=5 \
    --endpoint-url ${endpoint_url}

if [ $? -eq 0 ]; then
    echo "User Table Created"
else
    echo FAIL
fi

echo "Adding GSI on username field"

aws dynamodb update-table \
    --table-name User \
    --attribute-definitions AttributeName=username,AttributeType=S \
    --global-secondary-index-updates \
        "[{\"Create\":{\"IndexName\": \"username-index\",\"KeySchema\":[{\"AttributeName\":\"username\",\"KeyType\":\"HASH\"}], \
        \"ProvisionedThroughput\": {\"ReadCapacityUnits\": 10, \"WriteCapacityUnits\": 5      },\"Projection\":{\"ProjectionType\":\"ALL\"}}}]" \
    --endpoint-url ${endpoint_url}

if [ $? -eq 0 ]; then
    echo "GSI Created"
else
    echo FAIL
fi

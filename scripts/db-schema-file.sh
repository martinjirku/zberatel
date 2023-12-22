#!/bin/bash

# Load environment variables from .env file
if [[ -f ./.env ]]; then
    export $(cat .env | xargs)
fi

# Generate schema file
pg_dump -U $DB_USER -h $DB_ADDR -p $DB_PORT -d $DB_NAME -s > ./db/schema.sql

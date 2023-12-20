#!/bin/bash

# Parse arguments
while [ "$#" -gt 0 ]; do
  case "$1" in
    --name=*)
      name="${1#*=}"
      ;;
    --db=*)
      db="${1#*=}"
      ;;
    --user=*)
      user="${1#*=}"
      ;;
    --password=*)
      password="${1#*=}"
      ;;
    --port=*)
      port="${1#*=}"
      ;;
    *)
      printf "***************************\n"
      printf "* Error: Invalid argument.*\n"
      printf "***************************\n"
      exit 1
  esac
  shift
done


if [ $(docker ps -a -q -f name=$name) ]; then
    docker start $name
else
    docker run --name $name -e POSTGRES_DB=$db -e POSTGRES_USER=$user -e POSTGRES_PASSWORD=$password -p $port:5432 -d postgres
fi
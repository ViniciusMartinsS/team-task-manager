#!/usr/bin/env bash

until nc -z -v -w30 mysql 3306
do
  echo "Waiting for database connection..."
  sleep 5
done

/task-manager/database
/task-manager/service

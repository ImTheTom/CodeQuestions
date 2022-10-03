#!/usr/bin/env bash

echo 'Database password is test'

docker exec -it ${PWD##*/}_lc-db_1 mysql -u leetcode -p

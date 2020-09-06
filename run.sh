#!/bin/bash

run() {
    source config/.env
    docker run -p $PORT:$PORT --env-file='config/.env' github-activity-log
}

run
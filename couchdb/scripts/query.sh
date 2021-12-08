#!/bin/bash

curl -X POST http://admin:password@127.0.0.1:5984/deneme/_find \
-H "Content-Type: application/json" \
-d '{ "selector": { "merhaba": { "$eq": "buraya bakarlar" }}}'
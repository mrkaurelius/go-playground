#!/bin/bash

curl -X PUT http://admin:password@127.0.0.1:5984/deneme

curl -X GET http://admin:password@127.0.0.1:5984/_all_dbs

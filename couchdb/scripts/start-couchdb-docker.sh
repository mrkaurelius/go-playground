#!/bin/bash

docker run -d --name couchdb-pg -p 5984:5984 -e COUCHDB_USER=admin -e COUCHDB_PASSWORD=password apache/couchdb:3.2.0
#!/bin/bash

# random uuid
# docId=$(cat /proc/sys/kernel/random/uuid)



curl -X PUT http://admin:password@127.0.0.1:5984/deneme/292929 -d \
'{"merhaba" : "buraya bakarlar", "counter": "1"}'
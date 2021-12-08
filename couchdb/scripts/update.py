#!/usr/bin/env python3

import requests
import json


db = "deneme" # database name (database must exits, see create_db.sh)
doc_id = "29" # doc id
url = "http://admin:password@127.0.0.1:5984/{}/{}".format(db, doc_id) # request url
print(url)

# doc
doc = {
    "merhaba": "buraya bakarlar",
    "counter": "1"
}

# Add document
r = requests.put(url, data=json.dumps(doc)) # doc in request body
print(r.json())

# Get _rev (revision)
r = requests.get(url)
print(r.json())
rev_id = r.json()["_rev"]
print("fetched revision: {}".format(rev_id))


# Updated doc
updatedDoc = {
    "_rev": rev_id,
    "merhaba": "buraya baktilar",
    "counter": "2"
}

# Update doc
update_r = requests.put(url, data=json.dumps(updatedDoc))
print(update_r.json())

# Fetch updated doc
url = "http://admin:password@127.0.0.1:5984/deneme/292929"
r = requests.get(url)
print(r.json())
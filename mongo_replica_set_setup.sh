#!/bin/bash

echo "setting up replica set"

mongo --host mongodb-rs-node-1:27017 <<EOF
  var cfg = {
    "_id": "rs0",
    "members": [
      {
        "_id": 0,
        "host": "mongodb-rs-node-1:27017",
      },
      {
        "_id": 1,
        "host": "mongodb-rs-node-2:27017",
      },
      {
        "_id": 2,
        "host": "mongodb-rs-node-3:27017",
      }
    ]
  };
  rs.initiate(cfg);
EOF
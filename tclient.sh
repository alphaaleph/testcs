#!/bin/bash

#read all the data files
for file in ./data/*
do
    echo "upload $file"
    curl -X POST http://localhost:8282/upload --data-binary @"$file"
done

curl -X POST http://localhost:8282/upload --data-binary @"./data/file7.csv"
curl -X POST http://localhost:8282/done



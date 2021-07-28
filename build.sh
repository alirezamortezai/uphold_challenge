#!/bin/bash

go test .
if [ $? -eq 0 ]; then
    echo "Tests passed!"  
    rm ./pad 2> /dev/null
    go build -o pad  .
    if [ $? -eq 0 ]; then
        echo "The application was built -> ./pad"
        exit 0
    else
        echo "Build failed"
        exit -1
    fi
else
    echo "Tests failed!"
    exit -1
fi
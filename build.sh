#!/bin/bash

PROJECT_NAME=jeanfo_mix

BUILD_DIR="build"
OUTPUT_DIR="$BUILD_DIR/output"


echo ">> cleaning old build directory..."
rm -rf $BUILD_DIR

echo ">> ceating new build directories..."
mkdir -p $OUTPUT_DIR

echo ">> building the project..."
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/$PROJECT_NAME" cmd/main.go

echo ">> copying cofig files..."
mkdir -p $OUTPUT_DIR/config
cp -r config/*.yaml $OUTPUT_DIR/config/

echo
echo ">> build complete. output directory structure:"
tree $BUILD_DIR

echo
echo ">> build completed successfully! ouput in $OUTPUT_DIR"
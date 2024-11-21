#!/bin/bash

PROJECT_NAME=jeanfo_mix

BUILD_DIR="build"
OUTPUT_DIR="$BUILD_DIR/output"


echo ">> cleaning old build directory..."
rm -rf $OUTPUT_DIR/$PROJECT_NAME

echo ">> ceating new build directories..."
mkdir -p $OUTPUT_DIR

echo ">> building the project..."
# 动态链接库方式编译
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/$PROJECT_NAME" cmd/main.go

echo ">> copying cofig files..."
mkdir -p $OUTPUT_DIR/config
# cp -r config/*.yaml $OUTPUT_DIR/config/
for file in config/*.yaml; do
    if [ -f $file ]; then
        filename=$(basename $file)
        cp $file $OUTPUT_DIR/config/${filename}.samle
    fi
done

echo
echo ">> build complete. output directory structure:"
tree $BUILD_DIR

echo
echo ">> build completed successfully! ouput in $OUTPUT_DIR"
#!/bin/bash

PROJECT_NAME_BASE=jeanfo_mix
PROJECT_NAME=$PROJECT_NAME_BASE-`date "+%Y%m%d%H%M%S"`

BUILD_DIR="build"
OUTPUT_SUB_DIR=output
OUTPUT_DIR="$BUILD_DIR/$OUTPUT_SUB_DIR"


echo ">> cleaning old build directory..."
rm -rf $OUTPUT_DIR/$PROJECT_NAME

echo ">> ceating new build directories..."
mkdir -p $OUTPUT_DIR

echo ">> building the project..."
# 动态链接库方式编译
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/$PROJECT_NAME" cmd/main.go


# if real executable file does not exists, make soft_link to newest output
if [ ! -L "$BUILD_DIR/$PROJECT_NAME_BASE" ];then
    ln -sF $OUTPUT_SUB_DIR/$PROJECT_NAME $BUILD_DIR/$PROJECT_NAME_BASE
fi


# copy newest config file sample
echo ">> copying config files..."
mkdir -p $OUTPUT_DIR/config
# cp -r config/*.yaml $OUTPUT_DIR/config/
for file in config/*.yaml; do
    if [ -f $file ]; then
        filename=$(basename $file)
        cp $file $OUTPUT_DIR/config/${filename}.samle
    fi
done


# rm old and unlinked output files
LINKED_FILE_PATH=`readlink $BUILD_DIR/$PROJECT_NAME_BASE`
LINKED_FILE=`basename $LINKED_FILE_PATH`
for file in $OUTPUT_DIR/$PROJECT_NAME_BASE*;do
    if [[ "$file" != *$LINKED_FILE* && "$file" != *$PROJECT_NAME* ]];then
        rm $file
    fi
done


echo
echo ">> build complete. output directory structure:"
tree $BUILD_DIR

echo
echo ">> build completed successfully! ouput in $OUTPUT_DIR"
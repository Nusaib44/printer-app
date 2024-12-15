#!/bin/bash

# Set output directory
OUTPUT_DIR="./build"
APP_NAME="printer-manager"

# Ensure output directory exists
mkdir -p $OUTPUT_DIR

# Platforms to build for
PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

echo "Building for multiple platforms..."

for PLATFORM in "${PLATFORMS[@]}"
do
    OS=$(echo $PLATFORM | cut -d '/' -f 1)
    ARCH=$(echo $PLATFORM | cut -d '/' -f 2)
    OUTPUT_FILE="$OUTPUT_DIR/$APP_NAME-$OS-$ARCH"

    if [ "$OS" == "windows" ]; then
        OUTPUT_FILE+='.exe'
    fi

    echo "Building for $OS/$ARCH..."
    GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT_FILE ./cmd/app
    if [ $? -ne 0 ]; then
        echo "Failed to build for $PLATFORM"
        exit 1
    fi
done

echo "Build completed. Binaries are in $OUTPUT_DIR/"

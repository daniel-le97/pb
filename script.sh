#!/bin/bash

# Define your Go application name
APP_NAME="your-app-name"

# Define the output directory
RELEASES_DIR="releases"

# Create the releases directory if it doesn't exist
mkdir -p "$RELEASES_DIR"

# Define the targets you want to build for
TARGETS=(
    "linux/amd64"
    "linux/arm"
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    # Add more targets as needed
)

# Function to log messages
log() {
    echo "[INFO] $1"
}

# Build for each target
for target in "${TARGETS[@]}"; do
    # Split the target into OS and architecture
    IFS="/" read -r -a parts <<< "$target"
    OS="${parts[0]}"
    ARCH="${parts[1]}"

    # Set the output binary name
    OUTPUT_BINARY="$RELEASES_DIR/$APP_NAME-$OS-$ARCH"

    # Build the Go app for the target
    log "Building $APP_NAME for $target..."
    env GOOS="$OS" GOARCH="$ARCH" go build -o "$OUTPUT_BINARY"

    # Add executable permissions (Linux only)
    if [ "$OS" == "linux" ]; then
        chmod +x "$OUTPUT_BINARY"
    fi

    log "Built $APP_NAME for $target"
done

log "Build process completed!"

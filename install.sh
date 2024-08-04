#!/bin/sh
set -e

# Determine system architecture
ARCH=$(uname -m)
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Set download URL based on architecture
DOWNLOAD_URL="https://github.com/theduql/duql/releases/latest/download/duql-${OS}-${ARCH}"

# Download the binary
curl -L -o duql "${DOWNLOAD_URL}"

# Make it executable
chmod +x duql

# Move to a directory in PATH
if [ "${OS}" = "darwin" ]; then
    # macOS
    sudo mv duql /usr/local/bin/
else
    # Linux
    sudo mv duql /usr/local/bin/
fi

echo "DUQL has been installed successfully!"
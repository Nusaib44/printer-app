#!/bin/bash

# Set project name
PROJECT_NAME="project-root"

# Create root directory
mkdir -p $PROJECT_NAME
cd $PROJECT_NAME

# Initialize Go module
go mod init $PROJECT_NAME

# Create main project structure
mkdir -p cmd/app
mkdir -p internal/{printers,config,platform,logging}
mkdir -p pkg/{api,common}
mkdir -p config
mkdir -p scripts
mkdir -p test
mkdir -p build

# Create main.go entry point
cat << 'EOF' > cmd/app/main.go
package main

import (
	"log"
	"os"

	"$PROJECT_NAME/internal/config"
	"$PROJECT_NAME/pkg/api"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Start API server
	log.Println("Starting server...")
	err = api.StartServer(cfg)
	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
EOF

# Create placeholders for key files
touch internal/printers/{service.go,api.go,types.go,utils.go}
touch internal/config/{config.go,defaults.go}
touch internal/platform/{windows.go,linux.go,macos.go}
touch internal/logging/logger.go
touch pkg/api/{handlers.go,middlewares.go,router.go}
touch pkg/common/{errors.go,response.go,utils.go}
touch config/{config.json,env.example}
touch scripts/{build.sh,run.sh,test.sh}
touch test/{printers_test.go,api_test.go}

# Create .gitignore
cat << 'EOF' > .gitignore
# Binaries
build/

# Environment files
.env

# Logs
*.log

# OS generated
.DS_Store
Thumbs.db
EOF

# Create README.md
cat << 'EOF' > README.md
# $PROJECT_NAME

## Overview
This is a Go-based application for managing printers across different platforms.

## Features
- Printer discovery
- Print job management
- Cross-platform support (Windows, macOS, Linux)

## Structure
- \`cmd/\`: Main application entry point
- \`internal/\`: Business logic
- \`pkg/\`: Shared packages
- \`config/\`: Configuration files
- \`scripts/\`: Build and test scripts
- \`test/\`: Unit and integration tests

## Usage
1. Build the application:
    \`\`\`bash
    bash scripts/build.sh
    \`\`\`

2. Run the application:
    \`\`\`bash
    ./build/app-linux
    \`\`\`
EOF

# Print completion message
echo "Project structure for '$PROJECT_NAME' created successfully!"

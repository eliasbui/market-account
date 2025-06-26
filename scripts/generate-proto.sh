#!/bin/bash

# gRPC Code Generation Script for Marketplace Microservices
# This script generates gRPC code for Go, C#, Python, and Node.js services

set -e

echo "🚀 Starting gRPC code generation for all services..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo -e "${RED}❌ protoc could not be found. Please install Protocol Buffers compiler.${NC}"
    echo "Install instructions:"
    echo "  macOS: brew install protobuf"
    echo "  Ubuntu: sudo apt install protobuf-compiler"
    exit 1
fi

echo -e "${BLUE}📋 protoc version: $(protoc --version)${NC}"

# Create output directories
echo -e "${YELLOW}📁 Creating output directories...${NC}"

# Go service directories
mkdir -p services/user-service-go/proto/{user,product,payment}

# C# service directories  
mkdir -p services/product-service-dotnet/Infrastructure/Proto

# Python service directories
mkdir -p services/payment-service-python/infrastructure/proto

# Node.js service directories
mkdir -p services/api-gateway-node/src/infrastructure/proto

echo -e "${GREEN}✅ Directories created${NC}"

# Generate Go code
echo -e "${YELLOW}🔨 Generating Go gRPC code...${NC}"

# Check if Go protoc plugins are installed
if ! command -v protoc-gen-go &> /dev/null; then
    echo -e "${RED}❌ protoc-gen-go not found. Installing...${NC}"
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo -e "${RED}❌ protoc-gen-go-grpc not found. Installing...${NC}"
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Generate Go proto files
protoc --go_out=services/user-service-go/proto/user --go_opt=paths=source_relative \
       --go-grpc_out=services/user-service-go/proto/user --go-grpc_opt=paths=source_relative \
       proto/user/user.proto

protoc --go_out=services/user-service-go/proto/product --go_opt=paths=source_relative \
       --go-grpc_out=services/user-service-go/proto/product --go-grpc_opt=paths=source_relative \
       proto/product/product.proto

protoc --go_out=services/user-service-go/proto/payment --go_opt=paths=source_relative \
       --go-grpc_out=services/user-service-go/proto/payment --go-grpc_opt=paths=source_relative \
       proto/payment/payment.proto

echo -e "${GREEN}✅ Go code generated${NC}"

# Generate C# code
echo -e "${YELLOW}🔨 Generating C# gRPC code...${NC}"

# Check if C# protoc plugin is available
if ! command -v grpc_csharp_plugin &> /dev/null; then
    echo -e "${YELLOW}⚠️  grpc_csharp_plugin not found. Using protoc built-in C# support...${NC}"
fi

# Generate C# proto files
protoc --csharp_out=services/product-service-dotnet/Infrastructure/Proto \
       --grpc_out=services/product-service-dotnet/Infrastructure/Proto \
       --plugin=protoc-gen-grpc=/usr/local/bin/grpc_csharp_plugin \
       proto/user/user.proto proto/product/product.proto proto/payment/payment.proto 2>/dev/null || \
protoc --csharp_out=services/product-service-dotnet/Infrastructure/Proto \
       proto/user/user.proto proto/product/product.proto proto/payment/payment.proto

echo -e "${GREEN}✅ C# code generated${NC}"

# Generate Python code
echo -e "${YELLOW}🔨 Generating Python gRPC code...${NC}"

# Generate Python proto files
python -m grpc_tools.protoc --python_out=services/payment-service-python/infrastructure/proto \
                            --grpc_python_out=services/payment-service-python/infrastructure/proto \
                            --proto_path=. \
                            proto/user/user.proto proto/product/product.proto proto/payment/payment.proto

echo -e "${GREEN}✅ Python code generated${NC}"

# Generate Node.js code
echo -e "${YELLOW}🔨 Generating Node.js gRPC code...${NC}"

# Check if Node.js grpc-tools is available
if ! npm list -g grpc-tools &> /dev/null; then
    echo -e "${YELLOW}⚠️  grpc-tools not found globally. Installing...${NC}"
    npm install -g grpc-tools
fi

# Generate Node.js proto files
npx grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:services/api-gateway-node/src/infrastructure/proto \
    --grpc_out=grpc_js:services/api-gateway-node/src/infrastructure/proto \
    --proto_path=. \
    proto/user/user.proto proto/product/product.proto proto/payment/payment.proto

echo -e "${GREEN}✅ Node.js code generated${NC}"

# Summary
echo -e "${BLUE}📊 Generation Summary:${NC}"
echo -e "${GREEN}✅ Go service: gRPC code generated in services/user-service-go/proto/${NC}"
echo -e "${GREEN}✅ C# service: gRPC code generated in services/product-service-dotnet/Infrastructure/Proto/${NC}"
echo -e "${GREEN}✅ Python service: gRPC code generated in services/payment-service-python/infrastructure/proto/${NC}"
echo -e "${GREEN}✅ Node.js service: gRPC code generated in services/api-gateway-node/src/infrastructure/proto/${NC}"

echo -e "${GREEN}🎉 All gRPC code generation completed successfully!${NC}"
echo -e "${BLUE}💡 Next steps:${NC}"
echo "  1. Update service implementations to use generated gRPC clients/servers"
echo "  2. Configure service discovery and load balancing"
echo "  3. Implement error handling and retry logic"
echo "  4. Add health checks and monitoring" 
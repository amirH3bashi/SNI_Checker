#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${GREEN}🔧 Installing SNI Checker...${NC}"

# Detect package manager
detect_package_manager() {
    if command -v apt &>/dev/null; then
        echo "apt"
    elif command -v dnf &>/dev/null; then
        echo "dnf"
    elif command -v yum &>/dev/null; then
        echo "yum"
    elif command -v pacman &>/dev/null; then
        echo "pacman"
    elif command -v apk &>/dev/null; then
        echo "apk"
    else
        echo ""
    fi
}

PM=$(detect_package_manager)

if [ -z "$PM" ]; then
    echo -e "${RED}❌ Unsupported package manager. Install Go manually: https://golang.org/doc/install${NC}"
    exit 1
fi

# Install dependencies
echo -e "${GREEN}📦 Installing dependencies...${NC}"

case "$PM" in
    apt)
        sudo apt update
        sudo apt install -y git curl build-essential iputils-ping wget
        ;;
    dnf)
        sudo dnf install -y git curl gcc make iputils wget
        ;;
    yum)
        sudo yum install -y git curl gcc make iputils wget
        ;;
    pacman)
        sudo pacman -Sy --noconfirm git curl base-devel iputils wget
        ;;
    apk)
        sudo apk add --no-cache git curl build-base iputils wget
        ;;
esac

# Install Go if not available
if ! command -v go &>/dev/null; then
    echo -e "${GREEN}⬇️ Installing Go...${NC}"
    GO_VERSION="1.22.3"
    wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    rm go${GO_VERSION}.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
fi

# Clone the repository
echo -e "${GREEN}📥 Cloning repository...${NC}"
rm -rf SNI_Checker
git clone https://github.com/amirH3bashi/SNI_Checker.git
cd SNI_Checker

# Build the binary
echo -e "${GREEN}🔨 Building project...${NC}"
go mod tidy
go build -o SNI_Checker main.go

# Run the binary
echo -e "${GREEN}✅ Installation complete!${NC}"
echo -e "${GREEN}🚀 Running SNI Checker...${NC}"
./SNI_Checker

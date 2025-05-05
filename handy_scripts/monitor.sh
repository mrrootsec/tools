#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" &> /dev/null
}

# Function to install Go if it's not installed
install_go() {
    echo "[+] - Golang is not installed. Installing..."
    sudo apt-get update
    sudo apt-get install -y golang
}

# Check if Golang is installed
if ! command_exists go; then
    install_go
else
    echo "[+] - Golang is already installed"
    go version
fi

# Check and install Go tools if they are not already installed
for tool in "gungnir" "notify" "anew"; do
    if ! command_exists "$tool"; then
        echo "[+] - $tool is not installed. Installing..."
        case "$tool" in
            "gungnir")
                go install github.com/g0ldencybersec/gungnir/cmd/gungnir@latest
                ;;
            "notify")
                go install github.com/projectdiscovery/notify/cmd/notify@latest
                ;;
            "anew")
                go install github.com/tomnomnom/anew@latest
                ;;
        esac
    else
        echo "[+] - $tool is already installed"
    fi
done

echo "[+] - Running gungnir..."

# Continuous monitoring pipeline
gungnir -r domains.txt | \
stdbuf -oL anew all_domains.txt | \
stdbuf -oL notify -silent -id subs -provider discord | \
stdbuf -oL awk '{ print strftime("%Y-%m-%d %H:%M:%S")","$0 }' | \
tee -a "daily_domains_$(date +%Y-%m-%d).txt"

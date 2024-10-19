# Use Ubuntu 20.04 as the base image
FROM ubuntu:20.04

# Set environment variable to non-interactive to prevent prompts
ENV DEBIAN_FRONTEND=noninteractive

# Update package list and install dependencies
RUN apt-get update && apt-get install -y \
    libcap-dev \
    asciidoc-base \
    libsystemd-dev \
    pkg-config \
    git \
    build-essential
    
# Clone the isolate repository from GitHub
RUN git clone https://github.com/ioi/isolate.git /opt/isolate

# Set working directory to the isolate folder
WORKDIR /opt/isolate

# Build and install isolate
RUN make && make install

# Clean up APT cache to reduce image size
RUN apt-get clean && rm -rf /var/lib/apt/lists/*


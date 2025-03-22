FROM jenkins/jenkins:lts

USER root

# Install required tools and dependencies
RUN apt-get update && apt-get install -y \
    git \
    wget \
    tar \
    curl \
    apt-transport-https \
    ca-certificates \
    software-properties-common && \
    rm -rf /var/lib/apt/lists/*

# Install Docker 
RUN apt-get update && apt-get install -y docker.io
RUN usermod -aG docker jenkins

# Pre-install Go
ENV GO_VERSION=1.24.1
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Add Go to PATH
ENV PATH="/usr/local/go/bin:${PATH}"

# Adjust permissions for /usr/local
RUN chown -R jenkins:jenkins /usr/local

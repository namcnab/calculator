FROM jenkins/jenkins:lts

USER root

# Install required tools and dependencies
RUN apt-get update && apt-get install -y \
    git \
    wget \
    tar \
    curl \
    gnupg2 \
    lsb-release && \
    rm -rf /var/lib/apt/lists/*

# Install Docker CLI
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list && \
    apt-get update && apt-get install -y docker-ce-cli && \
    rm -rf /var/lib/apt/lists/*

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin/

# Pre-install Go
ENV GO_VERSION=1.24.1
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Adjust permissions for /usr/local
RUN chown -R jenkins:jenkins /usr/local

# Create the docker group
RUN groupadd docker 

# Add jenkins user to docker group
RUN gpasswd -a jenkins docker

# Add Go to PATH
ENV PATH="/usr/local/go/bin:${PATH}"

USER jenkins
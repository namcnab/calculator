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

# Install Docker CLI
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    echo "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list && \
    apt-get update && apt-get install -y docker-ce-cli && \
    rm -rf /var/lib/apt/lists/*

# Pre-install Go
ENV GO_VERSION=1.24.1
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Add Go to PATH
ENV PATH="/usr/local/go/bin:${PATH}"

# Adjust permissions for /usr/local
RUN chown -R jenkins:jenkins /usr/local

# Add Jenkins user to the existing docker group or create it if it doesn't exist
RUN if ! getent group docker; then groupadd docker; fi && \
    usermod -aG docker jenkins

# Ensure the Jenkins user has access to the Docker socket
RUN chmod 666 /var/run/docker.sock || true

USER jenkins
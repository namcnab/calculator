# Calculator Application

This application performs basic mathematics. The main purpose of this project is build a CI/CD pipeline.

## Tech Stack

- **Source Control**: Github
- **Continuous Integration (CI) Server**: Jenkins
- **Artifact Repository**: Dockerhub
- **Deployment Automation**: Kubernetes
- **Security**: SonarQube
- **Configuration Management**: Helm
- **Monitor**: Prometheus
- **Observability**: EFK Stack

### Github

- Stores application source code and configuration files
- Provides version control and history tracking
- Enables collaborative development through pull requests
- Triggers CI/CD pipelines on code changes
- Manages branching strategies and release workflows

### Jenkins

- Automates build processes when code changes are detected
- Executes automated tests to validate code quality
- Compiles code and packages applications
- Orchestrates the entire CI/CD pipeline workflow
- Provides visibility into build and deployment statuse

### Dockerhub

- Stores and manages build artifacts and dependencies
- Provides version control for binaries and packages
- Ensures consistent deployment artifacts across environments
- Scans artifacts for security vulnerabilities
- Manages Docker images and other container formats

### Kubernetes

- Orchestrates containerized application deployment
- Provides scaling, load balancing, and self-healing capabilities
- Manages application lifecycle across environments
- Enables declarative configuration for applications
- Supports rolling updates and deployment strategies

### SonarQube

- Performs static code analysis to identify code quality issues
- Detects security vulnerabilities in source code
- Enforces coding standards and best practices
- Tracks code coverage from automated tests
- Prevents security issues early in the development cycle

### Helm

- Packages Kubernetes applications as charts
- Manages complex application deployments
- Provides templating for environment-specific configurations
- Enables versioning and rollback of deployments
- Simplifies application updates and management

### Prometheus

- Collects and stores time-series metrics data
- Monitors application and infrastructure health
- Provides alerting based on metric thresholds
- Supports service discovery for dynamic environments
- Enables real-time monitoring of application performance

### EFK (Elasticsearch, Fluentd, Kibana) Stack

- Aggregates logs from all application components
- Provides centralized log storage and indexing
- Enables powerful log searching and filtering
- Creates visualizations and dashboards for log analysis
- Helps troubleshoot issues by correlating logs across services

## Stack Integration

Here's how each of the CI/CD components integrate with each other to form a complete pipeline:

### Source Control (GitHub) integrates with:

- **Jenkins**: Triggers builds via webhooks when code is pushed
- **SonarQube**: Provides code for static analysis and receives quality gate results
- **Helm**: Stores Helm charts as code for Kubernetes deployments
- **Kubernetes**: Provides application code and deployment manifests

### CI Server (Jenkins) integrates with:

- **GitHub**: Pulls source code and reports build status
- **JFrog**: Publishes and retrieves artifacts
- **SonarQube**: Initiates code scans and receives quality reports
- **Kubernetes**: Deploys applications using kubectl or Helm
- **Prometheus**: Exposes build metrics and job status
- **Helm**: Executes chart deployments within pipeline stages

### Artifact Repository (JFrog) integrates with:

- **Jenkins**: Stores and provides build artifacts
- **Kubernetes**: Supplies container images for deployment
- **SonarQube**: Scans stored artifacts for vulnerabilities
- **Helm**: Stores Helm charts for application deployment

### Deployment Automation (Kubernetes) integrates with:

- **Helm**: Consumes charts for application deployment
- **Prometheus**: Exposes metrics endpoints for monitoring
- **EFK Stack**: Forwards container and application logs
- **JFrog**: Pulls container images during deployment

### Security (SonarQube) integrates with:

- **GitHub**: Reviews code quality in pull requests
- **Jenkins**: Provides quality gates for pipeline progression
- **JFrog**: Scans artifacts for security issues

### Configuration Management (Helm) integrates with:

- **GitHub**: Stores Helm charts as code
- **Jenkins**: Used in deployment stages
- **Kubernetes**: Deploys templated applications
- **JFrog**: Stores and retrieves chart packages

### Monitoring (Prometheus) integrates with:

- **Kubernetes**: Scrapes metrics from pods and nodes
- **Jenkins**: Monitors build metrics and job health
- **EFK Stack**: Correlates metrics with logs for troubleshooting

### Observability (EFK Stack) integrates with:

- **Kubernetes**: Collects pod and application logs
- **Prometheus**: Correlates logs with metrics for incident analysis
- **Jenkins**: Captures build logs and pipeline execution details

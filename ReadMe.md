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

## Stack Integration

Here's how each of the CI/CD components integrate with each other to form a complete pipeline:

### Source Control (GitHub) integrates with:

- **Jenkins**: Triggers builds via webhooks when code is pushed
- **SonarQube**: Provides code for static analysis and receives quality gate results
- **Helm**: Stores Helm charts as code for Kubernetes deployments
- **Kubernetes**: Provides application code and deployment manifests

### CI Server (Jenkins) integrates with:

Jenkins Endpoint: http://localhost:8080

- **GitHub**: Pulls source code and reports build status
- **Docker Hub**: Publishes and retrieves artifacts
- **SonarQube**: Initiates code scans and receives quality reports
- **Kubernetes**: Deploys applications using kubectl or Helm
- **Prometheus**: Exposes build metrics and job status
- **Helm**: Executes chart deployments within pipeline stages

### Artifact Repository (Docker Hub) integrates with:

Docker Hub Repo: namcnab/calculator-api

- **Jenkins**: Stores and provides build artifacts
- **Kubernetes**: Supplies container images for deployment
- **SonarQube**: Scans stored artifacts for vulnerabilities
- **Helm**: Stores Helm charts for application deployment

### Deployment Automation (Kubernetes) integrates with:

Kubernetes Endpoint: https://host.docker.internal:51358
Rancher Endpoint: https://localhost:8443/

- **Helm**: Consumes charts for application deployment
- **Prometheus**: Exposes metrics endpoints for monitoring
- **EFK Stack**: Forwards container and application logs
- **Docker Hub**: Pulls container images during deployment

### Security (SonarQube) integrates with:

- **GitHub**: Reviews code quality in pull requests
- **Jenkins**: Provides quality gates for pipeline progression
- **Docker Hub**: Scans artifacts for security issues

### Configuration Management (Helm) integrates with:

- **GitHub**: Stores Helm charts as code
- **Jenkins**: Used in deployment stages
- **Kubernetes**: Deploys templated applications
- **Docker Hub**: Stores and retrieves chart packages

### Monitoring (Prometheus) integrates with:

- **Kubernetes**: Scrapes metrics from pods and nodes
- **Jenkins**: Monitors build metrics and job health
- **EFK Stack**: Correlates metrics with logs for troubleshooting

### Observability (EFK Stack) integrates with:

Elasticsearch Endpoint: http://localhost:9200
Kibana Endpoint: http://localhost:5601

- **Kubernetes**: Collects pod and application logs
- **Prometheus**: Correlates logs with metrics for incident analysis
- **Jenkins**: Captures build logs and pipeline execution details

pipeline {
    agent any

    tools {
        git 'Default' // Use the default Git installation configured in Jenkins
    }
    
    environment {
        GO_VERSION = "1.20" // Specify the Go version
        DOCKER_IMAGE = "calculator-api:latest" // Docker image name
    }

    stages {
        stage('Checkout') {
            steps {
                // Clone the repository
                checkout scm
            }
        }

        stage('Setup Go Environment') {
            steps {
                // Verify Go installation
                sh '''
                go version
                '''
            }
        }

        stage('Test') {
            steps {
                // Run unit tests
                sh '''
                go test -v ./...
                '''
            }
        }

        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build -v -o calculator-api ./cmd'
            }
        }

        stage('Build Docker Image') {
            steps {
                // Build a Docker image for the application
                sh '''
                docker build -t $DOCKER_IMAGE .
                '''
            }
        }

        stage('Push Docker Image') {
            steps {
                // Push the Docker image to a registry (e.g., Docker Hub)
                withCredentials([usernamePassword(credentialsId: 'docker-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    docker tag $DOCKER_IMAGE namcnab/$DOCKER_IMAGE
                    docker push namcnab/$DOCKER_IMAGE
                    '''
                }
            }
        }

        stage('Deploy') {
            steps {
                // Deploy the application (e.g., using Kubernetes or SSH)
                sh '''
                echo "Deploying application..."
                # Example: kubectl apply -f deployment.yaml
                '''
            }
        }

        stage('Package') {
            steps {
                // Package the application (optional)
                sh 'tar -czf app.tar.gz ./'
            }
        }
    }

    post {
        always {
            // Clean up workspace
            cleanWs()
        }
    }
}
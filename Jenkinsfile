pipeline {
    agent any

    environment {
        GO_VERSION = "1.20" // Specify the Go version
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
                // Install Go and set up the environments
                sh '''
                wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
                tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
                export PATH=$PATH:/usr/local/go/bin
                go version
                '''
            }
        }

        stage('Build') {
            steps {
                // Build the Go application
                sh 'go build -v ./...'
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
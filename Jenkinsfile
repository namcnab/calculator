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
                // Verify Go installation
                sh '''
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

        stage('Test') {
            steps {
                // Run tests
                sh 'go test -v ./...'
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
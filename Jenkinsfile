pipeline {
    agent any

    tools {
        git 'Default' // Use the default Git installation configured in Jenkins
    }
    
    environment {
        GO_VERSION = "1.20" // Specify the Go version
        DOCKER_IMAGE = "calculator-api:latest" // Docker image name defined
    }

    stages {
       /* stage('Checkout') {
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

        stage('Deploy to Kubernetes') {
            steps {
                // Deploy the Docker image to a Kubernetes cluster 
                
               
            }
        }*/

        stage('Debug Kubernetes') {
            steps {
                withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
                    sh '''
                    echo "Using KUBECONFIG: $KUBECONFIG"
                    kubectl config view || echo "Failed to view config"
                    kubectl get pods || echo "Failed to get pods"
                    kubectl get services || echo "Failed to get services"
                    kubectl get deployments || echo "Failed to get deployments"
                    '''
                }
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
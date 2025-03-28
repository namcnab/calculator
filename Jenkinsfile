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
                withKubeConfig([credentialsId: 'k8_token', serverUrl: 'https://host.docker.internal:51358']) {
                    // Remove existing deployment and service
                    sh 'kubectl delete deployment calculator-api-deployment --ignore-not-found=true'
                    sh 'kubectl delete service calculator-api-service --ignore-not-found=true'

                    sh 'kubectl apply -f ./configs/kubernetes/namespace.yaml'
                    sh 'kubectl apply -f ./configs/kubernetes/clusterrole.yaml'
                    sh 'kubectl apply -f ./configs/kubernetes/clusterrolebinding.yaml'
                    sh 'kubectl apply -f ./configs/kubernetes/jenkins-token.yaml'
                    sh 'kubectl apply -f ./configs/kubernetes/calculator-api-deployment.yaml'
                    sh 'kubectl apply -f ./configs/kubernetes/calculator-api-service.yaml'
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
pipeline {
    agent any

    tools {
        git 'Default' // Use the default Git installation configured in Jenkins
    }
    
    environment {
        GO_VERSION = "1.20" // Specify the Go version
        CALCULATOR_DOCKER_IMAGE = "calculator-api:latest" // Docker image name defined
        FLUENTD_DOCKER_IMAGE = "fluentd-elasticsearch:latest"
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

        stage('Build Go App') {
            steps {
                sh 'make build_go'
            }
        }

        stage('Build Docker Images') {
            steps {
                sh 'make build_docker_images'
            }
        }

        stage('Push Docker Images') {
            steps {
                // Push the Docker image to a registry (e.g., Docker Hub)
                withCredentials([usernamePassword(credentialsId: 'docker-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                    echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
                    docker tag $CALCULATOR_DOCKER_IMAGE namcnab/$CALCULATOR_DOCKER_IMAGE
                    docker push namcnab/$CALCULATOR_DOCKER_IMAGE
                    '''
                }
            }
        }

        stage('Setup Kubernetes Environment'){
            steps {
                withKubeConfig([credentialsId: 'k8_token', serverUrl: 'https://host.docker.internal:51358']) {
                    sh 'make setup_kubernetes_env'  
                }
            }
        }

         stage('Deploys EFK Stack to Kubernetes') {
             steps {
                withKubeConfig([credentialsId: 'k8_token', serverUrl: 'https://host.docker.internal:51358']) {        
                   sh 'make deploy_efk_stack'
                }
             }
        }

        stage('Deploys Calculator API to Kubernetes') {
             steps {
                withKubeConfig([credentialsId: 'k8_token', serverUrl: 'https://host.docker.internal:51358']) {
                   sh 'make deploy_calc_api'
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
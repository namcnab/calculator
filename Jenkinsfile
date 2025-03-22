pipeline {
    agent {
        kubernetes {
            label 'k8s-agent'
            defaultContainer 'jnlp'
            yaml """
                apiVersion: v1
                kind: Pod
                metadata:
                labels:
                    app: jenkins-agent
                spec:
                containers:
                    - name: jnlp
                    image: jenkins/inbound-agent:latest
                    args: ['\$(JENKINS_SECRET)', '\$(JENKINS_NAME)']
                    - name: docker
                    image: docker:20.10.24
                    command:
                        - cat
                    tty: true
                    volumeMounts:
                        - name: docker-sock
                        mountPath: /var/run/docker.sock
                    - name: kubectl
                    image: bitnami/kubectl:latest
                    command:
                        - cat
                    tty: true
                volumes:
                    - name: docker-sock
                    hostPath:
                        path: /var/run/docker.sock
                """
                        }
                    }

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
                container('docker') {
                    // Build a Docker image for the application
                    sh '''
                    docker build -t $DOCKER_IMAGE .
                    '''
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                container('docker') {
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
        }

        stage('Deploy') {
            steps {
                container('kubectl') {
                    // Deploy the application using Kubernetes
                    sh '''
                    echo "Deploying application to Kubernetes..."
                    kubectl apply -f configs/kubernetes/calculator-api-deployment.yaml
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
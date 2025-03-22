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
        }*/

        stage('Deploy to Kubernetes') {
            steps {
               script {
                  kubeconfig(
                    caCertificate: 'LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURCakNDQWU2Z0F3SUJBZ0lCQVRBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwdGFXNXAKYTNWaVpVTkJNQjRYRFRJMU1ETXlNVEU0TkRnMU1Wb1hEVE0xTURNeU1ERTRORGcxTVZvd0ZURVRNQkVHQTFVRQpBeE1LYldsdWFXdDFZbVZEUVRDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTUx2Ck5nMDczZ2hCVW80TlpyZjl4WmR6UWR6VG1MYzZaRkE5M1pDeUFjUnZkNUgxRWZJSUJMTngyQmtzdm9lZXpkR2EKOU5UY3VDS0FENVRXQmdIK2RVQm5KbmR2dHd1cGZPTCtuSGVNaVZuWDBFbUtNWGFFbnlOdE5uVkhhb0k5L29ubAp3WVllbFlFNEhLT3BXb0trTUFjVXJTSjRkMVFlY2Vyb29JZFgrZlRPWlIyNzJiUkh1bG96Sno0R1U4cUhJdkFNCmZINEgrc1NYR2kyTVdtMzVjQjVHaU9QN2MyOUJ0dzJKZ05Bb1NHRnphcFYzd0xnY0QvY3BhS3NTQnVPZS9ZLzIKd1AzVWFkbmhVVE1SclJQMDFBdzF5d2lXd2tIbU0wcUM1cjZtZ1JMankrREhDL241RFhCenRRTVgwTzBaSnZ4cQowMVduOS9sMFowaGxLMkR3ckpVQ0F3RUFBYU5oTUY4d0RnWURWUjBQQVFIL0JBUURBZ0trTUIwR0ExVWRKUVFXCk1CUUdDQ3NHQVFVRkJ3TUNCZ2dyQmdFRkJRY0RBVEFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQjBHQTFVZERnUVcKQkJSVW50SzB5WjZmR3dpTFd0Y0RYcnRmODFzRUhqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFIZFdIREloSwo4YTNoVkVFREwvTGNxWXVJZitURlZkQlVWVkord0Jqb3FFeXdYTWVHQ28rNFphSVR0ZnZ6Z0srN21STmRUeW05CnFPOW10WG1qK3poaGhLTXNlYTYweHNra1lBU0ZpNHJEM0JJeGVRbFRHYmE2L2w4VFk5U25nY3UxVkxIRFF5WnAKZ3FoV1hQSnlIVHF6S0RjdWY2ZlN4dlh4OWZmVUY2Q3d3YytZMktCQ0lENXJPeGlmU0dTbGhDOHE3SlBkd2FkVQpYU1kzM1lvelZDWkgyaUdqMThFV000NkgwbjhIaGFRbjJtQVF6ZnV5OHI0MWp4MHlRMnJxbFAraDI4NzBvL3VkCjRWMlNZMUJVWHBrWmpPV0JOK2g0MTNDbUxtM1Vad3FTRzFlRmRYYVdva3dkWkwvbWZTMEczbERsaFAwaVdmaDYKYUh4ZG9VeUZkTG13Rmc9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==', 
                    credentialsId: 'kubeconfig', 
                    serverUrl: 'https://127.0.0.1:51358') {
                    sh ''
                }
                 
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
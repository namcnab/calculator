# VARIABLES
CALCULATOR_DOCKER_IMAGE = calculator-api
DOCKER_TAG = latest

# CI/CD TASKS
start_cicd: start_jenkins start_kubernetes start_rancher

start_jenkins:
	docker compose -f ./configs/jenkins/docker-compose-jenkins.yaml down
	docker compose -f ./configs/jenkins/docker-compose-jenkins.yaml build
	docker compose -f ./configs/jenkins/docker-compose-jenkins.yaml up -d

start_kubernetes:
	minikube start

start_rancher:
	docker compose -f ./configs/kubernetes/rancher/docker-compose-rancher.yaml down
	docker compose -f ./configs/kubernetes/rancher/docker-compose-rancher.yaml up -d

start_sonarqube:
	docker compose -f ./configs/sonarqube/docker-compose-sonarqube.yaml down
	docker compose -f ./configs/sonarqube/docker-compose-sonarqube.yaml up

build_go:
	go build -v -o calculator-api ./cmd

build_docker_images:
	docker build -t $(CALCULATOR_DOCKER_IMAGE):$(DOCKER_TAG) .

setup_kubernetes_env:
	kubectl apply -f ./configs/kubernetes/namespace.yaml
	kubectl apply -f ./configs/kubernetes/clusterrole.yaml      
	kubectl apply -f ./configs/kubernetes/clusterrolebinding.yaml
	kubectl apply -f ./configs/kubernetes/jenkins-token.yaml

revert_kubernetes_env:
	kubectl delete -f ./configs/kubernetes/jenkins-token.yaml
	kubectl delete -f ./configs/kubernetes/clusterrolebinding.yaml
	kubectl delete -f ./configs/kubernetes/clusterrole.yaml
	kubectl delete -f ./configs/kubernetes/namespace.yaml

deploy_efk_stack:
	kubectl apply -f ./configs/efk/k8s/

revert_efk_kubernetes:
	kubectl delete -f ./configs/efk/k8s/

deploy_calc_api:                    
	kubectl apply -f ./configs/kubernetes/calculator-api-deployment.yaml
	kubectl apply -f ./configs/kubersnetes/calculator-api-service.yaml

revert_calc_api:
	kubectl delete -f ./configs/kubernetes/calculator-api-deployment.yaml
	kubectl delete -f ./configs/kubernetes/calculator-api-service.yaml

# DOCKER TASKS
start_efk_services:
	chmod +x ./configs/efk/start-efk-services.sh
	./configs/efk/start-efk-services.sh




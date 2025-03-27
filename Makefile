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
	helm repo list | grep -q "rancher-stable" || helm repo add rancher-stable https://releases.rancher.com/server-charts/stable
	kubectl get namespace cattle-system >/dev/null 2>&1 || kubectl create namespace cattle-system
	helm repo list | grep -q "jetstack" || helm repo add jetstack https://charts.jetstack.io
	helm repo update
	helm list -n cert-manager | grep -q "cert-manager" \
        || helm install cert-manager jetstack/cert-manager \
            --namespace cert-manager \
            --create-namespace \
            --set installCRDs=true
	[ -d rancher ] && rm -rf rancher || true
	helm pull rancher-stable/rancher --untar
	sed -i '' 's/kubeVersion: < 1.32.0-0/kubeVersion: \">=1.20.0 <1.33.0\"/' ./rancher/Chart.yaml
	helm list -n cattle-system | grep -q "rancher" \
        && echo "Rancher is already installed. Skipping installation." \
        || (helm install rancher ./rancher \
            --namespace cattle-system \
            --set hostname=rancher.my.org \
            --set bootstrapPassword=admin \
			--set cleanup.finalizers=false)
	kubectl -n cattle-system rollout status deploy/rancher --timeout=10m
	kubectl get svc -n cattle-system | grep -q "rancher-lb" \
        || kubectl expose deployment rancher --name=rancher-lb --port=443 --type=LoadBalancer -n cattle-system
	kubectl get svc -n cattle-system
	nohup kubectl port-forward service/rancher-lb 9443:443 -n cattle-system > port-forward.log 2>&1 &
	rm -rf rancher

stop_rancher:
	kubectl delete deployment rancher -n cattle-system || true
	
	kubectl delete namespace cattle-system || true

restart_rancher_lb:
	kubectl delete svc rancher-lb -n cattle-system || true
	kubectl expose deployment rancher --name=rancher-lb --port=443 --type=LoadBalancer -n cattle-system
	nohup kubectl port-forward service/rancher-lb 9443:443 -n cattle-system > port-forward.log 2>&1 &

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
	helm install efk-stack ./configs/kubernetes/helm/efk-stack \
    --namespace observability \
	--create-namespace \
    --set service.port=9200 \
	

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




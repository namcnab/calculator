build_fluentd:
	docker build --tag fluentd-elasticsearch ./configs/efk/fluentd

start_efk_services:
	chmod +x ./configs/efk/start-efk-services.sh
	./configs/efk/start-efk-services.sh

start_cicd:
	docker compose -f ./configs/jenkins/docker-compose-jenkins.yaml up -d
	
build_go_app:
	docker build --tag calculator-api .

start_go_app:
	docker compose -f docker-compose-go-app.yaml down
	docker compose -f docker-compose-go-app.yaml up -d

run_pipeline: start_efk_services start_go_app

clean_run_pipeline: build_fluentd start_efk_services build_go_app start_go_app


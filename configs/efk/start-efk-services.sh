#!/bin/bash

# Start all services
docker-compose -f ./configs/efk/docker-compose-efk.yaml down
docker-compose -f ./configs/efk/docker-compose-efk.yaml up -d

# Wait for Elasticsearch to be ready
printf "Waiting for Elasticsearch to be ready...\n"
until curl -s http://localhost:9200/_cluster/health | grep -q '"status":"green"\|"status":"yellow"'; do
  sleep 2
done

# Load the index template
printf "Elasticsearch is ready. Loading the index template...\n"
curl -X PUT "http://localhost:9200/_index_template/fluentd_template" \
  -H "Content-Type: application/json" \
  -d @./configs/efk/elasticsearch/index_template.json

printf "\n"

if [ $? -eq 0 ]; then
  printf "Index template applied successfully.\n"
else
  printf "Failed to apply index template\n."
fi

printf "All services are up and running.\n"
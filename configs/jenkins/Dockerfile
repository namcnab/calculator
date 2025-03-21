FROM jenkins/jenkins:lts

USER root
RUN apt-get update && apt-get install -y git wget
# Adjust permissions for /usr/local
RUN chown -R jenkins:jenkins /usr/local
USER jenkins
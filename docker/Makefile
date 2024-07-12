# docker user name
NAMESPACE := sarodesparsh

all: notification-server-docker

golang-base-docker:
	docker build -t $(NAMESPACE)/golang-base -f ./golang-base/Dockerfile ../
	docker push $(NAMESPACE)/golang-base

notification-server-docker: golang-base-docker
	docker build -t $(NAMESPACE)/notification-server -f ./notification-server/Dockerfile ../ --build-arg="NAMESPACE=$(NAMESPACE)"
	docker push $(NAMESPACE)/notification-server

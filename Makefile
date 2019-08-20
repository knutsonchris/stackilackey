.PHONY: all test clean start-test-servers stop-test-servers go-test

env_vars = STACKILACKEY_USERNAME=$$(docker exec -it $$(cat dockerImageName) cat /root/stacki-ws.cred | jq '.[0] .username' --raw-output)\
           STACKILACKEY_PASSWORD=$$(docker exec -it $$(cat dockerImageName) cat /root/stacki-ws.cred | jq '.[0] .key' --raw-output)\
		   STACKILACKEY_FRONTEND_IP=localhost:8080


all: test clean

start-test-servers:
	docker run -d --rm -p 8080:80 --mount source=develop,target=/root -P --privileged stacki/frontend-centos:05.03.00.00 > dockerImageName
	@echo "wait for the stacki frontend to come up"
	sleep 70

stop-test-servers:
	docker stop $$(cat dockerImageName)

clean:
	- rm dockerImageName

go-test:
	- $(env_vars) go test -cover -p 1 ./...
	
test: start-test-servers go-test stop-test-servers

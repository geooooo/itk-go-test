dev:
	cd app && \
		host="localhost" \
		port="8080" \
		apiVersion="v1" \
		dbUser="geo" \
		dbPassword="123" \
		dbName="wallets" \
		dbReset="yes" \
		go run main.go
test:
	cd internal/server && go test -v ./...

build:
	docker build -t itk-go-test-app .

run:
	docker rm itk-go-test-app && docker run --name itk-go-test-app -d -p 8080:8080 itk-go-test-app
exec:
	docker exec -it itk-go-test-app /bin/sh
stop:
	docker stop itk-go-test-app

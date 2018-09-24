build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .
	docker build -t drhelius/helloworld-go-microservice .


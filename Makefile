build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user .
docker:
	docker build -t joinc/user:0.1 -t joinc/user:latest . 

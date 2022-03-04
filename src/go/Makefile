build: 
	docker build -f ../../go.Dockerfile --tag server ../..

run: 
	docker run --rm -d --name server -p 4355:4355 server

local: 
	go run main.go
hello:
	echo "hello world"


all: generate	server
	@echo "This generates the gqlgen code resolvers"
	

generate:
	gqlgen generate
server:
	go run server.go

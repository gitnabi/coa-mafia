build:
	go mod download -x

run:
	go build client/main.go && ./main
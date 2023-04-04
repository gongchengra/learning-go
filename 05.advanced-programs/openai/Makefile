.DEFAULT_GOAL := run

BINARY_NAME=web5.exe

build:
	go build -o ${BINARY_NAME} web5.go

run:
	pkill ${BINARY_NAME} || true
	go build -o ${BINARY_NAME} web5.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}


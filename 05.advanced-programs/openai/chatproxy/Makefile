.DEFAULT_GOAL := run

BINARY_NAME=web6.exe

build:
	go build -o ${BINARY_NAME} main.go content.go user.go chat.go

run:
	pkill ${BINARY_NAME} || true
	go build -o ${BINARY_NAME} main.go content.go user.go chat.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

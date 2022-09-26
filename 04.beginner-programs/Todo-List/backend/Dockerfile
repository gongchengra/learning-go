#FROM golang:1.15.0
#
## Set the Current Working Directory inside the container
#WORKDIR /app
#
#RUN export GO111MODULE=on
#
## Copy go mod and sum files
#COPY go.mod go.sum ./
#
## Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download
#
#COPY . . 
#
## Build the application
#RUN go build -o main .
#
#ENV HOST=mongo
#
## Expose port 3000 to the outside world
#EXPOSE 3000
#
## Command to run the executable
#CMD ["./main"]
#
FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /server
COPY main .
#COPY ./src /server/
# ENV AUTH_USERNAME alan
# ENV AUTH_PASSWORD threeWalk
EXPOSE 8080
CMD ["./main"]


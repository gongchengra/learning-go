# Method 0: Use golang image
FROM golang:1.15.0
COPY . .
# Build the application
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./main"]

# Method 1: Link so
# FROM alpine:latest
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# COPY main /
# EXPOSE 8080
# ENTRYPOINT ["/main"]

# Method 2: CGO_ENABLED=0 go build
# FROM alpine:latest
# COPY FileUpload /
# EXPOSE 8080
# ENTRYPOINT ["/FileUpload"]

# Method 3: CGO_ENABLED=0 go build
# FROM scratch
# COPY FileUpload /
# CMD ["/FileUpload"]

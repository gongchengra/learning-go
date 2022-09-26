# builder image
# FROM golang:alpine3.16 as builder
# RUN mkdir /build
# COPY main.go /build/
# COPY go.mod /build/
# WORKDIR /build
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o lastpass .

# generate clean, final image for end users
FROM scratch
# COPY --from=builder /build/lastpass .
COPY lastpass .
COPY lastpass.log .
EXPOSE 8080
ENTRYPOINT [ "./lastpass" ]

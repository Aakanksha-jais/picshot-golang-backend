FROM alpine:latest

RUN mkdir -p /src/build
WORKDIR  /src/build
RUN apk add --no-cache tzdata ca-certificates
COPY build .
COPY configs configs
EXPOSE 8002
CMD ["./main"]

# docker build -t picshot-app .
# env GOOS=linux GOARCH=arm go build main.go
# docker run -p 8002:8002 picshot-app
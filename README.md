# picshot-api

This project is backend (written in golang) for a photo blog application named `PicShOt`.

![picShOt](https://i.ibb.co/864mJCt/Whats-App-Image-2021-02-13-at-12-03-18-AM.jpg)

## Server Setup

> Pre-requisites: Install Golang and Docker on your system.

Clone this repository:
```shell
git clone git@github.com:Aakanksha-jais/picshot-golang-backend.git
```

From the root directory of the project, run the following commands:
```shell
go mod download
```
*This will download all the project dependencies.*

To set up the database containers, run the following:
```shell
docker-compose up
```

To run the server:
```shell
go run main.go
```
Hit Ctrl+C to stop the server.

To stop the database containers, run the following:
```shell
docker-compose down
```

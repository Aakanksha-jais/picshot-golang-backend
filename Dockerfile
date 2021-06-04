FROM golang
WORKDIR /picshot-golang-backend
COPY . .
RUN go mod download
RUN go build
EXPOSE 8000
CMD ["./picshot-golang-backend"]

# docker build -t aakankshajais/picshot-api .
# docker run -it -p 8000:8000 aakankshajais/picshot-api

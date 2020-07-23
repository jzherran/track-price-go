FROM golang:1.14

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
EXPOSE 8080
ENV APP_DIR="/app"
CMD ["/app/run.sh"]
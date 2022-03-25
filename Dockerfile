FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-backend

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-backend /docker-backend

EXPOSE 8000

ENTRYPOINT ["/docker-backend"]
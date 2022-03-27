FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

ENTRYPOINT ["/bin/sh"]

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /docker-backend ./cmd/backend

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY configs ./configs
COPY --from=build /docker-backend /docker-backend

EXPOSE 8000

ENTRYPOINT ["/docker-backend"]
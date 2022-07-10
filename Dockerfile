FROM alpine:latest AS base
WORKDIR /app

FROM golang:1.18.3-alpine AS build
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /src/build/app

FROM base AS final
WORKDIR /app
COPY --from=build /src/build .
RUN chmod +x ./app
EXPOSE 3000
ENTRYPOINT ["./app"]
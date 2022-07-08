FROM alpine:latest AS base
WORKDIR /app

FROM go:latest AS build
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./build/app main.go

FROM base AS final
WORKDIR /app
COPY --from=build /src/build/ .
EXPOSE 3000
CMD [ "app" ]
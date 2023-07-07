FROM golang:alpine AS build-env
WORKDIR /app
COPY . .
RUN go build -o mingzhu

FROM golang:alpine
WORKDIR /app
COPY --from=build-env /app/mingzhu .
CMD ["./mingzhu"]

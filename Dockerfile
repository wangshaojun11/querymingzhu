FROM golang:alpine AS build-env
WORKDIR /app
COPY . .
RUN go build -o mingzhu

FROM golang:alpine
WORKDIR /app
COPY --from=build-env /app/mingzhu .
COPY ./templates/* /app/templates/
COPY ./assets/ /app/assets/
CMD ["./mingzhu"]

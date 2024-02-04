FROM golang:1.21.5-alpine AS build
WORKDIR /restaurant
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/restaurant 

FROM alpine:3.18.4
RUN apk add --no-cache ca-certificates
WORKDIR /restaurant
COPY --from=build /bin/restaurant /bin/restaurant

EXPOSE 8080
ENTRYPOINT ["/bin/restaurant"]

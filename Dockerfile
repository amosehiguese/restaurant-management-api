FROM golang:1.21.5-alpine AS build
WORKDIR /go/src/restaurant
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/restaurant 

FROM scratch
COPY --from=build /go/bin/restaurant /bin/restaurant
ENTRYPOINT ["/bin/restaurant"]
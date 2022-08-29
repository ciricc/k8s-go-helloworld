FROM golang:1.19.0-alpine as build
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build main.go

# Second stage
FROM alpine:latest as production
COPY --from=build /app/main .
EXPOSE 8080

CMD ["./main"]
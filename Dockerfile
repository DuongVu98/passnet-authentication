FROM golang:1.15.2-alpine3.12 AS build
WORKDIR /app
COPY . .
RUN go build -o ./dist/main src/main/main.go

FROM alpine:latest
ENV ENV=vagrant
WORKDIR /app
COPY --from=build /app/dist ./dist
COPY --from=build /app/env ./env
RUN ls -al dist
RUN pwd
CMD ["./dist/main"]
EXPOSE 1323
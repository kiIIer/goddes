FROM golang:1.17.7-alpine3.15 as build

WORKDIR /app
COPY . .
RUN go install

FROM alpine:3.15.0

COPY --from=build ./../go/bin/goddes .
ENTRYPOINT [ "./goddes" ]
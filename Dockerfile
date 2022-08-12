FROM golang:1.19 as builder

WORKDIR /app

ADD . .


RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/artifact-server

FROM scratch

COPY --from=builder /app/bin/artifact-server /bin/artifact-server

EXPOSE 1234

ENTRYPOINT ["/bin/artifact-server"]
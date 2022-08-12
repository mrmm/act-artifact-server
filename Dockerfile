FROM golang:1.19 as build

WORKDIR /app

ADD . .

RUN go mod tidy &&\
  go build -o bin/artifact-server &&\
  chmod +x bin/artifact-server

FROM scratch

COPY --from=build /app/bin/artifact-server /artifact-server

EXPOSE 1234

ENTRYPOINT ["/artifact-server"]
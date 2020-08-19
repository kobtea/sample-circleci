FROM golang:1.14 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/app
COPY . .
RUN go build

FROM alpine
COPY --from=builder /go/src/app/sample-circleci /sample-circleci
ENTRYPOINT ["/sample-circleci"]

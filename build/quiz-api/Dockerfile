FROM golang:1.18-alpine as builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN apk update && \
    apk add git &&\
    go build ./cmd/quiz-api

FROM alpine
COPY --from=builder /build/quiz-api /opt/app/
ENTRYPOINT ["/opt/app/quiz-api"]
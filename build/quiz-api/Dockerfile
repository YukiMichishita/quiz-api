FROM golang:1.18-alpine as builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
#RUN go run ./cmd/quiz-api/main.go
#RUN go build -ldflags '-s -w' ./cmd/quiz-api

#FROM alpine
#COPY --from=builder /build/quiz-api /opt/app/
#ENTRYPOINT ["/opt/app/quiz-api"]
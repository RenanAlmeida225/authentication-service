FROM golang:1.23.2-alpine as builder

ENV GOPRIVATE=github.com/PABX-Connect/*

RUN apk add --no-cache git openssh-client

RUN mkdir -p /root/.ssh && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

ARG SSH_PRIVATE_KEY
RUN echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /app/connect .

############## STAGE ##############

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/connect /app/connect

RUN mkdir -p 4-infra/database/postgres/migrations
RUN mkdir -p 2-uc/languages
RUN mkdir -p 2-uc/templates
RUN mkdir -p 1-api/swagger

COPY ./4-infra/database/postgres/migrations/. ./4-infra/database/postgres/migrations/
COPY ./2-uc/languages/. ./2-uc/languages/
COPY ./2-uc/templates/. ./2-uc/templates/
COPY ./1-api/swagger/swagger.json ./1-api/swagger/swagger.json

CMD ["/app/connect"]
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o program main.go


FROM postgres:alpine

WORKDIR /app

ENV POSTGRES_PASSWORD=1

COPY --from=builder /app/program .
COPY --from=builder /app/init.sql /docker-entrypoint-initdb.d/init.sql
COPY --chmod=755 entrypoint.sh .

CMD ["/app/entrypoint.sh"]
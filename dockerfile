FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o program main.go


FROM docker:cli

WORKDIR /app


COPY --from=builder /app/program-search /app
COPY --from=builder /app/init.sql /app



RUN docker pull postgres:alpine
RUN docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=1 -e POSTGRES_DB=postgres --name postgres --restart=always -v $(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql postgres:alpine

CMD ["/app/program"]
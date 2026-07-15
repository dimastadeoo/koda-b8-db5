FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o program main.go


FROM postgres:alpine

WORKDIR /app

COPY --from=builder /app/program /app
COPY --from=builder /app/init.sql /docker-entrypoint-initdb.d/init.sql

# RUN apk add --no-cache docker-cli
# RUN docker run -v /var/run/docker.sock:/var/run/docker.sock myimage


# RUN docker pull postgres:alpine
# RUN docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=1 -e POSTGRES_DB=postgres --name postgres --restart=always -v $(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql postgres:alpine

CMD ["/app/program"]
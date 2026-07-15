# ERD Data Contact List

## Berikut merupakan tampilan ERD dengan mermaid untuk contact list

### Akses Aplikasi

```sh
docker pull ghcr.io/dimastadeoo/koda-b8-db5:latest
```

```mermaid

erDiagram

contact {
    bigint id PK
    string fullname
    string no_hp UK
    string email UK
    timestamp created_at
    timestamp updated_at
}

```
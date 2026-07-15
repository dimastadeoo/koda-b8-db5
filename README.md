# ERD Data Contact List

## Berikut merupakan tampilan ERD dengan mermaid untuk contact list

### Akses Aplikasi
Untuk menjalankan aplikasi bisa pull dari package dan coba jalankan dengan -it supaya aplikasi berjalan di terminal

```sh
docker pull ghcr.io/dimastadeoo/koda-b8-db5:latest
docker run -it ghcr.io/dimastadeoo/koda-b8-db5:latest
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
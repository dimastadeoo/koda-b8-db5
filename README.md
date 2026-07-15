# ERD Data Contact List

## Berikut merupakan tampilan ERD dengan mermaid untuk contact list

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
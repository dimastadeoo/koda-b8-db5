package models

import (
	"context"
	"time"
	"github.com/jackc/pgx/v5"
)

type ListContact struct {
	Id    int
	Fullname  string
	No_Hp int
	Email string
	Created_At time.Time
	Updated_At time.Time
}

func GetAllData(conn *pgx.Conn) ([]ListContact, error) {

	rows, _ := conn.Query(context.Background(), `
		SELECT id, fullname, no_hp, email, created_at, updated_at
		FROM list_contact
	`)
	lists, err := pgx.CollectRows(rows, pgx.RowToStructByName[ListContact])
	return lists, err
}

func AddDataList(data ListContact, conn *pgx.Conn) (ListContact, error){
	rows, _ := conn.Query(context.Background(), `
		INSERT INTO list_contact  (fullname, no_hp, email)
		VALUES ($1, $2, $3)
		RETURNING id, fullname, no_hp, email, created_at, updated_at
	`, data.Fullname, data.No_Hp, data.Email)

	list, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[ListContact])

	return *list, err
}

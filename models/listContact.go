package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type ListContact struct {
	Id         int
	Fullname   string
	No_Hp      int
	Email      string
	Created_At time.Time
	Updated_At time.Time
}

func GetAllData(conn *pgx.Conn) ([]ListContact, error) {

	rows, _ := conn.Query(context.Background(), `
		SELECT id, fullname, no_hp, email, created_at, updated_at
		FROM list_contact
	`)
	defer rows.Close()

	lists, err := pgx.CollectRows(rows, pgx.RowToStructByName[ListContact])
	return lists, err
}

func AddDataList(data ListContact, conn *pgx.Conn) (ListContact, error) {
	rows, _ := conn.Query(context.Background(), `
		INSERT INTO list_contact  (fullname, no_hp, email)
		VALUES ($1, $2, $3)
		RETURNING id, fullname, no_hp, email, created_at, updated_at
	`, data.Fullname, data.No_Hp, data.Email)
	defer rows.Close()

	list, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[ListContact])

	return list, err
}

func GetDataByEmail(email string, conn *pgx.Conn) (ListContact, error) {
	rows, _ := conn.Query(context.Background(), `
		SELECT id, fullname, no_hp, email, created_at, updated_at 
		FROM list_contact WHERE email = $1
	`, email)
	defer rows.Close()

	list, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[ListContact])

	return list, err

}

func UpdateDataList(id int, data ListContact, conn *pgx.Conn) (ListContact, error) {
	rows, _ := conn.Query(context.Background(), `
		UPDATE list_contact 
		SET fullname = $1, 
		no_hp = $2, 
		email = $3,
		updated_at = NOW()
		WHERE id = $4
		RETURNING id, fullname, no_hp, email, created_at, updated_at
	`, data.Fullname, data.No_Hp, data.Email, id)
	defer rows.Close()

	list, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[ListContact])

	return list, err
}

func DeleteDataList(id int, conn *pgx.Conn) error {
	_, err := conn.Query(context.Background(), `
		DELETE FROM list_contact 
		WHERE id = $1
		RETURNING id, fullname, no_hp, email, created_at, updated_at
	`, id)

	return err
}

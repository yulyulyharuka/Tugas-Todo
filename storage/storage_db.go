package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/yulyulyharuka/todo2/model"
)

type database struct {
	conn *sql.DB
}

func newConnection() database {
	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	CreateTable(db)
	return database{
		conn: db,
	}
}

func CreateTable(conn *sql.DB) error {
	query := `CREATE TABLE todo(id integer, title VARCHAR(200), description VARCHAR(200), createdAt timestamp);`
	_, err := conn.Exec(query)
	if err != nil {
		return err
	}
	log.Println("table todo created")
	return err
}

func (o database) Create(obj model.Todo) error {
	query := `INSERT INTO todo(id, title, description, createdat)
				VALUES ($1, $2, $3, $4);`
	_, err := o.conn.Exec(query, obj.ID, obj.Title, obj.Description, obj.CreatedAt)
	log.Println("data added")
	return err
}

func (o database) Detail(id int32) (model.Todo, error) {
	var result model.Todo
	query := fmt.Sprintf("SELECT id, title, description, createdat FROM todo WHERE id=%d;", id)
	err := o.conn.QueryRow(query).Scan(&result.ID, &result.Title, &result.Description, &result.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (o database) Delete(id int32) error {
	query := `DELETE FROM todo WHERE id=$1;`
	_, err := o.conn.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data deleted")
	return nil
}

func (o database) List() ([]model.Todo, error) {
	query := "SELECT id, title, description, createdat FROM todo LIMIT 10;"
	rows, err := o.conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

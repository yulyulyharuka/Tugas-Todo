package main

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/yulyulyharuka/todo2/model"
	"github.com/yulyulyharuka/todo2/storage"
)

// type Customer struct {
// 	ID        int
// 	Name      string
// 	Address   string
// 	Phone     string
// 	Birthdate string
// }

func main() {
	// initiate
	var memStore = storage.GetStorage(storage.StorageTypeDB)

	// create new todo
	obj := []model.Todo{model.Todo{
		ID:          1,
		Title:       "first",
		Description: "First Todo",
		CreatedAt:   time.Now(),
	}, model.Todo{
		ID:          2,
		Title:       "second",
		Description: "Second Todo",
		CreatedAt:   time.Now(),
	}, model.Todo{
		ID:          3,
		Title:       "third",
		Description: "Third Todo",
		CreatedAt:   time.Now(),
	}}

	for _, o := range obj {
		if err := memStore.Create(o); err != nil {
			log.Fatal(err)
		}
	}

	// get detail
	todo, err := memStore.Detail(2)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d : %s", todo.ID, todo.Description)

	// get list
	list, err := memStore.List()
	if err != nil {
		log.Fatal(err)
	}

	for _, elmt := range list {
		log.Printf("%d : %s", elmt.ID, elmt.Description)
	}
}

// Database section
// func connectDB() *sql.DB {
// 	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=db sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }

// func InsertDB(db *sql.DB, customer Customer) error {
// 	query := `INSERT INTO customers(id, name, address, phone)
// 				VALUES ($1, $2, $3, $4);`
// 	_, err := db.Exec(query, customer.ID, customer.Name, customer.Address, customer.Phone)
// 	log.Println("data added")
// 	return err
// }

// func List(db *sql.DB) ([]Customer, error) {
// 	query := "SELECT id,name,address, phone, birthdate FROM customers LIMIT 10;"
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var customers []Customer
// 	for rows.Next() {
// 		var customer Customer
// 		err := rows.Scan(&customer.ID, &customer.Name, &customer.Address, &customer.Phone, &customer.Birthdate)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		customers = append(customers, customer)
// 	}

// 	return customers, nil
// }

// func Get(db *sql.DB, id int) (Customer, error) {
// 	var customer Customer
// 	query := fmt.Sprintf("SELECT * FROM customers WHERE ID=%d;", id)
// 	err := db.QueryRow(query).Scan(&customer.ID, &customer.Name, &customer.Address, &customer.Phone, &customer.Birthdate)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return customer, nil
// }

// func Update(db *sql.DB, customer Customer) error {
// 	query := "UPDATE customers SET name=$2, address=$3, phone=$4, birthdate=$5 WHERE id=$1;"
// 	_, err := db.Exec(query, customer.ID, customer.Name, customer.Address, customer.Phone, customer.Birthdate)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("data updated")
// 	return nil
// }

// func Delete(db *sql.DB, id int) error {
// 	query := "DELETE FROM customers WHERE id=$1;"
// 	_, err := db.Exec(query, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("data deleted")
// 	return nil
// }

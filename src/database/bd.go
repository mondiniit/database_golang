package my_postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"math/rand"
	"github.com/jackc/pgx/v5"
)

const USERNAME = "bexsy"
const PASSWORD = "b3xsy"
const PORT = "5432"
const URI = "localhost"
const DB = "bexsydb"

// urlExample := "postgres://username:password@localhost:5432/database_name"
var DATABASE_URL string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", USERNAME, PASSWORD, URI, PORT, DB)

// If there's an error, panic
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// It connects to a database, checks if the day is 6, and if it is not, it gets the day solution and
// updates the day result table.
func Connection() string {

	fmt.Printf("Connect to bd ... ")
	db, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("success !")
	}
	
	defer db.Close(context.Background())

	daySolution := GetDaySolution(db)
	return daySolution
}

// It returns the day of the month of the last time the solution was updated
func GetDaySolution(db *pgx.Conn) (string) {
	type Row struct {
		Id		int
		Solucao	string
		Data	time.Time
	}
	var r Row
	_ = db.QueryRow(context.Background(), "SELECT * FROM dayresult WHERE ID=1").Scan(&r.Id, &r.Solucao, &r.Data)
	if (r.Data.Day() !=  time.Now().Day()) {
		r.Solucao = GetNewDaySolution(db)
		UpdateTableSolution(db, r.Solucao)
	}
	return r.Solucao
}

// It updates the dayresult table with the solution of the day
func UpdateTableSolution(db *pgx.Conn, daySolution string) {
	sqlStatement := fmt.Sprintf(`UPDATE dayresult set solucao='%s', data=NOW() WHERE id=1`, daySolution)
	fmt.Println(sqlStatement)
	_, err := db.Exec(context.Background(), sqlStatement)
	checkErr(err)
}

// It connects to the database, queries the table, and returns a random solution
func GetNewDaySolution(db *pgx.Conn) string {
	rows, err := db.Query(context.Background(), "SELECT * FROM equacoes")
	checkErr(err)
	type Row struct {
		Id      int
		Solucao string
	}
	var rowSlice []Row
	for rows.Next() {
		var r Row
		err := rows.Scan(&r.Id, &r.Solucao)
		checkErr(err)
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	nResults := len(rowSlice)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(nResults)
	return rowSlice[index].Solucao
}
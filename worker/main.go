package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/sijms/go-ora/v2"
)

const (
	concurrency = 10
	numRecords  = 1000000
)

func worker(jobs <-chan int, results chan<- int, db *sql.DB) {
	for range jobs {
		var count int
		err := db.QueryRow("SELECT customer_id FROM customer order by created_date  OFFSET 10 ROWS FETCH NEXT 10 ROWS ONLY").Scan(&count)
		if err != nil {
			fmt.Println("get err", err)
			continue
		}
		results <- count
	}
}

var dbParams = map[string]string{
	"username": "acs_customer_db",
	"password": "acs_customer",
	"server":   "10.101.40.17",
	"port":     "1521",
	"service":  "coredb",
}

func main() {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]
	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	defer db.Close()

	jobs := make(chan int, numRecords)
	results := make(chan int, numRecords)

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results, db)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 0; i < numRecords; i++ {
		jobs <- i
	}
	close(jobs)

	for result := range results {
		fmt.Println(result)
	}
}

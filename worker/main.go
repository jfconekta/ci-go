package main

import (
	"time"

	"../common"
	_ "github.com/lib/pq"
)

// sleep 10 seconds and execute the query
func loop() int {
	for {
		time.Sleep(10 * time.Second)

		common.InsertDB()

	}

}

func main() {
	loop()
}

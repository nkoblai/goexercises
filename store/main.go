package main

import (
	_ "github.com/lib/pq"
)

func main() {
	a := App{}
	a.Initialize(
		host, port, user, password, dbname,
	)
	a.Run(":8080")
}

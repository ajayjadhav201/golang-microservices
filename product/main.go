package main

import "github.com/ajayjadhav201/product/database"

func main() {
	//
	_ = database.NewDatabase(database.GetConnection(), true)

}

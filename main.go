package main

import (
	"ginpet/src"
	"ginpet/src/db"
)

func main() {
	src.RunApp()
	db.Close()
}

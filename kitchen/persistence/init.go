package persistence

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
  var err error
  db, err = sql.Open("sqlite3", "../database/data")
  
  if err != nil {
    log.Fatal(err.Error())
    return
  }
  if err := db.Ping(); err != nil {
    fmt.Println("Error al hacer ping a la base de datos:", err)
    return
  }
}

package persistence

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
  var err error
  db, err = sql.Open("sqlite3", "persistence/data")
  if err != nil {
    log.Fatal(err.Error())
    return
  }
  if err := db.Ping(); err != nil {
    fmt.Println("Error al hacer ping a la base de datos:", err)
    return
  }
  // Consulta para obtener los nombres de todas las tablas
  rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table';")
  if err != nil {
    fmt.Println("Error al ejecutar la consulta:", err)
    return
  }
  defer rows.Close()

  // Variable para almacenar los nombres de las tablas
  var tableName string

  // Iterar sobre los resultados
  for rows.Next() {
    // Escanear el nombre de la tabla desde la fila actual
    if err := rows.Scan(&tableName); err != nil {
      fmt.Println("Error al escanear fila:", err)
      return
    }
    // Imprimir el nombre de la tabla
    fmt.Println("Tabla encontrada:", tableName)
  }

  // Manejar errores después del iterador
  if err := rows.Err(); err != nil {
    fmt.Println("Error final después de iterar filas:", err)
    return
  }

  fmt.Println("Consulta de tablas exitosa.")
}

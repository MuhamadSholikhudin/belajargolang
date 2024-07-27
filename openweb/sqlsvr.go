package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"

    _ "github.com/denisenkom/go-mssqldb"
)

func main() {
    server := "10.10.100.246"
    port := 1033 // Port default SQL Server
    user := "sa"
    password := "pastibisa123"
    database := "attendance_machine_database_3"

    // Setup koneksi string
    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        server, user, password, port, database)

    // Membuat koneksi ke SQL Server
    db, err := sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatal("Koneksi Gagal:", err.Error())
    }
    defer db.Close()

    // Uji koneksi
    ctx := context.Background()
    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal("Koneksi Gagal:", err.Error())
    }

    fmt.Println("Koneksi Berhasil!")

    // Di sini Anda bisa menjalankan query dan melakukan operasi lain ke database.
}

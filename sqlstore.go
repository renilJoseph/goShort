package main

import(
        "fmt"
        "database/sql"
        _ "github.com/lib/pq"
)

const (
        host = "localhost"
        port = 5432
        user = "postgres"
        password = "reducer"
        dbname = "postgres"
)

func main(){
            psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
            host, port, user, password, dbname)

            db, err := sql.Open("postgres", psqlinfo)
            if err != nil{
                panic(err)    
            }
            fmt.Println("succes")
            defer db.Close()
}

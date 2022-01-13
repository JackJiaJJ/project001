package api

import (
        "database/sql"
        "fmt"
)

func DBOpen(dbtype, dbinfo string) (*sql.DB, error) {
        db, err := sql.Open(dbtype, dbinfo)
        if err != nil {
                fmt.Println("Failed to open", dbtype)
                return nil, err
        }

        return db, nil
}

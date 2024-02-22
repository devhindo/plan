package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Notification struct {
    Title   string
    Message string
    Time    string
}

func SaveToDatabase(notifications []Notification, db *sql.DB) error {
    stmt, err := db.Prepare("INSERT INTO notifications(title, message, time) values(?,?,?)")
    if err != nil {
        return err
    }
    for _, notification := range notifications {
        _, err = stmt.Exec(notification.Title, notification.Message, notification.Time)
        if err != nil {
            return err
        }
    }
    return nil
}
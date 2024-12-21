package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect_db() (*sql.DB, error) {
    db, err := sql.Open("mysql", "dev:my_cool_secret@tcp(localhost:3306)/mydb")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func create_user(username, email string) error {
    db, err := connect_db()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("INSERT INTO users (username, email) VALUES (?, ?)",
                     username, email)
    if err != nil {
        return err
    }

    return nil
}

func read_user(user_id int) (string, string, error) {
    db, err := connect_db()
    if err != nil {
        return "", "", err
    }
    defer db.Close()


    var username, email string
    err = db.QueryRow("SELECT username, email FROM users where id=?",
                      user_id).Scan(&username, &email)
    if err != nil {
        return "", "", err
    }

    return username, email, nil
}

func update_user(user_id int, new_email string) error {
    db, err := connect_db()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("UPDATE users SET email=? WHERE id=?", new_email, user_id)
    if err != nil {
        return err
    }

    return nil
}

func delete_user(user_id int) error {
    db, err := connect_db()
    if err != nil {
        return err
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM users WHERE id=?", user_id)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    err := create_user("Some One", "some@some.com")
    if err != nil {
        fmt.Println("Error creating user", err)
    } else {
        fmt.Println("User creation success")
    }

    username, email, err := read_user(1)
    if err != nil {
        fmt.Println("Error reading user", err)
    } else {
        fmt.Printf("Username: %s, Email: %s\n", username, email)
    }

    err = update_user(1, "newmail@some.com")
    if err != nil {
        fmt.Println("Error updating user", err)
    } else {
        fmt.Println("User updated successfully")
    }

    err = delete_user(1)
    if err != nil {
        fmt.Println("Error deleting user", err)
    } else {
        fmt.Println("User deleted successfully")
    }
}

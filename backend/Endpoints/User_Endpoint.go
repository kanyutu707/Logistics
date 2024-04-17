package endpoints

import (
	"fmt"

	connection "example.com/backend/Connection"
)

type Users struct {
    ID         int64
    Email      string
    First_name string
    Last_name  string
	Password string
}


func FetchUsers() ([]Users, error) {
	db := connection.Database_connection()
	defer db.Close()

	var users []Users

	rows, err := db.Query("SELECT id, first_name, last_name, email, password FROM Users")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.ID, &user.First_name, &user.Last_name, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("failed to scan user row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over user rows: %v", err)
	}

	return users, nil
}

func AddUser(user Users) (int64, error) {
    db := connection.Database_connection()
    result, err := db.Exec("INSERT INTO Users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)", user.First_name, user.Last_name, user.Email, user.Password)
    if err != nil {
        return 0, fmt.Errorf("AddUser: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("AddUser: %v", err)
    }
    return id, nil
}

func UpdateUser(userID string, updatedUser Users) error {
    db := connection.Database_connection()
    defer db.Close()

    
    stmt, err := db.Prepare("UPDATE Users SET first_name=?, last_name=?, email=?, password=? WHERE id=?")
    if err != nil {
        return fmt.Errorf("failed to prepare update statement: %v", err)
    }
    defer stmt.Close()

    
    _, err = stmt.Exec(updatedUser.First_name, updatedUser.Last_name, updatedUser.Email,updatedUser.Password, userID)
    if err != nil {
        return fmt.Errorf("failed to execute update statement: %v", err)
    }

    return nil
}

func DeleteUser(userID string) error {
    db := connection.Database_connection()
    defer db.Close()

    
    stmt, err := db.Prepare("DELETE FROM Users WHERE id=?")
    if err != nil {
        return fmt.Errorf("failed to prepare delete statement: %v", err)
    }
    defer stmt.Close()

   
    _, err = stmt.Exec(userID)
    if err != nil {
        return fmt.Errorf("failed to execute delete statement: %v", err)
    }

    return nil
}


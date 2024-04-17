package endpoints

import (
	"fmt"

	connection "example.com/backend/Connection"
)

type Clients struct{
	ID int64
	User_id string
}

func FetchClients()([]Clients, error){
	db:=connection.Database_connection()
	defer db.Close()

	var clients []Clients

	rows, err:=db.Query("SELECT id, user_id from Clients")

	if err!=nil{
		return nil, fmt.Errorf("failed to fetch client: %v ", err)
	}
	defer rows.Close()
	for rows.Next(){
		var client Clients
		if err:=rows.Scan(&client.ID, &client.User_id); err!=nil{
			return nil, fmt.Errorf("failed to scan client rows: %v ", err)
		}
		clients=append(clients, client)
	}
	if err:=rows.Err(); err!=nil{
		return nil, fmt.Errorf("error iterating over the client rows: %v ", err)
	}
	return clients, nil
}

func AddClient(client Clients)(int64, error){
	db:=connection.Database_connection()
	result, err:=db.Exec("INSERT INTO Clients (user_id) VALUES (?)", client.User_id)
	if err!=nil{
		return 0, fmt.Errorf("AddClient: %v", err)
	}
	id, err:=result.LastInsertId()
	if err!=nil{
		return 0, fmt.Errorf("AddClient: %v", err)
	}
	return id, nil
}

func UpdateClient(clientId string, updateClient Clients) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("UPDATE Clients SET user_id=? where id=?")
	if err!=nil{
		return fmt.Errorf("failed to update the clients statement: %v", err)
	}
	defer stmt.Close()
	_, err=stmt.Exec(updateClient.User_id, clientId)
	if err!=nil{
		return fmt.Errorf("failed to execute the update statement on the clients table: %v ", err)
	}
	return nil
}

func DeleteClient(clientID string) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("DELETE FROM Clients WHERE id=?")

	if(err!=nil){
		return fmt.Errorf("failed to prepare delete statement: %v", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(clientID)

	if(err!=nil){
		return fmt.Errorf("failed to execute delete statement: %v", err)
	}
	return nil
}
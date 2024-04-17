package endpoints

import (
	"fmt"

	connection "example.com/backend/Connection"
)

type Trips struct{
	ID int64
	Client_id int64
	Driver_id int64
	Start string
	End string
}

func FetchTrips()([]Trips, error){
	db:=connection.Database_connection()
	defer db.Close()

	var trips []Trips
	rows, err:=db.Query("SELECT id, client_id, driver_id, start, end FROM Trips")
	if err!=nil{
		return nil, fmt.Errorf("failed to fetch trips: %v ", err)
	}
	defer rows.Close()
	for rows.Next(){
		var trip Trips
		if err:=rows.Scan(&trip.ID,&trip.Client_id, &trip.Driver_id, &trip.Start, &trip.End); err!=nil{
			return nil, fmt.Errorf("failed to scan trip rows: %v ", err)
		}
		trips=append(trips, trip)
	}
	if err:=rows.Err(); err!=nil{
		return nil, fmt.Errorf("error iterating over the trip rows: %v ", err)
	}
	return trips, nil
}

func AddTrip(trip Trips)(int64, error){
	db:=connection.Database_connection()
	result, err:=db.Exec("INSERT INTO Trips (client_id, driver_id, start, end) VALUES (?, ?, ?, ?)", trip.Client_id, trip.Driver_id, trip.Start, trip.End)

	if err!=nil{
		return 0, fmt.Errorf("AddTrip: %v", err)
	}
	id, err:=result.LastInsertId()
	if err!=nil{
		return 0, fmt.Errorf("AddTrip: %v", err)
	}
	return id, nil
}

func UpdateTrip(tripID string, updateTrip Trips) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("UPDATE Trips set client_id=?, driver_id=?, start=?, end=? WHERE id=?")

	if err!=nil{
		return fmt.Errorf("failed to prepare the update statement: %v ", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(updateTrip.Client_id, updateTrip.Driver_id, updateTrip.Start, updateTrip.End, tripID)

	if err!=nil{
		return fmt.Errorf("failed to execute update statement: %v ",err)
	}
	return nil
}

func DeleteTrip(tripID string) error{
	db:=connection.Database_connection()
	defer db.Close()
	stmt, err:=db.Prepare("DELETE FROM Trips WHERE id=?")

	if err!=nil{
		return fmt.Errorf("failed to prepare delete statment: %v ", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(tripID)

	if err!=nil{
		return fmt.Errorf("failed to execute delete statement: %v ", err)
	}
	return nil
}
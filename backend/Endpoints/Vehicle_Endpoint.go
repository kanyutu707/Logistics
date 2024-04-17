package endpoints

import (
	"fmt"

	connection "example.com/backend/Connection"
)

type Vehicles struct{
	ID int64
	Vehicle_number string
	Max_weight string
}

func FetchVehicles()([]Vehicles, error){
	db:=connection.Database_connection()
	defer db.Close()

	var vehicles []Vehicles;

	rows, err:=db.Query("SELECT id, vehicle_number, max_weight From Vehicles")

	if err!=nil{
		return nil, fmt.Errorf("failed to fetch vehicle: %v ", err)
	}
	defer rows.Close()
	for rows.Next(){
		var vehicle Vehicles
		if err:=rows.Scan(&vehicle.ID, &vehicle.Vehicle_number, &vehicle.Max_weight); err!=nil{
			return nil, fmt.Errorf("failed to scan vehicle rows: %v ", err)
		}
		vehicles=append(vehicles, vehicle)
	}
	if err:=rows.Err(); err!=nil{
		return nil, fmt.Errorf("error iterating over the vehicle rows: %v ", err)
	}
	return vehicles, nil
}

func AddVehicle(vehicle Vehicles)(int64, error){
	db:=connection.Database_connection()
	result, err:=db.Exec("INSERT INTO Vehicles (vehicle_number, max_weight) VALUES(?, ?)", vehicle.Vehicle_number, vehicle.Max_weight)
	if err!=nil{
		return 0, fmt.Errorf("AddVehicle: %v", err)
	}
	id, err:=result.LastInsertId()
	if err !=nil{
		return 0, fmt.Errorf("AddVehicle: %v", err)
	}
	return id, nil
}

func UpdateVehicle(vehicleID string, updateVehicle Vehicles) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("UPDATE Vehicles SET vehicle_number=?, max_weight=? WHERE id=?")
	if err!=nil{
		return fmt.Errorf("failed to update the statement: %v", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(updateVehicle.Vehicle_number, updateVehicle.Max_weight, vehicleID)
	if err!=nil{
		return fmt.Errorf("failed to execute update statement: %v", err)
	}
	return nil
}

func DeleteVehicle(vehicleID string) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("DELETE FROM Vehicles WHERE id=?")
	if(err!=nil){
		return fmt.Errorf("failed to prepare delete statement: %v", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(vehicleID)

	if err!=nil{
		return fmt.Errorf("failed to execute delete statement: %v", err)
	}
	return nil
}
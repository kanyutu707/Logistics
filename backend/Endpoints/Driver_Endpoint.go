package endpoints

import (
	"fmt"

	connection "example.com/backend/Connection"
)

type Drivers struct{
	ID int64
	User_id int64
	Vehicle_id string
	Unit_cost string
}

func FetchDrivers()([]Drivers, error){
	db:=connection.Database_connection()
	defer db.Close()

	var drivers []Drivers
	rows, err:=db.Query("SELECT id, user_id, vehicle_id, unit_cost FROM Drivers")
	if err!=nil{
		return nil, fmt.Errorf("failed to fetch drivers: %v ", err)
	}
	defer rows.Close()

	for rows.Next(){
		var driver Drivers
		if err:=rows.Scan(&driver.ID, &driver.User_id, &driver.Vehicle_id, &driver.Unit_cost); err!=nil{
			return nil, fmt.Errorf("failed to scan driver row: %v ", err)
		}
		drivers=append(drivers, driver)
	}
	if err:=rows.Err(); err!=nil{
		return nil, fmt.Errorf("error iterating over client rows: %v ", err)
	}
	return drivers, nil
}

func AddDriver(driver Drivers)(int64, error){
	db:=connection.Database_connection()
	result, err:=db.Exec("INSERT INTO Drivers (user_id, vehicle_id, unit_cost) VALUES (?, ?, ?)", driver.User_id, driver.Vehicle_id, driver.Unit_cost)
	if err!=nil{
		return 0, fmt.Errorf("AddDriver: %v ", err)
	}
	id, err:=result.LastInsertId()
	if err!=nil{
		return 0, fmt.Errorf("AddDriver: %v", err)
	}
	return id, nil
}

func UpdateDriver(driverID string, updateDriver Drivers) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("UPDATE Drivers set user_id=?, vehicle_id=?, unit_cost=? WHERE id=?")
	if err!=nil{
		return fmt.Errorf("failed to prepare the update statement: %v ", err)
	}
	defer stmt.Close()

	_,err=stmt.Exec(updateDriver.User_id, updateDriver.Vehicle_id, updateDriver.Unit_cost, driverID)
	if err!=nil{
		return fmt.Errorf("failed to execute update statement: %v ", err)
	}
	return nil
}

func DeleteDriver(driverID string) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("DELETE FROM Drivers WHERE id=?")
	if err!=nil{
		return fmt.Errorf("failed to prepare delete statement: %v ", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(driverID)
	if err!=nil{
		return fmt.Errorf("failed to execute delete statement: %v ", err)
	}
	return nil
}
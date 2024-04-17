package endpoints

import (
	"fmt"

	connection "example.com/backend/Connection"
)


type Payments struct{
	ID int64
	Trip_id int64
	Amount string
}

func FetchPayment()([]Payments, error){
	db:=connection.Database_connection()
	defer db.Close()

	var payments []Payments;

	rows, err:=db.Query("SELECT id, trip_id, amount FROM Payments")

	if err!=nil{
		return nil, fmt.Errorf("failed to fetch trips: %v ", err)
	}
	defer rows.Close()

	for rows.Next(){
		var payment Payments
		if err:=rows.Scan(&payment.ID, &payment.Trip_id, &payment.Amount); err!=nil{
			return nil, fmt.Errorf("failed to scan payment rows: %v", err)
		}
		payments=append(payments, payment)
	}

	if err:=rows.Err(); err!=nil{
		return nil, fmt.Errorf("error iterating over the payment rows: %v", err)
	}
	return payments, nil

}

func AddPayment(payment Payments)(int64, error){
	db:=connection.Database_connection()
	result, err:=db.Exec("INSERT INTO Payments (trip_id, amount) VALUES (?, ?)", payment.Trip_id, payment.Amount)
	
	if err!=nil{
		return 0, fmt.Errorf("AddPayment: %v", err)
	}
	id, err:=result.LastInsertId()
	if err!=nil{
		return 0, fmt.Errorf("AddPayment: %v", err)
	}
	return id, nil
}

func UpdatePayment(paymentId string, updatePayment Payments) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("UPDATE Payments set trip_id=?, amount=? where id=?")

	if err!=nil{
		return fmt.Errorf("failed to prepare the update statement: %v ", err)
	}
	defer stmt.Close()
	_,err=stmt.Exec(updatePayment.Trip_id, updatePayment.Amount, paymentId)

	if err!=nil{
		return fmt.Errorf("failed to execute update statement: %v ", err)
	}
	return nil

}
func DeletePayment(paymentId string) error{
	db:=connection.Database_connection()
	defer db.Close()

	stmt, err:=db.Prepare("DELETE FROM Payments WHERE id=?")

	if err!=nil{
		return fmt.Errorf("failed to prepare delete statement: %v ", err)
	}
	defer stmt.Close()

	_, err=stmt.Exec(paymentId)

	if err!=nil{
		return fmt.Errorf("failed to execute delete statement: %v ", err)
	}
	return nil
}
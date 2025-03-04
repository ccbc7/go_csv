package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Customer represents a customer record
type Customer struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	City        string
	State       string
	ZipCode     string
	Country     string
}

// GenerateRandomCustomer generates a random customer record
func GenerateRandomCustomer(id int) Customer {
	return Customer{
		ID:          id,
		FirstName:   fmt.Sprintf("FirstName%d", id),
		LastName:    fmt.Sprintf("LastName%d", id),
		Email:       fmt.Sprintf("user%d@example.com", id),
		PhoneNumber: fmt.Sprintf("555-000-%04d", id),
		Address:     fmt.Sprintf("%d Maple St", id),
		City:        fmt.Sprintf("City%d", id),
		State:       "CA",
		ZipCode:     fmt.Sprintf("900%02d", id%100),
		Country:     "USA",
	}
}

// CreateCSVFile creates a CSV file with the specified number of customer records
func CreateCSVFile(filePath string, numRecords int) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"ID", "First Name", "Last Name", "Email", "Phone Number", "Address", "City", "State", "Zip Code", "Country"})

	// Write customer records
	for i := 1; i <= numRecords; i++ {
		customer := GenerateRandomCustomer(i)
		record := []string{
			strconv.Itoa(customer.ID),
			customer.FirstName,
			customer.LastName,
			customer.Email,
			customer.PhoneNumber,
			customer.Address,
			customer.City,
			customer.State,
			customer.ZipCode,
			customer.Country,
		}
		writer.Write(record)
	}

	return nil
}

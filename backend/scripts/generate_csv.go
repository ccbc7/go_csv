package main

import (
	"fmt"
	"project/utils"
)

func main() {
	// CSVファイルの作成
	numRecords := 10000
	filePath := fmt.Sprintf("./data/sample_data_%d.csv", numRecords)
	err := utils.CreateCSVFile(filePath, numRecords)
	if err != nil {
		fmt.Printf("Failed to create CSV file: %v\n", err)
		return
	}
	fmt.Println("CSV file created successfully")
}

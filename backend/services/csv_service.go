package services

import (
	"project/models"
	"project/repositories"
	"encoding/csv"
	"os"
)

type ICsvService interface {
	ProcessCsv() error
}

type CsvService struct {
	repository repositories.ICsvRepository
	filePath   string
}

func NewCsvService(repository repositories.ICsvRepository, filePath string) ICsvService {
	return &CsvService{repository: repository, filePath: filePath}
}

func (s *CsvService) ProcessCsv() error {
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		csvData := models.Csv{
			FirstName:   record[1],
			LastName:    record[2],
			Email:       record[3],
			PhoneNumber: record[4],
			Address:     record[5],
			City:        record[6],
			State:       record[7],
			ZipCode:     record[8],
			Country:     record[9],
		}
		_, err := s.repository.CreateCsv(csvData)
		if err != nil {
			return err
		}
	}
	return nil
}

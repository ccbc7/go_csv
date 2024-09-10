package services

import (
	"encoding/csv"
	"os"
	"sync"

	"project/models"
	"project/repositories"
)

// インターフェースを定義
type ICsvService interface {
	ProcessCsv() error
}

// 構造体を定義
type CsvService struct {
	repository repositories.ICsvRepository
	filePath   string
}

// コンストラクタを定義
func NewCsvService(repository repositories.ICsvRepository, filePath string) ICsvService {
	return &CsvService{repository: repository, filePath: filePath}
}

// ワーカー関数
func worker(jobs <-chan []string, results chan<- error, repository repositories.ICsvRepository, wg *sync.WaitGroup) {
	defer wg.Done()
	for record := range jobs {
		// CSVデータを構造体に変換
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
		// リポジトリ層のCreateCsvメソッドを呼び出し,以降の処理はリポジトリ層に委ねる
		_, err := repository.CreateCsv(csvData)
		results <- err
	}
}

// CSVファイルを読み込み、リポジトリ層のCreateCsvメソッドにデータを渡す
func (s *CsvService) ProcessCsv() error {
	// CSVファイルを開く
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	// 関数の終了時にファイルを閉じる
	defer file.Close()

	// CSVファイルを読み込む
	reader := csv.NewReader(file)

	// CSVファイルのヘッダーを読み込む
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// ワーカーの数を設定
	const numWorkers = 30
	// ジョブキューと結果キューを作成
	jobs := make(chan []string, len(records)-1)
	// エラーを格納するチャネル
	results := make(chan error, len(records)-1)
	var wg sync.WaitGroup

	// ワーカーを起動, numWorkersの数だけgoroutineを起動
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(jobs, results, s.repository, &wg)
	}

	// jobsチャネルcsvレコードを送信
	for _, record := range records[1:] {
		jobs <- record // jobsにrecordを送信
	}
	close(jobs)

	// ワーカーの終了を待つ
	wg.Wait()
	close(results)

	// エラーチェック
	for err := range results {
		if err != nil {
			return err
		}
	}

	return nil
}

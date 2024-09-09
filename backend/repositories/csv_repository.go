package repositories

import (
	"project/models"

	"gorm.io/gorm"
)

type ICsvRepository interface {
	// 引数はmodels.Csv型、戻り値はmodels.Csv型とerror型
	CreateCsv(csv models.Csv) (models.Csv, error)
}

/*
* 単一責任の原則に従い、CsvRepository 構造体はデータベース操作のロジックを持つにとどめる
 */
type CsvRepository struct {
	db *gorm.DB // データベース接続
}

/*
* CsvRepositoryインスタンスを生成する関数
* 単一責任の原則に従い、NewCsvRepository 関数はCsvRepository 構造体の初期化のみを行うにとどめる
 */
func NewCsvRepository(db *gorm.DB) ICsvRepository {
	return &CsvRepository{db: db}
}

/*
* メソッド内で、引数で受け取ったmodels.Csv型の値をデータベースに保存する
* rはレシーバーとして構造体のポインタを受け取る
 */
func (r *CsvRepository) CreateCsv(csv models.Csv) (models.Csv, error) {
	err := r.db.Create(&csv).Error
	return csv, err
}

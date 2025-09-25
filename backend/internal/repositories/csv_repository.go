package repositories

import (
	"project/internal/ent"
)

type CsvRepository interface {
	// 引数はEntクライアント、戻り値はエラーのみ
	ProcessCsvData() error
}

/*
* 単一責任の原則に従い、CsvRepository 構造体はデータベース操作のロジックを持つにとどめる
 */
type csvRepository struct {
	client *ent.Client // Entクライアント
}

/*
* CsvRepositoryインスタンスを生成する関数
* 単一責任の原則に従い、NewCsvRepository 関数はCsvRepository 構造体の初期化のみを行うにとどめる
 */
func NewCsvRepository(client *ent.Client) CsvRepository {
	return &csvRepository{client: client}
}

/*
* メソッド内で、CSVデータを処理してデータベースに保存する
* rはレシーバーとして構造体のポインタを受け取る
 */
func (r *csvRepository) ProcessCsvData() error {
	// CSVデータの処理ロジックをここに実装
	// 現在は空の実装
	return nil
}

package seeders

import (
	"context"
	"time"

	"project/internal/ent"
)

func SeedMovies(client *ent.Client) error {
	movies := []ent.Movie{
		{
			Title:       "インセプション",
			Description: "夢の中の夢を操る特殊な技術を持つ泥棒の物語",
			Duration:    148,
			Genre:       "SF・アクション",
			Director:    "クリストファー・ノーラン",
			Cast:        "レオナルド・ディカプリオ, マリオン・コティヤール",
			ReleaseDate: time.Date(2010, 7, 16, 0, 0, 0, 0, time.UTC),
			Rating:      "PG-13",
			Score:       8.8,
		},
		{
			Title:       "パルプ・フィクション",
			Description: "複数の物語が絡み合う犯罪映画の傑作",
			Duration:    154,
			Genre:       "犯罪・ドラマ",
			Director:    "クエンティン・タランティーノ",
			Cast:        "ジョン・トラボルタ, サミュエル・L・ジャクソン",
			ReleaseDate: time.Date(1994, 10, 14, 0, 0, 0, 0, time.UTC),
			Rating:      "R",
			Score:       8.9,
		},
		{
			Title:       "君の名は。",
			Description: "時空を超えた恋愛アニメーション",
			Duration:    106,
			Genre:       "アニメ・ロマンス",
			Director:    "新海誠",
			Cast:        "神木隆之介, 上白石萌音",
			ReleaseDate: time.Date(2016, 8, 26, 0, 0, 0, 0, time.UTC),
			Rating:      "G",
			Score:       8.4,
		},
		{
			Title:       "アベンジャーズ/エンドゲーム",
			Description: "マーベル・シネマティック・ユニバースの集大成",
			Duration:    181,
			Genre:       "アクション・SF",
			Director:    "アンソニー・ルッソ, ジョー・ルッソ",
			Cast:        "ロバート・ダウニー・Jr, クリス・エヴァンス",
			ReleaseDate: time.Date(2019, 4, 26, 0, 0, 0, 0, time.UTC),
			Rating:      "PG-13",
			Score:       8.4,
		},
		{
			Title:       "千と千尋の神隠し",
			Description: "宮崎駿監督によるファンタジーアニメーション",
			Duration:    125,
			Genre:       "アニメ・ファンタジー",
			Director:    "宮崎駿",
			Cast:        "柊瑠美, 入野自由",
			ReleaseDate: time.Date(2001, 7, 20, 0, 0, 0, 0, time.UTC),
			Rating:      "G",
			Score:       8.6,
		},
	}

	for _, movieData := range movies {
		_, err := client.Movie.Create().
			SetTitle(movieData.Title).
			SetDescription(movieData.Description).
			SetDuration(movieData.Duration).
			SetGenre(movieData.Genre).
			SetDirector(movieData.Director).
			SetCast(movieData.Cast).
			SetReleaseDate(movieData.ReleaseDate).
			SetRating(movieData.Rating).
			SetScore(movieData.Score).
			Save(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

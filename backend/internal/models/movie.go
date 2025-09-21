package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Movie holds the schema definition for the Movie entity.
type Movie struct {
	ent.Schema
}

// Fields of the Movie.
func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Comment("映画のタイトル"),
		field.String("description").Optional().Comment("映画の説明"),
		field.Int("duration").Positive().Comment("上映時間（分）"),
		field.String("genre").Optional().Comment("ジャンル"),
		field.String("director").Optional().Comment("監督"),
		field.String("cast").Optional().Comment("キャスト"),
		field.Time("release_date").Optional().Comment("公開日"),
		field.String("rating").Optional().Comment("評価"),
		field.Float("score").Optional().Comment("スコア"),
	}
}

// Edges of the Movie.
func (Movie) Edges() []ent.Edge {
	return nil
}

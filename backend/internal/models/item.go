package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("price").Positive(),
		field.String("description").Optional(),
		field.Bool("sold_out").Default(false),
		field.Int("user_id").Positive(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("items").
			Field("user_id").
			Unique().
			Required(),
	}
}

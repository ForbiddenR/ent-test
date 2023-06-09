package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "customer_pet"},
	}
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return nil
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}

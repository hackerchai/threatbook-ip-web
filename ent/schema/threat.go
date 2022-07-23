package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Threat holds the schema definition for the Threat entity.
type Threat struct {
	ent.Schema
}

func (Threat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "threat"},
	}
}

// Fields of the Threat.
func (Threat) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("ip").Unique(),
		field.String("threat_id_info"),
		field.Int("domain_count"),
		field.Int("tag_count"),
		field.Int("itel_count"),
		field.Int("judge"),
		field.Bool("poc"),
		field.Int("source"),
		field.Int64("ctime"),
	}
}

// Edges of the Threat.
func (Threat) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash_id").
			Annotations(entsql.Annotation{Size: 250}).
			NotEmpty().
			Unique(),
		field.String("title").
			Annotations(entsql.Annotation{Size: 250}).
			NotEmpty(),
		field.String("git_url").
			NotEmpty(),
		field.String("dir_path").
			NotEmpty(),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}

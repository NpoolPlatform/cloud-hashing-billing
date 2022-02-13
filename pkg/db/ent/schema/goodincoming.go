package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// GoodIncoming holds the schema definition for the GoodIncoming entity.
type GoodIncoming struct {
	ent.Schema
}

// Fields of the GoodIncoming.
func (GoodIncoming) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("good_id", uuid.UUID{}),
		field.UUID("coin_type_id", uuid.UUID{}),
		field.UUID("account_id", uuid.UUID{}),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the GoodIncoming.
func (GoodIncoming) Edges() []ent.Edge {
	return nil
}

// Indexes of the GoodIncoming.
func (GoodIncoming) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "coin_type_id").
			Unique(),
	}
}

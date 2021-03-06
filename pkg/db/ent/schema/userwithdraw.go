package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// UserWithdraw holds the schema definition for the UserWithdraw entity.
type UserWithdraw struct {
	ent.Schema
}

// Fields of the UserWithdraw.
func (UserWithdraw) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.String("name"),
		field.String("message"),
		field.JSON("labels", []string{}),
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

// Edges of the UserWithdraw.
func (UserWithdraw) Edges() []ent.Edge {
	return nil
}

// Indexes of the UserWithdraw.
func (UserWithdraw) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "user_id", "coin_type_id", "account_id").
			Unique(),
	}
}

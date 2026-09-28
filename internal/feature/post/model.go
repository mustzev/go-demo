package post

import "github.com/uptrace/bun"

type Post struct {
	bun.BaseModel `bun:"table:posts"`

	ID        int    `bun:"id,pk,autoincrement" json:"id"`
	Title     string `bun:"title,notnull" json:"title"`
	Body      string `bun:"body,notnull" json:"body"`
	Published bool   `bun:"published,notnull" json:"published"`
}

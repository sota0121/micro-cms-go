package db

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"users,alias:user"`

	ID       int64  `bun:"type:serial,pk,autoincrement"`
	Name     string `bun:"type:varchar(255),notnull"`
	Email    string `bun:"type:varchar(255),notnull"`
	Password string `bun:"type:varchar(255),notnull"`
}

type Tag struct {
	bun.BaseModel `bun:"tags,alias:tag"`

	ID   int64  `bun:"type:serial,pk,autoincrement"`
	Name string `bun:"type:varchar(255),notnull"`
}

type Article struct {
	bun.BaseModel `bun:"articles,alias:article"`

	ID      int64  `bun:"type:serial,pk,autoincrement"`
	Title   string `bun:"type:varchar(255),notnull"`
	Content string `bun:"type:varchar(255),notnull"`
	UserId  int64  `bun:"type:integer,notnull"`
}

type ArticleTag struct {
	bun.BaseModel `bun:"article_tags,alias:article_tag"`

	ArticleId int64 `bun:"type:integer,notnull"`
	TagId     int64 `bun:"type:integer,notnull"`
}

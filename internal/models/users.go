package models

type User struct {
	ID   uint64 `db:"id" json:"-"`
	Name string `db:"name" json:"name"`
	Age  int64  `db:"age" json:"age"`
}

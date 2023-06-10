package psql

import (
	"airbnd/internal/models"
	"log"

	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, usr *models.User) error {
	// ctx, cancel := context.WithCancel(ctx)
	// cancel()

	// ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	// cancel()

	_, err := r.db.QueryContext(ctx, "INSERT INTO users VALUES ($1, $2)", usr.Name, usr.Age)
	if err != nil {
		log.Print("error: %v", err)
		return err
	}

	return nil
	// defer rows.Close()

	// var users []*models.User
	// for rows.Next() {
	// var user models.User
	// if rows.Scan(
	// 		&user.ID, &user.Name, &user.Age,
	// 	);

	// 	users = append(users, &user)
	// }

	// return users

}

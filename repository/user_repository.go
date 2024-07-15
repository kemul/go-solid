package repository

import (
	"database/sql"
	"go-solid/domain"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

func NewPostgresUserRepository(Conn *sql.DB) domain.UserRepository {
	return &postgresUserRepository{Conn: Conn}
}

func (r *postgresUserRepository) GetByID(id int64) (*domain.User, error) {
	query := "SELECT user_id, username, email FROM users WHERE user_id = $1"
	row := r.Conn.QueryRow(query, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *postgresUserRepository) Store(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := r.Conn.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	return err
}

package user

import (
	outermodel "auth/internal/model"
	"auth/internal/repository"
	innerlmodel "auth/internal/repository/user/model"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

const (
	tableName = "users"

	idUserColumn        = "id"
	nameUserColumn      = "username"
	emailUserColumn     = "email"
	roleUserColumn      = "role"
	passwordHash        = "password"
	createdAtUserColumn = "created_at"
	updatedAtUserColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Get(ctx context.Context, id int64) (*outermodel.User, error) {
	builder := sq.Select(idUserColumn, nameUserColumn, emailUserColumn, passwordHash, roleUserColumn, createdAtUserColumn, updatedAtUserColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": id}).
		Limit(1)
	log.Printf("Before build sql User with id=%d, get", id)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var userModel innerlmodel.User
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&userModel.Id,
		&userModel.UserInfo.Name,
		&userModel.UserInfo.Email,
		&userModel.UserInfo.Password,
		&userModel.UserInfo.Role,
		&userModel.CreatedAt,
		&userModel.UpdatedAt)
	if err != nil {
		return nil, err
	}
	log.Printf("User with id=%d, get records[name: %s, email: %s]", id, userModel.UserInfo.Name, userModel.UserInfo.Email)

	// Converter from db model to service model
	return &outermodel.User{
		Id: userModel.Id,
		UserInfo: outermodel.UserInfo{
			Name:     userModel.UserInfo.Name,
			Email:    userModel.UserInfo.Email,
			Password: userModel.UserInfo.Password,
			Role:     userModel.UserInfo.Role,
		},
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}, nil
}

func (r *repo) Create(ctx context.Context, user *outermodel.UserInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameUserColumn, emailUserColumn, passwordHash, roleUserColumn).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return userID, err
	}

	return userID, nil
}

func (r *repo) Update(ctx context.Context, id int64, user *outermodel.UserInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameUserColumn, user.Name).
		Set(emailUserColumn, user.Email).
		Set(roleUserColumn, user.Role).
		Set(updatedAtUserColumn, sql.NullTime{Time: time.Now(), Valid: true}).
		Where(sq.Eq{idUserColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Error updating user with id=%d: %v", id, err)
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	log.Printf("Starting to delete user with id=%d", id)
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idUserColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		log.Printf("Error building SQL query for delete: %v", err)
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Error deleting user with id=%d: %v", id, err)
		return err
	}

	return nil
}

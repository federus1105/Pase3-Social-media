package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/federus1105/socialmedia/internals/pkg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Register(ctx context.Context, user models.UserRegister) (models.UserRegister, error) {
	// Start transaction
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Println("Failed to begin transaction:", err)
		return models.UserRegister{}, err
	}
	defer tx.Rollback(ctx)

	// Hash password
	hc := pkg.NewHashConfig()
	hc.UseRecommended()
	hashedPassword, err := hc.GenHash(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return models.UserRegister{}, err
	}

	sql := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING userid, email, password`
	values := []any{user.Email, hashedPassword}
	var newUser models.UserRegister
	err = r.db.QueryRow(ctx, sql, values...).Scan(&newUser.Id, &newUser.Email, &newUser.Password)
	if err != nil {
		log.Println("Failed to insert into users: ", err.Error())
		return models.UserRegister{}, err
	}
	// create user_id di table account
	account := `INSERT INTO account (user_id) VALUES ($1)`

	_, err = tx.Exec(ctx, account, newUser.Id)
	if err != nil {
		log.Println("Failed to insert empty account:", err)
		return models.UserRegister{}, err
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		return models.UserRegister{}, err
	}

	return newUser, nil
}

func (a *AuthRepository) Login(rctx context.Context, email string) (models.UserAuth, error) {
	sql := `SELECT userid, email, password FROM users WHERE email = $1`

	var user models.UserAuth
	if err := a.db.QueryRow(rctx, sql, email).Scan(&user.Id, &user.Email, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return models.UserAuth{}, errors.New("user not found")
		}
		log.Println("Internal Server Error.\nCause: ", err.Error())
		return models.UserAuth{}, err
	}
	return user, nil
}

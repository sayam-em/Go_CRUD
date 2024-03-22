package handler

import "github.com/jackc/pgx/v5/pgxpool"

type handler struct {
	store pgxpool
}
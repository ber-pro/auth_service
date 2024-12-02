package config

import (
	"errors"
	"os"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

const pgDSN = "PG_DSN"

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(pgDSN)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}
	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (pgCfg *pgConfig) DSN() string {
	return pgCfg.dsn
}

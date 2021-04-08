package db

import (
	"context"
)

const (
	LevelSerializable    = "serializable"
	LevelRepeatableRead  = "repeatable read"
	LevelReadCommitted   = "read committed"
	LevelReadUncommitted = "read uncommitted"
)

type (
	DB interface {
		Close() error
		Exec(query string, args ...interface{}) (Result, error)
		Query(query string, args ...interface{}) (Rows, error)
		QueryRow(query string, args ...interface{}) Row
		BeginTx(ctx context.Context, isolationLevel string) (Tx, error)
		ErrNoRows() error
		ErrGetCode(err error) string
	}

	Tx interface {
		ExecContext(ctx context.Context, query string, args ...interface{}) (Result, error)
		QueryContext(ctx context.Context, query string, args ...interface{}) (Rows, error)
		QueryRowContext(ctx context.Context, query string, args ...interface{}) Row
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}

	Result interface {
		RowsAffected() (int64, error)
	}

	Rows interface {
		Close() error
		Err() error
		Next() bool
		Scan(dest ...interface{}) error
	}

	Row interface {
		Scan(dest ...interface{}) error
	}

	Scannable interface {
		Scan(dest ...interface{}) error
	}

	ConvertParameters interface {
		ConvertParameters(string, []interface{}) (string, []interface{})
	}
)

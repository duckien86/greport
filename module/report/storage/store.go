package reportstorage

import (
	"github.com/ClickHouse/clickhouse-go/v2"
)

type sqlStore struct {
	db clickhouse.Conn
}

func NewSQLStore(db clickhouse.Conn) *sqlStore {
	return &sqlStore{db: db}
}

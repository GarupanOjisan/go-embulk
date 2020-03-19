package mysql

import (
	"bytes"
	"database/sql"
	"io"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/sqltocsv"
)

type MySQL struct {
	dsn   string
	query string
	args  []interface{}
}

// dsn: refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
func NewMySQLInput(dsn, query string, args ...interface{}) *MySQL {
	return &MySQL{
		dsn:   dsn,
		query: query,
		args:  args,
	}
}

func (m *MySQL) Read(p []byte) (n int, err error) {
	db, err := m.connect()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	rows, err := db.Query(m.query, m.args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	data, err := sqltocsv.WriteString(rows)
	if err != nil {
		return 0, err
	}

	src := bytes.NewBuffer([]byte(data))
	dst := bytes.NewBuffer(p)
	writes, err := io.Copy(dst, src)
	return int(writes), err
}

func (m *MySQL) connect() (*sql.DB, error) {
	return sql.Open("mysql", m.dsn)
}

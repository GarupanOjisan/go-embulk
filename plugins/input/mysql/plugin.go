package mysql

import (
	"database/sql"
	"github.com/joho/sqltocsv"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	dsn   string
	query string
	args  []interface{}

	data []byte
	pos  int
}

// dsn: refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
func NewMySQLInput(dsn, query string, args ...interface{}) *MySQL {
	return &MySQL{
		dsn:   dsn,
		query: query,
		args:  args,
		data:  nil,
		pos:   0,
	}
}

func (m *MySQL) Read(p []byte) (n int, err error) {
	src := m.data[m.pos:]
	for n = 0; n < len(p) && n < len(src); n++ {
		p[n] = src[n]
	}
	m.pos += n

	if m.pos == len(m.data) {
		return n, io.EOF
	}
	return n, nil
}

func (m *MySQL) Init() error {
	db, err := sql.Open("mysql", m.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(m.query, m.args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	data, err := sqltocsv.WriteString(rows)
	if err != nil {
		return err
	}
	m.data = []byte(data)

	return nil
}

func (m *MySQL) Finalize() error {
	return nil
}

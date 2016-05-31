package example

import (
	"database/sql"
)

//go:generate gdbinder -type=TestStruct
type TestStruct struct {
	ID    int            `db:"id" hoge:"fuga"`
	Col1  sql.NullString `db:"col1"`
	Col2  sql.NullString `db:"col2"`
	Col3  sql.NullString `db:"col3"`
	Col4  sql.NullString `db:"col4"`
	Col5  sql.NullString `db:"col5"`
	Col6  sql.NullString `db:"col6"`
	Col7  sql.NullString `db:"col7"`
	Col8  sql.NullString `db:"col8"`
	Col9  sql.NullString `db:"col9"`
	Col10 sql.NullString `db:"col10"`
}

//go:generate gdbinder -type=TestStruct2
type TestStruct2 struct {
	ID    int            `db:"id"`
	Col1  sql.NullString `db:"col1"`
	Col2  sql.NullString `db:"col2"`
	Col3  sql.NullString `db:"col3"`
	Col4  sql.NullString `db:"col4"`
	Col5  sql.NullString `db:"col5"`
	Col6  sql.NullString `db:"col6"`
	Col7  sql.NullString `db:"col7"`
	Col8  sql.NullString `db:"col8"`
	Col9  sql.NullString `db:"col9"`
	Col10 sql.NullString `db:"col10"`
}

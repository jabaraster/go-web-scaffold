package model

import (
    "os"

    _ "github.com/mattn/go-sqlite3"
    _ "github.com/lib/pq"
    "github.com/naoina/genmai"

    "../env"
)

var (
    db *genmai.DB
)

type Duplicate interface {
    ColumnName() string
    ColumnValue() interface{}
}

type duplicate struct {
    columnName string
    columnValue interface{}
}
func (e *duplicate) ColumnName() string {
    return e.columnName
}
func (e *duplicate) ColumnValue() interface{} {
    return e.columnValue
}

func NewDuplicate(columnName string, columnValue interface{}) Duplicate {
    return &duplicate{columnName: columnName, columnValue: columnValue}
}

type NotFound interface {
}
type notFound struct {
}

func NewNotFound() NotFound {
    return &notFound{}
}

func createTable(tableType interface{}) {
    if err := db.CreateTableIfNotExists(tableType); err != nil {
        panic(err)
    }
}

func init() {
    var err error
    db, err = createDb()
    if err != nil {
        panic(err)
    }

    db.SetLogOutput(os.Stdout)

    createTable(&Order{})
}

func createDb() (*genmai.DB, error) {
    switch env.DbKind() {
        case env.DbKindSQLite3:
            return genmai.New(&genmai.SQLite3Dialect{}, "./curtain.db")
        case env.DbKindPostgreSQL:
            return genmai.New(&genmai.PostgresDialect{}, "host=" + env.PostgresHost() + " dbname=" + env.PostgresDbName() + " user=" + env.PostgresUser() + " password=" + env.PostgresPassword() + " sslmode=" + env.PostgresSslMode())
    }
    panic(env.DbKind())
}

func tx() {
    if err := recover(); err != nil {
        db.Rollback()
    } else {
        db.Commit()
    }
}

func beginTx() {
    if err := db.Begin(); err != nil {
        panic(err)
    }
}

type id struct {
    Id int64
    Name string
}

func checkNameDuplicateForInsert(tableType interface{}, name string) bool {
    var c int64
    if err := db.Select(&c, db.Count(), db.From(tableType), db.Where("name","=",name)); err != nil {
        panic(err)
    }
    return c == 0
}

func checkNameDuplicateForUpdate(tableType interface{}, updateTargetId int64, name string) bool {
    var c int64
    if err := db.Select(&c, db.Count(), db.From(tableType), db.Where("name","=",name).And(db.Where("id","<>",updateTargetId))); err != nil {
        panic(err)
    }
    return c == 0
}

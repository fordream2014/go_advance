package dao

import "database/sql"

type MysqlBase struct {
}

// mock query row
func (db *MysqlBase) QueryRow() (map[string]interface{}, error) {
	return nil, sql.ErrNoRows
}

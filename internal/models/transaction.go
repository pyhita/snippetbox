package models

import "database/sql"

// 事务测试
type ExampleModel struct {
	DB *sql.DB
}

func (m *ExampleModel) ExampleTransaction() error {

	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO ....")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE ....")
	if err != nil {
		return err
	}

	return tx.Commit()
}

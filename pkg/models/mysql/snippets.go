package mysql

import (
	"database/sql"

	"dshurubtsov.com/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Method for create new snippet to database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {

	// SQL query for instert data to base
	stmt := `INSERT INTO snippets (title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// method LastInsertId() so get last ID inserted entry from table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

package mysql

import (
	"database/sql"
	"errors"

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
	// SQL query for get data of one entry
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Use the QueryRow() method to execute the SQL query,
	// passing the untrusted id variable as the value for the placeholder
	// Returns a pointer to an sql.Row object that contains the record data.
	row := m.DB.QueryRow(stmt, id)

	// init pointer on new struct Snippet{}
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// return error from model ErrNoRecord
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}

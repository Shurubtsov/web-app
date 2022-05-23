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

	// method LastInsertId() so get last ID inserted note from table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Method for Get snippet from ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	// SQL query for get data of one note
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

// Method for get Latest note
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {

	// SQL query for get latest data from table snippets
	stmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	// Use the Query() method to execute our SQL query.
	// In response, we will receive sql.Rows, which contains the result of our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// Delay call rows.Close() to sure what results close right
	defer rows.Close()

	// init slice for storage objects models.Snippet
	var snippets []*models.Snippet

	for rows.Next() {
		// pointer to new struct Snippet
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

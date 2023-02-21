// Filename: questions.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// The Question model will represent a single question in our DB
type Question struct {
	ID int64
	Content string
	CreatedAt time.Time
}
// The QuestionModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type QuestionModel struct {
	DB *sql.DB
}

func (m *QuestionModel) Insert(body string) (int64, error) {
	var id int64
	statement := 
	            `
							INSERT INTO questions(content)
							VALUES($1)
							RETURNING question_id
	            `
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package sqlc

import (
	"time"

	zero "gopkg.in/guregu/null.v4/zero"
)

type Integration struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt zero.Time `json:"deleted_at"`
}

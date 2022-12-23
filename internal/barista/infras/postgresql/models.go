// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgresql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type BaristaBaristaOrder struct {
	ID       uuid.UUID
	ItemType int32
	ItemName string
	TimeUp   time.Time
	Created  time.Time
	Updated  sql.NullTime
}

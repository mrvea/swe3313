// Package types defines some types shared by other subpackages, idea for a better name anyone?
package types

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DBReader interface {
	GetSource() string
}

type DBWriter interface {
}

// Nullstring is a wrapper for sql.Nullstring that implements json.Marshaler
type NullString struct {
	sql.NullString
}

func (n NullString) MarshalJSON() (b []byte, err error) {
	if !n.Valid {
		b, err = json.Marshal(nil)
	} else {
		b, err = json.Marshal(n.String)
	}
	return b, err
}

// Set sets the value to the string value of value if possible, otherwise sets Valid to false
// Panics if value is not nil or string
func (n *NullString) Set(value interface{}) {
	if value == nil {
		n.String = ""
		n.Valid = false
		return
	}
	v := value.(string)
	if v == "" {
		n.Valid = false
		n.String = v
		return
	}
	n.String = v
	n.Valid = true
}

type NullTime struct {
	mysql.NullTime
}

func (n NullTime) Set(value interface{}) {
	if value == nil {
		n.Valid = false
		return
	}
	v := value.(time.Time)
	n.Time = v
	n.Valid = true
}

// MarshalJSON implements the json.Marshaller interface, and formats all dates as YYYY-MM-DD
func (n NullTime) MarshalJSON() (b []byte, err error) {
	if !n.Valid {
		b, err = json.Marshal(nil)
	} else {
		b, err = json.Marshal(n.Time.Format("2006-01-02"))
	}
	return b, err
}
func (n NullString) getString() string {
	return n.String
}

type TimeStamper interface {
	CreateStamper
	ModifyStamper
}

type CreateStamper interface {
	SetCreated()
}

type ModifyStamper interface {
	SetModified()
	GetModified() time.Time
}

type Deleter interface {
	Delete()
}

type Starter interface {
	SetStart(string) error
	GetStart() time.Time
}

type Ender interface {
	SetEnd(string) error
	GetEnd() time.Time
}

type StartEnder interface {
	Starter
	Ender
}

type RoomBlock interface {
	StartEnder
	GetQuantity() int
}

package model

import (
	"sync"

	authboss "gopkg.in/authboss.v1"

	log "github.com/class/pizza/logger"
	"github.com/class/pizza/types"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const (
	USER_TABLE        = "users"
	USER_INSERT_QUERY = `
		INSERT INTO ` + USER_TABLE + `
		VALUES(
			:role,
			:email,
			:firlst_name,
			:last_name,
			:password
		)
	`
)

var (
	userTableCache              *userTable
	userTableCacheInstantiation = sync.Once{}
)

type User struct {
	DefaultModel

	Role               string `db:"role_slug"`
	FirstName          string `db:"first_name"`
	LastName           string `db:"last_name"`
	Password           string
	Email              string
	RecoverToken       types.NullString `db:"recover_token" json:"-"`
	RecoverTokenExpiry types.NullTime   `db:"recover_token_expiry" json:"-"`

	DefaultDeleter
	DefaultTimeStamper
	DefaultAuditor
}

func NewUser() *User {
	u := User{}
	u.source = UserTable()
	return &u
}

type userTable struct {
	DefaultTable
}

func UserTable() *userTable {
	userTableCacheInstantiation.Do(func() {
		userTableCache = &userTable{
			DefaultTable{USER_TABLE},
		}
	})
	return userTableCache
}

func (U *User) Populate(data map[string]interface{}) *User {
	U.SetFromData(&U.Email, data, "Email")
	U.SetFromData(&U.Role, data, "Role")
	U.SetFromData(&U.FirstName, data, "FirstName")
	U.SetFromData(&U.LastName, data, "LastName")

	return U
}

func (U *User) Insert(optionals ...interface{}) (err error) {
	var tx *sqlx.Tx
	commit := false

	optionals = getFirstOptional(tx, optionals)
	if tx == nil {
		tx, err = db.Beginx()
		if err != nil {
			return
		}
		commit = true
	}
	_, err = tx.NamedExec(USER_INSERT_QUERY, U)
	if err != nil {
		if commit {
			tx.Rollback()
		}
		return errors.Wrap(err, "Failed to Insert New User")
	}

	if commit {
		err = tx.Commit()
		if err != nil {
			return errors.Wrap(err, "Failed to Commit transaction in New User Insert")
		}
	}
	return nil
}

func (u *userTable) Insert(data map[string]interface{}) (err error) {

	tx, err := db.Beginx()
	if err != nil {
		return
	}
	U := NewUser().Populate(data)

	err = U.Insert(tx)
	if err != nil {

		tx.Rollback()

		return errors.Wrap(err, "Failed to Insert New User")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "Failed to Commit transaction in New User Insert")
	}

	return nil
}

//Interface with authboss register module
func (u *userTable) Create(key string, attr authboss.Attributes) error {
	data := make(map[string]interface{}, 0)
	data["Email"] = attr["email"]
	data["Password"] = attr["password"]
	err := u.Insert(data)
	return err
}

func (u *userTable) Get(email string) (interface{}, error) {
	U, err := u.GetByEmail(email)
	if err != nil {
		log.Info("Error getting user: ", err)
		// TODO tj invalidate existing login if exists
		return nil, authboss.ErrUserNotFound
	}
	return U, err
}
func (u *userTable) GetByEmail(email string) (*User, error) {
	U := NewUser()
	err := db.Get(U, "SELECT * FROM users WHERE email = ? AND `active` = 1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	return U, err
}

//Put is used by authboss to update information in the database from authboss, currently only used for password recovery
func (u *userTable) Put(key string, attr authboss.Attributes) error {
	U, err := u.GetByEmail(key)
	if err != nil {
		return err
	}
	U.SetFromData(&U.Password, attr, "password")
	U.SetFromData(&U.RecoverToken, attr, "recover_token")
	U.SetFromData(&U.RecoverTokenExpiry, attr, "recover_token_expiry")

	_, err = db.NamedExec(
		`UPDATE users SET
			password = :password,
			recover_token = :recover_token,
			recover_token_expiry = :recover_token_expiry,
			modified = :modified
		WHERE id = :id
		`, U)
	return err
}
func (u *userTable) RecoverUser(recoverToken string) (interface{}, error) {
	U := NewUser()
	log.Debug("Recover User with Token")
	log.Debug(recoverToken)
	err := db.Get(U, "SELECT * FROM users WHERE recover_token = ? LIMIT 1", recoverToken)
	log.Debug(err)
	if err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return U, err
}

package model

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/class/pizza/constants"
	"github.com/class/pizza/env"
	log "github.com/class/pizza/logger"
	"github.com/class/pizza/types"
	"github.com/eyesore/gomigrate"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	// RetryInterval is the delay in seconds between connection retries against the database.
	RetryInterval time.Duration = 2

	// ErrDatabaseConnectionFailed indicates that database discover failed after specified timeout.
	ErrDatabaseConnectionFailed = errors.New("Unable to establish database connection.")

	// ErrNoIDForPopulate indicates that populate cannot be performed on a model with a zero value ID
	ErrNoIDForPopulate = errors.New("ID is not defined for this model.")

	// ErrCannotEnumerateFields
	ErrCannotEnumerateFields = errors.New("Trying to enumerate fields of something other that a struct or ptr to a struct.")
)

type DefaultTimeStamper struct {
	Created  time.Time
	Modified time.Time
}

func (dts *DefaultTimeStamper) SetCreated() {
	dts.Created = time.Now()
}

func (dts *DefaultTimeStamper) SetModified() {
	dts.Modified = time.Now()
}

func (dts *DefaultTimeStamper) GetModified() time.Time {
	return dts.Modified
}

func (dts *DefaultTimeStamper) SetTimestamps() {
	dts.SetCreated()
	dts.SetModified()
}

type DefaultAuditor struct {
	CreatorID  int `db:"creator_id"`
	ModifierID int `db:"modifier_id"`
}

func (a *DefaultAuditor) SetCreator(u *User) {
	a.CreatorID = u.ID
}

func (a *DefaultAuditor) SetModifier(u *User) {
	a.ModifierID = u.ID
}

func (a *DefaultAuditor) SetCreateModifier(u *User) {
	a.SetCreator(u)
	a.SetModifier(u)
}

type DefaultDeleter struct {
	Active bool `db:"active"`
}

func (d *DefaultDeleter) Delete() {
	d.Active = false
}

type DefaultStartEnder struct {
	Start types.NullTime `db:"start"`
	End   types.NullTime `db:"end"`
}

func (s *DefaultStartEnder) SetStart(start string) error {
	validFormats := []string{
		time.RFC1123,
		time.RFC3339,
	}
	var err error
	for _, f := range validFormats {
		s.Start.Time, err = time.Parse(f, start)
		if err == nil {
			s.Start.Valid = true
			break
		}
	}
	return err
}

func (s *DefaultStartEnder) SetEnd(end string) error {
	validFormats := []string{
		time.RFC1123,
		time.RFC3339,
	}
	var err error
	for _, f := range validFormats {
		s.End.Time, err = time.Parse(f, end)
		if err == nil {
			s.End.Valid = true
			break
		}
	}
	return err
}

func (s *DefaultStartEnder) GetStart() time.Time {
	return s.Start.Time
}

func (s *DefaultStartEnder) GetEnd() time.Time {
	return s.End.Time
}

type Model interface {
	GetSource() Table
	GetID() int
}

type DefaultModel struct {
	source Table
	ID     int `db:"id"`
}

func (m *DefaultModel) GetSource() Table {
	return m.source
}

func (m *DefaultModel) GetID() int {
	return m.ID
}

func (m *DefaultModel) Validate() error {
	// TODO
	return nil
}

func (m *DefaultModel) Set(dest interface{}, value interface{}) {
	switch t := dest.(type) {
	case *int:
		m.setInt(t, value)
	case *bool:
		m.setBool(t, value)
	case *string:
		m.setString(t, value)
	case *types.NullString:
		m.setNullString(t, value)
	case *types.NullTime:
		m.setNullTime(t, value)
	case *float64:
		m.setFloat64(t, value)
	default:
		log.Infof("Illegal type %v passed to Model.Set.  Does this need to be implemented?", t)
	}
}

// SetFromData sets a value on the model if it exists in data, otherwise it does nothing
func (m *DefaultModel) SetFromData(dest interface{}, data map[string]interface{}, key string) {
	if val, ok := data[key]; ok {
		m.Set(dest, val)
	}
}

// TODO sanitize in mutators
func (m *DefaultModel) setInt(dest *int, value interface{}) {
	v := int(value.(float64))
	*dest = v
}

func (m *DefaultModel) setString(dest *string, value interface{}) {
	v := value.(string)
	*dest = v
}

func (m *DefaultModel) setBool(dest *bool, value interface{}) {
	v := value.(bool)
	*dest = v
}

func (m *DefaultModel) setFloat64(dest *float64, value interface{}) {
	v := value.(float64)
	*dest = v
}

func (m *DefaultModel) setNullString(dest *types.NullString, value interface{}) {
	dest.Set(value)
}

func (m *DefaultModel) setNullTime(dest *types.NullTime, value interface{}) {
	dest.Set(value)
}

func Populate(m Model) error {
	id := m.GetID()
	if id == 0 {
		return errors.New("Populate: " + m.GetSource().GetName())
	}
	err := db.Get(m, "SELECT * FROM "+m.GetSource().GetName()+" WHERE id=?", id)
	return err
}

func (d *DefaultModel) Populate() error {
	id := d.ID
	if id == 0 {
		return ErrNoIDForPopulate
	}
	err := db.Get(d, "SELECT * FROM "+d.source.GetName()+" WHERE id=?", id)
	return err
}

func GetNameForFK(id string, FK string) (string, error) {
	table := ""
	switch FK {
	case "Attendee_Category_ID":
		table = "attendee_categories"
	default:
		table = strings.ToLower(strings.TrimSuffix(FK, "_ID")) + "s"
	}
	type Name struct {
		Name string
	}
	name := Name{}

	err := db.Get(&name, "SELECT name FROM "+table+" WHERE id=?", id)
	if err != nil {
		return "", err
	}

	return name.Name, nil
}

// TODO cache fields at application level per type
// Working but Not currently used
func GetAllFields(s interface{}) (fields []interface{}, err error) {
	v := reflect.Indirect(reflect.ValueOf(s))
	t := reflect.TypeOf(s)
	if v.Kind() != reflect.Struct {
		return nil, ErrCannotEnumerateFields
	}
	for i := 0; i < v.NumField(); i++ {
		var sf reflect.StructField
		if t.Kind() == reflect.Ptr {
			sf = t.Elem().Field(i)
		} else {
			sf = t.Field(i)
		}
		fv := reflect.Indirect(v.Field(i))
		// skip unexported fields
		if !fv.CanInterface() || !fv.CanAddr() {
			continue
		}
		if sf.Anonymous {
			embeddedFields, err := GetAllFields(fv.Addr().Interface())
			if err != nil {
				// TODO possibly just skip the field, maybe saving the skipped fields
				return nil, err
			}
			fields = append(fields, embeddedFields...)
			continue
		}
		if !fv.CanAddr() {
			return nil, err
		}
		fields = append(fields, fv.Addr().Interface()) // ensure that this is settable
	}
	return fields, err
}

type Table interface {
	GetName() string
}

type DefaultTable struct {
	Name string
}

func (t *DefaultTable) GetName() string {
	return t.Name
}
func (t *DefaultTable) Select(fields ...string) sq.SelectBuilder {
	if len(fields) == 0 {
		fields = append(fields, "*")
	}

	return sq.Select(fields...).From(t.GetName())
}

func getLastInsertID(tx *sqlx.Tx) (ID int, err error) {
	query := "SELECT LAST_INSERT_ID();"
	var row *sql.Row
	if tx == nil {
		row = db.QueryRow(query)
	} else {
		row = tx.QueryRow(query)
	}
	err = row.Scan(&ID)

	return
}

func SetDatasource(d *sqlx.DB) {
	d.SetConnMaxLifetime(1200 * time.Second)
	db = d
}

// DiscoverDatabase tests a connection pool until it is able to verify the connection
func DiscoverDatabase(cp *sqlx.DB, timeout time.Duration) bool {
	var err error
	var timeSpent time.Duration = 0

	for err = cp.Ping(); err != nil && timeSpent < timeout; {
		log.Info("Waiting for a database connection...")
		time.Sleep(RetryInterval * time.Second)
		err = cp.Ping()
		timeSpent += RetryInterval
	}

	if err != nil {
		log.Info("No database connection: ", err)
		return false
	}
	return true
}

func Migrate(rollback bool, howMany int) error {
	log.Debug("migrating")
	appEnv, err := env.New("orb")
	if err != nil {
		return err
	}
	if !DiscoverDatabase(appEnv.DB, 60) {
		return ErrDatabaseConnectionFailed
	}

	migrator, err := gomigrate.NewMigrator(appEnv.DB.DB, gomigrate.Mysql{}, "./schema/migrations")
	if err != nil {
		return err
	}

	if !rollback {
		log.Debug("Doing forward migration.")
		err = migrator.Migrate()
	} else {
		log.Debug("Rolling back", howMany, "migrations.")
		err = migrator.RollbackN(howMany)
	}

	if err != nil {
		return err
	}
	return nil
}

func getFirstOptional(dest interface{}, opt []interface{}) []interface{} {
	if len(opt) > 0 {
		SetValue(dest, opt[0])
		opt = opt[1:]
	}
	return opt
}
func SetValue(dest interface{}, value interface{}) {
	// log.Debug(value)
	switch t := dest.(type) {
	case *int:
		if v, ok := value.(int); ok {
			*t = v
		}
	case *string:
		if v, ok := value.(string); ok {
			*t = v
		}
	case *map[string]interface{}:
		if v, ok := value.(map[string]interface{}); ok {
			*t = v
		}
	case *bool:
		if v, ok := value.(bool); ok {
			*t = v
		}
	case *sqlx.Tx:
		if v, ok := value.(sqlx.Tx); ok {
			*t = v
		}
	case *time.Time:
		if v, ok := value.(string); ok {
			validFormats := constants.ValidFormats
			var err error
			for _, f := range validFormats {
				dest, err = time.Parse(f, v)
				if err != nil {
					log.Debug(err)
				}
			}
		}
	default:
		log.Debug("Destination type is not the same as value")
	}
}

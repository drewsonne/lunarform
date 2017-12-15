package database

// Record is an untyped, schemaless record type which all
// record types embed and implement
type Record map[string]interface{}

// Key returns a record's Key
func (r Record) Key() string {
	return r["Key"].(string)
}

// Type returns a record's type
func (r Record) Type() string {
	return r["Type"].(string)
}

// Driver represents a low level storage serialiser/ deserialiser
// This is wrapped in the Database
type Driver interface {
	Create(recordType, key string, doc interface{}) error
	Read(recordType, key string, i interface{}) error
	Update(recordType, key string, doc interface{}) error
	Delete(recordType, key string) error

	Ping() error
	Close() error
}

// BufferedDriver represents a low level storage which exposes a Bytes method to allow
// serialization to disk
type BufferedDriver interface {
	Driver
	Bytes() ([]byte, error)
}

// Database stores data for terraform server
type Database struct {
	driver Driver
}

// NewDatabaseWithDriver creates a new Database struct with
func NewDatabaseWithDriver(driver Driver) Database {
	return Database{
		driver: driver,
	}
}
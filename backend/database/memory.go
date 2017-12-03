package database

import (
	"encoding/json"
	"fmt"
)

func NewMemoryDatabase() (MemoryDatabase, error) {
	return MemoryDatabase{
		collections: make(map[string]string),
	}, nil
}

type MemoryDatabase struct {
	collections map[string]string
}

func (md MemoryDatabase) Close() error {
	return nil
}

func (md MemoryDatabase) Ping() error {
	return nil
}

func (md *MemoryDatabase) Create(recordType, key string, doc interface{}) error {
	if md.exists(recordType, key) {
		return fmt.Errorf("%q %q already exists", recordType, key)
	}

	md.collections[md.key(recordType, key)] = md.serialize(doc)

	return nil
}

func (md MemoryDatabase) Read(recordType, key string, i interface{}) error {
	if !md.exists(recordType, key) {
		return fmt.Errorf("%q %q doesn't exist", recordType, key)
	}

	md.deserialize(md.collections[md.key(recordType, key)], i)

	return nil
}

func (md MemoryDatabase) Update(recordType, key string, doc interface{}) error {
	if !md.exists(recordType, key) {
		return fmt.Errorf("%q %q doesn't exist", recordType, key)
	}

	md.collections[md.key(recordType, key)] = md.serialize(doc)

	return nil
}

func (md MemoryDatabase) Delete(recordType, key string) error {
	delete(md.collections, md.key(recordType, key))

	return nil
}

func (md MemoryDatabase) key(recordType, key string) string {
	return fmt.Sprintf("%s %s", recordType, key)
}

func (md MemoryDatabase) exists(recordType, key string) (ok bool) {
	_, ok = md.collections[md.key(recordType, key)]

	return
}

// The following exist because we expect to be able to serialise to/ from
// a pointer to a struct. The original implementation, for example, of
// Read() looked like:
//
// func (md MemoryDatabase) Read(recordType, key string, i interface{}) error {
//     if !md.exists(recordType, key) {
//         return fmt.Errorf("%q %q doesn't exist", recordType, key)
//     }
//
//     i = md.collections[md.key(recordType, key)]
//
//     return nil
// }
//
// this meant we were never really updating the i that was passed in (scoping, I assume)
// and so we weren't seeing the data we expected.
// Instead, then, we're going to do json

func (md MemoryDatabase) serialize(i interface{}) string {
	v, _ := json.Marshal(i)

	return string(v)
}

func (md MemoryDatabase) deserialize(s string, i interface{}) error {
	return json.Unmarshal([]byte(s), i)
}
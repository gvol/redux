package redo

import (
  "encoding/json"
)

func (p Prerequisite) encode() ([]byte, error) {
	return json.Marshal(p)
}

func decodePrerequisite(b []byte) (Prerequisite, error) {
	var p Prerequisite
	return p, json.Unmarshal(b, &p)
}

func (m Metadata) encode() ([]byte, error) {
	return json.Marshal(m)
}

// Get returns a database record decoded into the specified type.
func (f *File) Get(key string, obj interface{}) (bool, error) {
	data, found, err := f.db.Get(key)
	if err == nil && found {
		err = json.Unmarshal(data, &obj)
	}
	return found, err
}

// Put stores a database record.
func (f *File) Put(key string, obj interface{}) error {
	if b, err := json.Marshal(obj); err != nil {
		return err
	} else {
		return f.db.Put(key, b)
	}
}


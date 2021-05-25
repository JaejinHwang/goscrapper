package dictionarys

import "errors"

//create Dictionary type
type Dictionary map[string]string

var errNotfound = errors.New("cannot find an input word in a dictionary")

//method Search returns value of input key, and returns error if it doesn't exist
func (d Dictionary) Search (key string) (string, error) {
	value, exist := d[key]
	if exist {
		return value, nil
	}
	return "", errNotfound
}
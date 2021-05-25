package dictionarys

import "errors"

//create Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("cannot find an input word in a dictionary")
	errWordExsist = errors.New("input word already exists in a dictionary")
	errCantUpdate = errors.New("cannot update value of non-exist key")
	errCantDelete = errors.New("cannot delete value of non-exist key")
)

//method Search returns value of input key, and returns error if it doesn't exist
func (d Dictionary) Search (key string) (string, error) {
	value, exist := d[key]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

//method Add adds input word into dictionary and return error if input word already exists
func (d Dictionary) Add (key, value string) error {
	_, error := d.Search(key)
	switch error {
	case errNotFound:
		d[key] = value
	case nil:
		return errWordExsist
	}
	return nil
}

//method Update changes value of input key by input value in dictionary
func (d Dictionary) Update (key, value string) error {
	_, error := d.Search(key)
	switch error {
	case nil:
		d[key] = value
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//method Delete erase input key and value of it in dictionary
func (d Dictionary) Delete (key string) error {
	_, error := d.Search(key)
	switch error {
	case nil:
		delete(d, key)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
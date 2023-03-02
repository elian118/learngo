package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// Go -> map[key] -> 매핑된 값, 키 존재 여부 등 두 개 변수 반환
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

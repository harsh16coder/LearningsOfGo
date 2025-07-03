package mapstests

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

func (d Dictionary) Add() {

}

func (d Dictionary) Delete() {

}

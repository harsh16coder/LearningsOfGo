package mapstests

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return nil
	}
	return nil
}

func (d Dictionary) Delete() {

}

package main

type Dictionary map[string]string

const (
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrNotFound         = DictionaryErr("the word does not exist")
	ErrWordDoesNotExist = DictionaryErr("the word to update does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

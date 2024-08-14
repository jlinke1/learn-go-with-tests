package main

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("definition for word already exists")
	ErrWordDoesNotExist = DictionaryErr("the word to update does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return res, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

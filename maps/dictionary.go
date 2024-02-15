// a map value is a pointer to a runtime.hmap structure
// so when passing a map to func/method you are copying it
// but just the pointer part, not underlying DS
//
// a map can be "nil". But attempt to write to it causes a
// runtime panic. Therefore, never init an empty map var.
// WRONG: var m map[string]string USE:var dictionary = map[string]string{} OR var dictionary = make(map[string]string)
package main

type Dictionary map[string]string
type DictionaryErr string


var (
	ErrNotFound =		DictionaryErr("could not find the word you were lookong for")
	ErrWordExists =		DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist =	DictionaryErr("cannot update word because it does not exist")
)

//method to Search for word in Dictionary type
func (d Dictionary) Search(word string) (string, error) {

	//def=value retrieved & ok=bool(check it word exist)
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//method to Add word in Dictionary type
func (d Dictionary) Add(word, definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition //add value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

//since we assign update as an error in tests we have to
//return error type
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word) //builtin delete function"(map, key)"
}

//an error wrapper with constant error "DictionaryErr"
func (e DictionaryErr) Error() string {
	return string(e)
}

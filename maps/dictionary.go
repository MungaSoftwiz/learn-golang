package maps

// Dictionary type is special. It can only be a comparable type
// language spe: https://golang.org/ref/spec#Comparison_operators
type Dictionary map[string]string

// error type that implements "error" interface
type DictionaryErr string

// we made errors constant so requires us to create our own
// DictionaryErr type which implements "error" interface
// article: https://dave.cheney.net/2016/04/07/constant-errors
const (
	ErrNotFound         = DictionaryErr("could not find the word you were lookong for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// method to Search for word in Dictionary type
func (d Dictionary) Search(word string) (string, error) {

	//def=value retrieved & ok=bool(check it word exist)
	// we are using an interesting property of the map lookup.
	// It can return 2 values.(Allows us to diff a word that exist
	//and one that doesn't have a definition)
	definition, ok := d[word] // (value, exists(bool))
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// method to Add word in Dictionary type
func (d Dictionary) Add(word, definition string) error {
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

// since we assign update as an error in tests we have to
// return error type
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

// an error wrapper with constant error "DictionaryErr"
func (e DictionaryErr) Error() string {
	return string(e)
}

/* An interesting property of maps is that you can modify them
without passing as an address to it (e.g &myMap). Makes them feel
like a reference type. But they are not. They are value types.

As Dave Cheney describes:
https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
		A map is a pointer to a runtime.hmap structure

So when you pass a map to a function/method, you are indeed copying it,
but just the pointer part, not the underlying data structure that contains the data.

A map can be "nil". But attempt to write to it causes a
runtime panic. Therefore, never initialize an empty map var.
-> WRONG: var m map[string]string
-> USE:var dictionary = map[string]string{} OR var dictionary = make(map[string]string)
NB: Both approaches create an empty hash map and point dictionary at it.
which ensures that you will never get a runtime panic.*/

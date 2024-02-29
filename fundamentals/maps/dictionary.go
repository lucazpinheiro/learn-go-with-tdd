package maps

/*
A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map
when reading, but attempts to write to a nil map will cause a runtime panic. You can read more
about maps

	var m map[string]string

Instead, you can initialize an empty map like we were doing above, or use the make keyword to
create a map for you:

	var dictionary = map[string]string{}

// OR

	var dictionary = make(map[string]string)

Both approaches create an empty hash map and point dictionary at it. Which ensures that you
will never get a runtime panic.
*/
type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
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
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		/*
			We could reuse ErrNotFound and not add a new error. However, it is often better to have
			a precise error for when an update fails.

			Having specific errors gives you more information about what went wrong.
			Here is an example in a web app:

			You can redirect the user when ErrNotFound is encountered,
			but display an error message when ErrWordDoesNotExist is encountered.
		*/
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

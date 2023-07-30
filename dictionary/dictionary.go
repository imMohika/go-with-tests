package dictionary

import "errors"

type Dictionary map[string]string

var ErrUnknownKey = errors.New("key doesn't exist in dictionary")
var ErrExistingKey = errors.New("key already exists in dictionary")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrUnknownKey
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err != ErrUnknownKey {
		return ErrExistingKey
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, newValue string) error {
	_, err := d.Search(key)

	if err != nil {
		return err
	}

	d[key] = newValue
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

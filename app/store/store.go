package store

import "errors"

var store = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

func Put(key, val string) error {
	store[key] = val

	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Del(key string) error {
	delete(store, key)

	return nil
}

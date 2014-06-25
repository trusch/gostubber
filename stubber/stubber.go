package stubber

import (
	"errors"
)

var stubs map[string][]byte = make(map[string][]byte)

func registerStub(name string, data []byte) {
	stubs[name] = data
}

func Get(name string) ([]byte, error) {
	if stub, ok := stubs[name]; ok {
		return stub, nil
	}
	return nil, errors.New("couldn't find stub \"" + name + "\"")
}

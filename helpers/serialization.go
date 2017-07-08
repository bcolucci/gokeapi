package helpers

import (
	"bytes"
	"encoding/gob"
)

func Serialize(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(key); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func Unserialize(data []byte, v interface{}) {
	reader := bytes.NewReader(data)
	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(v); err != nil {
		panic(err)
	}
}

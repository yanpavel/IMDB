package storage

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Key struct {
	Id int64
}

type Value struct {
	Measure   int64
	IsDeleted bool
	CreatedAt int64
}

func (k *Key) Marshal() []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.LittleEndian, k)
	if err != nil {
		fmt.Println(err)
	}
	return buffer.Bytes()
}

func (k *Key) Unmarshal(data []byte) {
	binary.Decode(data, binary.LittleEndian, k)
}

func (v *Value) Marshal() []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.LittleEndian, v)
	if err != nil {
		fmt.Println(err)
	}
	return buffer.Bytes()
}

func (v *Value) Unmarshal(data []byte) {
	binary.Read(bytes.NewReader(data), binary.LittleEndian, v)
}

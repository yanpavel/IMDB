package storage

import (
	"bytes"
	"encoding/binary"
	"time"
)

type Key struct {
	Id int64
}

type Value struct {
	Name      string
	IsDeleted bool
	CreatedAt time.Time
}

func (k *Key) Marshal() []byte {
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.LittleEndian, k)
	return buffer.Bytes()
}

func (k *Key) Unmarshal(data []byte) {
	binary.Decode(data, binary.LittleEndian, k)
}

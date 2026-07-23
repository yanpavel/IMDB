package storage

import (
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/yanpavel/imdb/bloom"
)

type MemTable struct {
	Peoples     map[Key]Value
	BloomFilter bloom.BloomFilter
}

func NewRepository(bloomFilter bloom.BloomFilter) *MemTable {
	return &MemTable{
		BloomFilter: bloomFilter,
		Peoples:     make(map[Key]Value),
	}
}

func (r *MemTable) Add(k Key, v Value) {
	v.CreatedAt = time.Now().UnixNano()
	r.Peoples[k] = v
	d := k.Marshal()
	r.BloomFilter.Add(d)
}

func (r *MemTable) Delete(k Key) {
	v := Value{IsDeleted: true}
	r.Add(k, v)
}

func (r *MemTable) GetByKey(k Key) *Value {
	v, exists := r.Peoples[k]
	if !exists {
		return nil
	}

	return &v
}

func (r *MemTable) Flush(path string) error {
	if _, err := os.OpenRoot(path); err != nil {
		return errors.Wrapf(err, "invalid directory %s", path)
	}

	timestamp := time.Now().Format("20060102150405")
	file_keys, err := os.Create("sstable_keys" + timestamp)
	if err != nil {
		return errors.Wrap(err, "create sstable_keys file error")
	}

	file_values, err := os.Create("sstable_values" + timestamp)
	if err != nil {
		return errors.Wrap(err, "create sstable_values file error")
	}

	defer file_values.Close()
	defer file_keys.Close()

	var bytesWritten int

	for k, v := range r.Peoples {
		vbytes := v.Marshal()
		bytesWrittenLocal, err := file_values.Write(vbytes)
		if err != nil {
			fmt.Println("error of value writing")
		}
		bytesWritten += bytesWrittenLocal
		file_values.Seek(int64(bytesWritten), io.SeekStart)

		keyBytes := k.Marshal()

		writeKeyData := func(data []byte) {

			length := uint16(len(data))
			binary.Write(file_keys, binary.LittleEndian, length)
			file_keys.Write(data)
			binary.Write(file_keys, binary.LittleEndian, int64(bytesWritten))
			file_keys.Write([]byte(strconv.Itoa(bytesWritten)))
		}

		writeKeyData(keyBytes)
	}

	return nil
}

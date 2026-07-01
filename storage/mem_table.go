package storage

import (
	"sync"
	"time"

	"github.com/yanpavel/imdb/bloom"
)

type MemTable struct {
	Peoples     sync.Map
	BloomFilter bloom.BloomFilter
}

func NewRepository(bloomFilter bloom.BloomFilter) *MemTable {
	return &MemTable{
		BloomFilter: bloomFilter,
	}
}

func (r *MemTable) Add(k Key, v Value) {
	v.CreatedAt = time.Now()
	r.Peoples.Store(k, v)
	d := k.Marshal()
	r.BloomFilter.Add(d)
}

func (r *MemTable) Delete(k Key) {
	v := Value{IsDeleted: true}
	r.Add(k, v)
}

func (r *MemTable) GetByKey(k Key) *Value {
	v, exists := r.Peoples.Load(k)
	if !exists {
		return nil
	}

	value, ok := v.(Value)
	if !ok {
		return nil
	}

	if value.IsDeleted {
		return nil
	}

	return &value
}

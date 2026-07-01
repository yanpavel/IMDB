// Echo2 выводит аргументы командной строки
package main

import (
	"fmt"

	"github.com/yanpavel/imdb/bloom"
	"github.com/yanpavel/imdb/storage"
)

func main() {
	// repo := bloom.New(1000) // нужно вычислять длину массива
	// data := []byte("1")
	// repo.Add(data)
	// // Добавить количество хэш-функций(2)
	// data2 := []byte("1")
	// value := repo.Contains(data2)
	// fmt.Print(value)
	bloomFilter := bloom.New(10)
	r := storage.NewRepository(*bloomFilter)
	k := storage.Key{Id: 1500}
	v := storage.Value{Name: "Joe"}
	r.Add(k, v)
	d := r.BloomFilter.Marshal()
	r.BloomFilter.Unmarshal(d)
	fmt.Print(r.BloomFilter)
}

package bloom

import (
	"encoding/binary"
	"encoding/json"
	"hash/adler32"
	"hash/fnv"
	"log"
	"math"
)

type BloomFilter struct {
	len        uint
	bitMap     []bool
	hashAmount uint
}

func New(len uint) *BloomFilter {
	m := calculateBitMapLen(len)
	k := uint(math.Round((float64(m) / float64(len)) * math.Log(2)))
	return &BloomFilter{
		len:        len,
		bitMap:     make([]bool, m),
		hashAmount: k,
	}
}

func (b *BloomFilter) Marshal() []byte {
	data, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (b *BloomFilter) Unmarshal(data []byte) {
	binary.Decode(data, binary.LittleEndian, b)
}

func (b *BloomFilter) Add(data []byte) {
	keys := b.calculatePosition(data)
	for _, key := range keys {
		b.bitMap[key] = true
	}
}

func (b *BloomFilter) Contains(data []byte) bool {
	keys := b.calculatePosition(data)
	for _, key := range keys {
		if b.bitMap[key] == false {
			return false
		}
	}
	return true
}

func calculateBitMapLen(len uint) uint {
	p := 0.01
	m := -(float64(len) * math.Log(p)) / math.Sqrt(math.Log(2))
	return uint(m)
}

func (b *BloomFilter) calculatePosition(data []byte) []uint {
	hash1 := fnv.New32a()
	hash2 := adler32.New()
	positions := make([]uint, b.hashAmount)
	var pos uint

	hash1.Write(data)
	hash2.Write(data)

	hashA := hash1.Sum32()
	hashB := hash2.Sum32()

	for i := range b.hashAmount {
		pos = (uint(hashA) * uint(hashB+2)) % uint(len(b.bitMap))
		positions[i] = pos
	}

	return positions
}

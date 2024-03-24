package bloomfilter

import "sync"

type Filter struct {
	bitMap  []uint8
	bitSize int64
	rw      sync.RWMutex
}

const numHashFunctions = 5

func New(bitSize int64) *Filter {
	filter := Filter{
		bitSize: bitSize,
		bitMap:  make([]uint8, bitSize),
	}

	return &filter
}

func (filter *Filter) Put(v interface{}) {
	/*
		Quoted from the book: Advanced Algorithms and Data Structures

		Ideally, we would need k different independent hash functions, so that no two indices
		are duplicated for the same value. It is not easy to design a large number of independent
		hash functions, but we can get good approximations. There are a few solutions
		commonly used:
		-  	Use a parametric hash function H(i). This meta-function, which is a generator
			of hash functions, takes in as input an initial value i and outputs a hash function
			Hi=H(i). During the Bloom filter’s initialization, we can create k of this
			functions, H0 to Hk-1, by calling the generator H on k different (and usually random)
			https://github.com/csorchard/SimpleBloomFilter/blob/74e19e12b3efb8b6627157097881c054a75c26b2/bloom_filter/BloomFilter.go#L35-L39
		- 	Use a single hash H function but initialize a list L of k random (and unique)
			values. For each entry key that is inserted/searched, create k values by adding
			or appending L[i] to key, and then hash them using H. [Done here]
		-   Use double or triple hashing. h(x), h(h(x)), h(h(h(x))), etc., are all different

	*/
	var hash1 = hash(v)
	var hash2 = hash1 << 32

	filter.rw.Lock()
	defer filter.rw.Unlock()

	for i := 1; i < numHashFunctions; i++ {
		nextHash := hash1 + uint64(i)*hash2
		if nextHash < 0 {
			nextHash = -nextHash
		}

		// should be in the range of 0 to filter.bitSize
		index := nextHash % uint64(filter.bitSize)
		filter.bitMap[index] = 1
	}
}

func (filter *Filter) MightContains(v interface{}) bool {
	var hash1 = hash(v)
	var hash2 = hash1 << 32

	filter.rw.RLock()
	defer filter.rw.RUnlock()

	for i := 1; i < numHashFunctions; i++ {
		nextHash := hash1 + uint64(i)*hash2
		if nextHash < 0 {
			nextHash = -nextHash
		}

		index := nextHash % uint64(filter.bitSize)
		if filter.bitMap[index] == 0 {
			return false
		}
	}

	return true
}

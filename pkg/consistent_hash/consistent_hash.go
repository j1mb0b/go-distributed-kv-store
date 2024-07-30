// Implement consistent hashing to distribute keys across nodes.
// Create resiliency by distributing the keys and effciency of the hashing system.
// Lets distribute the data across a set of nodes.

// Idea is depending on the data size we can distribute (buffer the data)
// Accross our hash ring of nodes and scale them.

package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32 // Store only positive numbers

// Map type.
type Map struct {
	hash Hash
	replicas int
	keys []int
	hashMap map[int]string
	mu sync.RWMutex // Allow multiple readers to access the shared resource simultaneously
}

// Initializer of Map
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash: fn,
		hashMap: make(map[int]string),
	}
	if m.hash == nil {
		// Generate a 32-bit hash.
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Modify Map and add varible string of keys to hashMap.
func (m *Map) Add(keys ...string) {
	// Prevent concurrent modifications
	// to avoid race conditions, implement the write lock.
	m.mu.Lock()
	defer m.mu.Unlock()
	// hashMap should be assigned 
	for _, key := range keys {
		// Generate multiple hash values for each key to distribute the load.
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Map given key to a specific node
func (m *Map) Get (key string) string {
	// Implement the read lock.
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := len(m.keys)
	if keys == 0 {
		return ""
	}

	// Compute the has value for given key.
	// Convert to int to locate the key on the hash ring.
	hash := int(m.hash([]byte(key)))

	// Binary search.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	// Handle index out of range.
	if idx == len(m.keys) {
		idx = 0
	}
	return m.hashMap[m.keys[idx]]
}
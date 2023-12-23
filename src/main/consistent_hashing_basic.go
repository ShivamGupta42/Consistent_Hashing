package main

import (
	"errors"
	"hash/crc32"
	"sort"
	"sync"
)

// Node represents a node with a host string and a hash ID.
type Node struct {
	Host   string
	HashID uint32
}

// GenHash generates a CRC32 hash for a given string ID.
func GenHash(ID string) uint32 {
	return crc32.ChecksumIEEE([]byte(ID))
}

// NewNode creates a new Node with a host and a calculated hash ID.
func NewNode(Host string) *Node {
	return &Node{Host: Host, HashID: GenHash(Host)}
}

// Nodes is a slice of Node pointers.
type Nodes []*Node

// Implementation of the sort.Interface for Nodes.
func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].HashID < n[j].HashID }

// Ring represents a collection of Nodes with mutex for concurrent access.
type Ring struct {
	sync.Mutex
	Nodes Nodes
}

// NewRing creates a new Ring with an empty Nodes slice.
func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

// addNode adds a new Node to the Ring if it does not exist and sorts the Nodes.
func (r *Ring) addNode(Host string) {
	r.Lock()
	defer r.Unlock()

	// Check if node exists before adding.
	for _, node := range r.Nodes {
		if node.Host == Host {
			return
		}
	}

	r.Nodes = append(r.Nodes, NewNode(Host))
	sort.Sort(r.Nodes)
}

// searchInsertion finds the position for a new node or an existing node.
func (r *Ring) searchInsertion(Host string) int {
	searchFunc := func(i int) bool {
		return r.Nodes[i].HashID >= GenHash(Host)
	}

	return sort.Search(r.Nodes.Len(), searchFunc)
}

// removeNode removes a Node from the Ring, returning an error if not found.
func (r *Ring) removeNode(Host string) error {
	r.Lock()
	defer r.Unlock()

	i := r.searchInsertion(Host)

	// Check if node exists before removal.
	if i >= r.Nodes.Len() || r.Nodes[i].Host != Host {
		return errors.New("Node not found")
	}

	r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)
	return nil
}

// Get returns the host of the node responsible for a given key.
func (r *Ring) Get(Key string) string {

	searchFunc := func(i int) bool {
		return r.Nodes[i].HashID >= crc32.ChecksumIEEE([]byte(Key))
	}

	i := sort.Search(r.Nodes.Len(), searchFunc)

	// Wrap around if necessary.
	if i >= r.Nodes.Len() {
		i = 0
	}

	return r.Nodes[i].Host
}

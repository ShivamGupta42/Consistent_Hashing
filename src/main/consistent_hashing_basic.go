package main

import (
	"errors"
	"hash/crc32"
	"sort"
	"sync"
)

/* Node Operations */
type Node struct {
	Host   string
	HashID uint32
}

func GenHash(ID string) uint32 {
	return crc32.ChecksumIEEE([]byte(ID))
}

func NewNode(Host string) *Node {
	return &Node{Host: Host, HashID: GenHash(Host)}
}

/* Create a sortable Nodes Array */
type Nodes []*Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].HashID < n[j].HashID }

/* Ring Operations  */
type Ring struct {
	sync.Mutex
	Nodes Nodes
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

func (r *Ring) addNode(Host string) {
	r.Lock()
	defer r.Unlock()

	// If node already exists then no-op
	for _, node := range r.Nodes {
		if node.Host == Host {
			return
		}
	}

	r.Nodes = append(r.Nodes, NewNode(Host))
	sort.Sort(r.Nodes)
}

func (r *Ring) search(Host string) int {
	searchFunc := func(i int) bool {
		return r.Nodes[i].HashID >= GenHash(Host)
	}

	return sort.Search(r.Nodes.Len(), searchFunc)
}

func (r *Ring) removeNode(Host string) error {
	r.Lock()
	defer r.Unlock()

	i := r.search(Host)

	if i >= r.Nodes.Len() || r.Nodes[i].Host != Host {
		return errors.New("Node not found")
	}

	r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)
	return nil
}

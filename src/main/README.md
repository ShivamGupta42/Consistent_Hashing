# Consistent Hashing Ring Implementation in Go

This repository contains a Go implementation of a consistent hashing ring. 
It's a simplistic example to distribute data across a cluster of nodes

## Features

- **Node Management:** Add and remove nodes dynamically from the hash ring.
- **Consistent Hashing:** Uses CRC32 hashing to distribute nodes evenly.
- **Concurrency Safe:** Thread-safe operations using mutex locks.
- **Simple Key Lookup:** Retrieve the host responsible for a given key.

## Code Structure

- `Node`: Represents a node in the ring with a host and a hash ID.
- `Nodes`: A slice of `Node` pointers, implementing the sort interface.
- `Ring`: Manages the collection of nodes with concurrency-safe operations.

## Functions

- `GenHash(ID string) uint32`: Generates a CRC32 hash for a given string ID.
- `NewNode(Host string) *Node`: Creates and returns a new Node.
- `NewRing() *Ring`: Creates a new empty Ring.
- `addNode(Host string)`: Adds a new Node to the Ring.
- `removeNode(Host string) error`: Removes a Node from the Ring.
- `Get(Key string) string`: Retrieves the host responsible for a given key.

## Usage

To use this consistent hashing implementation, follow these combined steps:

```go
// Step 1: Create a New Ring
ring := NewRing()

// Step 2: Adding Nodes to the Ring
ring.addNode("192.168.1.1")
ring.addNode("192.168.1.2")

// Step 3: Retrieving a Node for a Given Key
key := "someKey"
host := ring.Get(key)
fmt.Println("Host for key", key, "is", host)

// Step 4: Removing a Node from the Ring
err := ring.removeNode("192.168.1.1")
if err != nil {
    fmt.Println("Error removing node:", err)
} else {
    fmt.Println("Node removed successfully")
}
```


## Problems

-  Key LookUp time is O(LogN) where N is the number of servers. Can it be O(1)?
-  No Function to store keys on different nodes and replicate them for fault-tolerance
-  A large number of keys can still be mapped to a single server (Hot shard Problem)
-  No way to allocate more keys to a server with higher capacity
-  No way of dynamically re-allocating keys in case a server experiences higher load
-  Key remapping is not explictly solved in case of addition or removal of a node





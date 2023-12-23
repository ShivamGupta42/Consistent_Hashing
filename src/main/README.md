# Basic Consistent Hashing Implementation 

This repository contains a Go implementation of a consistent hashing ring. It uses CRC32 for hashing and provides basic functionalities to add, search, and remove nodes in the ring.

## Features

- **CRC32 Hashing**: Uses CRC32 algorithm for generating node hashes.
- **Concurrency Safe**: Thread-safe operations using mutex locks.
- **Sortable Nodes**: Nodes in the ring are sortable based on their hash values.

## Structure

- `Node`: Represents a node in the ring.
- `Ring`: Represents the consistent hashing ring.

## Usage

### Node Operations

- **Creating a New Node**
  - `NewNode(host string) *Node`: Creates a new `Node` with the given host identifier.

### Ring Operations

- **Creating a New Ring**
  - `NewRing() *Ring`: Initializes a new empty ring.

- **Adding a Node to the Ring**
  - `(*Ring) addNode(host string)`: Adds a new node to the ring with the given host identifier.

- **Searching for a Node**
  - `(*Ring) search(host string) int`: Searches for a node in the ring and returns its index.

- **Removing a Node from the Ring**
  - `(*Ring) removeNode(host string) error`: Removes a node from the ring by its host identifier.

## Example

```go
package main

func main() {
    ring := NewRing()

    // Add nodes
    ring.addNode("192.168.1.1")
    ring.addNode("192.168.1.2")

    // Remove a node
    err := ring.removeNode("192.168.1.1")
    if err != nil {
        // handle error
    }

    // Search for a node
    index := ring.search("192.168.1.3")
    // index is the position where the node would be if it were in the ring
}
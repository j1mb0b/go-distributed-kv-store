# NOTE

This is a learning exercise and is under development.

# Consistent Hashing with Replication

This project implements a consistent hashing ring with a specified replication factor. The goal is to distribute data items across nodes in a distributed system in a predictable and balanced manner, ensuring high availability and load balancing.

## Overview

This project utilizes consistent hashing to distribute data across a set of nodes. Each data item is replicated according to a defined replication factor, which ensures fault tolerance and balanced load distribution.

## Features

- **Consistent Hashing:** Distributes data items across nodes using a hash ring.
- **Replication:** Replicates each data item to a specified number of nodes for fault tolerance.
- **Deterministic Assignment:** Places replicas on nodes in a predictable manner based on hash values.

## Usage

### Initializing the Hash Ring

To initialize a consistent hashing ring with a replication factor, use the following code snippet:

```go
package main

import (
    "fmt"
    "github.com/stathat/consistent"
)

func main() {
    // Create a new consistent hash ring with a replication factor of 3
    hashRing := consistent.New(3, nil)

    // Add nodes to the hash ring
    hashRing.Add("Node1")
    hashRing.Add("Node2")
    hashRing.Add("Node3")
    hashRing.Add("Node4")

    // Hash a key to determine its position in the ring
    key := "dataItem"
    primaryNode, replicas := hashRing.Get(key)

    // Print results
    fmt.Printf("Primary Node: %s\n", primaryNode)
    fmt.Printf("Replicas: %v\n", replicas)
}
```

## TODO
- Add Error Handling: Implement error handling for cases where nodes are not added properly or data  items cannot be hashed.
- Improve Documentation: Provide detailed documentation on the algorithms used and how to extend the project.
- Test Cases: Write unit tests to verify the correctness of the hashing and replication logic.
- Dynamic Node Management: Implement functionality to dynamically add or remove nodes from the hash ring and handle data rebalancing.
- Performance Optimization: Optimize the performance of the hashing algorithm and replica assignment, especially for large numbers of nodes and data items.
- Integration with Real-World Systems: Explore integration with real-world distributed systems and databases to test the implementation in practical scenarios.

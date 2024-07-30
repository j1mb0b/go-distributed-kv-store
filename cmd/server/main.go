package main

import (
	"encoding/json"
	"go-distributed-kv-store/pkg/consistenthash"
	"go-distributed-kv-store/pkg/node"
	"net/http"
	"sync"
)

var (
    hashRing *consistenthash.Map
    nodes    = make(map[string]*node.Node)
    mu       sync.RWMutex
)


func init() {
    // Initialize nodes and add them to the hash ring
    addNode("node1")
    addNode("node2")
    addNode("node3")
}

func addNode(id string) {
    n := node.NewNode(id)
    mu.Lock()
    nodes[id] = n
    mu.Unlock()
    hashRing.Add(id)
}

func main() {
    // Here we construct the hash ring with a replication factor.
    // We create a consistent hashing ring with a speciffied number of replicas. 
    // The deterministic nature of the replica assignment depends on the following:
    // a. The hash key assignment is determined based on the hashed value of the data 
    //    item and the node positions on the ring.
    // b. The placement of replicas is determined by the position of the hash key on the ring and 
    //    the order of nodes following this position.
    // c. The hash ring is initialized with a replication factor of 3, 
    //    meaning each data item will have 3 replicas across nodes.
    
    hashRing = consistenthash.New(3, nil)
    http.HandleFunc("/put", putHandler)
    http.HandleFunc("/get", getHandler)
    http.ListenAndServe(":8080", nil)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    value := r.URL.Query().Get("value")
    if key == "" || value == "" {
        http.Error(w, "missing key or value", http.StatusBadRequest)
        return
    }

    nodeID := hashRing.Get(key)
    mu.RLock()
    node := nodes[nodeID]
    mu.RUnlock()
    node.Put(key, value)

    w.WriteHeader(http.StatusOK)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    if key == "" {
        http.Error(w, "missing key", http.StatusBadRequest)
        return
    }

    nodeID := hashRing.Get(key)
    mu.RLock()
    node := nodes[nodeID]
    mu.RUnlock()
    value, ok := node.Get(key)
    if !ok {
        http.Error(w, "key not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"value": value})
}
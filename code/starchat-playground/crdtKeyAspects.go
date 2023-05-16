// Write a Go program that demonstrates the key aspects of a CRDT

// Here’s an example implementation of a Conflict-Free Replicated Data Type
// (CRDT) using Golang. This particular implementation uses Operational
// Transformation (OT), where operations are represented as JSON objects with
// specific fields representing different types of edits such as “set”,
// “delete” or “increment”. The data type also includes methods for merging
// concurrent changes made by multiple replicas.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Define the datatype
type Counter struct {
	Value int                      `json:"value"` // Holds the counter value
	Ops   []map[string]interface{} `json:"ops"`   // List of all operations performed on the counter
}

// Implement functions to apply operations
func (c *Counter) ApplyOp(op map[string]interface{}) error {
	switch op["type"] {
	case "set":
		c.Value = int(op["value"].(float64))
	case "inc":
		c.Value += int(op["amount"].(float64))
	default:
		return errors.New("Invalid operation")
	}
	c.Ops = append(c.Ops, op)
	return nil
}

// Merge function used to merge concurrent changes from different replicas
func (c *Counter) Merge(other *Counter) bool {
	merged := false
	for _, op := range other.Ops {
		err := c.ApplyOp(op)
		if err != nil {
			log.Fatal(err)
		} else {
			merged = true
		}
	}
	return merged
}

// Example usage
func main() {
	var c1, c2 Counter
	ops1 := []map[string]interface{}{
		{"type": "set", "value": 1},
		{"type": "inc", "amount": 2}}
	for _, op := range ops1 {
		_ = c1.ApplyOp(op)
	}
	jsonStr1, _ := json.MarshalIndent(&c1, "", " ")
	fmt.Printf("%v\n", string(jsonStr1))

	ops2 := []map[string]interface{}{
		{"type": "inc", "amount": 5},
		{"type": "set", "value": 10}}
	for _, op := range ops2 {
		_ = c2.ApplyOp(op)
	}
	jsonStr2, _ := json.MarshalIndent(&c2, "", " ")
}

// How does a CRDT synchronize concurrent writes?

// In the case of two or more users editing the same document, it is important
// to ensure that all changes are properly merged. This process is called
// synchronization and ensures that each user has an up-to-date view of the
// shared data without overwriting their own changes with those from others.
// There are several approaches for achieving this goal in distributed systems
// such as cloud computing environments. One approach is based on using
// conflict resolution algorithms like Operational Transformation (OT) or
// Conflict-free Replicated Data Types (CRDT). These techniques allow multiple
// clients to edit documents simultaneously while preserving consistency and
// resolving conflicts when necessary. Another approach uses peer-to-peer
// communication protocols like BitTorrent which allows nodes to exchange
// information about updates made by other peers. In either scenario, the key
// challenge lies in how to merge conflicting edits so that everyone remains
// consistent at any given point in time.

// What other example applications of Operational transformation in software exist?
//
// Here’s another example application of OT algorithm : https://github.com/ottypes/docs#example

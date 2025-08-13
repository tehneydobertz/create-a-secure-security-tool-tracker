package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

type Tracker struct {
	ID        string
	ToolName  string
	ToolType  string
	Last Seen time.Time
}

var trackers = make(map[string]Tracker)

func generateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}

func createTracker(toolName, toolType string) Tracker {
	id := generateID()
	tracker := Tracker{
		ID:       id,
		ToolName: toolName,
		ToolType: toolType,
		LastSeen: time.Now(),
	}
trackers[id] = tracker
	return tracker
}

func updateLastSeen(id string) {
 trackers[id].LastSeen = time.Now()
}

func main() {
	tracker1 := createTracker("Wireshark", "Network Analyzer")
	tracker2 := createTracker("JohnTheRipper", "Password Cracker")

	fmt.Printf("Tracker 1: %+v\n", tracker1)
	fmt.Printf("Tracker 2: %+v\n", tracker2)

	updateLastSeen(tracker1.ID)
	updateLastSeen(tracker2.ID)

	fmt.Printf("Updated Tracker 1: %+v\n", trackers[tracker1.ID])
	fmt.Printf("Updated Tracker 2: %+v\n", trackers[tracker2.ID])

	// Example of hashing a tracker's details for secure storage
	trackerHash := sha256.Sum256([]byte(tracker1.ToolName + tracker1.ToolType + tracker1.LastSeen.String()))
	fmt.Printf("Tracker 1 Hash: %x\n", trackerHash)
}
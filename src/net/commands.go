package net

import "github.com/Fantom-foundation/go-lachesis/src/poset"

// SyncRequest initiates a synchronization request
type SyncRequest struct {
	FromID uint64
	Known  map[uint64]int64
}

// SyncResponse is a response to a SyncRequest request
type SyncResponse struct {
	FromID    uint64
	SyncLimit bool
	Events    []poset.WireEvent
	Known     map[uint64]int64
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// EagerSyncRequest after an initial sync to quickly catch up
type EagerSyncRequest struct {
	FromID uint64
	Events []poset.WireEvent
}

// EagerSyncResponse response to an EagerSyncRequest
type EagerSyncResponse struct {
	FromID  uint64
	Success bool
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// FastForwardRequest request to start a fast forward catch up
type FastForwardRequest struct {
	FromID uint64
}

// FastForwardResponse response with the snapshot data for fast forward request
type FastForwardResponse struct {
	FromID   uint64
	Block    poset.Block
	Frame    poset.Frame
	Snapshot []byte
}

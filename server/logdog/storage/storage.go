// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package storage

import (
	"errors"

	"github.com/luci/luci-go/common/logdog/types"
)

var (
	// ErrExists returned if an attempt is made to overwrite an existing record.
	ErrExists = errors.New("storage: record exists")

	// ErrDoesNotExist returned if an attempt is made to read a record that
	// doesn't exist.
	ErrDoesNotExist = errors.New("storage: record does not exist")

	// ErrBadData is an error returned when the stored data is invalid.
	ErrBadData = errors.New("storage: bad data")
)

// PutRequest describes adding a single storage record to BigTable.
type PutRequest struct {
	// Path is the stream path to retrieve.
	Path types.StreamPath
	// Index is the entry's stream index.
	Index types.MessageIndex

	// Value is the contents of the cell to add.
	Value []byte
}

// GetRequest is a request to retrieve a series of LogEntry records.
type GetRequest struct {
	// Path is the stream path to retrieve.
	Path types.StreamPath
	// Index is the entry's stream index.
	Index types.MessageIndex

	// Limit is the maximum number of records to return before stopping iteration.
	// If zero, no maximum limit will be applied.
	//
	// The Storage instance may return fewer records than the supplied Limit as an
	// implementation detail.
	Limit int
}

// GetCallback is invoked for each record in the Get request. If it returns
// false, iteration should stop.
type GetCallback func(types.MessageIndex, []byte) bool

// Storage is an abstract LogDog storage implementation. Interfaces implementing
// this may be used to store and retrieve log records by the collection service
// layer.
//
// All of these methods must be synchronous and goroutine-safe.
//
// All methods may return errors.Transient errors if they encounter an error
// that may be transient.
type Storage interface {
	// Close shuts down this instance, releasing any allocated resources.
	Close()

	// Writes log record data to storage.
	//
	// If the data already exists, ErrExists will be returned.
	Put(*PutRequest) error

	// Get invokes a callback over a range of sequential LogEntry records.
	//
	// These log entries will be returned in order (e.g., seq(Rn) < seq(Rn+1)),
	// but, depending on ingest, may not be contiguous.
	//
	// The underlying Storage implementation may return fewer records than
	// requested based on availability or implementation details; consequently,
	// receiving fewer than requsted records does not necessarily mean that more
	// records are not available.
	//
	// Returns nil if retrieval executed successfully, ErrDoesNotExist if
	// the requested stream does not exist, and an error if an error occurred
	// during retrieval.
	Get(*GetRequest, GetCallback) error

	// Tail invokes a callback over a range of contiguous LogEntry records,
	// starting with the largest entry.
	//
	// The contiguous space is started from the specified Index. If a Limit is
	// supplied, the tail record will begin at most Limit records from the Index.
	//
	// Returns nil if retrieval executed successfully, ErrDoesNotExist if
	// the requested stream does not exist, and an error if an error occurred
	// during retrieval.
	Tail(*GetRequest, GetCallback) error

	// Purges a stream and all of its data from the store.
	//
	// If the requested stream doesn't exist, ErrDoesNotExist will be returned.
	Purge(types.StreamPath) error
}

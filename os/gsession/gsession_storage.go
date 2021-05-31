// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gsession

import "time"

type Storage interface {
	// Get retrieves session value with given key.
	// It returns nil if the key does not exist in the session.
	Get(key string) interface{}
	// GetMap retrieves all key-value pairs as map from storage.
	GetMap() map[string]interface{}
	// GetSize retrieves the size of key-value pairs from storage.
	GetSize(id string) int

	// Set sets key-value session pair to the storage.
	Set(key string, value interface{}) error
	// SetMap batch sets key-value session pairs with map to the storage.
	SetMap(data map[string]interface{}) error

	// Remove deletes key with its value from storage.
	Remove(key string) error
	// RemoveAll deletes all key-value pairs from storage.
	RemoveAll() error

	// GetSession returns the session data bytes for given session id.
	// The parameter specifies the TTL for this session.
	// It returns nil if the TTL is exceeded.
	GetSession(id string, ttl time.Duration) map[string]interface{}
	// SetSession updates the content for session id.
	// Note that the parameter <content> is the serialized bytes for session map.
	SetSession(id string, data map[string]interface{}) error

	// UpdateTTL updates the TTL for specified session id.
	UpdateTTL(id string) error
}

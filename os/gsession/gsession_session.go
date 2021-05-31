// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gsession

import (
	"time"

	"github.com/snail007/gf/container/gtype"

	"github.com/snail007/gf/container/gmap"
	"github.com/snail007/gf/container/gvar"
	"github.com/snail007/gf/os/gtime"
	"github.com/snail007/gf/util/gconv"
)

// Session struct for storing single session data.
type Session struct {
	id      string          // Session id.
	data    *gmap.StrAnyMap // Session data.
	dirty   *gtype.Bool     // Used to mark session is modified.
	manager *Manager        // Parent manager.
}

// init does the delay initialization for session.
// It here to initialization real session if necessary.
func (s *Session) init() {
	if s.dirty == nil {
		s.dirty = gtype.NewBool(false)
	}
	if len(s.id) > 0 && s.data == nil {
		if data := s.manager.storage.GetSession(s.id, s.manager.ttl); data != nil {
			if s.data = gmap.NewStrAnyMapFrom(data, true); s.data == nil {
				panic("session restoring failed for id:" + s.id)
			}
			return
		} else {
			// Invalid or expired session id,
			// it should create a new one.
			s.id = ""
		}
	}
	if len(s.id) == 0 {
		s.id = NewSessionId()
	}
	if s.data == nil {
		s.data = gmap.NewStrAnyMap(true)
	}
}

// Set sets key-value pair to this session.
func (s *Session) Set(key string, value interface{}) error {
	s.init()
	if err := s.manager.storage.Set(key, value); err != nil {
		if err == ErrorDisabled {
			s.data.Set(key, value)
		} else {
			return err
		}
	}
	s.dirty.Set(true)
	return nil
}

// Sets batch sets the session using map.
func (s *Session) Sets(data map[string]interface{}) error {
	s.init()
	if err := s.manager.storage.SetMap(data); err != nil {
		if err == ErrorDisabled {
			s.data.Sets(data)
		} else {
			return err
		}
	}
	s.dirty.Set(true)
	return nil
}

// Remove removes key along with its value from this session.
func (s *Session) Remove(key string) error {
	if len(s.id) == 0 {
		return nil
	}
	s.init()
	if err := s.manager.storage.Remove(key); err != nil {
		if err == ErrorDisabled {
			s.data.Remove(key)
		} else {
			return err
		}
	}
	s.dirty.Set(true)
	return nil
}

// Clear is alias of RemoveAll.
func (s *Session) Clear() error {
	return s.RemoveAll()
}

// RemoveAll deletes all key-value pairs from this session.
func (s *Session) RemoveAll() error {
	if len(s.id) == 0 {
		return nil
	}
	s.init()
	if err := s.manager.storage.RemoveAll(); err != nil {
		if err == ErrorDisabled {
			s.data.Clear()
		} else {
			return err
		}
	}
	s.dirty.Set(true)
	return nil
}

// Id returns the session id for this session.
// It create and returns a new session id if the session id is not passed in initialization.
func (s *Session) Id() string {
	s.init()
	return s.id
}

// Map returns all data as map.
// Note that it's using value copy internally for concurrent-safe purpose.
func (s *Session) Map() map[string]interface{} {
	if len(s.id) > 0 {
		s.init()
		if data := s.manager.storage.GetMap(); data != nil {
			return data
		}
		return s.data.Map()
	}
	return nil
}

// Size returns the size of the session.
func (s *Session) Size() int {
	if len(s.id) > 0 {
		s.init()
		if size := s.manager.storage.GetSize(s.id); size >= 0 {
			return size
		}
		return s.data.Size()
	}
	return 0
}

// Contains checks whether key exist in the session.
func (s *Session) Contains(key string) bool {
	s.init()
	return s.Get(key) != nil
}

// IsDirty checks whether there's any data changes in the session.
func (s *Session) IsDirty() bool {
	if s.dirty == nil {
		return false
	}
	return s.dirty.Val()
}

// Close closes current session and updates its ttl in the session manager.
// If this session is dirty, it also exports it to storage.
//
// NOTE that this function must be called ever after a session request done.
func (s *Session) Close() {
	if len(s.id) > 0 && s.data != nil {
		if s.manager.storage != nil {
			if s.dirty.Cas(true, false) {
				s.data.RLockFunc(func(m map[string]interface{}) {
					if err := s.manager.storage.SetSession(s.id, m); err != nil {
						panic(err)
					}
				})
			} else {
				if err := s.manager.storage.UpdateTTL(s.id); err != nil {
					panic(err)
				}
			}
		}
		s.manager.UpdateSessionTTL(s.id, s)
	}
}

// Get retrieves session value with given key.
// It returns <def> if the key does not exist in the session if <def> is given,
// or else it return nil.
func (s *Session) Get(key string, def ...interface{}) interface{} {
	if len(s.id) == 0 {
		return nil
	}
	s.init()
	if v := s.manager.storage.Get(key); v != nil {
		return v
	}
	if v := s.data.Get(key); v != nil {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

func (s *Session) GetVar(key string, def ...interface{}) *gvar.Var {
	return gvar.New(s.Get(key, def...), true)
}

func (s *Session) GetString(key string, def ...interface{}) string {
	return gconv.String(s.Get(key, def...))
}

func (s *Session) GetBool(key string, def ...interface{}) bool {
	return gconv.Bool(s.Get(key, def...))
}

func (s *Session) GetInt(key string, def ...interface{}) int {
	return gconv.Int(s.Get(key, def...))
}

func (s *Session) GetInt8(key string, def ...interface{}) int8 {
	return gconv.Int8(s.Get(key, def...))
}

func (s *Session) GetInt16(key string, def ...interface{}) int16 {
	return gconv.Int16(s.Get(key, def...))
}

func (s *Session) GetInt32(key string, def ...interface{}) int32 {
	return gconv.Int32(s.Get(key, def...))
}

func (s *Session) GetInt64(key string, def ...interface{}) int64 {
	return gconv.Int64(s.Get(key, def...))
}

func (s *Session) GetUint(key string, def ...interface{}) uint {
	return gconv.Uint(s.Get(key, def...))
}

func (s *Session) GetUint8(key string, def ...interface{}) uint8 {
	return gconv.Uint8(s.Get(key, def...))
}

func (s *Session) GetUint16(key string, def ...interface{}) uint16 {
	return gconv.Uint16(s.Get(key, def...))
}

func (s *Session) GetUint32(key string, def ...interface{}) uint32 {
	return gconv.Uint32(s.Get(key, def...))
}

func (s *Session) GetUint64(key string, def ...interface{}) uint64 {
	return gconv.Uint64(s.Get(key, def...))
}

func (s *Session) GetFloat32(key string, def ...interface{}) float32 {
	return gconv.Float32(s.Get(key, def...))
}

func (s *Session) GetFloat64(key string, def ...interface{}) float64 {
	return gconv.Float64(s.Get(key, def...))
}

func (s *Session) GetBytes(key string, def ...interface{}) []byte {
	return gconv.Bytes(s.Get(key, def...))
}

func (s *Session) GetInts(key string, def ...interface{}) []int {
	return gconv.Ints(s.Get(key, def...))
}

func (s *Session) GetFloats(key string, def ...interface{}) []float64 {
	return gconv.Floats(s.Get(key, def...))
}

func (s *Session) GetStrings(key string, def ...interface{}) []string {
	return gconv.Strings(s.Get(key, def...))
}

func (s *Session) GetInterfaces(key string, def ...interface{}) []interface{} {
	return gconv.Interfaces(s.Get(key, def...))
}

func (s *Session) GetTime(key string, format ...string) time.Time {
	return gconv.Time(s.Get(key), format...)
}

func (s *Session) GetGTime(key string, format ...string) *gtime.Time {
	return gconv.GTime(s.Get(key), format...)
}

func (s *Session) GetDuration(key string, def ...interface{}) time.Duration {
	return gconv.Duration(s.Get(key, def...))
}

func (s *Session) GetMap(key string, tags ...string) map[string]interface{} {
	return gconv.Map(s.Get(key), tags...)
}

func (s *Session) GetMapDeep(key string, tags ...string) map[string]interface{} {
	return gconv.MapDeep(s.Get(key), tags...)
}

func (s *Session) GetMaps(key string, tags ...string) []map[string]interface{} {
	return gconv.Maps(s.Get(key), tags...)
}

func (s *Session) GetMapsDeep(key string, tags ...string) []map[string]interface{} {
	return gconv.MapsDeep(s.Get(key), tags...)
}

func (s *Session) GetStruct(key string, pointer interface{}, mapping ...map[string]string) error {
	return gconv.Struct(s.Get(key), pointer, mapping...)
}

func (s *Session) GetStructDeep(key string, pointer interface{}, mapping ...map[string]string) error {
	return gconv.StructDeep(s.Get(key), pointer, mapping...)
}

func (s *Session) GetStructs(key string, pointer interface{}, mapping ...map[string]string) error {
	return gconv.Structs(s.Get(key), pointer, mapping...)
}

func (s *Session) GetStructsDeep(key string, pointer interface{}, mapping ...map[string]string) error {
	return gconv.StructsDeep(s.Get(key), pointer, mapping...)
}

// (c) Copyright 2015 JONNALAGADDA Srinivas
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flagon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

// CompOp enumerates the possible comparison operators that can be
// used when searching entities.
type CompOp uint8

const (
	CompOpEquals CompOp = iota
	CompOpLessThan
	CompOpLessThanEquals
	CompOpGreaterThan
	CompOpGreaterThanEquals
	CompOpPrefix
	CompOpSuffix
	CompOpContains
)

// SearchFn accepts a (key, entity) tuple, and answers `true` if they
// satisfy the predicate.
type SearchFn func(uint64, Entity) bool

// SearchOpts are the possible options that application code can
// specify for a search.  Not all entity types may understand and
// honour all options.
type SearchOpts struct {
	// Operator to use for searching.
	Operator CompOp
	// Key where search should begin.  If not found, search begins at
	// the first key that is greater than the given key.
	StartAt uint64
	// Maximum number of results to answer; `0` for unlimited.
	Limit uint64
	// Fields to deserialise and make available to the predicate.
	// Judicious use of this could increase search efficiency,
	// particularly for large entities.  `nil` deserialises the entire
	// object, and is hence expensive.
	Fields []int
}

// EntityKey holds the globally-unique ID of an instance within its
// entity type.
type EntityKey struct {
	id uint64
}

// ID answers the globally-unique ID of an instance within its entity
// type.
func (k EntityKey) ID() uint64 {
	return k.id
}

// Key answers the globally-unique ID of an instance within its entity
// type in a serialisable form.
func (k EntityKey) Key() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	err := binary.Write(buf, binary.BigEndian, k.id)
	if err != nil {
		log.Printf("error writing key: %s", err)
		return nil
	}

	return buf.Bytes()
}

// fromKey reads and sets the ID of this key instance from the given
// serialised form of a key.
func (k *EntityKey) fromKey(by []byte) error {
	r := bytes.NewReader(by)
	err := binary.Read(r, binary.BigEndian, &k.id)
	if err != nil {
		return err
	}

	return nil
}

// Entity specifies the methods that data entities should implement to
// be recognised by the entity types defined as their resources.
type Entity interface {
	// ID answers a globally-unique identifier for this instance
	// within its entity type.
	ID() uint64
	// Key answers the globally-unique ID of an instance within its
	// entity type in a serialisable form.
	Key() []byte
	// TypeName answers the name of this instance's entity type.
	TypeName() string

	fmt.Stringer // for debugging
}

// EntityType specifies the methods that data entity types should
// implement to handle instances of the resources they represent.
type EntityType interface {
	// Name answers the user-defined name of this entity type.
	Name() string
	// Get looks up the table for the entity having the given ID, and
	// answers the same if found.
	Get(uint64) (Entity, error)
	// Put creates - or updates - the given entity in the table.
	Put(Entity) error
	// Delete removes the entity having the given ID from the table,
	// if found.
	Delete(uint64) error
	// Search iterates through the table, passing each (key, entity)
	// tuple to the provided predicate.
	Search(SearchOpts, SearchFn) ([]uint64, error)
}

// EntityTypeDefn captures the necessary information for defining and
// dealing with instances of specific entity types.

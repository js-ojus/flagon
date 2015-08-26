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
	"encoding/binary"
	"fmt"
	"io"
)

// FieldType enumerates the possible data types of fields that can be
// defined as parts of entity types.
type FieldType uint8

const (
	FieldTypeUnknown FieldType = iota
	FieldTypeBool
	FieldTypeInt8
	FieldTypeInt16
	FieldTypeInt32
	FieldTypeInt64
	FieldTypeUint8
	FieldTypeUint16
	FieldTypeUint32
	FieldTypeUint64
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeTime
	FieldTypeString
	FieldTypeBytes
	FieldTypeReference
	FieldTypeLink
	FieldTypeCollection
)

// IsValidFieldType answers `true` if the given type is recognised,
// `false` otherwise.
func IsValidFieldType(t FieldType) bool {
	switch t {
	case FieldTypeBool,
		FieldTypeInt8,
		FieldTypeInt16,
		FieldTypeInt32,
		FieldTypeInt64,
		FieldTypeUint8,
		FieldTypeUint16,
		FieldTypeUint32,
		FieldTypeUint64,
		FieldTypeFloat32,
		FieldTypeFloat64,
		FieldTypeTime,
		FieldTypeString,
		FieldTypeBytes,
		FieldTypeReference, // strong reference
		FieldTypeLink,      // weak reference
		FieldTypeCollection:
		return true
	default:
		return false
	}
}

// FieldDefn captures the necessary information for defining and
// dealing with fields and their data.
type FieldDefn struct {
	ftype FieldType // type of the data in this field
	id    int       // unique ID within its entity type
	name  string    // name of the field
}

// Field is the building block of an entity.  It is identified by the
// ID of its field definition, and stores the actual content of
// user-supplied data.
type Field interface {
	ID() int

	io.ReaderFrom
	io.WriterTo
}

// basicField defines the common core of all fields.
type basicField struct {
	id uint8
}

// ID answers the unique identifier of this field within its entity
// type definition.
func (f basicField) ID() int {
	return int(f.id)
}

// setID sets the unique identifier of this field within its entity
// type definition.
func (f *basicField) setID(id int) error {
	if id <= 0 {
		return fmt.Errorf("negative or zero ID: %d", id)
	}

	f.id = uint8(id)
	return nil
}

// FieldBool represents a boolean value.
type FieldBool struct {
	basicField
	value uint8
}

// Get answers this field's value.
func (f *FieldBool) Get() bool {
	if f.value > 0 {
		return true
	}
	return false
}

// Set sets the given value to this field's storage.
func (f *FieldBool) Set(v bool) {
	if v {
		f.value = 1
	} else {
		f.value = 0
	}
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldBool) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldBool) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// FieldInt8 represents a boolean value.
type FieldInt8 struct {
	basicField
	value int8
}

// Get answers this field's value.
func (f *FieldInt8) Get() int8 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldInt8) Set(v int8) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldInt8) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldInt8) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// FieldInt16 represents a boolean value.
type FieldInt16 struct {
	basicField
	value int16
}

// Get answers this field's value.
func (f *FieldInt16) Get() int16 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldInt16) Set(v int16) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldInt16) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 2, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldInt16) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 2, nil
}

// FieldInt32 represents a boolean value.
type FieldInt32 struct {
	basicField
	value int32
}

// Get answers this field's value.
func (f *FieldInt32) Get() int32 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldInt32) Set(v int32) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldInt32) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 4, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldInt32) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 4, nil
}

// FieldInt64 represents a boolean value.
type FieldInt64 struct {
	basicField
	value int64
}

// Get answers this field's value.
func (f *FieldInt64) Get() int64 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldInt64) Set(v int64) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldInt64) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 8, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldInt64) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 8, nil
}

// FieldUint8 represents a boolean value.
type FieldUint8 struct {
	basicField
	value uint8
}

// Get answers this field's value.
func (f *FieldUint8) Get() uint8 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldUint8) Set(v uint8) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldUint8) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldUint8) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// FieldUint16 represents a boolean value.
type FieldUint16 struct {
	basicField
	value uint16
}

// Get answers this field's value.
func (f *FieldUint16) Get() uint16 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldUint16) Set(v uint16) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldUint16) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 2, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldUint16) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 2, nil
}

// FieldUint32 represents a boolean value.
type FieldUint32 struct {
	basicField
	value uint32
}

// Get answers this field's value.
func (f *FieldUint32) Get() uint32 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldUint32) Set(v uint32) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldUint32) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 4, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldUint32) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 4, nil
}

// FieldUint64 represents a boolean value.
type FieldUint64 struct {
	basicField
	value uint64
}

// Get answers this field's value.
func (f *FieldUint64) Get() uint64 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldUint64) Set(v uint64) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldUint64) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 8, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldUint64) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 8, nil
}

// FieldFloat32 represents a boolean value.
type FieldFloat32 struct {
	basicField
	value float32
}

// Get answers this field's value.
func (f *FieldFloat32) Get() float32 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldFloat32) Set(v float32) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldFloat32) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 4, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldFloat32) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 4, nil
}

// FieldFloat64 represents a boolean value.
type FieldFloat64 struct {
	basicField
	value float64
}

// Get answers this field's value.
func (f *FieldFloat64) Get() float64 {
	return f.value
}

// Set sets the given value to this field's storage.
func (f *FieldFloat64) Set(v float64) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldFloat64) ReadFrom(r io.Reader) (int64, error) {
	err := binary.Read(r, binary.BigEndian, &f.value)
	if err != nil {
		return 0, err
	}

	return 8, nil
}

// WriteTo conforms to `io.WriteTo`.
func (f *FieldFloat64) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, f.value)
	if err != nil {
		return 0, err
	}

	return 8, nil
}

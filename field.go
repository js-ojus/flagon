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
	"log"
	"time"
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
//
// N.B. A field's unique ID is represented as `uint8`.  Hence, there
// is a limit of 255 fields per entity type.  In practice, much
// smaller entity types are recommended.
type FieldDefn struct {
	Ftype FieldType // type of the data in this field
	ID    uint8     // unique ID within its entity type
	Name  string    // name of the field
}

// Field is the building block of an entity.  It is identified by the
// ID of its field definition, and stores the actual content of
// user-supplied data.
type Field interface {
	ID() uint8

	io.ReaderFrom
	io.WriterTo
}

// basicField defines the common core of all fields.
type basicField struct {
	id uint8
}

// ID answers the unique identifier of this field within its entity
// type definition.
func (f basicField) ID() uint8 {
	return f.id
}

// setID sets the unique identifier of this field within its entity
// type definition.
func (f *basicField) setID(id uint8) error {
	if id == 0 {
		return fmt.Errorf("zero ID: %d", id)
	}

	f.id = id
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

// Set sets the given value in this field's storage.
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

// FieldInt8 represents an 8-bit integer value.
type FieldInt8 struct {
	basicField
	value int8
}

// Get answers this field's value.
func (f *FieldInt8) Get() int8 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldInt16 represents a 16-bit integer value.
type FieldInt16 struct {
	basicField
	value int16
}

// Get answers this field's value.
func (f *FieldInt16) Get() int16 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldInt32 represents a 32-bit integer value.
type FieldInt32 struct {
	basicField
	value int32
}

// Get answers this field's value.
func (f *FieldInt32) Get() int32 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldInt64 represents a 64-bit integer value.
type FieldInt64 struct {
	basicField
	value int64
}

// Get answers this field's value.
func (f *FieldInt64) Get() int64 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldUint8 represents an unsigned 8-bit integer value.
type FieldUint8 struct {
	basicField
	value uint8
}

// Get answers this field's value.
func (f *FieldUint8) Get() uint8 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldUint16 represents an unsigned 16-bit integer value.
type FieldUint16 struct {
	basicField
	value uint16
}

// Get answers this field's value.
func (f *FieldUint16) Get() uint16 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldUint32 represents an unsigned 32-bit integer value.
type FieldUint32 struct {
	basicField
	value uint32
}

// Get answers this field's value.
func (f *FieldUint32) Get() uint32 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldUint64 represents an unsigned 64-bit integer value.
type FieldUint64 struct {
	basicField
	value uint64
}

// Get answers this field's value.
func (f *FieldUint64) Get() uint64 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldFloat32 represents a 32-bit float value.
type FieldFloat32 struct {
	basicField
	value float32
}

// Get answers this field's value.
func (f *FieldFloat32) Get() float32 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldFloat64 represents a 64-bit float value.
type FieldFloat64 struct {
	basicField
	value float64
}

// Get answers this field's value.
func (f *FieldFloat64) Get() float64 {
	return f.value
}

// Set sets the given value in this field's storage.
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

// FieldTime represents a time value.
//
// N.B. Time values are converted to UTC before serialisation, to
// enable standardised search and comparison.  Hence, applications
// should adjust them to the desired time zones before use.
type FieldTime struct {
	basicField
	value time.Time
}

// Get answers this field's value.
func (f *FieldTime) Get() time.Time {
	return f.value
}

// Set sets the given value in this field's storage.
func (f *FieldTime) Set(v time.Time) {
	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldTime) ReadFrom(r io.Reader) (int64, error) {
	by := make([]byte, 0, 15)
	n, err := r.Read(by)
	if n != 15 && err != nil {
		return int64(n), err
	}

	err = f.value.UnmarshalBinary(by)
	if err != nil {
		return 0, err
	}

	return 15, nil
}

// WriteTo conforms to `io.WriterTo`.
func (f *FieldTime) WriteTo(w io.Writer) (int64, error) {
	by, err := f.value.UTC().MarshalBinary()
	if err != nil {
		return 0, err
	}

	n, err := w.Write(by)
	if n != 15 && err != nil {
		return int64(n), err
	}

	return 15, nil
}

// FieldString represents a string value.
//
// N.B. The length of a string field is represented as `uint16`, and
// is hence limited to 65535 bytes.  Trying to set string values
// larger than that will be ignored, but will be logged.  Use files
// for larger strings.
type FieldString struct {
	basicField
	value string
}

// Get answers this field's value.
func (f *FieldString) Get() string {
	return f.value
}

// Set sets the given value in this field's storage.
func (f *FieldString) Set(v string) {
	if len(v) > 65535 {
		log.Printf("string length exceeds maximum limit of 65535")
		return
	}

	f.value = v
}

// ReadFrom conforms to `io.ReaderFrom`.
func (f *FieldString) ReadFrom(r io.Reader) (int64, error) {
	l := uint16(len(f.value))
	err := binary.Read(r, binary.BigEndian, &l)
	if err != nil {
		return 0, err
	}
	if l == 0 {
		return 0, io.EOF
	}

	by := make([]byte, 0, l)
	n, err := r.Read(by)
	if err != nil || n != int(l) {
		return int64(n), err
	}

	f.value = string(by)
	return int64(n), nil
}

// WriteTo conforms to `io.WriterTo`.
func (f *FieldString) WriteTo(w io.Writer) (int64, error) {
	if f.value == "" {
		return 0, nil
	}

	by := []byte(f.value)
	l := uint16(len(by))
	err := binary.Write(w, binary.BigEndian, l)
	if err != nil {
		return 0, err
	}

	n, err := w.Write(by)
	if err != nil || n != int(l) {
		return int64(n), err
	}

	return int64(n), nil
}

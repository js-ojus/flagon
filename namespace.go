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
	"regexp"
	"sync"
)

// nsNameRegexp holds the compiled regular expression that validates
// namespace names.
var nameRegexp *regexp.Regexp

// Namespace provides a logical grouping of related data.
//
// Similar data that needs to be grouped differently can use a
// different namespace.  Namespaces enable isolation of data, as well
// as help in improving search performance by limiting the amount of
// data to be traversed.
type Namespace struct {
	name string // application-visible name

	mutex   sync.RWMutex
	buckets []string // buckets in this namespace
}

// NewNamespace creates and registers a namespace with `flagon`.
//
// Namespace names must begin with an ASCII letter or a digit, can
// contain ASCII letters, digits, hyphens and underscores, and must
// not end with a hyphen or an underscore.  Names must be of length 2
// or more.  Names must be lowercase.
func NewNamespace(name string) (*Namespace, error) {
	if name == "" {
		return nil, ErrNameEmpty
	}
	if !nameRegexp.MatchString(name) {
		return nil, ErrNameInvalid
	}

	return &Namespace{name: name, buckets: make([]string, 0, 1)}, nil
}

// Name answers the name of this namespace.
func (ns *Namespace) Name() string {
	return ns.name
}

// Buckets answers a copy of the names of the buckets in this
// namespace.
func (ns *Namespace) Buckets() []string {
	ns.mutex.RLock()
	defer ns.mutex.RUnlock()

	bs := make([]string, len(ns.buckets))
	copy(bs, ns.buckets)
	return bs
}

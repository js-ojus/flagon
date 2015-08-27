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

package storage

import "errors"

var (
	// ErrPathEmpty is answered when a non-empty path was expected but
	// an empty path was given.
	ErrPathEmpty = errors.New("empty path given")

	// ErrPathNotAbsolute is answered when an absolute path was
	// expected but a relative path was given.
	ErrPathNotAbsolute = errors.New("given path is not an absolute one")
)

var (
	// ErrNameEmpty is answered when an unexpected empty name is
	// provided.
	ErrNameEmpty = errors.New("empty name given")
)

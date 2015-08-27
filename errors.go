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
	"errors"
	"fmt"
)

var (
	// ErrNameEmpty is answered when an unexpected empty name is
	// provided.
	ErrNameEmpty = errors.New("empty name given")

	// ErrNameInvalid is answered when the given name does not conform
	// to naming rules of `flagon`.
	ErrNameInvalid = fmt.Errorf("name should conform to `%s`", nameRegexp.String())

	// ErrNameExists is answered when a unique name was expected, but
	// an existing name was provided.
	ErrNameExists = errors.New("name already exists")

	// ErrNameUnknown is answered when an existing name was expected,
	// but an unknown name was provided.
	ErrNameUnknown = errors.New("unknown name given")
)

var (
	// ErrFieldTypeUnknown is answered when an unrecognised field type
	// is specified when defining a field.
	ErrFieldTypeUnknown = errors.New("unknown field type specified")
)

var (
	// ErrIdentifierZero is answered when an ID was expected, but a
	// zero value was provided.
	ErrIdentifierZero = errors.New("zero ID value given")
)

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
	"log"
	"regexp"
)

func init() {
	// Set log format.
	f := log.Flags()
	log.SetFlags(f | log.Llongfile)

	// Initialisations.
	nsNameRegexp = regexp.MustCompile("^[a-z0-9][a-z0-9\\-_]*[a-z0-9]$")
}

<!--
   (c) Copyright 2015 JONNALAGADDA Srinivas

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
-->

[![Build Status](https://travis-ci.org/js-ojus/flagon.svg?branch=master)](https://travis-ci.org/js-ojus/flagon)

## flagon
A small document store (built on BoltDB) written in Go (golang).

`flagon` is a good fit for semi-structured data that requires ACID compliance across documents.  It provides a thin structure to enable easy storage and retrieval.

`flagon` automatically maintains all old revisions of all user documents.  Each document revision is stamped with the ID of the user creating or modifying it, date-time, _etc_.  This metadata can be used for audit purposes.

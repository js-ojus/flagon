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

import (
	"os"
	"path"
	"sync"

	"github.com/boltdb/bolt"
)

const (
	// Database directory inside the storage path.
	dbdir = "db"

	// Database file name inside the database directory.
	dbname = "flagon.db"

	// System catalogue bucket name inside the database.
	dbsysname = "_sys"

	// System catalogue namespace definitions bucket name.
	dbnsdefsname = "nsdefs"

	// System catalogue entity type definitions bucket name.
	dbetdefsname = "etdefs"
)

var (
	// The database singleton instance protector.
	onceDB sync.Once

	// The database singleton instance.
	theDB DB
)

// Database initialisation error, if any.
var dberr error

// DB represents a BoltDB database.  All of `flagon` uses a single
// database.
//
// Internally, each entity type has its own bucket per namespace in
// which its instances have to be stored.
type DB struct {
	db *bolt.DB // handle to the underlying BoltDB database
}

// InitDB creates and initialises a database inside the given base
// storage directory path.  This path should be an absolute path.
func InitDB(p string) error {
	if p == "" {
		return ErrPathEmpty
	}
	if !path.IsAbs(p) {
		return ErrPathNotAbsolute
	}

	storageDir = p
	initDB()

	return dberr
}

func initDB() {
	dberr = nil

	// Create the directories.
	fn := path.Join(storageDir, dbdir)
	err := os.MkdirAll(fn, 0700)
	if err != nil {
		dberr = err
		return
	}

	// Create the BoltDB database.
	fn = path.Join(fn, dbname)
	theDB.db, dberr = bolt.Open(fn, 0600, nil)
	if dberr != nil {
		return
	}
	defer theDB.db.Close()

	// Set up `flagon` system catalogues.
	dberr = theDB.db.Update(func(tx *bolt.Tx) error {
		var err error
		sys, err := tx.CreateBucketIfNotExists([]byte(dbsysname))
		if err != nil {
			return err
		}
		_, err = sys.CreateBucketIfNotExists([]byte(dbnsdefsname))
		if err != nil {
			return err
		}
		_, err = sys.CreateBucketIfNotExists([]byte(dbetdefsname))
		if err != nil {
			return err
		}

		return nil
	})
}

// DbInstance opens the underlying BoltDB database, and answers the
// singleton DB instance.
func DbInstance() (*DB, error) {
	onceDB.Do(instance)
	if dberr != nil {
		return nil, dberr
	}

	return &theDB, nil
}

func instance() {
	fn := path.Join(storageDir, dbdir, dbname)
	theDB.db, dberr = bolt.Open(fn, 0600, nil)
	if dberr != nil {
		return
	}
}

// Close closes the underlying BoltDB database.
func (db *DB) Close() error {
	return theDB.db.Close()
}

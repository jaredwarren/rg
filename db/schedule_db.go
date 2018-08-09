package db

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	rg "github.com/jaredwarren/rg/gen/schedule"
)

// ScheduleStore ...
type ScheduleStore interface {
	New(bucket string) (string, error)
	FetchAll(bucket string) (data []*rg.Schedule, err error)
	Fetch(bucket, id string) (data *rg.Schedule, err error)
	Save(bucket, id string, schedule *rg.Schedule) error
	Delete(bucket, id string) error
}

// ErrNotFound is the error returned when attempting to load a record that does
// not exist.
var ErrNotFound = fmt.Errorf("missing record")

// ScheduleDB is the database driver.
type ScheduleDB struct {
	// client is the ScheduleDB client.
	client *bolt.DB
}

// NewScheduleDB creates a ScheduleDB DB database driver given an underlying client.
func NewScheduleDB(client *bolt.DB) (ScheduleStore, error) {
	err := client.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("SCHEDULE"))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("SETTINGS"))
		return err
	})
	if err != nil {
		return nil, err
	}
	return &ScheduleDB{client}, nil
}

// New returns a unique ID for the given bucket.
func (b *ScheduleDB) New(bucket string) (string, error) {
	var sid string
	err := b.client.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		id, err := bkt.NextSequence()
		if err != nil {
			return err
		}
		sid = fmt.Sprintf("%s%s", "schedule_", strconv.FormatUint(id, 10))
		return nil
	})
	return sid, err
}

// Save writes the record to the DB and returns the corresponding new ID.
// data must contain a value that can be marshaled by the encoding/json package.
func (b *ScheduleDB) Save(bucket, id string, schedule *rg.Schedule) error {
	buf, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	return b.client.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		return bkt.Put([]byte(id), buf)
	})
}

// Update writes the record to the DB and returns the corresponding new ID.
func (b *ScheduleDB) Update(bucket, id string, schedule *rg.Schedule) error {
	buf, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	return b.client.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		return bkt.Put([]byte(id), buf)
	})
}

// Delete deletes a record by ID.
func (b *ScheduleDB) Delete(bucket, id string) error {
	return b.client.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(id))
	})
}

// Fetch reads a record by ID.
func (b *ScheduleDB) Fetch(bucket, id string) (data *rg.Schedule, err error) {
	data = &rg.Schedule{}
	err = b.client.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		v := bkt.Get([]byte(id))
		if v == nil {
			return ErrNotFound
		}
		return json.Unmarshal(v, data)
	})
	return
}

// FetchAll returns all the records in the given bucket.
func (b *ScheduleDB) FetchAll(bucket string) (data []*rg.Schedule, err error) {
	data = []*rg.Schedule{}
	err = b.client.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(bucket))
		if bkt != nil {
			bkt.ForEach(func(_, v []byte) error {
				s := &rg.Schedule{}
				err := json.Unmarshal(v, s)
				if err != nil {
					return err
				}
				data = append(data, s)

				return nil
			})
		}
		return nil
	})
	return
}

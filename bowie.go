
/* * * * * * * * * * * * * * * * * * * * * * * *
  *
 *
  *   Just For my Idol
 *
  *   David Bowie, Ziggy Stardust.
 *
  *   udtrokia
 *
  *
 * * * * * * * */

package Bowie

import (
	"os"
//	"fmt"
	"log"
	"github.com/boltdb/bolt"
)

type Bowie struct {
	moonage string
	daydream os.FileMode
}

type Ash struct {
	Comet []byte `json:"comet"`
	Orbit []byte `json:"orbit"`
}

type Asher struct {
	Comet string `json:"comet"`
	Orbit string `json:"orbit"`
}


func Ziggy(moonage string, daydream os.FileMode) Bowie {
	return Bowie{
		moonage: moonage,
		daydream: daydream,
	};
}

func (david *Bowie) Mars() *bolt.DB {
	// Open the database.	
	db, err := bolt.Open(david.moonage, david.daydream, nil);
	if err != nil { log.Fatal(err)};
	
	return db;
}

func (david *Bowie) Star(Comet []byte, Orbit []byte, fiveyears bool) {
	db := david.Mars();
	
	// Insert data into a bucket.
	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("blackstar"));
		if err != nil { b = tx.Bucket([]byte("blackstar"))}

		velvet, _ := b.NextSequence();
		if(fiveyears == true){ Comet = []byte(string(velvet)) };
		if err := b.Put(Comet, Orbit);
		err != nil { return err }

		return nil
	});
	err != nil { log.Fatal(err) };

	// Close database to release file lock.
	if err := db.Close();
	err != nil { log.Fatal(err) };
}

func (david *Bowie) Oddity(fn func([]Asher)) {
	db := david.Mars();
	
	// Insert data into a bucket.
	if err := db.Update(func(tx *bolt.Tx) error  {
		b, err := tx.CreateBucket([]byte("blackstar"));
		if err != nil {
			b = tx.Bucket([]byte("blackstar"));
		}

		// Iterate over items in sorted key order.
		ashes := []Asher{};
		if err := b.ForEach(func(k []byte, v []byte) error {
			ashes = append(ashes,Asher{
				Comet: string(k[:]),
				Orbit: string(v[:]),
			})
			return nil
		});
		err != nil { return err };
		fn(ashes);
		return nil
	});
	err != nil { log.Fatal(err) };

	// Close database to release file lock.
	if err := db.Close();
	err != nil { log.Fatal(err) };
}

func (david *Bowie) LiftOff() {
	db := david.Mars();
	defer os.Remove(db.Path());
}

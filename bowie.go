
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
	"log"
	"github.com/boltdb/bolt"
)

type Bowie struct {
	moonage string
	daydream os.FileMode
}

type Ashes struct {
	comet []byte
	orbit []byte
}

func Ziggy(moonage string, daydream os.FileMode) Bowie {
	return Bowie{moonage: moonage, daydream: daydream};
}

func (david *Bowie) Mars() *bolt.DB {
	// Open the database.	
	db, err := bolt.Open("my.db", 0666, nil);
	if err != nil { log.Fatal(err)};
	
	return db;
}

func (david *Bowie) Star(comet []byte, orbit []byte) {
	db := david.Mars();
	
	// Insert data into a bucket.
	if err := db.Update(func(tx *bolt.Tx) error  {
		b, err := tx.CreateBucket([]byte("blackstar"));
		if err != nil {
			b = tx.Bucket([]byte("blackstar"));
		}
		if err := b.Put(comet, orbit);
		err != nil { return err }

		return nil
	});
	err != nil { log.Fatal(err) };

	// Close database to release file lock.
	if err := db.Close();
	err != nil { log.Fatal(err) };
}

func (david *Bowie) Oddity(fn func([]Ashes)) {
	db := david.Mars();
	
	// Insert data into a bucket.
	if err := db.Update(func(tx *bolt.Tx) error  {
		b, err := tx.CreateBucket([]byte("blackstar"));
		if err != nil {
			b = tx.Bucket([]byte("blackstar"));
		}

		// Iterate over items in sorted key order.
		Ashess := []Ashes{};
		if err := b.ForEach(func(k, v []byte) error {
			//fmt.Printf("A %s is %s.\n", k, v);
			Ashess = append(Ashess, Ashes{comet: k, orbit: v});
			return nil
		});
		err != nil { return err };
		fn(Ashess);
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

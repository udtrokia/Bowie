
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
	path string
	mode os.FileMode
}

type Pair struct {
	key []byte
	value []byte
}

func Ziggy(path string, mode os.FileMode) Bowie {
	return Bowie{path: path, mode:mode};
}

func (david *Bowie) Mars() *bolt.DB {
	// Open the database.	
	db, err := bolt.Open("my.db", 0666, nil);
	if err != nil { log.Fatal(err)};
	
	return db;
}

func (david *Bowie) Star(k []byte, v []byte) {
	db := david.Mars();
	
	// Insert data into a bucket.
	if err := db.Update(func(tx *bolt.Tx) error  {
		b, err := tx.CreateBucket([]byte("animals"));
		if err != nil {
			b = tx.Bucket([]byte("animals"));
		}
		if err := b.Put(k, v);
		err != nil { return err }

		return nil
	});
	err != nil { log.Fatal(err) };

	// Close database to release file lock.
	if err := db.Close();
	err != nil { log.Fatal(err) };
}

func (david *Bowie) Oddity(fn func([]Pair)) {
	db := david.Mars();
	
	// Insert data into a bucket.
	if err := db.Update(func(tx *bolt.Tx) error  {
		b, err := tx.CreateBucket([]byte("animals"));
		if err != nil {
			b = tx.Bucket([]byte("animals"));
		}

		// Iterate over items in sorted key order.
		Pairs := []Pair{};
		if err := b.ForEach(func(k, v []byte) error {
			//fmt.Printf("A %s is %s.\n", k, v);
			Pairs = append(Pairs, Pair{key: k, value: v});
			return nil
		});
		err != nil { return err };
		fn(Pairs);
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

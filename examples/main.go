/*
"encoding/json"
out, err := json.Marshal(a);
if err != nil { panic(err) };
fmt.Println(string(out));
*/

package main;

import (
	"fmt"
	"github.com/udtrokia/bowie"
)

func main() {                                                                                                                                                                                             
	ziggy := Bowie.Ziggy("halo", 0666);
	ziggy.Star([]byte("This is Major Tom to Ground Control,"), []byte("I'm stepping though the door."));
	ziggy.Oddity(func (pairs []Bowie.Pair){
		fmt.Printf("\n%s\n\n", pairs[0]);
	});
	ziggy.LiftOff();
} 

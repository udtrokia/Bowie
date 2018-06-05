
/*
  *
 *   Ziggy Stardust
  * 
 */

package main;

import (
	"fmt"
	"github.com/udtrokia/bowie"
)

func main(){
	ziggy := Bowie.Ziggy("halo", 0666);
	ziggy.Star(
		[]byte("This is Major Tom to Ground Control,"),
		[]byte("I'm stepping though the door."));
	
	ziggy.Oddity(func (pairs []Bowie.Ashes){
		fmt.Printf("\n%s\n\n", pairs[0]);
	});
	
	ziggy.LiftOff();
}

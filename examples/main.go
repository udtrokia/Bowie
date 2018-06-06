
/*
  *
 *   Ziggy Stardust
  * 
 */

package main;

import (
	"fmt"
	"github.com/udtrokia/Bowie"
)

func main(){
	ziggy := Bowie.Ziggy("halo", 0666);
	ziggy.Star(
		[]byte("This is Major Tom to Ground Control,"),
		[]byte("I'm stepping though the door."),
	false);
	
	ziggy.Oddity(func (pairs []Bowie.Ash){
		fmt.Printf("\n%s\n\n", pairs[0]);
	});

//	print(ziggy.FiveYears());
	
	ziggy.LiftOff();
}

package main
/*
   - Earth: orbital period 365.25 Earth days, or 31557600 seconds
   - Mercury: orbital period 0.2408467 Earth years
   - Venus: orbital period 0.61519726 Earth years
   - Mars: orbital period 1.8808158 Earth years
   - Jupiter: orbital period 11.862615 Earth years
   - Saturn: orbital period 29.447498 Earth years
   - Uranus: orbital period 84.016846 Earth years
   - Neptune: orbital period 164.79132 Earth years
*/
import (
	"fmt"
	//"math"
)

//main driver 
func main()  {
	fmt.Print("\n",Age("Mercury", 1000000000))
	fmt.Print("\n",Age("Venus", 1000000000))
	fmt.Print("\n",Age("Test", 1000000000))
	fmt.Print("\n",Age("Neptune", 1000000000))
}

func Age(p string, s float64) float64 {
	switch p {
	case "Earth":
		x := s / 31557600
		return x
		//To do - need to figure out how to round up
		//return math.Ceil((x*100)/100)
	case "Mercury":
		fmt.Print("Mercury Matched")
	case "Venus":
		fmt.Print("Venus Matched")
	case "Mars":
		fmt.Print("Mars Matched")
	case "Saturn":
		fmt.Print("Saturn Matched")
	case "Uranus":
		fmt.Print("Uranus Matched")
	case "Neptune":
		fmt.Print("Neptune Matched")
	}



	return 0 
}
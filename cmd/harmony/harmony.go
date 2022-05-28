package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/poolpOrg/go-harmony/chords"
)

func main() {
	flag.Parse()
	//	for _, note := range "CDEFGAB" {
	for _, chord := range flag.Args() {
		c, err := chords.Parse(chord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("chord", c.Name())
		for _, note := range c.Notes() {
			fmt.Println(note.Name())
		}

	}

	/*


		fmt.Println("7 dim", n.Interval(intervals.DiminishedSeventh).Name())
		fmt.Println("7 min", n.Interval(intervals.MinorSeventh).Name())
		fmt.Println("7 maj", n.Interval(intervals.MajorSeventh).Name())

		fmt.Println("octave", n.Interval(intervals.Octave).Name())
	*/
}

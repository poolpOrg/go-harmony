package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/poolpOrg/go-harmony/notes"
)

func main() {
	flag.Parse()

	for _, name := range flag.Args() {
		n, err := notes.Parse(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n, n.Frequency())
	}

	//	for _, note := range "CDEFGAB" {
	/*
		for _, scale := range flag.Args() {
			c, err := scales.Parse(scale)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("scale", c.Name())
			for _, note := range c.Notes() {
				fmt.Println(note.Name())
			}
			for _, chord := range c.Chords() {
				fmt.Println(chord.Name(), chord)
			}

		}
	*/
	/*


		fmt.Println("7 dim", n.Interval(intervals.DiminishedSeventh).Name())
		fmt.Println("7 min", n.Interval(intervals.MinorSeventh).Name())
		fmt.Println("7 maj", n.Interval(intervals.MajorSeventh).Name())

		fmt.Println("octave", n.Interval(intervals.Octave).Name())
	*/
}

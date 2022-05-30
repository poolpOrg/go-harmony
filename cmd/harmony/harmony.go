package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/poolpOrg/go-harmony/instruments"
	"github.com/poolpOrg/go-harmony/tunings"
)

func main() {
	var opt_note string
	var opt_chord string
	var opt_scale string
	var opt_notes string

	flag.StringVar(&opt_note, "note", "", "note")
	flag.StringVar(&opt_chord, "chord", "", "chord")
	flag.StringVar(&opt_scale, "scale", "", "scale")
	flag.StringVar(&opt_notes, "notes", "", "notes")

	flag.Parse()

	instrument := instruments.NewInstrument(tunings.A440)

	if opt_note != "" {
		n, err := instrument.Note(opt_note)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n.Name(), n.Frequency())
	}

	if opt_chord != "" {
		c, err := instrument.Chord(opt_chord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.Name())
		for _, n := range c.Notes() {
			fmt.Println("  ", n.Name(), n.Frequency())
		}
	}

	if opt_scale != "" {
		s, err := instrument.Scale(opt_scale)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s.Name())
		for _, n := range s.Notes() {
			fmt.Println("  ", n.Name(), n.Frequency())
			// plug here chord construction for this degree
		}
	}

	if opt_notes != "" {
		c, err := instrument.Notes(strings.Split(opt_notes, ",")...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.Name())
		for _, n := range c.Notes() {
			fmt.Println("  ", n.Name(), n.Frequency())
		}
	}
}

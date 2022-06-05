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
	var opt_distance string

	flag.StringVar(&opt_note, "note", "", "note")
	flag.StringVar(&opt_chord, "chord", "", "chord")
	flag.StringVar(&opt_scale, "scale", "", "scale")
	flag.StringVar(&opt_notes, "notes", "", "notes")
	flag.StringVar(&opt_distance, "distance", "", "distance")

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
			fmt.Printf("%8s: %3s %.02f\n", c.Root().Distance(n).Name(), n.Name(), n.Frequency())
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

		fmt.Println("Triads:")
		for _, c := range s.Triads() {
			fmt.Println("  ", c.Name())
		}

		fmt.Println("Sevenths:")
		for _, c := range s.Sevenths() {
			fmt.Println("  ", c.Name())
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

	if opt_distance != "" {
		c, err := instrument.Distances(strings.Split(opt_distance, ",")...)
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range c {
			fmt.Println("  ", n.Name())
		}
	}

}

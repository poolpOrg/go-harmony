package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/poolpOrg/go-harmony/instruments"
	"github.com/poolpOrg/go-harmony/scales"
	"github.com/poolpOrg/go-harmony/tunings"
)

func main() {
	var opt_note string
	var opt_chord string
	var opt_scale string
	var opt_notes string
	var opt_distance string
	var opt_progression string

	flag.StringVar(&opt_note, "note", "", "note")
	flag.StringVar(&opt_chord, "chord", "", "chord")
	flag.StringVar(&opt_scale, "scale", "", "scale")
	flag.StringVar(&opt_notes, "notes", "", "notes")
	flag.StringVar(&opt_distance, "distance", "", "distance")
	flag.StringVar(&opt_progression, "progression", "", "progression")
	flag.Parse()

	instrument := instruments.NewInstrument(tunings.A440)

	if opt_note != "" {
		n, err := instrument.Note(opt_note)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n.Name(), n.Octave(), n.Frequency())
	}

	if opt_chord != "" {
		c, err := instrument.Chord(opt_chord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c.Name())
		for _, n := range c.Notes() {
			fmt.Printf("%8s: %3s %d %.02f\n", c.Root().Distance(n).Name(), n.Name(), n.Octave(), n.Frequency())
		}

		scales := scales.FromChord(c)
		for _, scale := range scales {
			for _, scaleNote := range scale.Notes() {
				found := false
				for _, chordNote := range c.Notes() {
					//\033[1;36m%s\033[0m"
					if chordNote.Inharmonic(scaleNote) {
						fmt.Printf("\033[1;31m%s\033[0m\t", chordNote.Name())
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("%s\t", scaleNote.Name())
				}
			}
			if len(scale.Notes()) < 8 {
				for i := 0; i < 8-len(scale.Notes()); i++ {
					fmt.Printf("\t")
				}
			}
			fmt.Println(scale.Name())
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

	if opt_progression != "" {
		tmp := strings.Split(opt_progression, ":")

		chord := tmp[0]
		progression := tmp[1]

		s, err := instrument.Scale(chord)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Triads:")
		for _, prog := range strings.Split(progression, ",") {
			offset := 0
			switch prog {
			case "I":
				offset = 0
				break
			case "II":
				offset = 1
				break
			case "III":
				offset = 2
				break
			case "IV":
				offset = 3
				break
			case "V":
				offset = 4
				break
			case "VI":
				offset = 5
				break
			case "VII":
				offset = 6
				break
			}
			fmt.Println(prog, "\t", s.Triads()[offset].Name())
		}
		fmt.Println()

		fmt.Println("Sevenths:")
		for _, prog := range strings.Split(progression, ",") {
			offset := 0
			switch prog {
			case "I":
				offset = 0
				break
			case "II":
				offset = 1
				break
			case "III":
				offset = 2
				break
			case "IV":
				offset = 3
				break
			case "V":
				offset = 4
				break
			case "VI":
				offset = 5
				break
			case "VII":
				offset = 6
				break
			}
			fmt.Println(prog, "\t", s.Sevenths()[offset].Name())
		}
	}

}

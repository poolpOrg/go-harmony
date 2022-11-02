package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/instruments"
	"github.com/poolpOrg/go-harmony/intervals"
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
	var opt_interval string

	flag.StringVar(&opt_note, "note", "", "note")
	flag.StringVar(&opt_chord, "chord", "", "chord")
	flag.StringVar(&opt_scale, "scale", "", "scale")
	flag.StringVar(&opt_notes, "notes", "", "notes")
	flag.StringVar(&opt_distance, "distance", "", "distance")
	flag.StringVar(&opt_progression, "progression", "", "progression")
	flag.StringVar(&opt_interval, "interval", "", "interval")
	flag.Parse()

	instrument := instruments.NewInstrument(tunings.A440)

	if opt_note != "" {
		n, err := instrument.Note(opt_note)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n.OctaveName(), "=", n.Frequency())

		fmt.Println()
		structures := chords.Structures()
		chordsList := make([]string, 0)
		for _, structure := range structures {
			nc, err := chords.Parse(fmt.Sprintf("%s%s", n.Name(), structure.Name()))
			if err != nil {
				panic(err)
			}
			chordsList = append(chordsList, nc.Name())
		}
		fmt.Println(strings.Join(chordsList, ", "))
	}

	if opt_chord != "" {
		c, err := instrument.Chord(opt_chord)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-20s", c.Name()+":")
		for _, n := range c.Notes() {
			fmt.Printf("%8s: %-3s", c.Root().Distance(n).Name(), n.Name())
		}
		fmt.Println()
		fmt.Println()

		if !strings.Contains(c.Name(), "/") {
			for i := 1; i < len(c.Notes()); i++ {
				inversion, err := chords.Parse(fmt.Sprintf("%s/%s", c.Name(), c.Notes()[i].Name()))
				if err != nil {
					panic(err)
				}
				fmt.Printf("%-20s", fmt.Sprintf("%d inversion:", i))
				for _, n := range inversion.Notes() {
					fmt.Printf("%8s: %-3s", c.Root().Distance(n).Name(), n.Name())
				}
				fmt.Println()
			}
			fmt.Println()
		}

		if c.Relative() != nil {

			fmt.Printf("%-20s", fmt.Sprintf("relative %s:", c.Relative().Name()))
			for _, n := range c.Relative().Notes() {
				fmt.Printf("%8s: %-3s", c.Relative().Root().Distance(n).Name(), n.Name())
			}
			fmt.Println()
			fmt.Println()
		}

		scales := scales.FromChord(c)
		for _, scale := range scales {
			fmt.Printf("%20s (%d)\t", scale.Name(), scale.NotesInChord(c))
			for _, scaleNote := range scale.Notes() {
				found := false
				for _, chordNote := range c.Notes() {
					//\033[1;36m%s\033[0m"
					if chordNote.Inharmonic(scaleNote) {
						fmt.Printf("\033[1;32m%s\033[0m\t", chordNote.Name())
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("%s\t", scaleNote.Name())
				}
			}
			fmt.Println()
		}

	}

	if opt_scale != "" {
		s, err := instrument.Scale(opt_scale)
		if err != nil {
			log.Fatal(err)
		}

		colors := []string{
			"\033[1;32m%16s\033[0m",
			"\033[1;33m%16s\033[0m",
			"\033[1;32m%16s\033[0m",
			"\033[1;33m%16s\033[0m",
			"\033[1;31m%16s\033[0m",
			"\033[1;32m%16s\033[0m",
			"\033[1;31m%16s\033[0m",
		}

		roman := []string{"(T) I", "(SD) II", "(T) III", "(SD) IV", "(D) V", "(T) VI", "(D) VII"}
		fmt.Printf("%20s", strings.Repeat(" ", 20))
		for offset, n := range roman {
			fmt.Printf(colors[offset], n)
		}
		fmt.Println()

		notes := s.Notes()
		fmt.Printf("%-20s", notes[0].Name()+" "+s.Name()+" scale:")
		for offset, n := range notes[0 : len(notes)-1] {
			fmt.Printf(colors[offset], n.Name())
		}
		fmt.Println()

		if len(s.Notes()) >= 7 {
			fmt.Printf("%-20s", "diatonic triads:")
			for offset, c := range s.Triads() {
				fmt.Printf(colors[offset], c.Name())
			}
			fmt.Println()

			fmt.Printf("%-20s", "diatonic sevenths:")
			for offset, c := range s.Sevenths() {
				fmt.Printf(colors[offset], c.Name())
			}
			fmt.Println()
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

	if opt_interval != "" {
		interval, err := intervals.Parse(opt_interval)
		if err != nil {
			panic(err)
		}
		fmt.Println(interval.Name(), interval)
	}

}

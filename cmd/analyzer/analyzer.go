package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/poolpOrg/go-harmony/chords"
	"github.com/poolpOrg/go-harmony/notes"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	var opt_track int

	flag.IntVar(&opt_track, "track", 0, "track to follow")
	flag.Parse()

	for _, filename := range flag.Args() {
		channels := make(map[int]map[int64]map[string]uint8)
		smf.ReadTracks(filename).Do(
			func(te smf.TrackEvent) {
				if _, exists := channels[te.TrackNo]; !exists {
					channels[te.TrackNo] = make(map[int64]map[string]uint8)
				}

				if opt_track != 0 {
					if opt_track != te.TrackNo {
						return
					}
				}

				if te.Message.IsMeta() {
					fmt.Printf("[%v] @%vms %s\n", te.TrackNo, te.AbsMicroSeconds/1000, te.Message.String())
				} else {
					var ch, key, vel uint8
					switch te.Message.Type().String() {
					case "NoteOn":
						if _, exists := channels[te.TrackNo][te.AbsTicks]; !exists {
							channels[te.TrackNo][te.AbsTicks] = make(map[string]uint8)
						}
						te.Message.GetNoteStart(&ch, &key, &vel)
						channels[te.TrackNo][te.AbsTicks][midi.Note(key).String()] = key

					case "NoteOff":
						keys := make([]int64, 0)
						for key := range channels[te.TrackNo] {
							if key < te.AbsTicks {
								keys = append(keys, key)
							}
						}

						sort.SliceStable(keys, func(i, j int) bool {
							return keys[i] < keys[j]
						})

						for _, key := range keys {
							activeNotes := make([]string, 0)
							for note, _ := range channels[te.TrackNo][key] {
								activeNotes = append(activeNotes, note)
							}

							sort.SliceStable(activeNotes, func(i, j int) bool {
								return channels[te.TrackNo][key][activeNotes[i]] < channels[te.TrackNo][key][activeNotes[j]]
							})

							switch len(activeNotes) {
							case 1:
								fmt.Println("note:", activeNotes[0])
							case 2:
								root, _ := notes.Parse(activeNotes[0])
								target, _ := notes.Parse(activeNotes[1])
								fmt.Println(root.Distance(*target).Name(), "interval:", root.OctaveName(), target.OctaveName())
							case 3:
								fallthrough
							case 4:
								chordNotes := make([]notes.Note, 0)
								for _, noteName := range activeNotes {
									n, _ := notes.Parse(noteName)
									chordNotes = append(chordNotes, *n)
								}

								c := chords.FromNotes(chordNotes)
								//c.SetRoot(*c.Root().Interval(intervals.PerfectFourth))

								//for _, interval := range c.Structure() {
								//fmt.Println(interval.Name())
								//}
								fmt.Println(c.Name())
								//fmt.Println(c.Relative().Name())
							}

							delete(channels[te.TrackNo], key)
						}

					default:
						//fmt.Printf("[%v] %s %s\n", te.TrackNo, te.Message.Type().String(), te.Message)
					}

				}
			},
		)
	}
}

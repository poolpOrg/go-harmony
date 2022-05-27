package main

import (
	"fmt"
	"log"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

func main() {
	n, err := notes.NoteByName("Cb")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("1", n.Interval(intervals.PerfectUnison).Name())

	fmt.Println("1 aug", n.Interval(intervals.AugmentedUnison).Name())
	fmt.Println("2 min", n.Interval(intervals.MinorSecond).Name())

	fmt.Println("2 maj", n.Interval(intervals.MajorSecond).Name())
	fmt.Println("2 aug", n.Interval(intervals.AugmentedSecond).Name())
	fmt.Println("3 min", n.Interval(intervals.MinorThird).Name())
	fmt.Println("3 maj", n.Interval(intervals.MajorThird).Name())
	fmt.Println("4", n.Interval(intervals.PerfectFourth).Name())
	fmt.Println("4 aug", n.Interval(intervals.AugmentedFourth).Name())
	fmt.Println("5 dim", n.Interval(intervals.DiminishedFifth).Name())
	fmt.Println("5", n.Interval(intervals.PerfectFifth).Name())
	fmt.Println("5 aug", n.Interval(intervals.AugmentedFifth).Name())
	fmt.Println("6 min", n.Interval(intervals.MinorSixth).Name())
	fmt.Println("6 maj", n.Interval(intervals.MajorSixth).Name())

	fmt.Println("7 dim", n.Interval(intervals.DiminishedSeventh).Name())
	fmt.Println("7 min", n.Interval(intervals.MinorSeventh).Name())
	fmt.Println("7 maj", n.Interval(intervals.MajorSeventh).Name())

	fmt.Println("octave", n.Interval(intervals.Octave).Name())

}

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/poolpOrg/go-harmony/intervals"
	"github.com/poolpOrg/go-harmony/notes"
)

func main() {
	flag.Parse()
	//	for _, note := range "CDEFGAB" {
	for _, note := range flag.Args() {
		n, err := notes.Parse(string(note))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(intervals.PerfectUnison.Name(), n.Interval(intervals.PerfectUnison).Name())
		fmt.Println(intervals.AugmentedUnison.Name(), n.Interval(intervals.AugmentedUnison).Name())
		fmt.Println(intervals.MinorSecond.Name(), n.Interval(intervals.MinorSecond).Name())
		fmt.Println(intervals.MajorSecond.Name(), n.Interval(intervals.MajorSecond).Name())
		fmt.Println(intervals.AugmentedSecond.Name(), n.Interval(intervals.AugmentedSecond).Name())
		fmt.Println(intervals.MinorThird.Name(), n.Interval(intervals.MinorThird).Name())
		fmt.Println(intervals.MajorThird.Name(), n.Interval(intervals.MajorThird).Name())
		fmt.Println(intervals.PerfectFourth.Name(), n.Interval(intervals.PerfectFourth).Name())
		fmt.Println(intervals.AugmentedFourth.Name(), n.Interval(intervals.AugmentedFourth).Name())
		fmt.Println(intervals.DiminishedFifth.Name(), n.Interval(intervals.DiminishedFifth).Name())
		fmt.Println(intervals.PerfectFifth.Name(), n.Interval(intervals.PerfectFifth).Name())
		fmt.Println(intervals.AugmentedFifth.Name(), n.Interval(intervals.AugmentedFifth).Name())
		fmt.Println(intervals.MinorSixth.Name(), n.Interval(intervals.MinorSixth).Name())
		fmt.Println(intervals.MajorSixth.Name(), n.Interval(intervals.MajorSixth).Name())
		fmt.Println(intervals.DiminishedSeventh.Name(), n.Interval(intervals.DiminishedSeventh).Name())
		fmt.Println(intervals.MinorSeventh.Name(), n.Interval(intervals.MinorSeventh).Name())
		fmt.Println(intervals.MajorSeventh.Name(), n.Interval(intervals.MajorSeventh).Name())
		fmt.Println(intervals.Octave.Name(), n.Interval(intervals.Octave).Name())
	}

	/*


		fmt.Println("7 dim", n.Interval(intervals.DiminishedSeventh).Name())
		fmt.Println("7 min", n.Interval(intervals.MinorSeventh).Name())
		fmt.Println("7 maj", n.Interval(intervals.MajorSeventh).Name())

		fmt.Println("octave", n.Interval(intervals.Octave).Name())
	*/
}

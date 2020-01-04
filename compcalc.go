
/* Created by Pavel Raykov aka 'rabbit' / 2018-11-28 (c) */
/* vim: set ai tabstop=4 expandtab shiftwidth=4 softtabstop=4 filetype=go */

package main

import (
	"fmt"
	"math"
	"flag"
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

const (
	version float64 = 0.1
)

var (
	progname string = filepath.Base(os.Args[0])
	db DB
	dbFile string
    freq float64
    noColor bool
	onlyMatched bool
    color_matched string = "\x1b[1;37;43m"
)

// YAML configuration file structure
type DB struct {
	Range struct {
		Lower float64
		Higher float64
	}

	Tonearms []Tonearm
	Cartridges []Cart
	Headshells []Headshell
}

// Tonearm object structure
type Tonearm struct {
    Name string
    Weight float64
}

// Cartridge object structure
type Cart struct {
    Name string
    Weight float64
    Compliance float64
}

// Headshell object structure
type Headshell struct {
    Name string
    Weight float64
}

// Cartridge's method
func (c *Cart) ShowCartridgeInfo() {
    fmt.Printf("===> Name: %-10s | Weight: %.1fg | Compliance: %.1fcm/dyne\n", c.Name, c.Weight, c.Compliance)
}

func Compatibility(t *Tonearm, c *Cart, h *Headshell) {

    freq = 159/math.Sqrt((t.Weight + c.Weight + h.Weight) * c.Compliance)
	hl := "\t"
	color := ""
	matched := false

	if freq > db.Range.Lower && freq < db.Range.Higher {
		if noColor {
			hl = hl + "\b\b* "
		} else {
			color = color_matched
		}
		matched = true
	}

	if onlyMatched && !matched {
		return
	}

	fmt.Printf("%sFrequency: %s%.3f Hz\x1b[0m (%s)\n", hl, color, freq, h.Name)
}

func parseCliOptions() {
	flag.BoolVar(&noColor, "nocolor", false, "disable color output")
	flag.BoolVar(&onlyMatched, "matched", false, "show only matched cartridges")
	flag.StringVar(&dbFile, "database", "compcalc.yaml", "database file location")

	// Custom usage information
    flag.Usage = func() {
		fmt.Printf("\n# Resonance Frequency Calculator v%.1f\n\n", version)
        fmt.Printf("Usage: %s [options]\n\n", progname)
        flag.PrintDefaults()
    }

    flag.Parse()
}

func main() {
	// Parse command line options
	parseCliOptions()

	// Read database file
    rawData, err := ioutil.ReadFile(dbFile)

	if err != nil {
		log.Fatalf("Can't read database file: %s\n", dbFile)
	}

	// Parse database file
	err = yaml.Unmarshal(rawData, &db)

	if err != nil {
		log.Fatalf("Can't parse database file %s: %v\n", dbFile, err)
	}

    // Do the rest
	for _, cart := range db.Cartridges {
		cart.ShowCartridgeInfo()

		for _, tonearm := range db.Tonearms {
			fmt.Printf("  - w/Tonearm: %s (%.1fg)\n", tonearm.Name, tonearm.Weight)
			for _, hs := range db.Headshells {
					Compatibility(&tonearm, &cart, &hs)
			}
		}
	}
}

// EOF

package main

import (
	"os"

	kserver "github.com/Kserver"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

var (
	mode      string
	noRepeat  = false
	repeat    = true
	firstPass = true
)

func main() {

	if !firstPass {
		os.Args = []string{}
	}

	if len(os.Args) == 1 {
		modeSelect()
	} else {
		switch os.Args[1] {
		case "-help":
			help()
		case "-notes":
			notes()
		case "-currency", "-currencyconvert":
			CurrencyConvert(noRepeat)
		case "-ksorter":
			ksorter_integration()
		case "-ls", "-listfromdirectory":
			listFromDirectory()
		case "-funny":
			funny(noRepeat)
		case "-pricer":
			pricer(noRepeat)
		case "-fibonacci":
			fibonacciLuncher(noRepeat)
		case "-kserver":
			kserver.StartKserver()
		// case "-time", "-timezoneconverter":
		// 	timeZoneConvert()
		default:
			log.Warn("Invalid argument provided, opening normally...")
			modeSelect()
		}
	}

	if firstPass {
		firstPass = false
	}

}

func modeSelect() {

	huh.NewSelect[string]().
		Options(
			huh.NewOption("notes", "notes"),
			huh.NewOption("currency convert", "currency convert"),
			huh.NewOption("open Ksorter", "open Ksorter"),
			huh.NewOption("list from directory", "list from directory"),
			huh.NewOption("open Kserver", "open Kserver"),
			huh.NewOption("funny", "funny"),
			huh.NewOption("pricer", "pricer"),
			huh.NewOption("fibonacci", "fibonacci"),
			huh.NewOption("exit", "exit"),
		).
		Title("app mode:").
		Value(&mode).
		Run()

	switch mode {
	case "sys info":
		sysInfo()
	case "pricer":
		pricer(repeat)
	case "notes":
		notes()
	case "list from directory":
		listFromDirectory()
	// case "messenger":
	// 	messenger()
	case "funny":
		funny(repeat)
	case "currency convert":
		CurrencyConvert(repeat)
	case "fibonacci":
		fibonacciLuncher(repeat)
	// case "time zone converter":
	// 	timeZoneConvert()
	case "test":
		test()
	case "open Ksorter":
		ksorter_integration()
	case "open Kserver":
		kserver.StartKserver()
	case "exit":
		log.Error("exiting...")
		return
	}

}

func test() { // used only for new feature testing

}

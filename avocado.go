package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func getEnv(key string, defaultVal string) string {
	/*
		If key has ampty value, return default string defined as a defaultVal

		return string
	*/
	_ret := os.Getenv(key)
	if len(_ret) == 0 {
		return defaultVal
	}
	return _ret
}

var interval int

func main() {
	pwr_pin := flag.Int("pwr-pin", 18, "Power PIN where DC will be enaboled")
	measure_pin := flag.Int("pin", 4, "Input PIN ID")
	flag.IntVar(&interval, "interval", 30, "Timeout betwean measurement")
	flag.Parse()

	// Initialize GPIO
	log.Println("Initialize GPIO....")
	log.Println("Power PIN: ", *pwr_pin)
	log.Println("Read PIN:  ", *measure_pin)

	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	// Deregister in all cases
	defer rpio.Close()

	// Initialize pins
	p_pwr := rpio.Pin(*pwr_pin)
	p_read := rpio.Pin(*measure_pin)
	p_read.Input()

	// Measure every 30s
	for {
		time.Sleep(time.Duration(interval) * time.Second) // Wait 10 sec
		p_pwr.High()                                      // Enable measure detector
		time.Sleep(1 * time.Second)                       // Wait 2 secs

		log.Println("AVOCADO: measure pin ", *measure_pin)
		if p_read.Read() == 1 {
			lokiMsg(fmt.Sprintf("AVOCADO: Pin %d needs more watter", *measure_pin))
		}
		time.Sleep(1 * time.Second) // Wait 2 secs
		p_pwr.Low()

	}
}

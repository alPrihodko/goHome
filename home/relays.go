package home

import (
	"log"
	//"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

var r = raspi.NewRaspiAdaptor("raspi")
var relHeatMotor1 = gpio.NewLedDriver(r, "led", "11")
var relHeatMotor2 = gpio.NewLedDriver(r, "led", "12")

var relHeat1 = gpio.NewLedDriver(r, "led", "13")
var relHeat2 = gpio.NewLedDriver(r, "led", "15")

func init() {
	log.Println("init relays")
}

/*
ToggleHeatMotor1 switches state for heatmotor
*/
func ToggleHeatMotor1() error {
	log.Println("toggle heat motor r1")
	return relHeatMotor1.Toggle()
}

/*
ToggleHeatMotor2 switches state for heatmotor
*/
func ToggleHeatMotor2() error {
	log.Println("toggle heat motor r2")
	return relHeatMotor2.Toggle()
}

/*
ToggleHeat1 switches state for heatmotor
*/
func ToggleHeat1() error {
	log.Println("toggle heat r1")
	return relHeat1.Toggle()
}

/*
ToggleHeat2 switches state for heatmotor
*/
func ToggleHeat2() error {
	log.Println("toggle heat r2")
	return relHeat2.Toggle()
}

/*
OnHeatMotor1 switches state for heatmotor
*/
func OnHeatMotor1() error {
	log.Println("toggle heat motor r1")
	return relHeatMotor1.On()
}

/*
OnHeatMotor2 switches state for heatmotor
*/
func OnHeatMotor2() error {
	log.Println("toggle heat motor r2")
	return relHeatMotor2.On()
}

/*
OnHeat1 switches state for heatmotor
*/
func OnHeat1() error {
	log.Println("toggle heat r1")
	return relHeat1.On()
}

/*
OnHeat2 switches state for heatmotor
*/
func OnHeat2() error {
	log.Println("toggle heat r2")
	return relHeat2.On()
}

/*
OffHeatMotor1 switches state for heatmotor
*/
func OffHeatMotor1() error {
	log.Println("toggle heat motor r1")
	return relHeatMotor1.Off()
}

/*
OffHeatMotor2 switches state for heatmotor
*/
func OffHeatMotor2() error {
	log.Println("toggle heat motor r2")
	return relHeatMotor2.Off()
}

/*
OffHeat1 switches state for heatmotor
*/
func OffHeat1() error {
	log.Println("toggle heat r1")
	return relHeat1.Off()
}

/*
OffHeat2 switches state for heatmotor
*/
func OffHeat2() error {
	log.Println("toggle heat r2")
	return relHeat2.Off()
}

/*
GetRelHeat1 returns rely status
*/
func getRelHeat1() bool {
	return relHeat1.State()
}

/*
GetRelHeat2 returns rely status
*/
func GetRelHeat2() bool {
	return relHeat2.State()
}

/*
GetRelMotor1 returns rely status
*/
func GetRelMotor1() bool {
	return relHeatMotor1.State()
}

/*
GetRelMotor2 returns rely status
*/
func GetRelMotor2() bool {
	return relHeatMotor2.State()
}

package main

import "fmt"

const usixteenbitmax float64 = 645535
const kmh_multiple float64 = 1.60934

type Car struct {
	// define the struct
	gas_pedal uint16 // 0 -> 65535
	brake_pedal uint16
	steering_input int16 // -32k -> 32k
	top_speed float64

	// methods (value receivers [read], pointer receivers [write])
}

// Value receiver method
func (c Car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed/usixteenbitmax)
}

func (c Car) mph() float64{
	return float64(c.gas_pedal) * (c.top_speed/usixteenbitmax/kmh_multiple)
}

// Pointer receiver
func (c *Car) newTopSpeed(newSpeed float64){
	c.top_speed = new_speed
}

func main(){
	// reference the struct
	aNewCar := Car{65000,1,0,125.50}

	fmt.Println(aNewCar)
	fmt.Println(aNewCar.kmh())
	fmt.Println(aNewCar.mph())

	aNewCar.newTopSpeed(500)
	fmt.Println(aNewCar.kmh())
	fmt.Println(aNewCar.mph())
}
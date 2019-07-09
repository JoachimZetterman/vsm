package main

import (
	"fmt"

	. "./cmd"
)

//Example implementation of vsm.go
//
//to run:
//go run main.go
func main() {
	vechile := NewVechile("scooter1", Ready)
	fmt.Printf("name of vechile: %s \n", vechile.Name)
	fmt.Printf("init state: %s \n\n", vechile.State)

	fmt.Println("Example of events:")

	//Assumes authentication has been done...
	err := OnRiding(EndUser, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	//System user type used for automatic requests that is not handled by enduser nor hunter
	err = OnBatteryLow(System, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	err = OnBounty(System, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	err = OnCollected(Hunter, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	err = OnDropped(Hunter, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	err = OnReady(Hunter, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	err = OnUnknown(System, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

	//OverrideEvent can trigger any state of the vechile
	err = OverrideEvent(Admin, ServiceMode, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

}

package main

import (
	"fmt"

	. "./cmd"
)

//Example implementation of vsm.go
func main() {
	vechile := NewVechile("scooter1", Ready)
	fmt.Println(vechile.Name)
	fmt.Println(vechile.State)

	err := OnRiding(EndUser, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

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

	err = OverrideEvent(Admin, ServiceMode, vechile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vechile.State)

}

package vsm

import (
	"errors"
)

//User type
type User int

//State type
type State int

//enum for Users
const (
	EndUser User = iota
	Hunter
	System //Used to trigger automatic states by "overhead orchestrating system"
	Admin
)

//enum for States
const (
	Ready State = iota
	Riding
	BatteryLow
	Bounty
	Collected
	Dropped
	Unknown
	ServiceMode
	Terminated
)

//To string methond for State
func (s State) String() string {
	names := [...]string{"ready", "riding", "battery-low", "bounty", "collected", "dropped", "unknown", "service-mode", "terminated"}
	return names[s]
}

//To string methond for User
func (u User) String() string {
	names := [...]string{"end-user", "hunter", "system", "admin"}
	return names[u]
}

//Vechile - Representes a generic vechile
type Vechile struct {
	Name    string
	Battery int
	State   State
}

//NewVechile - Creates new vechile object cointaning battery and state
func NewVechile(name string, state State) *Vechile {
	v := &Vechile{
		Name:    name,
		Battery: 100,
		State:   state,
	}
	return v
}

//OverrideEvent for overrinding states as admin
func OverrideEvent(user User, state State, v *Vechile) error {
	if user == Admin {
		v.State = state
		return nil
	}
	return errors.New("Not valid user")
}

//OnReady state event
func OnReady(user User, v *Vechile) error {
	if user == EndUser || user == Hunter {
		//Checks if trasition is valid
		if v.State == Dropped || v.State == Riding {
			v.State = Ready
			return nil
		}
		return errors.New("Not valid state trasition")
	}
	return errors.New("Not valid user")
}

//OnRiding state event
func OnRiding(user User, v *Vechile) error {
	if user == EndUser {
		//Checks if trasition is valid
		if v.State == Ready {
			v.State = Riding
			return nil
		}
		return errors.New("Not valid state trasition")
	}
	return errors.New("Not valid user")
}

//OnBatteryLow state event
func OnBatteryLow(user User, v *Vechile) error {

	//TODO: ready => batter-low should be valid state?
	if user == System {
		//Checks if trasition is valid
		//TODO: Check if battery is under 20%?
		if v.State == Riding {
			v.State = BatteryLow
			return nil
		}
		return errors.New("Not valid state trasition")
	}
	return errors.New("Not valid user")
}

//OnBounty state event
func OnBounty(user User, v *Vechile) error {

	//TODO: Which system? dev, preprod prod?
	//Another event which comfirms time is after 9:30? busniess logic in event or keep logic on consumer of the API?
	if user == System {
		//Checks if trasition is valid
		if v.State == BatteryLow || v.State == Ready {
			v.State = Bounty
			return nil
		}
		return errors.New("Not valid state trasition")
	}
	return errors.New("Not valid user")
}

//OnCollected state event
func OnCollected(user User, v *Vechile) error {
	if user == Hunter {
		//Checks if trasition is valid
		if v.State == Bounty {
			v.State = Collected
			return nil
		}
		return errors.New("Not valid state trasition")
	}
	return errors.New("Not valid user")
}

//OnDropped state event
func OnDropped(user User, v *Vechile) error {
	if user == Hunter {
		//Checks if trasition is valid
		if v.State == Collected {
			v.State = Dropped
			return nil
		}
		return errors.New("Not valid state trasition")
	}
	return errors.New("Not valid user")
}

//OnUnknown state event
func OnUnknown(user User, v *Vechile) error {
	if user == System {
		v.State = Unknown
		return nil
	}
	return errors.New("Not valid user")
}

//OnService state event
func OnService(user User, v *Vechile) error {
	if user == Admin {
		v.State = ServiceMode
		return nil
	}
	return errors.New("Not valid user")
}

//OnTerminated state event
func OnTerminated(user User, v *Vechile) error {
	if user == Admin {
		v.State = Terminated
		return nil
	}
	return errors.New("Not valid user")
}

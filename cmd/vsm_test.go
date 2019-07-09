package vsm

import (
	"testing"
)

//Basic unit tests for vsm.go, not complete coverage but tests expected outcome
//
//instructions:
//in ./cmd, do go test

func TestOnReady(t *testing.T) {
	v := NewVechile("test", Riding)
	OnReady(EndUser, v)
	if v.State != Ready {
		t.Fail()
	}

	v = NewVechile("test", Dropped)
	OnReady(Hunter, v)
	if v.State != Ready {
		t.Fail()
	}
}

func TestOnRiding(t *testing.T) {
	v := NewVechile("test", Ready)
	OnRiding(EndUser, v)
	if v.State != Riding {
		t.Fail()
	}
}

func TestOnBatteryLow(t *testing.T) {
	v := NewVechile("test", Riding)
	OnBatteryLow(System, v)
	if v.State != BatteryLow {
		t.Fail()
	}
}

func TestOnBounty(t *testing.T) {
	v := NewVechile("test", BatteryLow)
	OnBounty(System, v)
	if v.State != Bounty {
		t.Fail()
	}

	v = NewVechile("test", Ready)
	OnBounty(System, v)
	if v.State != Bounty {
		t.Fail()
	}
}

func TestOnCollected(t *testing.T) {
	v := NewVechile("test", Bounty)
	OnCollected(Hunter, v)
	if v.State != Collected {
		t.Fail()
	}
}

func TestOnDropped(t *testing.T) {
	v := NewVechile("test", Collected)
	OnDropped(Hunter, v)
	if v.State != Dropped {
		t.Fail()
	}
}

func TestOnUnknown(t *testing.T) {
	v := NewVechile("test", Ready)
	OnUnknown(System, v)
	if v.State != Unknown {
		t.Fail()
	}
}

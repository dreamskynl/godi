package godi

import (
	"fmt"
	"testing"
)

func NewBuildingUnitService1(says string, bye string) IBuildingUnitService {
	return &BuildingUnitService{says: says, bye: bye}
}

func NewBuildingUnitService2(says string, bye string) (IBuildingUnitService, error, error) {
	return &BuildingUnitService{says: says, bye: bye}, nil, nil
}

func NewBuildingUnitService3(says string, bye string) (int, error) {
	return 2, nil
}

func NewBuildingUnitService4(says string, bye string) (IBuildingUnitService, int) {
	return &BuildingUnitService{says: says, bye: bye}, 0
}

func NewBuildingUnitService5(says string, bye string) (int, int) {
	return 0, 0
}

type IBuildingUnitService interface {
	SayHello()
}

type BuildingUnitService struct {
	says string
	bye  string
}

func NewBuildingUnitService(says string, bye string) (IBuildingUnitService, error) {
	return &BuildingUnitService{says: says, bye: bye}, nil
}

func (b *BuildingUnitService) SayHello() {
	fmt.Println(b.says)
}

func (b *BuildingUnitService) SayGoodbye() {
	fmt.Println(b.bye)
}

func TestRegister(t *testing.T) {
	c := New()

	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, "BuildingUnitServicePtr says hello"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, 1233); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, false); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, nil); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye", 112); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, "BuildingUnitServicePtr says hello", 112); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, 112, "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, 112, 112); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, true, false); err != nil {
		t.Error(err)
	}

	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService1, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService2, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService3, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService4, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService5, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}
}

func TestMustResolve(t *testing.T) {
	c := New()

	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}

	b := c.MustResolve(&BuildingUnitService{}).(*BuildingUnitService)
	(*b).SayHello()
}

func TestMustResolveAs(t *testing.T) {
	c := New()

	if err := c.Register(&BuildingUnitService{}, NewBuildingUnitService, "BuildingUnitServicePtr says hello", "BuildingUnitServicePtr says goodbye"); err != nil {
		t.Error(err)
	}

	b := c.MustResolveAsInstance(&BuildingUnitService{}).(BuildingUnitService)
	b.SayHello()
}

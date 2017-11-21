package shelters

import "fmt"

type Shelter struct {
	ID      uint
	Address string
}

func (c Shelter) String() string {
	return fmt.Sprintf("%v %s\n", c.ID, c.Address)
}

package pavilions

import "fmt"

type Pavilion struct {
	Number uint
	Name   string
	ID     uint
	Count  uint
}

func (c Pavilion) String() string {
	return fmt.Sprintf("%v %s %v %v\n", c.Number, c.Name, c.ID, c.Count)
}

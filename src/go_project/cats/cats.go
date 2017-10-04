package Cats
import (
	"fmt"
	"encoding/xml"
)

type Cat struct{
	XMLName xml.Name `xml:"cat"`
	Health int		`xml:"health"`
	Name string		`xml:"name"`
	Color string	`xml:"color"`
}
//**Cat to string** it is cool
func(c Cat) String() string{
	return fmt.Sprintf("%s\n%v\n%s",c.Name, c.Health, c.Color)
}

type Cats struct{
	XMLName xml.Name	`xml:"Data"`
	Cats []Cat			`xml:"cat"`
}

func (data Cats) String() string{
	var result string
	for _, v := range data.Cats{
		result += v.String()
		result +="\n\n"
	}
	result = result[0:len(result)-1]
	return result
}

//пакет предназначен для представления коллекций цветов
package colors
import (
	"encoding/xml"
)

type Colors struct{
	XMLName xml.Name	`xml:"Colors"`
	Colors []string		`xml:"color"`
}

func (data Colors) String() string{
	var result string
	for _, v := range data.Colors{
		result += v
		result +="\n\n"
	}
	result = result[0:len(result)-1]
	return result
}
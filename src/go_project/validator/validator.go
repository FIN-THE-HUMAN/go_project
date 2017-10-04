package validator
import (
	"fmt"
	"go_project/colors"
	"go_project/cats"
)

func IsAgeValid(age int) bool{
	if (age > 0 && age < 100) {return true}
	return false
}

func TestAgeAllCats(data Cats.Cats){
	fmt.Println(" --- validator ---")
	fmt.Println(" - Ages - ")
	for i := 0; i < len(data.Cats); i++ {
		tempHealth := data.Cats[i].Health
		fmt.Printf("%v valid %v",tempHealth,IsAgeValid(tempHealth))
		fmt.Println()
	}
}

func IsColorValid(TestColor string, AllColors colors.Colors) bool{
	for i:= 0; i < len(AllColors.Colors); i++{
		var temp = AllColors.Colors[i]
		if (TestColor == temp) {return true}
	}
	return false
}

func TestColorAllCats(data Cats.Cats, AllColors colors.Colors){
	fmt.Println(" --- validator ---")
	fmt.Println(" - Colors - ")
	for i := 0; i < len(data.Cats); i++ {
		tempColor := data.Cats[i].Color
		fmt.Printf("%s valid %v",tempColor,IsColorValid(tempColor, AllColors))
		fmt.Println()
	}
}
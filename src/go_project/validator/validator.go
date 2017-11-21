package validator
import (
	"fmt"
	"go_project/colors"
	"go_project/cats"
)

func IsHealthValid(health int) bool{
	if (health > 0 && health < 100) {return true}
	return false
}

func TestHealthAllCats(data Cats.Cats){
	fmt.Println(" --- validator ---")
	fmt.Println(" - Health - ")
	for i := 0; i < len(data.Cats); i++ {
		tempHealth := data.Cats[i].Health
		fmt.Printf("%v valid %v",tempHealth,IsHealthValid(tempHealth))
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

func GetCatValidatorResult(cat Cats.Cat, AllColors colors.Colors) string{
	var result string;
	if( IsHealthValid(cat.Health) && IsColorValid(cat.Color, AllColors)){
		result = fmt.Sprintf("%s is a good cat", cat.Name)
		return result;
	} 
		result = fmt.Sprintf("%s is not a good cat", cat.Name)
		return result;
}

func GetListCatValidatorResult(cats Cats.Cats, AllColors colors.Colors) []string{
	result := make([]string, len(cats.Cats))
	for i := 0; i < len(cats.Cats); i++{
		result[i] = GetCatValidatorResult(cats.Cats[i], AllColors)
	}
	return result;
}


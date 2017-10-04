package tests

import (
	"testing"
	"go_project/validator"
	"go_project/colors"
)

func TestValidatorAge_AgeIsValid(t *testing.T){
	if(!validator.IsAgeValid(20)){
		t.Error("Validator error", 20, "became invalid, but age valid in [0, 100]")
	}
}

func TestValidatorAge_AgeIsValidToLittle(t *testing.T){
	if(validator.IsAgeValid(-20)){
		t.Error("Validator error", -20, "became valid, but age valid in [0, 100]")
	}
}

func TestValidatorAge_AgeIsInValidToMuch(t *testing.T){
	if(validator.IsAgeValid(120)){
		t.Error("Validator error", 120, "became valid, but age valid in [0, 100]")
	}
}

func TestValidatorColor_ColorIsValid(t *testing.T){
	var color string = "White"
	var AllColors colors.Colors
	var array []string = []string{"White","Red","Magenta"}
	AllColors.Colors = array
	if(!validator.IsColorValid(color, AllColors)){
		t.Error("validator error", color, "became invalid, but it must be in valid colors array", AllColors)
	}
}

func TestValidatorColor_ColorIsInValid(t *testing.T){
	var color string = "Green"
	var AllColors colors.Colors
	var array []string = []string{"White","Red","Magenta"}
	AllColors.Colors = array
	if(validator.IsColorValid(color, AllColors)){
		t.Error("validator error", color, "became valid, but it must not be in valid colors array\n", AllColors)
	}
}
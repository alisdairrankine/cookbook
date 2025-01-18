package cookbook

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	ID     uuid.UUID
	Name   string
	Serves uint8

	Ingredients []Ingredient
	Steps       []CookingStep
}

type CookingStep struct {
	Description       string
	RequiredEquipment []Equipment
	Duration          time.Duration
}

type Unit string

const (
	UnitMillilitre Unit = "ml"
	UnitGram       Unit = "g"
)

type Ingredient struct {
	Amount float32
	Unit   Unit
	Name   string
}

type Equipment string

const (
	EquipmentOven  Equipment = "oven"
	EquipmentStove Equipment = "stove"
)

package cookbook

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	ID     uuid.UUID `json:"id,omitempty"`
	Name   string    `json:"name,omitempty"`
	Serves uint8     `json:"serves,omitempty"`

	Ingredients []Ingredient  `json:"ingredients,omitempty"`
	Steps       []CookingStep `json:"steps,omitempty"`
}

type CookingStep struct {
	Description       string        `json:"description,omitempty"`
	RequiredEquipment []Equipment   `json:"required_equipment,omitempty"`
	Duration          time.Duration `json:"duration,omitempty"`
}

type Unit string

const (
	UnitMillilitre Unit = "ml"
	UnitGram       Unit = "g"
)

type Ingredient struct {
	Amount float32 `json:"amount,omitempty"`
	Unit   Unit    `json:"unit,omitempty"`
	Name   string  `json:"name,omitempty"`
}

type Equipment string

const (
	EquipmentOven  Equipment = "oven"
	EquipmentStove Equipment = "stove"
)

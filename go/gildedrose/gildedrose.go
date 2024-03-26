package main

import (
    "github.com/andradecierdo/gilded-rose-kata-go/go/gildedrose"
)

type GildedRose struct {
	Items             []Item
	CalculatorBuilder ItemCalculatorBuilder
}

func NewGildedRose(items []IItem, calculatorBuilder ItemCalculatorBuilder) *GildedRose {
	return &GildedRose{
		Items:             items,
		CalculatorBuilder: calculatorBuilder,
	}
}

func (gr *GildedRose) UpdateQuality() {
	for _, item := range gr.Items {
		itemCalculator := gr.CalculatorBuilder.GetItemCalculator(item)
		itemCalculator.UpdateQuality()
	}
}


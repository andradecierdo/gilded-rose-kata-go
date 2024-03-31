package gildedrose

type GildedRose struct {
    Items             []*Item
    CalculatorBuilder ItemCalculatorBuilder
}

func NewGildedRose(items []*Item, calculatorBuilder ItemCalculatorBuilder) *GildedRose {
    return &GildedRose{
        Items:             items,
        CalculatorBuilder: calculatorBuilder,
    }
}

func (gr *GildedRose) UpdateQuality() {
    for index, item := range gr.Items {
        itemCalculator := gr.CalculatorBuilder.GetItemCalculator(item)
        itemCalculator.UpdateQuality()
        gr.Items[index] = itemCalculator.GetItem()
    }
}

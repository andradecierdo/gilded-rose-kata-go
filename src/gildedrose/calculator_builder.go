package gildedrose

type ItemCalculatorBuilder struct{}

func NewItemCalculatorBuilder() ItemCalculatorBuilder {
    return ItemCalculatorBuilder{}
}

func (icb *ItemCalculatorBuilder) GetItemCalculator(item *Item) ItemQualityCalculator {
	switch item.Type {
	case AgedBrie:
		return NewAgedBrieItem(item)
	case BackStagePasses:
		return NewBackStagePassesItem(item)
	case ConjuredManaCake:
		return NewConjuredManaCakeItem(item)
	case Sulfuras:
		return NewSulfurasItem(item)
	default:
		return NewRegularItem(item)
	}
}

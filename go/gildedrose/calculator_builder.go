package gildedrose

type ItemQualityCalculator interface {
	UpdateQuality()
}

type ItemCalculatorBuilder struct{}

func (icb *ItemCalculatorBuilder) GetItemCalculator(item models.IItem) ItemQualityCalculator {
	switch item.GetType() {
	case AgedBrie:
		return NewAgedBrieItem(item)
	case BackStagePassess:
		return NewBackStagePassesItem(item)
	case ConjuredManaCake:
		return NewConjuredManaCakeItem(item)
	case Sulfuras:
		return NewSulfurasItem(item)
	default:
		return NewRegularItem(item)
	}
}

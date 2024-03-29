package gildedrose

type ConjuredManaCakeItem struct {
    *RegularItem
}

func NewConjuredManaCakeItem(item *Item) *ConjuredManaCakeItem {
    cmci := &ConjuredManaCakeItem{
        RegularItem: NewRegularItem(item),
    }
    cmci.ItemQualityCalculator = cmci
    return cmci
}

func (cmci *ConjuredManaCakeItem) CalculateQualityAdjustment() int {
    return cmci.RegularItem.CalculateQualityAdjustment() * 2
}

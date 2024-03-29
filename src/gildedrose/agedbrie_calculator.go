package gildedrose

type AgedBrieItem struct {
    *ItemQualityManager
}

func NewAgedBrieItem(item *Item) *AgedBrieItem {
    abi := &AgedBrieItem{
        ItemQualityManager: NewItemQualityManager(item, 50, 0),
    }
    abi.ItemQualityCalculator = abi
    return abi
}

func (abi *AgedBrieItem) CalculateQualityAdjustment() int {
    if abi.SellIn < 0 {
        return 2
    }
    return 1
}

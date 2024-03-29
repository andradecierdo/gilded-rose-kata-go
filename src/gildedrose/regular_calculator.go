package gildedrose

type RegularItem struct {
    *ItemQualityManager
}

func NewRegularItem(item *Item) *RegularItem {
    ri := &RegularItem{
        ItemQualityManager: NewItemQualityManager(item, 50, 0),
    }
    ri.ItemQualityCalculator = ri
    return ri
}

func (ri *RegularItem) CalculateQualityAdjustment() int {
    if ri.SellIn < 0 {
        return -2
    }
    return -1
}
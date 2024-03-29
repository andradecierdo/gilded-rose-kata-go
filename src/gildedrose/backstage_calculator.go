package gildedrose

type BackStagePassesItem struct {
    *AgedBrieItem
}

func NewBackStagePassesItem(item *Item) *BackStagePassesItem {
    bspi := &BackStagePassesItem{
        AgedBrieItem: NewAgedBrieItem(item),
    }
    bspi.ItemQualityCalculator = bspi
    return bspi
}

func (bspi *BackStagePassesItem) CalculateQualityAdjustment() int {
    if bspi.SellIn < 0 {
        return -bspi.Quality
    }
    if bspi.SellIn < 5 {
        return 3
    }
    if bspi.SellIn < 10 {
        return 2
    }
    return bspi.AgedBrieItem.CalculateQualityAdjustment()
}

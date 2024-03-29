package gildedrose

type ItemQualityCalculator interface {
    CalculateQualityAdjustment() int
    CalculateQuality() int
    UpdateItemSellIn()
    UpdateQuality()
    GetItem() *Item
}

type QualityLimiter struct {
    Min, Max int
    Limit func(quality int) int
}

type ItemQualityManager struct {
    QualityLimiter
    ItemQualityCalculator
    *Item
}

func NewItemQualityManager(item *Item, max, min int) *ItemQualityManager {
    return &ItemQualityManager{
        Item: &Item{
            Name:    item.Name,
            SellIn:  item.SellIn,
            Quality: item.Quality,
            Type:    item.Type,
        },
        QualityLimiter: QualityLimiter{
            Min: min,
            Max: max,
        },
    }
}

func (iqm *ItemQualityManager) Limit(quality int) int {
    if quality < iqm.QualityLimiter.Min {
        return iqm.QualityLimiter.Min
    }
    if quality > iqm.QualityLimiter.Max {
        return iqm.QualityLimiter.Max
    }
    return quality
}

func(iqm *ItemQualityManager) GetItem() *Item {
    return iqm.Item
}

func (iqm *ItemQualityManager) UpdateItemSellIn() {
    iqm.SellIn--
}

func (iqm *ItemQualityManager) CalculateQualityAdjustment() int {
    return 0
}

func (iqm *ItemQualityManager) CalculateQuality() int {
    return iqm.Quality + iqm.ItemQualityCalculator.CalculateQualityAdjustment()
}

func (iqm *ItemQualityManager) UpdateQuality() {
    iqm.ItemQualityCalculator.UpdateItemSellIn()
    iqm.Quality = iqm.Limit(iqm.CalculateQuality())
}

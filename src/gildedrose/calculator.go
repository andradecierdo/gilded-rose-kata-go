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

// Regular Item
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

// Conjured Mana Cake Item
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

// Aged Brie Item
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

// Back Stage Passes Item
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

// Sulfuras Item
type SulfurasItem struct {
	*ItemQualityManager
}

func NewSulfurasItem(item *Item) *SulfurasItem {
    si := &SulfurasItem{
        ItemQualityManager: &ItemQualityManager{
            Item: &Item{
                Name:    item.Name,
                SellIn:  item.SellIn,
                Quality: item.Quality,
                Type:    item.Type,
            },
        },
    }
    si.ItemQualityCalculator = si
    return si
}

func (si *SulfurasItem) UpdateQuality() {} // Do nothing

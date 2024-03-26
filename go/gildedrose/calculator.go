package gildedrose

type Item interface {
	UpdateQuality()
}

type QualityManager interface {
	UpdateQuality()
}

type ItemQualityManager interface {
	UpdateItemSellIn()
	CalculateQuality() int
	CalculateQualityAdjustment() int
}

type QualityLimiter interface {
	Limit(quality int) int
}

type ItemQualityCalculator interface {
	IItem
	IQualityManager
}

type ItemQualityManager struct {
	Item
	max int
	min int
}

func NewItemQualityManager(item IItem, max, min int) *ItemQualityManager {
	return &ItemQualityManager{
		Item: Item{
			Name:    item.GetName(),
			SellIn:  item.GetSellIn(),
			Quality: item.GetQuality(),
			Type:    item.GetType(),
		},
		max: max,
		min: min,
	}
}

func (iqm *ItemQualityManager) Limit(quality int) int {
	if quality < iqm.min {
		return iqm.min
	}
	if quality > iqm.max {
		return iqm.max
	}
	return quality
}

func (iqm *ItemQualityManager) UpdateItemSellIn() {
	iqm.SellIn--
}

func (iqm *ItemQualityManager) CalculateQuality() int {
	return iqm.Quality + iqm.CalculateQualityAdjustment()
}

func (iqm *ItemQualityManager) UpdateQuality() {
	iqm.UpdateItemSellIn()
	iqm.Quality = iqm.Limit(iqm.CalculateQuality())
}

type RegularItem struct {
	*ItemQualityManager
}

func NewRegularItem(item IItem) *RegularItem {
	return &RegularItem{
		ItemQualityManager: NewItemQualityManager(item, 50, 0),
	}
}

func (ri *RegularItem) CalculateQualityAdjustment() int {
	if ri.SellIn < 0 {
		return -2
	}
	return -1
}

type ConjuredManaCakeItem struct {
	*RegularItem
}

func NewConjuredManaCakeItem(item IItem) *ConjuredManaCakeItem {
	return &ConjuredManaCakeItem{
		RegularItem: NewRegularItem(item),
	}
}

func (cmci *ConjuredManaCakeItem) CalculateQualityAdjustment() int {
	return cmci.RegularItem.CalculateQualityAdjustment() * 2
}

type AgedBrieItem struct {
	*ItemQualityManager
}

func NewAgedBrieItem(item IItem) *AgedBrieItem {
	return &AgedBrieItem{
		ItemQualityManager: NewItemQualityManager(item, 50, 0),
	}
}

func (abi *AgedBrieItem) CalculateQualityAdjustment() int {
	if abi.SellIn < 0 {
		return 2
	}
	return 1
}

type BackStagePassesItem struct {
	*AgedBrieItem
}

func NewBackStagePassesItem(item IItem) *BackStagePassesItem {
	return &BackStagePassesItem{
		AgedBrieItem: NewAgedBrieItem(item),
	}
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

type SulfurasItem struct {
	*Item
}

func NewSulfurasItem(item IItem) *SulfurasItem {
	return &SulfurasItem{
		Item: &Item{
			Name:    item.GetName(),
			SellIn:  item.GetSellIn(),
			Quality: item.GetQuality(),
			Type:    item.GetType(),
		},
	}
}

func (si *SulfurasItem) UpdateQuality() {} // Do nothing

package gildedrose

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

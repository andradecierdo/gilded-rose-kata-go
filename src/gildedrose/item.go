package gildedrose

type ItemType int

const (
    Sulfuras ItemType = iota
    AgedBrie
    BackStagePasses
    ConjuredManaCake
    RegularItemType
)

type Item struct {
    Name    string
    SellIn  int
    Quality int
    Type    ItemType
}

func NewItem(name string, sellIn int, quality int, itemType ItemType) *Item {
    return &Item{
        Name:    name,
        SellIn:  sellIn,
        Quality: quality,
        Type:    itemType,
    }
}

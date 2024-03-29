package main

import (
	"fmt"
	"os"
	"strconv"
    "gildedrose"
)

func main() {
    fmt.Println("OMGHAI!")

    var items = []*gildedrose.Item{
        {"+5 Dexterity Vest", 10, 20, gildedrose.RegularItemType},
        {"Aged Brie", 2, 0, gildedrose.AgedBrie},
        {"Elixir of the Mongoose", 5, 7, gildedrose.RegularItemType},
        {"Sulfuras, Hand of Ragnaros", 0, 80, gildedrose.Sulfuras},
        {"Sulfuras, Hand of Ragnaros", -1, 80, gildedrose.Sulfuras},
        {"Backstage passes to a TAFKAL80ETC concert", 15, 20, gildedrose.BackStagePasses},
        {"Backstage passes to a TAFKAL80ETC concert", 10, 49, gildedrose.BackStagePasses},
        {"Backstage passes to a TAFKAL80ETC concert", 5, 49, gildedrose.BackStagePasses},
        {"Conjured Mana Cake", 3, 6, gildedrose.ConjuredManaCake}, // <-- :O
    }
    cb := gildedrose.NewItemCalculatorBuilder()
    gl := gildedrose.NewGildedRose(items, cb)

    days := 2
    var err error
    if len(os.Args) > 1 {
        days, err = strconv.Atoi(os.Args[1])
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }
        days++
    }

    for day := 0; day < days; day++ {
        fmt.Printf("-------- day %d --------\n", day)
        fmt.Println("Name, SellIn, Quality")
        for i := 0; i < len(items); i++ {
            fmt.Println(items[i])
        }
        fmt.Println("")
        gl.UpdateQuality()
    }
}

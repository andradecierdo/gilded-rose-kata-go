package gildedrose_test

import (
	"testing"
    "gildedrose"
)

type ExpectedItems struct {
    SellIn  int
    Quality int
}

var items = []*gildedrose.Item{
    {"+5 Dexterity Vest", 10, 20, gildedrose.RegularItemType},
    {"Aged Brie", 2, 0, gildedrose.AgedBrie},
    {"Sulfuras, Hand of Ragnaros", 0, 80, gildedrose.Sulfuras},
    {"Backstage passes to a TAFKAL80ETC concert", 15, 20, gildedrose.BackStagePasses},
    {"Conjured Mana Cake", 3, 6, gildedrose.ConjuredManaCake},
}

var TestFunc = func(test *testing.T, expected []ExpectedItems, actual []*gildedrose.Item) {
    for index, item := range actual {
        if expected[index].SellIn != item.SellIn || expected[index].Quality != item.Quality  {
            test.Errorf(
                "Item: %s - Expected (Quality = %d, SellIn = %d), but got (Quality = %d, SellIn = %d)",
                 item.Name, expected[index].Quality, expected[index].SellIn, item.Quality, item.SellIn,
            )
        }
    }
}

var cb = gildedrose.NewItemCalculatorBuilder()
var gr = gildedrose.NewGildedRose(items, cb)

func Test_Gilded_Rose_Initialization(test *testing.T) {
    var expected_items = []ExpectedItems{
        { Quality: 20, SellIn: 10},
        { Quality: 0, SellIn: 2},
        { Quality: 80, SellIn: 0},
        { Quality: 20, SellIn: 15},
        { Quality: 6, SellIn: 3},
    }
    TestFunc(test, expected_items, gr.Items)
}

func Test_Gilded_Rose_Day_1(test *testing.T) {
    gr.UpdateQuality()
    var expected_items = []ExpectedItems{
        { Quality: 19, SellIn: 9},
        { Quality: 1, SellIn: 1},
        { Quality: 80, SellIn: 0},
        { Quality: 21, SellIn: 14},
        { Quality: 4, SellIn: 2},
    }
    TestFunc(test, expected_items, gr.Items)
}

func Test_Gilded_Rose_Day_2(test *testing.T) {
    gr.UpdateQuality()
    var expected_items = []ExpectedItems{
        { Quality: 18, SellIn: 8},
        { Quality: 2, SellIn: 0},
        { Quality: 80, SellIn: 0},
        { Quality: 22, SellIn: 13},
        { Quality: 2, SellIn: 1},
    }
    TestFunc(test, expected_items, gr.Items)
}

func Test_Gilded_Rose_Day_3(test *testing.T) {
    gr.UpdateQuality()
    var expected_items = []ExpectedItems{
        { Quality: 17, SellIn: 7},
        { Quality: 4, SellIn: -1},
        { Quality: 80, SellIn: 0},
        { Quality: 23, SellIn: 12},
        { Quality: 0, SellIn: 0},
    }
    TestFunc(test, expected_items, gr.Items)
}

func Test_Gilded_Rose_Day_4(test *testing.T) {
    gr.UpdateQuality()
    var expected_items = []ExpectedItems{
        { Quality: 16, SellIn: 6},
        { Quality: 6, SellIn: -2},
        { Quality: 80, SellIn: 0},
        { Quality: 24, SellIn: 11},
        { Quality: 0, SellIn: -1},
    }
    TestFunc(test, expected_items, gr.Items)
}

func Test_Gilded_Rose_Day_5(test *testing.T) {
    gr.UpdateQuality()
    var expected_items = []ExpectedItems{
        { Quality: 15, SellIn: 5},
        { Quality: 8, SellIn: -3},
        { Quality: 80, SellIn: 0},
        { Quality: 25, SellIn: 10},
        { Quality: 0, SellIn: -2},
    }
    TestFunc(test, expected_items, gr.Items)
}

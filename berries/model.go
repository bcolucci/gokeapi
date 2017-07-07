package berries

import (
	"encoding/json"
	"net/http"

	cache "github.com/patrickmn/go-cache"
)

type BerryFirmness struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type BerryFlavor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type BerryFlavorPotency struct {
	Flavor  BerryFlavor
	Potency uint `json:"potency"`
}

type BerryItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type BerryNaturalGiftType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Berry struct {
	ID               uint                 `json:"id"`
	Name             string               `json:"name"`
	GrowthTime       uint                 `json:"growth_time"`
	MaxHarvest       uint                 `json:"max_harvest"`
	NaturalGiftPower uint                 `json:"natural_gift_power"`
	NaturalGiftType  BerryNaturalGiftType `json:"natural_gift_type"`
	Size             uint                 `json:"size"`
	Smoothness       uint                 `json:"smoothness"`
	SoilDryness      uint                 `json:"soil_dryness"`
	Firmness         BerryFirmness        `json:"firmness"`
	Flavors          []BerryFlavorPotency `json:"flavors"`
	Item             BerryItem            `json:"item"`
}

func NewBerry(url string, mem *cache.Cache) *Berry {
	if entity, found := mem.Get(url); found {
		return entity.(*Berry)
	}
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var entity Berry
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&entity); err != nil {
		panic(err)
	}
	mem.Set(url, &entity, cache.DefaultExpiration)
	return &entity
}

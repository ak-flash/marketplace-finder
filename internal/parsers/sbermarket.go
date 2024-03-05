package sbermarket

import (
	"fmt"
	"marketplace-finder/config"
	"marketplace-finder/models"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/imroc/req/v3"
)

const (
	baseURL   = "https://megamarket.ru"
	searchURL = baseURL + "/api/mobile/v1/catalogService/catalog/search"
)

type MarketAnswer struct {
	Success bool   `json:"success"`
	Errors  []any  `json:"errors"`
	Total   string `json:"total"`
	Offset  string `json:"offset"`
	Limit   string `json:"limit"`
	Items   []struct {
		Goods struct {
			GoodsID    string `json:"goodsId"`
			Title      string `json:"title"`
			TitleImage string `json:"titleImage"`
			WebURL     string `json:"webUrl"`
			Brand      string `json:"brand"`
			Stocks     int    `json:"stocks"`
		} `json:"goods"`
		Price        int     `json:"price"`
		BonusPercent int     `json:"bonusPercent"`
		BonusAmount  int     `json:"bonusAmount"`
		Rating       float64 `json:"rating"`
		ReviewCount  int     `json:"reviewCount"`
		FinalPrice   int     `json:"finalPrice"`
	}
}

var client *req.Client

func init() {
	client = req.C().ImpersonateChrome()
}

func Check(target models.Target) []models.Product {

	search := target.Name

	var foundedProducts []models.Product
	var result = MarketAnswer{}

	jsonData := `{"requestVersion": 10, "limit": 42, "offset": 0, "collectionId": "", "selectedAssumedCollectionId": "", "isMultiCategorySearch": false, "searchByOriginalQuery": false, "selectedSuggestParams": [], "sorting": 1, "ageMore18": null,"showNotAvailable": false, "searchText": "` + search + `", "os": "ANDROID_OS"}`

	// Random wait time before request (For humanization)
	timeBetweenRequests := config.CfgValues.RandomSeconds
	timeout := rand.Intn(timeBetweenRequests)
	fmt.Println("Wait", timeout, "seconds until request for", target.Name)
	time.Sleep(time.Duration(timeout) * time.Second)

	// Send request to SberMegamarket
	err := client.Post(searchURL).SetBodyJsonString(jsonData).Do().Into(&result)

	if err != nil {
		config.LogFile.Println("Ошибка получения ответа: ", result)
		fmt.Println(err)
	}

	if !result.Success {
		config.LogFile.Println("Ошибка: ", result.Errors)
		fmt.Println(result)
	}

	total, _ := strconv.Atoi(result.Total)

	if total == 0 {
		target.SetError("not_found")
	}

	if total > 0 {

		fmt.Println(time.Now().Format("02/01 15:04:05"), "Информация получена: ", target.Name)
		//fmt.Println(result.Items)

		for _, item := range result.Items {

			difference := item.Price - item.BonusAmount

			if difference > 0 && difference < target.Price {

				p := models.Product{
					Name:              item.Goods.Title,
					Image:             item.Goods.TitleImage,
					Price:             item.Price,
					Bonuses:           item.BonusAmount,
					BonusesPercentage: item.BonusPercent,
					Link:              item.Goods.WebURL,
					VirtualPrice:      difference,
				}

				foundedProducts = append(foundedProducts, p)
			}

		}

		if foundedProducts != nil {
			sort.SliceStable(foundedProducts, func(i, j int) bool {
				return foundedProducts[i].VirtualPrice < foundedProducts[j].VirtualPrice
			})

			founded := len(foundedProducts)
			fmt.Println("Найдено товаров: " + strconv.Itoa(founded))
		}
	}

	return foundedProducts
}

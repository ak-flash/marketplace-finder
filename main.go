package main

import (
	_ "marketplace-finder/config"
	utils "marketplace-finder/internal"
	sbermarket "marketplace-finder/internal/parsers"
	"marketplace-finder/models"
	"sync"
)

func main() {

	models.ConnectDatabase()

	runParser()
}

func runParser() {

	targets := models.GetTargets()

	var wg sync.WaitGroup
	wg.Add(len(targets))

	for i, target := range targets {

		go func(target models.Target, i int) {

			//sbermarket.Check(item)
			foundedProducts := sbermarket.Check(target)
			//fmt.Println(foundedProducts)
			if len(foundedProducts) > 0 {
				utils.SendMessage(foundedProducts, &target)
			}

			target.UpdateCheckTime()

			wg.Done()

		}(target, i)

	}

	wg.Wait()
}

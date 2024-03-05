package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	_ "marketplace-finder/config"
	"marketplace-finder/controllers"
	utils "marketplace-finder/internal"
	sbermarket "marketplace-finder/internal/parsers"
	"marketplace-finder/models"
	"sync"
	"time"
)

func main() {

	models.ConnectDatabase()

	go func() {
		s := gocron.NewScheduler(time.UTC)
		// добавляем одну задачу на каждую минуту
		s.Cron("* * * * *").Do(runParser)
		// запускаем планировщик с блокировкой текущего потока
		s.StartBlocking()
	}()

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./web/dist", false)))
	r.GET("/getTargets", controllers.GetTargets)

	r.Run()

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
				message := utils.MakeProductNotifyText(foundedProducts, target.Name)
				utils.SendMessage(message)
			}

			target.UpdateCheckTime()

			wg.Done()

		}(target, i)

	}

	wg.Wait()
}

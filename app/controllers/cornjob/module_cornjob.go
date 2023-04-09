package controllers

import (
	// "fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
)

func CronJob(a *fiber.App) {
	d := cron.New()

	//SET ทุกๆ 1 วินาที 	=> "*/1 * * * * *"
	//SET ทุกๆ 1 ชั่วโมง 	=> "0 0 */1 * * *"
	//SET ทุกๆ เที่ยงคืน 	 => "0 0 0 * * *"
	//SET เฉพาะ 12:00 	  => "00 00 12 * *"

	// เอาไว้ทดสอบ
	// d.AddFunc("0 0 */1 * * *", func() {
	// 	fmt.Println("alert when 1 hour")
	// })

	d.Start()
}

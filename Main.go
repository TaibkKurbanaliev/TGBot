package main

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	os.Setenv("LANG", "en_US.UTF-8")
	bot, _ := CreateBot("Bot", "7440754965:AAHbEbi6FcTb5T1G_VE9wkoCO7LYReGlgwU", &tele.LongPoller{Timeout: 10 * time.Second})
	//err := db.Init("postgres", "123", "test", "disable")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// if bot == nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var user User
	// err = db.Select(2, "user", &user)
	// fmt.Print(user)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	//init parralel function
	// ticker := time.NewTicker(1 * time.Minute)
	// quit := make(chan struct{})

	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			SetCoinsRequest()
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()

	bot.Start()
	// log.Fatal("Wqe")
	// partitions, err := disk.Partitions(false)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, partition := range partitions {
	// 	fmt.Printf("Device: %v\n", partition.Device)
	// 	fmt.Printf("Mountpoint: %v\n", partition.Mountpoint)
	// 	fmt.Printf("Fstype: %v\n", partition.Fstype)
	// 	fmt.Printf("Opts: %v\n", partition.Opts)

	// 	// Get disk usage for the partition
	// 	usage, err := disk.Usage(partition.Mountpoint)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("Total: %v GB\n", usage.Total/1e9)
	// 	fmt.Printf("Used: %v GB\n", usage.Used/1e9)
	// 	fmt.Printf("Free: %v GB\n", usage.Free/1e9)
	// 	fmt.Printf("Percent: %v %%\n", usage.UsedPercent)

	// 	fmt.Println()
	// }
}

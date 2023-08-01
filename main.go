package main

import (
	"belajar_native/configs"
	"belajar_native/models"
	rabbitmq "belajar_native/rabbitMQ"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func main() {

	con, err := configs.ConnectRabbitMq()
	if err != nil {
		log.Fatal(err)
	}

	channel, err := con.Channel()
	if err != nil {
		log.Fatal(err)
	}
	rabbitmq.DeclareQueQue(channel)
	rabbitmq.DeclareExchangeAndBind(channel)
	go CheckDbFrequency1(channel)

	time.Sleep(1500 * time.Second)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on port 8080 ")
}

func CheckDbFrequency1(channel *amqp.Channel) {
	SptaChan := make(chan []models.TSpta)
	TSelectorChan := make(chan []models.TSelektor)
	TTimbanganChan := make(chan []models.TTimbang)
	MejaTebuChan := make(chan []models.TMejaTebu)
	AriChan := make(chan []models.TAri)
	VwSbhDataChan := make(chan []models.VwSbhData)
	go func() {
		for {
			db, err := configs.ConnectDbSatu()
			if err != nil {
				log.Fatal(err)
			}

			var TbLogsSyncProcess []models.TbLogsSyncProcessMtani
			err = db.Table("tb_logs_sync_process_mtani").Limit(50).Order("id asc").Find(&TbLogsSyncProcess).Error

			checktoSql := db.ToSQL(func(tb *gorm.DB) *gorm.DB {
				return tb.Table("tb_logs_sync_process_mtani").Limit(5).Order("id asc").Find(&TbLogsSyncProcess)
			})
			fmt.Println("checktoSql", checktoSql)

			if err != nil {
				log.Fatal("errrrrosssss", err)
			}
			var coubtData int64 = 0
			fmt.Println("TbLogsSyncProcess", TbLogsSyncProcess)
			for _, checkData := range TbLogsSyncProcess {
				if checkData.TTable == "t_spta" {
					var Spta []models.TSpta
					var VwSbhData []models.VwSbhData
					err = db.Table("t_spta").Where("id = ?", checkData.TId).Find(&Spta).Error
					if err != nil {
						log.Fatal(err)
					}

					for _, SptaData := range Spta {
						var VwSbhDataSingle models.VwSbhData
						err := db.Table("vw_sbh_data").Where("no_spat = ?", SptaData.NoSpat).Find(&VwSbhDataSingle).Error
						if err != nil {
							log.Fatal(err)
						}
						VwSbhData = append(VwSbhData, VwSbhDataSingle)
					}
					VwSbhDataChan <- VwSbhData
					SptaChan <- Spta
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}

				if checkData.TTable == "t_selektor" {
					var TSelector []models.TSelektor
					var VwSbhData []models.VwSbhData
					err = db.Table("vw_selektor_mtani").Where("id_selektor = ?", checkData.TId).Find(&TSelector).Error
					if err != nil {
						log.Fatal(err)
					}
					for _, TselektorData := range TSelector {
						var VwSbhDataSingle models.VwSbhData
						err := db.Table("vw_sbh_data").Where("no_spat = ?", TselektorData.NoSpat).Find(&VwSbhDataSingle).Error
						if err != nil {
							log.Fatal(err)
						}
						VwSbhData = append(VwSbhData, VwSbhDataSingle)
					}
					VwSbhDataChan <- VwSbhData
					TSelectorChan <- TSelector
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)

					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1

				}

				if checkData.TTable == "t_timbangan" {
					var TTimbangan []models.TTimbang
					var VwSbhData []models.VwSbhData
					err = db.Table("vw_timbangan_mtani").Where("id_timbangan = ?", checkData.TId).Find(&TTimbangan).Error
					if err != nil {
						log.Fatal(err)
					}
					for _, TTimbanganData := range TTimbangan {
						var VwSbhDataSingle models.VwSbhData
						err := db.Table("vw_sbh_data").Where("no_spat = ?", TTimbanganData.NoSpat).Find(&VwSbhDataSingle).Error
						if err != nil {
							log.Fatal(err)
						}
						VwSbhData = append(VwSbhData, VwSbhDataSingle)
					}
					VwSbhDataChan <- VwSbhData
					TTimbanganChan <- TTimbangan
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}
				if checkData.TTable == "t_meja_tebu" {
					var MejaTebu []models.TMejaTebu
					var VwSbhData []models.VwSbhData
					err = db.Table("vw_meja_tebu_mtani").Where("id_mejatebu = ?", checkData.TId).Find(&MejaTebu).Error
					if err != nil {
						log.Fatal(err)
					}
					for _, MejaTebuData := range MejaTebu {
						var VwSbhDataSingle models.VwSbhData
						err := db.Table("vw_sbh_data").Where("no_spat = ?", MejaTebuData.NoSpat).Find(&VwSbhDataSingle).Error
						if err != nil {
							log.Fatal(err)
						}
						VwSbhData = append(VwSbhData, VwSbhDataSingle)
					}
					VwSbhDataChan <- VwSbhData
					MejaTebuChan <- MejaTebu
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}

				if checkData.TTable == "t_ari" {
					var Ari []models.TAri
					var VwSbhData []models.VwSbhData
					err := db.Table("vw_ari_mtani").Where("id_ari = ?", checkData.TId).Find(&Ari).Error
					if err != nil {
						log.Fatal(err)
					}
					for _, AriData := range Ari {
						var VwSbhDataSingle models.VwSbhData
						err := db.Table("vw_sbh_data").Where("no_spat = ?", AriData.NoSpat).Find(&VwSbhDataSingle).Error
						if err != nil {
							log.Fatal(err)
						}
						VwSbhData = append(VwSbhData, VwSbhDataSingle)
					}
					VwSbhDataChan <- VwSbhData
					AriChan <- Ari
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}

			}
			fmt.Println("coubtData", coubtData)
			time.Sleep(5 * time.Second)
			fmt.Println("Send Data to RabbitMQ 1")
		}
	}()

	go func() {
		for results := range SptaChan {

			if len(results) > 0 {
				for _, result := range results {
					data, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}

					message := amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					}

					if err := channel.PublishWithContext(
						context.Background(),
						"t_spta",
						"",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}

			} else {
				log.Println("No Data 1")
			}
		}
	}()

	go func() {
		for result2 := range TSelectorChan {
			if len(result2) > 0 {
				for _, result := range result2 {
					data, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}

					message := amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					}

					if err := channel.PublishWithContext(
						context.Background(),
						"t_selektor",
						"",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}()
	go func() {
		for result2 := range TTimbanganChan {
			if len(result2) > 0 {
				for _, result := range result2 {
					data, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}

					message := amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					}

					if err := channel.PublishWithContext(
						context.Background(),
						"t_timbangan",
						"",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}()
	go func() {
		for result2 := range MejaTebuChan {
			if len(result2) > 0 {
				for _, result := range result2 {
					data, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}

					message := amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					}

					if err := channel.PublishWithContext(
						context.Background(),
						"t_meja_tebu",
						"",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}()
	go func() {
		for result2 := range AriChan {
			if len(result2) > 0 {
				for _, result := range result2 {
					data, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}

					message := amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					}

					if err := channel.PublishWithContext(
						context.Background(),
						"t_ari",
						"",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}()

	go func() {
		for result := range VwSbhDataChan {
			if len(result) > 0 {
				for _, result := range result {
					data, err := json.Marshal(result)
					if err != nil {
						log.Fatal(err)
					}

					message := amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					}

					if err := channel.PublishWithContext(
						context.Background(),
						"vw_sbh_data",
						"",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}()
}

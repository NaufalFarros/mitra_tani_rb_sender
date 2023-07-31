package main

import (
	"belajar_native/configs"
	"belajar_native/models"
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

	_, err = channel.QueueDeclare(
		"mitra_tani_spta",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"mitra_tani_spta",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"mitra_tani_selektor",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"mitra_tani_timbangan",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"mitra_tani_meja_tebu",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"mitra_tani_ari",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"t_spta_kuota_mtani",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"t_spta_kuota_kkw_mtani",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
		"t_spta_kuota_tot_mtani",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	go CehckDbFrequencySptaKuota(channel)

	// go CheckDbFrequency1(channel)
	time.Sleep(1500 * time.Second)
	// go CheckDbFrequency2(channel, wg)

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
					err = db.Table("t_spta").Where("id = ?", checkData.TId).Find(&Spta).Error
					if err != nil {
						log.Fatal(err)
					}
					SptaChan <- Spta
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}

				if checkData.TTable == "t_selektor" {
					var TSelector []models.TSelektor
					err = db.Table("vw_selektor_mtani").Where("id_selektor = ?", checkData.TId).Find(&TSelector).Error
					if err != nil {
						log.Fatal(err)
					}
					TSelectorChan <- TSelector
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)

					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1

				}

				if checkData.TTable == "t_timbangan" {
					var TTimbangan []models.TTimbang
					err = db.Table("vw_timbangan_mtani").Where("id_timbangan = ?", checkData.TId).Find(&TTimbangan).Error
					if err != nil {
						log.Fatal(err)
					}
					TTimbanganChan <- TTimbangan
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}
				if checkData.TTable == "t_meja_tebu" {
					var MejaTebu []models.TMejaTebu
					err = db.Table("vw_meja_tebu_mtani").Where("id_mejatebu = ?", checkData.TId).Find(&MejaTebu).Error
					if err != nil {
						log.Fatal(err)
					}
					MejaTebuChan <- MejaTebu
					db.Table("tb_logs_sync_done_mtani").Create(&checkData)
					db.Table("tb_logs_sync_process_mtani").Where("id = ?", checkData.ID).Delete(&checkData)
					coubtData += 1
				}

				if checkData.TTable == "t_ari" {
					var Ari []models.TAri
					err := db.Table("vw_ari_mtani").Where("id_ari = ?", checkData.TId).Find(&Ari).Error
					if err != nil {
						log.Fatal(err)
					}
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
						"",
						"mitra_tani_spta",
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
						"",
						"mitra_tani_selektor",
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
						"",
						"mitra_tani_timbangan",
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
						"",
						"mitra_tani_meja_tebu",
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
						"",
						"mitra_tani_ari",
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
func CehckDbFrequencySptaKuota(channel *amqp.Channel) {
	SptaKuotaChan := make(chan []models.TSptaKuota)
	SptaKuotaKkwChan := make(chan []models.TSptaKuotaKkw)
	SptaKuotaTotChan := make(chan []models.TSptaKuotaTot)
	var offsetSptaKuota, offsetSptaKuotaKkw, offsetSptaKuotaTot int
	go func() {
		for {
			db, err := configs.ConnectDbSatu()
			if err != nil {
				log.Fatal(err)
			}

			var TbSptaKuota []models.TSptaKuota
			err = db.Table("t_spta_kuota").Limit(50).Offset(offsetSptaKuota).Order("id asc").Find(&TbSptaKuota).Error
			if err != nil {
				log.Fatal(err)
			}

			var TbSptaKuotaKkw []models.TSptaKuotaKkw
			err = db.Table("t_spta_kuota_kkw").Limit(50).Offset(offsetSptaKuotaKkw).Order("id asc").Find(&TbSptaKuotaKkw).Error
			if err != nil {
				log.Fatal(err)
			}
			var TbSptaKuotaTot []models.TSptaKuotaTot
			err = db.Table("t_spta_kuota_tot").Limit(50).Offset(offsetSptaKuotaTot).Order("id asc").Find(&TbSptaKuotaTot).Error
			if err != nil {
				log.Fatal(err)
			}

			if len(TbSptaKuota) > 0 {
				SptaKuotaChan <- TbSptaKuota
				for _, data := range TbSptaKuota {
					// Buat struct untuk log deteksi
					logDeteksi := models.LogDeteksi{
						OriginalDataID:   data.ID,
						MigrateToMongoDb: false, // Ganti dengan status migrasi yang sesuai
					}

					// Cari data berdasarkan ID di tabel log untuk deteksi
					var existingData models.LogDeteksi
					err := db.Table("t_spta_kuota_migrate_log").Where("original_data_id = ?", data.ID).First(&existingData).Error
					if err != nil {
						// Jika data tidak ditemukan, buat data baru
						err = db.Table("t_spta_kuota_migrate_log").Create(&logDeteksi).Error
						if err != nil {
							log.Fatal(err)
						}
					} else {
						// Jika data ditemukan, update data
						err = db.Table("t_spta_kuota_migrate_log").Where("original_data_id = ?", data.ID).Updates(&logDeteksi).Error
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				offsetSptaKuota += len(TbSptaKuota)
			}
			if len(TbSptaKuotaKkw) > 0 {
				SptaKuotaKkwChan <- TbSptaKuotaKkw

				for _, data := range TbSptaKuotaKkw {
					// Buat struct untuk log deteksi
					logDeteksi := models.LogDeteksi{
						OriginalDataID:   data.ID,
						MigrateToMongoDb: false, // Ganti dengan status migrasi yang sesuai
					}

					// Cari data berdasarkan ID di tabel log untuk deteksi
					var existingData models.LogDeteksi

					err := db.Table("t_spta_kuota_kkw_migrate_log").Where("original_data_id = ?", data.ID).First(&existingData).Error
					if err != nil {
						// Jika data tidak ditemukan, buat data baru
						err = db.Table("t_spta_kuota_kkw_migrate_log").Create(&logDeteksi).Error
						if err != nil {
							log.Fatal(err)
						}
					} else {
						// Jika data ditemukan, update data

						err = db.Table("t_spta_kuota_kkw_migrate_log").Where("original_data_id = ?", data.ID).Updates(&logDeteksi).Error
						if err != nil {
							log.Fatal(err)
						}
					}

				}
				offsetSptaKuotaKkw += len(TbSptaKuotaKkw)
			}
			if len(TbSptaKuotaTot) > 0 {
				SptaKuotaTotChan <- TbSptaKuotaTot

				for _, data := range TbSptaKuotaTot {
					// Buat struct untuk log deteksi
					logDeteksi := models.LogDeteksi{
						OriginalDataID: data.ID,

						MigrateToMongoDb: false, // Ganti dengan status migrasi yang sesuai
					}

					// Cari data berdasarkan ID di tabel log untuk deteksi
					var existingData models.LogDeteksi

					err := db.Table("t_spta_kuota_tot_migrate_log").Where("original_data_id = ?", data.ID).First(&existingData).Error
					if err != nil {
						// Jika data tidak ditemukan, buat data baru
						err = db.Table("t_spta_kuota_tot_migrate_log").Create(&logDeteksi).Error
						if err != nil {
							log.Fatal(err)
						}
					} else {
						// Jika data ditemukan, update data

						err = db.Table("t_spta_kuota_tot_migrate_log").Where("original_data_id = ?", data.ID).Updates(&logDeteksi).Error
						if err != nil {
							log.Fatal(err)
						}
					}

				}
				offsetSptaKuotaTot += len(TbSptaKuotaTot)
			}

			time.Sleep(15 * time.Second)
			fmt.Println("Send Data to RabbitMQ Spta Kuota")
		}
	}()

	go func() {
		for results := range SptaKuotaChan {
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
						"",
						"t_spta_kuota_mtani",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			} else {
				log.Println("No Data Spta Kuota")
			}
		}
	}()

	go func() {
		for results := range SptaKuotaKkwChan {
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
						"",
						"t_spta_kuota_kkw_mtani",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			} else {
				log.Println("No Data Spta Kuota Kkw")
			}
		}
	}()

	go func() {
		for results := range SptaKuotaTotChan {
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
						"",
						"t_spta_kuota_tot_mtani",
						false,
						false,
						message,
					); err != nil {
						log.Fatal(err)
					}
				}
			} else {
				log.Println("No Data Spta Kuota Tot")
			}
		}
	}()
}

// func CheckDbFrequency2(channel *amqp.Channel, wg *sync.WaitGroup) {
// 	resultsChan := make(chan []models.TSpta)

// 	go func() {
// 		for {
// 			db, err := configs.ConnectDBDua()
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			var results []models.TSpta
// 			var result2 []models.
// 				err = db.Find(&results).Error
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			resultsChan <- results
// 			time.Sleep(10 * time.Second)

// 			fmt.Println("Send Data to RabbitMQ 2")
// 		}
// 	}()

// 	for results := range resultsChan {
// 		if len(results) > 0 {
// 			for _, result := range results {
// 				data, err := json.Marshal(result)
// 				if err != nil {
// 					log.Fatal(err)
// 				}

// 				message := amqp.Publishing{
// 					ContentType: "application/json",
// 					Body:        data,
// 				}

// 				if err := channel.PublishWithContext(
// 					context.Background(),
// 					"",
// 					"mitra_tani",
// 					false,
// 					false,
// 					message,
// 				); err != nil {
// 					log.Fatal(err)
// 				}
// 			}
// 			// delete data from db after send to rabbitmq
// 			// db, err := configs.ConnectDB()
// 			// if err != nil {
// 			// 	log.Fatal(err)
// 			// }

// 			// err = db.Delete(&results).Error
// 			// if err != nil {
// 			// 	log.Fatal(err)
// 			// }

// 		} else {
// 			log.Println("No Data 2")
// 		}
// 	}
// }

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
		"mitra_tani", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	go CheckDbFrequency(channel)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on port 8080 ")
}

func CheckDbFrequency(channel *amqp.Channel) {
	resultsChan := make(chan []models.TestSpta)

	go func() {
		for {
			db, err := configs.ConnectDB()
			if err != nil {
				log.Fatal(err)
			}

			var results []models.TestSpta

			err = db.Find(&results).Error
			if err != nil {
				log.Fatal(err)
			}

			resultsChan <- results
			time.Sleep(10 * time.Second)

			fmt.Println("Send Data to RabbitMQ")
		}
	}()

	for results := range resultsChan {
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
					"mitra_tani",
					false,
					false,
					message,
				); err != nil {
					log.Fatal(err)
				}
			}
			// delete data from db after send to rabbitmq
			// db, err := configs.ConnectDB()
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// err = db.Delete(&results).Error
			// if err != nil {
			// 	log.Fatal(err)
			// }

		} else {
			log.Println("No Data")
		}
	}
}

package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareExchangeAndBind(channel *amqp.Channel) error {
	err := channel.ExchangeDeclare(
		"vw_sbh_data",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"vw_sbh_data_in",
		"",
		"vw_sbh_data",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"vw_sbh_data_out",
		"",
		"vw_sbh_data",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"vw_sbh_data_notif",
		"",
		"vw_sbh_data",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare(
		"t_spta",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_spta_in",
		"",
		"t_spta",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = channel.QueueBind(
		"t_spta_out",
		"",
		"t_spta",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_spta_notif",
		"",
		"t_spta",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare(
		"t_selektor",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_selektor_in",
		"",
		"t_selektor",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_selektor_out",
		"",
		"t_selektor",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_selektor_notif",
		"",
		"t_selektor",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare(
		"t_timbangan",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_timbangan_in",
		"",
		"t_timbangan",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_timbangan_out",
		"",
		"t_timbangan",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_timbangan_notif",
		"",
		"t_timbangan",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare(
		"t_meja_tebu",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_meja_tebu_in",
		"",
		"t_meja_tebu",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_meja_tebu_out",
		"",
		"t_meja_tebu",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_meja_tebu_notif",
		"",
		"t_meja_tebu",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare(
		"t_ari",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_ari_in",
		"",
		"t_ari",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(
		"t_ari_out",
		"",
		"t_ari",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind(

		"t_ari_notif",
		"",
		"t_ari",
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func DeclareQueQue(channel *amqp.Channel) error {
	_, err := channel.QueueDeclare(
		"t_spta_in",
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
		"t_spta_out",
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
		"t_spta_notif",
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
		"t_selektor_in",
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
		"t_selektor_out",
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
		"t_selektor_notif",
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
		"t_timbangan_in",
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
		"t_timbangan_out",
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
		"t_timbangan_notif",
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
		"t_meja_tebu_in",
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
		"t_meja_tebu_out",
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
		"t_meja_tebu_notif",
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
		"t_ari_in",
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
		"t_ari_out",
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
		"t_ari_notif",
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
		"vw_sbh_data_in",
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
		"vw_sbh_data_out",
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
		"vw_sbh_data_notif",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

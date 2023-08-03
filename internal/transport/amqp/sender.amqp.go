package amqp

import (
	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
	amqp "github.com/rabbitmq/amqp091-go"
)

/**
Для работы с отправкой имейлов нам нужны следующие вещи
- название очереди - пока константа для упрощения
- создать саму очередь при запуске amqp транспорта
- структура для вычитывания из очереди - Consumer
- структура для записи в очередь - Publisher
- структура сообщения - подобная grpc
- структура конфигурации amqp - хост, порт, логин, пароль, connectionstring. Пока ограничимся одним consumer и producer для упрощения
Общая идея: транспортный слой принимает данные из сообщения и преобразует его в dto для взаимодействия с service-слоем
*/

type MailQueueMessage struct {
	to        []string
	subject   string
	variables map[string]interface{}
}

type Consumer struct {
	cfg *config.AMQP
	ch  <-chan amqp.Delivery
}

// попытка подключения к очереди
func (c *Consumer) Connect() error {
	return nil
}

// открытие канала доставки
func (c *Consumer) OpenChannel() error {
	return nil
}

// Получение канала доставки
func (c *Consumer) QueueDelivery() (<-chan amqp.Delivery, error) {
	return nil, nil
}

type Publisher struct {
}

func Test() {

	/** Вычитываем из очереди*/

	/** Создаём dto */
	/** Мапинг variables */
	vm := map[string]interface{}{"message": "test html template"}

	_ = dto.NewEmailDTO([]string{"not-real@example.com"}, "Very important subject", vm)

	/** Вызываем сервис mail */

}

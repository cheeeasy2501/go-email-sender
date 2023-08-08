package amqp

import (
	_ "fmt"

	_ "github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
	_ "github.com/cheeeasy2501/go-email-sender/pkg/errs"
	_ "github.com/rabbitmq/amqp091-go"
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

type IMailQueueMessage interface {
	To() []string
	Subject() string
	Variables() map[string]interface{}
}

type MailQueueMessage struct {
	to        []string
	subject   string
	variables map[string]interface{}
}

func (m *MailQueueMessage) To() []string {
	return m.to
}

func (m *MailQueueMessage) Subject() string {
	return m.subject
}

func (m *MailQueueMessage) Variables() map[string]interface{} {
	return m.variables
}

type Publisher struct {
}

func Run() {

	/** Вычитываем из очереди*/

	/** Создаём dto */
	vm := map[string]interface{}{"message": "test html template", "test": "abscscs", "test1": "qdddddd"}
	m := dto.NewEmailDTO([]string{"not-real@example.com"}, "Very important subject", vm)
	_ = m
	/** Мапинг variables */

	/** Вызываем сервис mail */

}

func MapVariables(m IMailQueueMessage) {

}

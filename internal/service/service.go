package service

type Services struct {
	SenderService ISenderService
}

func NewServices() *Services {
	return &Services{
		SenderService: NewSenderService(),
	}
}

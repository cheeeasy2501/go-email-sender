package dto

type IEmailDTO interface {
	To() []string
	Subject() string
	Variables() map[string]interface{}
}

type EmailDTO struct {
	// from      string
	to        []string
	subject   string
	variables map[string]interface{}
}

func NewEmailDTO(
	// from string, TODO: зачем нам знать откуда пришёл имеил, если само приложение должно знать об этом?
	to []string, subject string, variables map[string]interface{}) EmailDTO {

	return EmailDTO{
		// from:    from,
		to:        to,
		subject:   subject,
		variables: variables,
	}
}

// func (d *EmailDTO) From() string {
// 	return d.from
// }

func (d *EmailDTO) To() []string {
	return d.to
}

func (d *EmailDTO) Subject() string {
	return d.subject
}

func (d *EmailDTO) Variables() map[string]interface{} {
	return d.variables
}

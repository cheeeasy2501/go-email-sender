package dto

type EmailDTO struct {
	to      []string
	subject string
	body    string
}

func NewEmailDTO(to []string, subject, body string) EmailDTO {
	return EmailDTO{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (d *EmailDTO) To() []string {
	return d.to
}

func (d *EmailDTO) Subject() string {
	return d.subject
}

func (d *EmailDTO) Body() string {
	return d.body
}

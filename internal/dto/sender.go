package dto

type EmailDTO struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewEmailDTO(from string, to []string, subject, body string) EmailDTO {
	return EmailDTO{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (d *EmailDTO) From() string {
	return d.from
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

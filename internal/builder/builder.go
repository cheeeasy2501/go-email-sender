package builder

import (
	"github.com/cheeeasy2501/go-email-sender/internal/dto"
)

// TODO: придумать как собрать msg по RFC. Возможно builder
type MailBuilder struct {
}

func NewMailBuilder() *MailBuilder {
	return &MailBuilder{}
}

func (b *MailBuilder) Build(d dto.IEmailDTO) ([]byte, error) {
	madeBody := "TEST BODY HERE!"

	return []byte(
		"From: " + "test@example.com" + "\r\n" +
			"To: " + d.To()[0] + "\r\n" +
			"Subject: " + d.Subject() + "\r\n\r\n" +
			madeBody + "\r\n"), nil
}

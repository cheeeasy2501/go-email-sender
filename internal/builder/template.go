package builder

import (
	"fmt"

	"github.com/cheeeasy2501/go-email-sender/pkg/errs"
)

type ITemplate interface {
	Load() error
	SetVariables(v map[string]interface{})
	Path() string
	SetPath(string)
	ToHtml() string
}

type Template struct {
	p, t string
	v    map[string]interface{}
}

func NewTemplate(p string, v map[string]interface{}) Template {
	return Template{
		p: p,
		v: v,
	}
}

func NewFastTemplate(p string, v map[string]interface{}) (Template, error) {
	sn := "amqp"
	fn := "NewTemplate"
	m := "Can't create new template"

	if p == "" {
		return Template{}, errs.NewError(
			sn,
			fn,
			m,
			fmt.Errorf("Can't create new template - path variable is empty"),
		)
	}

	t := Template{
		p: p,
		v: v,
	}

	err := t.Load()
	if err != nil {
		return Template{}, errs.NewError(
			sn,
			fn,
			m,
			fmt.Errorf("Can't create new temlate - can't load template from path - %s", t.Path()),
		)
	}

	return t, nil
}

func (t *Template) SetVariables(v map[string]interface{}) *Template {
	t.v = v

	return t
}

func (t *Template) Path() string {
	return t.p
}

func (t *Template) SetPath(p string) *Template {
	t.p = p

	return t
}

func (t *Template) Load() error {
	// TODO: считываем из базы html
	t.t = "<div> HELLO WORLD </div>"

	return nil
}

func (t *Template) ToHtml() string {
	// TODO: мапим в шаблон, и возвращаем html
	return t.t
}

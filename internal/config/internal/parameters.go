package internal

import (
	"strings"

	"github.com/goexl/gox"
	"github.com/pangum/db/internal/config/internal/constant"
)

type Parameters map[string]any

func (p Parameters) String() string {
	values := make([]string, len(p))
	for key, value := range p {
		builder := gox.StringBuilder(key)
		if stringed := gox.ToString(value); "" != stringed {
			builder.Append(constant.Equal).Append(stringed)
		}
		values = append(values, builder.String())
	}

	return strings.Join(values, constant.And)
}

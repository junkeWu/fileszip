package fileszip

import "strings"

type DefaultUserHook struct {
	replacer *strings.Replacer
}

func (d *DefaultUserHook) TransPath(p Sources) string {
	return d.replacer.Replace(p.Url)
}

package time_helpers

import (
	"github.com/moisespsena-go/i18n-modular/i18nmod"
	path_helpers "github.com/moisespsena-go/path-helpers"
)

var (
	group = i18nmod.PkgToGroup(path_helpers.GetCalledDir())

	ErrBadFormat = i18nmod.Err(group + ".err_bad_format")

	ErrBadHour   = i18nmod.ErrData{Group: group + ".err_", Key: "bad_hour"}
	ErrBadMinute = i18nmod.ErrData{Group: group + ".err_", Key: "bad_minute"}
	ErrBadSecond = i18nmod.ErrData{Group: group + ".err_", Key: "bad_second"}
)

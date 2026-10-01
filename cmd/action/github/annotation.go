package github

import (
	"strings"

	"github.com/sethvargo/go-githubactions"
)

type AnnotationLevel int

const (
	AnnotationLevelError   AnnotationLevel = 0
	AnnotationLevelWarning AnnotationLevel = 1
	AnnotationLevelNotice  AnnotationLevel = 2
)

func (al AnnotationLevel) String() string {
	switch al {
	case AnnotationLevelError:
		return "error"
	case AnnotationLevelWarning:
		return "warning"
	case AnnotationLevelNotice:
		return "notice"
	}
	return ""
}

func Annotate(action *githubactions.Action, level AnnotationLevel, title string, message string, file string, line int) {
	msg := strings.ReplaceAll(message, "%", "%25")
	msg = strings.ReplaceAll(msg, "\r", "%0D")
	msg = strings.ReplaceAll(msg, "\n", "%0A")

	action.Infof("::%s file=%s,line=%d,title=%s::%s\n", level.String(), file, line, title, msg)
}

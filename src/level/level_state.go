package level

import "strings"

var (
	Trac  = "TRAC"
	Debug = "DEBUG"
	Info  = "INFO"
	Warn  = "WARN"
	Error = "ERROR"
	Crit  = "CRIT"
	Alrt  = "ALRT"
	Emer  = "EMER"
)

var currentLevelSate = Trac

func GetLevelState(contentLine string) string {
	updateLevelState(contentLine)
	return currentLevelSate
}

func updateLevelState(contentLine string) {
	if strings.Contains(contentLine, Trac) {
		currentLevelSate = Trac
	} else if strings.Contains(contentLine, Debug) {
		currentLevelSate = Debug
	} else if strings.Contains(contentLine, Info) {
		currentLevelSate = Info
	} else if strings.Contains(contentLine, Warn) {
		currentLevelSate = Warn
	} else if strings.Contains(contentLine, Error) {
		currentLevelSate = Error
	} else if strings.Contains(contentLine, Crit) {
		currentLevelSate = Crit
	} else if strings.Contains(contentLine, Alrt) {
		currentLevelSate = Alrt
	} else if strings.Contains(contentLine, Emer) {
		currentLevelSate = Emer
	}
}

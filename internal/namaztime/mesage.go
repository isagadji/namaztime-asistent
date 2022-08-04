package namaztime

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"
)

const (
	namazTimeMessageTemplate    = "До {{.Description}} намаза {{.TitleRu}} в {{.Time}}, {{.LeftText}} {{.TimeLeft}}"
	namazTimeMessageTemplateTTS = "До {{.Description}} намаза {{.TTS}} в {{.Time}}, {{.LeftText}} {{.TimeLeft}}"
)

var (
	declensionHours   = []string{"час", "часа", "часов"}
	declensionMinutes = []string{"минута", "минуты", "минут"}
	declensionLeft    = []string{"Остался", "Осталось", "Осталось"}
)

type TextDto struct {
	Description string
	LeftText    string
	TTS         string
	Time        string
	TimeLeft    string
	Title       string
	TitleRu     string
}

func (n *TextDto) New(title string, time time.Time, timeLeft time.Duration) *TextDto {
	namazTime := TextsMap[title]
	namazTime.TimeLeft = fmtDurationToText(timeLeft)
	namazTime.Time = time.Format("15:04")
	namazTime.LeftText = getDeclensionByNumber(time.Hour(), declensionLeft)

	return namazTime
}

func fmtDurationToText(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	mText := getDeclensionByNumber(int(m), declensionMinutes)
	hText := getDeclensionByNumber(int(h), declensionHours)

	if h <= 0 {
		return fmt.Sprintf("%d %s", m, mText)
	}
	return fmt.Sprintf("%d %s и %d %s", h, hText, m, mText)
}

func getDeclensionByNumber(number int, declensionForms []string) string {
	if len(declensionForms) < 3 {
		return ""
	}

	numberRune := []rune(strconv.Itoa(number))
	numberLen := len(numberRune)
	lastOneSymbol, _ := strconv.Atoi(string(numberRune[numberLen-1 : numberLen]))

	if numberLen > 1 {
		lastTwoSymbol, _ := strconv.Atoi(string(numberRune[numberLen-2 : numberLen]))
		if lastTwoSymbol >= 10 && lastTwoSymbol <= 20 {
			return declensionForms[2]
		}
	}

	if lastOneSymbol == 0 {
		return declensionForms[2]
	} else if lastOneSymbol < 2 {
		return declensionForms[0]
	} else if lastOneSymbol < 5 {
		return declensionForms[1]
	} else {
		return declensionForms[2]
	}
}

func getTextByTextDtoAndTemplate(namazTextDto *TextDto, textTemplate string) (*string, error) {
	tmpl, err := template.New("namaz-time").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	resultBuffer := &bytes.Buffer{}
	if err = tmpl.Execute(resultBuffer, namazTextDto); err != nil {
		return nil, err
	}

	result := resultBuffer.String()

	return &result, nil
}

var TextsMap = map[string]*TextDto{
	Fajr: {
		Description: "утреннего",
		LeftText:    "",
		TTS:         "Фаджр",
		Time:        "",
		TimeLeft:    "",
		Title:       "Fajr",
		TitleRu:     "Фаджр",
	},
	Dhuhr: {
		Description: "обеденного",
		LeftText:    "",
		TTS:         "Зухр",
		Time:        "",
		TimeLeft:    "",
		Title:       "Dhuhr",
		TitleRu:     "Зухр",
	},
	Asr: {
		Description: "послеобеденного",
		LeftText:    "",
		TTS:         "`Аср",
		Time:        "",
		TimeLeft:    "",
		Title:       "Asr",
		TitleRu:     "Аср",
	},
	Maghrib: {
		Description: "вечернего",
		LeftText:    "",
		TTS:         "М`агриб",
		Time:        "",
		TimeLeft:    "",
		Title:       "Maghrib",
		TitleRu:     "Магриб",
	},
	Isha: {
		Description: "Ночного",
		LeftText:    "",
		TTS:         "`Иша",
		Time:        "",
		TimeLeft:    "",
		Title:       "Isha",
		TitleRu:     "Иша",
	},
}

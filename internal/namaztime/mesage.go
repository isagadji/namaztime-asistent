package namaztime

import (
	"bytes"
	"fmt"
	"text/template"
	"time"
)

const (
	namazTimeMessageTemplate    = "До {{.Description}} намаза {{.TitleRu}} в {{.Time}}, осталось {{.TimeLeft}}"
	namazTimeMessageTemplateTTS = "До {{.Description}} намаза {{.TTS}} в {{.Time}}, осталось {{.TimeLeft}}"
)

type TextDto struct {
	Title       string
	TitleRu     string
	Description string
	Time        string
	TimeLeft    string
	TTS         string
}

func (n *TextDto) New(title string, time time.Time, timeLeft time.Duration) *TextDto {
	namazTime := TextsMap[title]
	namazTime.TimeLeft = fmtDurationToText(timeLeft)
	namazTime.Time = time.Format("15:04")

	return namazTime
}

func fmtDurationToText(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	mText := "минут"
	if m == 1 {
		mText = "минута"
	} else if m <= 4 {
		mText = "минуты"
	}

	hText := "часов"
	if h == 1 {
		hText = "час"
	} else if h <= 4 {
		hText = "часа"
	}

	if h <= 0 {
		return fmt.Sprintf("%d %s", m, mText)
	}
	return fmt.Sprintf("%d %s и %d %s", h, hText, m, mText)
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
		Title:       "Fajr",
		TitleRu:     "Фаджр",
		TTS:         "Фаджр",
		Description: "утреннего",
		Time:        "",
		TimeLeft:    "",
	},
	Dhuhr: {
		Title:       "Dhuhr",
		TitleRu:     "Зухр",
		TTS:         "Зухр",
		Description: "обеденного",
		Time:        "",
		TimeLeft:    "",
	},
	Asr: {
		Title:       "Asr",
		TitleRu:     "Аср",
		TTS:         "`Аср",
		Description: "послеобеденного",
		Time:        "",
		TimeLeft:    "",
	},
	Maghrib: {
		Title:       "Maghrib",
		TitleRu:     "Магриб",
		TTS:         "М`агриб",
		Description: "вечернего",
		Time:        "",
		TimeLeft:    "",
	},
	Isha: {
		Title:       "Isha",
		TitleRu:     "Иша",
		TTS:         "`Иша",
		Description: "Ночного",
		Time:        "",
		TimeLeft:    "",
	},
}

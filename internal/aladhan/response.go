package aladhan

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	//Timings map[string]string `json:"timings"`
	Timings Timings `json:"timings"`
	Date    Date    `json:"date"`
	Meta    Meta    `json:"meta"`
}

type Timings struct {
	Fajr    string `json:"Fajr"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

type Date struct {
	Readable  string   `json:"readable"`
	Timestamp string   `json:"timestamp"`
	Hijri     DateDate `json:"hijri"`
	Gregorian DateDate `json:"gregorian"`
}

type DateDate struct {
	Date    string `json:"date"`
	Format  string `json:"format"`
	Year    string `json:"year"`
	Day     string `json:"day"`
	WeekDay struct {
		En string `json:"en"`
		Ar string `json:"ar"`
	} `json:"week_day"`
	Month struct {
		Number int    `json:"number"`
		En     string `json:"en"`
		Ar     string `json:"ar"`
	} `json:"month"`
	Designation struct {
		Abbreviated string `json:"abbreviated"`
		Expanded    string `json:"expanded"`
	} `json:"designation"`
	HolyDays []*interface{} `json:"holydays,omitempty"`
}

type Meta struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Method    struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Params struct {
			Fajr int `json:"fajr"`
			Isha int `json:"isha"`
		} `json:"params"`
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"method"`
	Offset Offset `json:"offset"`
}

type Offset struct {
	Fajr    int `json:"fajr"`
	Dhuhr   int `json:"dhuhr"`
	Asr     int `json:"asr"`
	Maghrib int `json:"maghrib"`
	Isha    int `json:"isha"`
}

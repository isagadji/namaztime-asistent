package namaztime

type MarusyaResponse struct {
	//Данные для ответа пользователю.
	Response Response `json:"response"`

	//Данные о сессии.
	Session Session `json:"session"`

	//Версия протокола, текущая — 1.0.
	Version string `json:"version"`
}

func NewMarusyaResponse(msg *string, r *MarusyaRequest) MarusyaResponse {
	return MarusyaResponse{
		Response: Response{
			Text:       msg,
			Tts:        msg,
			EndSession: false,
		},
		Session: Session{
			SessionID: r.Session.SessionID,
			UserID:    r.Session.UserID,
			MessageID: r.Session.MessageID,
			SkillID:   r.Session.SkillID,
			AuthToken: r.Session.AuthToken,
		},
		Version: r.Version,
	}
}

type Session struct {
	//Уникальный идентификатор сессии, максимум 64 символа.
	SessionID string `json:"session_id"`

	//Идентификатор экземпляра приложения, в котором пользователь общается с Марусей, максимум 64 символа.
	UserID string `json:"user_id"`

	//Идентификатор сообщения в рамках сессии, максимум 8 символов. Инкрементируется с каждым следующим запросом.
	MessageID int `json:"message_id"`

	SkillID   *string `json:"skill_id"`
	AuthToken *string `json:"auth_token"`
}

type Response struct {
	//Текст, который следует показать и сказать пользователю, максимум 1 024 символа. Не должен быть пустым.
	//В тексте ответа можно указать переводы строк последовательностью «\n».
	//Если передать массив строк, то сообщения разобьются на баблы.
	Text *string `json:"text"`

	//Ответ в формате TTS (text-to-speech) (https://dev.vk.com/marusia/tts), максимум 1 024 символа.
	//Поддерживается расстановка ударений с помощью '+'.
	Tts *string `json:"tts"`

	//Кнопки (suggest), которые следует показать пользователю.
	//Кнопки можно использовать как релевантные ответу ссылки или подсказки для продолжения разговора.
	Buttons []*Button `json:"buttons,omitempty"`

	//Признак конца разговора:
	// • true — сессию следует завершить,
	// • false — сессию следует продолжить.
	EndSession bool `json:"end_session"`

	//Описание карточки — сообщения с различным контентом.
	//Подробнее о типах карточек и описание структур в специальном разделе https://dev.vk.com/marusia/cards.
	Card *Card `json:"card,omitempty"`

	//Команды. Поле позволяет передать несколько сообщений в нужном порядке.
	//На данный момент поддерживаются только карточки https://dev.vk.com/marusia/cards.
	Commands *[]string `json:"commands,omitempty"`
}

type Button struct {
	//Текст кнопки, максимум 64 символа.
	Title string `json:"title"`

	//URL, который откроется при нажатии на кнопку, максимум 1 024 байта.
	//Если свойство URL не указано, по нажатию на кнопку навыку будет отправлен текст кнопки.
	//Пока кнопки с URL не поддерживаются в приложении VK.
	URL string `json:"url"`

	//Любой JSON, который нужно отправить скиллу, если эта кнопка будет нажата, максимум 4 096 байт.
	Payload *Payload `json:"payload"`
}

type Payload struct {
}

type Card struct {
	//Тип карточки: BigImage
	Type string `json:"type"`

	//ID изображения из раздела Медиафайлы в настройках скилла.
	ImageID int `json:"image_id"`
}

type TTS struct {
}

type MarusyaRequest struct {
	Request Request `json:"request"`
	Session Session `json:"session"`
	Version string  `json:"version"`
	Meta    Meta    `json:"meta"`
}

type Meta struct {
	ClientID string `json:"client_id"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	CityRu   string `json:"_city_ru"`
}

type Request struct {
	Commands string `json:"commands"`
}

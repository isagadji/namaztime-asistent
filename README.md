# NamazTime
Скилл голосового ассистента, для получения информации об оставшемся времени до следующего намаза.
При запросе `Маруся запусти скилл время намаза` ответит `До обеденного намаза Зухр в 12:35, Осталось 9 часов и 43 минуты`

Скилл удобен при использовании с колонкой. 
Например: Капсула, Капсула мини или любая другая колонка с Марусей (пока) внутри.

В планах добавить поддержку Алисы от Яндекса.
Пока проблемы с определением города, откуда пришел запрос.
Есть идея как это обойти, будем решать.


#### Подерживаемые способы вызова:
- Маруся: `маруся запусти скилл время намаза`, `маруся запусти скилл намаз`, `маруся запусти скилл азан`
- Алиса: -

#### TODO:
1. ~~Add client for azan API~~
2. ~~Setup deploy~~
3. ~~Add http server metrics~~ 
4. Add http client metrics (client aladhan)
5. AddTests (in process..)
6. ...
7. Add support Yandex.Alisa


#### Links
API: https://aladhan.com/prayer-times-api#GetTimingsByCity

test webhook url: http://localhost:3000/webhook@8207853
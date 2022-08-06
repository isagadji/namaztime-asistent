# NamazTime
Скилл для голосового ассистента для информации о времени намаза.
При запросе `Маруся запусти скилл время намаза` ответит `До обеденного намаза Зухр в 12:35, Осталось 9 часов и 43 минуты` 

Сейчас поддерживается только асистент Маруся от VK.

В планах добавить поддержку Алисы от Яндекса. Пока проблемы с определением города откуда пришел запрос.
Есть идея как это обойти, будем решать.



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
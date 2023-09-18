Запустить сервер:

```(bash)
docker-compose up
```

Запустить клиент:

```(bash)
make build && make run
```

Детали:
* Игра начинается, когда подключатся 4 клиента к серверу.
* Роли выдаются рандомно и при старте сессии.
* Сервер может одновременно обслуживать много сессий.
* При подключении необходимо ввести имя(любая строка до '\n').
* В первый день можно либо сразу "лечь спать", либо обмениваться сообщениями.
* В первую ночь комиссар и мафия выбирают цель. Сразу после выбора цели, наступает день.
* Днем становится известно о жертве мафии. При желании комиссар может опубликовать данные проверенного игрока.
* Днем голосованием выбирают мафию, а результаты голосования становятся известны сразу после того как все будут готовы к ночи.
* Если голосованием удалось убить мафию или если мафия осталась 1 на 1 с мирным жителем, то игра прекращается и объявляются результаты.
* После объявления результатов сессия завершается. Для новой сессии нужно повторно подключить к серверу клиентов.


Технические детали:
Клиент на сервер отправляет "действия", а сервер может рассылать уведомлений клиентам. Общение построено вокруг двунаправленного стрима.

Типы действия, которые порождает клиент:
* Инициализация нового игрока;
* Голосование за мафию;
* Проверка роли(для комиссара);
* Публикация роли(для комиссара);
* Убийство(для мафии);
* Отправка сообщения;
* "Завершить дневную активность.


Типы уведомлений которые присылает сервер:
* Подключение нового клиента;
* Начало игры;
* О голосовании(кто за кого проголосовал);
* Результаты голосования;
* Сообщение в чате;
* Окончание игры.
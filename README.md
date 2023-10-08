# RESTFul API

## Структура
├───cmd<br />
├───internal<br />
│   ├───app<br />
│   ├───models<br />
│   ├───repository<br />
│   │ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  └───postgres<br />
│   ├───service<br />
│   └───transport<br />
└───pkg<br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───client<br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───postgresql<br />

- cmd - основная папка в которой содержиться точка входа в приложение (main.go)

- internal  - папки, которые используются для бизнес логики, модели и репозитории
  -  app         - папка содержащая app.go для запуска всего приложения

  -  models      - модели использующие во всем приложении

  -  repository  - папка репозитория для создании запросов в БД

  -  service     - основной сервис для бизнес логики

  -  transport   - транспортный слой для обработки запросов извне (handler)

- pkg       - папка содержащая клиента для соединения к БД

# APIs
##### [POST]  /consultation-requests
>   Создание запроса на консультацию

##### [POST]  /consultation-requests/{id}/recommendation
>   Создание рекомендации для конкретного запроса на консультацию

##### [GET]   /patient/{id}/recommendations
>   Получение списка рекомендаций для конкретного пациента
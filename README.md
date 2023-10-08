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

```
Body: 
{
    "first_name": "Daniel",
    "middle_name": "Parker",
    "last_name": "John",
    "DateOfBirst": "12.04.1995",
    "phone": "+334442211",
    "email": "anet@gj.com",
    "text": "My consultation request"
}
Response:
{
    "id": 2,
    "message": "success"
}
```
##### [POST]  /consultation-requests/{id}/recommendation
>   Создание рекомендации для конкретного запроса на консультацию

```
Response:
{
    "recommendations": [
        {
            "first_name": "Daniel",
            "middle_name": "Parker",
            "last_name": "John",
            "request_text": "My second",
            "recommendations": "Recommendation about health; Recommendation about knee"
        }
    ],
    "total": 1
}
```

##### [GET]   /patient/{id}/recommendations
>   Получение списка рекомендаций для конкретного пациента

```
Body: 
{
    "text": "Recommendation about knee"
}
Resonse:
{
    "id": 2,
    "message": "success"
}
```
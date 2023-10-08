# RESTFul API
### Как запустить

1. Нужно подключиться к БД и создать таблицы скопировав данные с db.sql
2. ```docker-compose up --build server``` - для запуска контейнеров

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
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───cache<br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───client<br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───postgresql<br />
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───util<br />


- cmd - основная папка в которой содержиться точка входа в приложение (main.go)

- internal  - папки, которые используются для бизнес логики, модели и репозитории
  -  app         - папка содержащая app.go для запуска всего приложения

  -  models      - модели использующие во всем приложении

  -  repository  - папка репозитория для создании запросов в БД

  -  service     - основной сервис для бизнес логики

  -  transport   - транспортный слой для обработки запросов извне (handler)

- pkg       - папка содержащая клиента для соединения к БД, кэш и утилиты (хелперы)

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

##### [GET]   /patient/{id}/recommendations
>   Получение списка рекомендаций для конкретного пациента, используется API OpenFDA для получения медицинский перпаратов (количество перпаратов с каждым запросом рандомны)

```
Response:
{
    "recommendations": [
        {
            "first_name": "Daniel",
            "middle_name": "Parker",
            "last_name": "John",
            "request_text": "My second",
            "recommendations": "Recommendation about health; Recommendation about knee",
            "medicine": [
                {
                    "product_ndc": "72606-003",
                    "brand_name": "Clobazam",
                    "active_ingredients": [
                        {
                            "name": "CLOBAZAM",
                            "strength": "10 mg/1"
                        }
                    ],
                    "route": [
                        "ORAL"
                    ],
                    "labeler_name": "CELLTRION USA, INC."
                }
            ]
        }
    ],
    "total": 1
}

{
    "recommendations": [
        {
            "first_name": "Daniel",
            "middle_name": "Parker",
            "last_name": "John",
            "request_text": "My second",
            "recommendations": "Recommendation about health; Recommendation about knee",
            "medicine": []
        }
    ],
    "total": 1
}

{
    "recommendations": [
        {
            "first_name": "Daniel",
            "middle_name": "Parker",
            "last_name": "John",
            "request_text": "My second",
            "recommendations": "Recommendation about health; Recommendation about knee",
            "medicine": [
                {
                    "product_ndc": "72606-003",
                    "brand_name": "Clobazam",
                    "active_ingredients": [
                        {
                            "name": "CLOBAZAM",
                            "strength": "10 mg/1"
                        }
                    ],
                    "route": [
                        "ORAL"
                    ],
                    "labeler_name": "CELLTRION USA, INC."
                },
                {
                    "product_ndc": "72698-811",
                    "brand_name": "Art of Sport Anti Dandruff Compete",
                    "active_ingredients": [
                        {
                            "name": "PYRITHIONE ZINC",
                            "strength": "100 g/mL"
                        }
                    ],
                    "route": [
                        "TOPICAL"
                    ],
                    "labeler_name": "AOS GROUP INC, THE"
                },
                {
                    "product_ndc": "72789-184",
                    "brand_name": "Tadalafil",
                    "active_ingredients": [
                        {
                            "name": "TADALAFIL",
                            "strength": "10 mg/1"
                        }
                    ],
                    "route": [
                        "ORAL"
                    ],
                    "labeler_name": "PD-Rx Pharmaceuticals, Inc."
                }
            ]
        }
    ],
    "total": 1
}
```
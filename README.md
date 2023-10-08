# RESTFul API

## Структура
├───cmd
├───internal
│   ├───app
│   ├───models
│   ├───repository
│   │ &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;  └───postgres
│   ├───service
│   └───transport
└───pkg
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───client
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;└───postgresql

- cmd - основная папка в которой содержиться точка входа в приложение (main.go)

- internal  - папки, которые используются для бизнес логики, модели и репозитории
  -  app         - папка содержащая app.go для запуска всего приложения

  -  models      - модели использующие во всем приложении

  -  repository  - папка репозитория для создании запросов в БД

  -  service     - основной сервис для бизнес логики

  -  transport   - транспортный слой для обработки запросов извне (handler)

- pkg       - папка содержащая клиента для соединения к БД
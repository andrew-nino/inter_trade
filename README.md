# Account management service

Микросервис для работы со строками: кешированием входящих строк по алгоритмам MD5 и SHA256, выдачей значения по ключу и удалением значения.

Используемые технологии:
- Redis (в качестве хранилища для бысрого доступа)
- PostgreSQL (в качестве основного хранилища данных)
- Docker (для запуска сервиса)
- Swagger (для документации API)
- Gin (веб фреймворк)
- golang-migrate/migrate (для миграций БД)
- sqlx (драйвер для работы с PostgreSQL)

Сервис был написан с Clean Architecture, что позволяет легко расширять функционал сервиса и тестировать его.
Также был реализован Graceful Shutdown для корректного завершения работы сервиса

# Getting Started

Для запуска сервиса без интеграции с Google Drive достаточно заполнить .env файл,
и определить режим запуска фреймворка Gin.

# Usage

Для просмотра всех доступных команд используйте `make help`

Собрать и запустить сервис можно с помощью команд `make build && make up`

Документацию после завпуска сервиса можно посмотреть по адресу `http://localhost:8080/swagger/index.html`
с портом 8080 по умолчанию

Документацию после завпуска сервиса можно посмотреть по адресу `http://localhost:8080/swagger/index.html`
с портом 8080 по умолчанию.

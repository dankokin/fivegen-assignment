# Hello, interviewee

### Intro

Сделай форк данного репозитория и выполни задание в соответствии с требованиями, указанными ниже. Git Flow обеспечивает плюс в карму.

### Требования:

Необходимо написать аналог файлообменника. Сервер должен принимать на ручку _/api/upload_ файл и возвращать в ответе укороченную ссылку формата _http://addr/{url}_, по которой, в дальнейшем, можно загрузить данный файл.

Требования к инструментам и библиотекам:
- Используй std пакет http для реализации вэб сервера
- Мы не ограничиваем тебя в выборе систем хранения данных, но если будешь использовать реляционную БД - выбирай postgres
- Используй docker-compose для оркестрации docker-контейнеров

Бонусные задачи:
- Продумать и реализовать алгоритм удаления "устаревших" файлов
- Использовать в реализации сервиса nginx для проксирования запросов на web сервер 

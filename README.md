# task_rest
test task for create REST API with Go

1) Запрос который умеет расшифровывать строки такого вида A2(A2BC3(F4B)D)E => [ output: ABBFBBFBBFBBDBBFBBFBBFBBDE ]
    Тело запроса будет примерно таким: {decrypt:"A2(A2BC3(F4B)D)E"}
2) Сделать запрос для шифрования строк
3) Получение истории, в теле поля Limit и Offset для получения истории по частям

БД: postgresql (развернуть в докере)
Сохраняем историю
Таблица: 
    Колонки: id, type, input, output
        type => encrypt|decrypt

Для реализации сервера использовать стандартный http пакет
Прикрутить логгер не из стандартной либы (можно zerolog, он в одном из проектов)
Добавить middleware для логгирования запросов в формате: от кого, метод, название запроса

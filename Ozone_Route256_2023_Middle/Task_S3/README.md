# S3. Контест: таблица результатов (SQL, 30 баллов)

- ограничение по времени на тест: *15 секунд*
- ограничение по памяти на тест: *1024 мегабайта*
- ввод: *стандартный ввод*
- вывод: *стандартный вывод*
***
Это необычная задача — вам надо написать **SQL**-запрос. В качестве решения вы должны отослать один запрос к базе данных, который возвращает требуемые данные. Запрос может содержать произвольное количество подзапросов, других конструкций, быть сколь угодно навороченным, но это должен быть один запрос (в нём не должна встречаться точка с запятой для разделения разных запросов).

При проверке вашего решения используется **PostgreSQL 15.1**. В качестве входных данных вам предоставляется дамп состояния базы данных. Обратите внимание, что время работы вашего решения на тесте включает восстановление состояния базы данных из дампа, но это время значительно меньше ограничения по времени. Вы можете использовать сервис [sqlfiddle](http://sqlfiddle.com/) как инструмент для запуска запросов.

В этой задаче вам предстоит написать запрос к базе данных простейшей системы проведения соревнований по программированию. Вы прямо сейчас участвуете в подобном соревновании. Время почувствовать себя в роли разработчика системы для проведения таких соревнований!

Напишите запрос к базе данных, который построит таблицу результатов для соревнования с *максимальным id*.

Вывод должен включать всех пользователей, кто сделал хотя бы одну попытку в этом соревновании. Вывод должен включать 5 колонок:

- **rank** — место пользователя в контесте (пользователи с одинаковыми результатами делят место);
- **user_id** — *id* пользователя;
- **user_name** — *name* пользователя;
- **problem_count** — количество решённых в контесте задач (если одна задача решена многократно, то всё-равно учитывается как одна задача);
- **latest_successful_submitted_at** — время, когда была решена последняя из решённых задач у этого пользователя (если одна задача решена многократно, то задача считается решённой в момент первого решения), иными словами, последний момент времени, когда у пользователя увеличилось количество решённых задач.

Строки следует сортировать по невозрастанию *problem_count*, затем по неубыванию *latest_successful_submitted_at*, затем по возрастанию *user_id*.

Пользователи, которые решили одинаковое количество задач (имеют равные *problem_count*) и имеют равные значения *latest_successful_submitted_at*, должны поделить одно место. Обратите внимание, что если несколько пользователей делят места, то в нумерации мест образуется разрыв. Например, если первое место делят два пользователя, то следующий пользователь должен получить место *3* (то есть последовательность мест имеет вид: *1,1,3*).
Внимательно ознакомьтесь с примерами вывода. Ваш запрос должен иметь в точности такой же вывод на примерах.

- **users** — пользователи системы (описываются двумя полями: *id* и *name*),
- **contests** — контесты в системе (описываются двумя полями: *id* и *name*),
- **problems** — задачи в системе, каждая задача принадлежит одному контесту (описываются тремя полями: *id*, *contest_id* и *code*, где *code* — это кодовое короткое название задачи),
- **submissions** — отосланные попытки решения задач, каждая попытка принадлежит одной задаче и одному пользователю (описываются 5 полями: *id*, *user_id*, *problem_id*, *success* и *submitted_at*, где *success* — это булевское значение была ли попытка успешной и *submitted_at* — дата-время, когда попытка была совершена).

Таким образом, **contests** и **problems** находятся в отношении «один ко многим», **submissions** и **users** находятся в отношении «многие к одному», **submissions** и **problems** находятся в отношении «многие к одному».

Изучите входные данные примера, чтобы подробно ознакомиться со схемой базы данных. Диаграмма ниже иллюстрирует схему базы данных.

### Входные данные
Входными данными в этой задаче является дамп базы данных. Вам он может быть полезен для ознакомления с состоянием базы данных для конкретного теста. В качестве решения вы должны отправить один **SQL**-запрос.

### Выходные данные
Ваш **SQL**-запрос должен вывести результаты соревнования с максимальным *id* в требуемом формате.

Внимательно ознакомьтесь с примерами вывода. Ваш запрос должен иметь в точности такой же вывод на примерах.

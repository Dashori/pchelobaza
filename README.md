## PCHELOBAZA

#### Слоган — Укуси меня пчела!
---- 
#### Цель работы — создать сервис для пчеловодов.
----
#### Краткий перечень функциональных требований:
- регистрация и аутентификация пчелочелов;
- добавление, редактирование и удаление контента (информация о себе/ферма/конференция/отзыв);
- просмотр ферм/конференций/отзывов/информации о пчелочелах;
- возможность оставить заявку на повышение прав.
-----

#### Use-case диаграмма системы
![use-case.svg](./images/use-case.svg)
-----

#### BPMN диаграмма основных бизнес-процессов
![bpmn.svg](./images/bpmn.svg)
-----


#### Примеры описания основных пользовательских сценариев:
- пользователь может зарегестрироваться как пчелочел;
- пользователь может зайти в систему и выйти из нее;
- пчелочел/пчеломастер/админ может изменять информацию о себе;
- пчелочел может отправить заявку, чтобы стать пчеломастером;
- админ может одобрить заявку пчелочела стать пчеломастером;
- пчелочел/пчеломастер может добавить/изменить информацию о своей ферме; 
- пчеломастер может добавить/изменить пчелоконференцию;
- пчелочел/мастер может записаться на пчелоконфцеренцию в случае, если на нее еще остались места или удалить свою запись;
- пчелочел/пчеломастер может добавить/удалить отзыв на конференцию.

-----

#### ER-диаграмма сущностей
![er.svg](./images/er.svg)

-----

#### Диаграмма БД
![db.png](./images/db.png)

-----

#### Компонентная диаграмма системы
![upper.svg](./images/upper.svg)

-----

#### Экраны будущего web-приложения на уровне черновых эскизов.
![first_ui_web.png](./images/first_ui_web.png)



### Нагрузочное тестирование 
#### 1 инстанс
```
Server Software:        nginx/1.24.0
Server Hostname:        localhost
Server Port:            8089

Document Path:          /api/v1/honey
Document Length:        3530 bytes

Concurrency Level:      1
Time taken for tests:   28.637 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      36740000 bytes
HTML transferred:       35300000 bytes
Requests per second:    349.20 [#/sec] (mean)
Time per request:       2.864 [ms] (mean)
Time per request:       2.864 [ms] (mean, across all concurrent requests)
Transfer rate:          1252.90 [Kbytes/sec] received

```
Connection Times (ms)
|          | min      | mean[+/-sd]|        | median   | max      |
|----------|----------|----------|----------|----------|----------|
|Connect:  | 0        | 0        |  1.3     |    0     |  129     |
| Processing:  |   2  |  3  | 4.3    |  2  |   169|
| Waiting:  |      2  |  2 |  3.6 |     2   |  169|
|Total:       |   2 |   3 |  4.5   |   3   |  169|

Percentage of the requests served within a certain time (ms)

| | |
|----------|----------|
  | 50%   |   3 |
  | 66%   |   3|
  | 75%   |   3|
 | 80%   |   3|
  | 90%   |   3|
  | 95%   |   3|
  | 98%   |   4|
  | 99%   |   4|
 | 100%   | 169 (longest request)|


![1.png](./images/1.png)


#### 3 инстанса
```
Server Software:        nginx/1.24.0
Server Hostname:        localhost
Server Port:            8089

Document Path:          /api/v1/honey
Document Length:        3530 bytes

Concurrency Level:      3
Time taken for tests:   16.646 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      36740000 bytes
HTML transferred:       35300000 bytes
Requests per second:    600.74 [#/sec] (mean)
Time per request:       4.994 [ms] (mean)
Time per request:       1.665 [ms] (mean, across all concurrent requests)
Transfer rate:          2155.39 [Kbytes/sec] received

```

Connection Times (ms)
|          | min      | mean[+/-sd]|        | median   | max      |
|----------|----------|----------|----------|----------|----------|
|Connect:  | 0        | 0        |   0.1    |    0     |   5     |
| Processing:|     3  |   5 | 2.2|      4 |     62|
| Waiting:   |     2  |   4 | 2.0|      4 |     42|
| Total:     |     3  |   5 | 2.2|      4 |     62|


Percentage of the requests served within a certain time (ms)
| | |
|----------|----------|
|  50% |      4 |
|  66% |      5 |
|  75% |      5 |
|  80% |      6 |
|  90% |      7 |
|  95% |      7 |
|  98% |      9 |
|  99% |     11 |
| 100% |     62 (longest request)| 

![3.png](./images/3.png)
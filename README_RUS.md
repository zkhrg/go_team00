# Team 00 - Go Boot camp

## Случайные пришельцы

## Содержание

1. [Глава I](#chapter-i) \
    1.1. [Основные правила](#general-rules)
2. [Глава II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Глава III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Глава IV](#chapter-iv) \
    4.1. [Task 00: Transmitter](#exercise-00-transmitter)
5. [Глава V](#chapter-v) \
    5.1. [Task 01: Anomaly Detection](#exercise-01-anomaly-detection)
6. [Глава VI](#chapter-vi) \
    6.1. [Task 02: Report](#exercise-02-report)
7. [Глава VII](#chapter-vii) \
    7.1. [Task 03: All Together](#exercise-03-all-together)
8. [Глава VIII](#chapter-viii) \
    8.1. [Reading](#reading)

<h2 id="chapter-i" >Глава I</h2>
<h2 id="general-rules" >Основные правила</h2>

* Твоя программа не должна закрываться неожиданно (выдавая ошибку при корректном вводе). Если это произойдет, твой проект будет считаться неработаспособным и получит 0 во время оценки.
* Мы рекомендуем тебе писать тесты для твоего проекта, даже если если они и не оцениваются. Это даст тебе возможность легко тестировать твою работу и работу твоих пиров. Ты убедишься что тесты очень полезны, во время защиты. Во время защиты ты свободен использовать свои тесты и/или тесты пира которого ты проверяешь.
* Отправляй свою работу в нужный git репозиторий. Работа будет оцениваться только из git репозитория.
* Если твой код использует сторонние зависимости, следует использовать [Go Modules](https://go.dev/blog/using-go-modules) для управления ими.

<h2 id="chapter-ii" >Глава II</h2>
<h2 id="rules-of-the-day" >Правила дня</h2>

* Пиши код только в `*.go` файлах и (в случае стронних зависимостей) `go.mod` + `go.sum`
* Твой код для этого задания должен собираться с использовния простого `go build`
* Все твои тесты должны запускаться стандартным вызовом `go test ./...`

<h2 id="chapter-iii" >Глава III</h2>
<h2 id="intro" >Intro</h2>

"Мы без понятия как это сделать!" - Луиза почти отчаялась. - "Корабль постоянно меняет частоту!"

Это была вторая бессонная ночь подряд для нее. Все указывало что эти инопрешеленцы пытались связаться с землянами, но главная проблема была это понимания друг друга.

Радиоприемник на столе внезапно включился: "Халперн докладывает. Наши агенты подключили кодирующий девайс к кораблю. Он собирает генерируемые частоты и может отправлять их в бинарном виде по сети. 

Луиза немедленно рванула поближе к экрану, но вероятно устройство передает данные только внутри зашированной военной сети, где никто из научных сотрудников не имеет доступа

На блестящего лингвиста было жалко смотреть. Но через пару минут он тряснула своей головой, словно она пытается прогрнать прочь мысли, и начала злобно вести переговоры с военными по рации. Одновременно ее рука яростно делала заметки на клочке бумаги.

Примерно через час после, она устало опустилась на кресло и просила рацию на стол. Затем оглядела команду.

"Тут кто-нибудь знает как программировать?" - она спросила. "Эти тупицы не хотят давать нам доступ к их устройству. Мы можем сделать что-то подобное и затем они согласятся подключить наш анализатор в их сеть. Но только если мы его протестируем сначала."

"Две или три руки поднялись неуверенно"

"Оки, оно использует нечто, называемое gRPC, что бы это не значило. Наш анализатор должен подключиться и принимать поток частот, смотреть на него и генерировать что-то типа отчета в PostgreSQL. Они мне дали формат данных"

Она встала и прошлась прошлась немного взад-вперед.

"Я поняла что анализирование полностью случайного сигнала это сложная задача. Жаль, что у нас нет больше информации"

И затем радиоприемным включился еще раз. И вещь, услышанная Луизой зажгло ее глаза энтузиазмом. Она глянула на команду и сказала еще одну вещь громким, триумфальным шепетом:

"Я думаю я знаю что делать! ЭТО ЖЕ НОРМАЛЬНОЕ РА`СПРЕДЕЛЕНИЕ!"

<h2 id="chapter-iv" >Глава IV</h2>
<h3 id="ex00">Task 00: Transmitter</h3>

"So, we have to reimplement this military device's protocol on our own." - Louise said. - "I've already mentioned that it uses gRPC, so let's do that."

She showed a basic schema of data types. Looks like each message consists of just three fields - 'session_id' as a string, 'frequency' as a double and also a current timestamp in UTC.

We don't know much about distribution here, so let's implement it in such way that whenever new client connects [expected value](https://en.wikipedia.org/wiki/Expected_value) and [standard deviation](https://en.wikipedia.org/wiki/Standard_deviation)" are picked at random. For this experiment, let's pick mean from [-10, 10] interval and standard deviation from [0.3, 1.5].

On each new connection server should generate a random UUID (sent as session_id) and new random values for mean and STD. All generated values should be written to a server log (stdout or file). After that it should send a stream of entries with fields explained above, where for each message 'frequency' would be a value picked at random (sampled) from a normal distribution with these standard deviation and expected value.

It is required to describe the schema in a *.proto* file and generate the code itself from it. Also, you shouldn't modify generated code manually, just import it.

<h2 id="chapter-v" >Глава V</h2>
<h3 id="ex01">Task 01: Anomaly Detection</h3>

"Now to the interesting part! While others are working on gRPC server, let's think of a client. I expect that gRPC client should be handled by the same guys writing the server to test it, so let's focus on a different thing. We need to detect anomalies in a frequency distribution!"

So, you know you're getting a stream of values. With each new incoming entry from a stream your code should be able to approximate mean and STD from the random distribution generated on a server. Of course it's not really possible to predict it looking only on 3-5 values, but after 50-100 it should be precise enough. Keep in mind that mean and STD are generated for each new connection, so you shouldn't restart the client during the process. Also, values shouldn't keep piling up in memory, so you may consider using sync.Pool for easy reuse.

While working on this task, you can temporarily forget about gRPC and test the code by just sending it a sequence of values to stdin.

Your client code should write into a log periodically, how many values are processed so far as well as predicted values of mean and STD.

After some time, when your client decides that the predicted distribution parameters are good enough (feel free to choose this moment by yourself), it should switch automatically into an Anomaly Detection stage. Here there is one more parameter which comes into play - an *STD anomaly coefficient*. So, your client should accept a command-line parameter (let it be '-k') with a float-typed coefficient.

An incoming frequency is considered an anomaly, if it differs from the expected value by more than *k \* STD* to any side (to the left or to the right, as the distribution is symmetric). You can read more about how it works by following links from Глава 4.

For now you should just write found anomalies into a log.

<h2 id="chapter-vi" >Глава VI</h2>
<h3 id="ex02">Task 02: Report</h3>

"As general knows nothing about our *sciency gizmo*, let's store all anomalies that we encounter in a database and then he'll be able to look at it through some interface they have" - Louise seems to be a lot more concerned about the data rather than the general.

So, let's learn how to write data entries to PostgreSQL. Usually it is considered a bad practice to just write plain SQL queries in code when dealing with highly secure environments (you can read about SQL Injections by following links from Глава 4). Let's use an ORM. In case of PostgreSQL there are two most obvious choices (these links are below as well), but you can choose any other. The main idea here is to not have any strings with SQL code in your sources.

You'll have to describe your entry (session_id, frequency and a timestamp) as a structure in Go and then use it together with ORM to map it into database columns.

<h2 id="chapter-vii" >Глава VII</h2>
<h3 id="ex03">Task 03: All Together</h3>

Okay, so when we have a transmitter, receiver, anomaly detection and ORM, we can plug things into one another and merge them into a full project.

So, if you start a server and a client (PostgreSQL should be already running on your machine), your client will connect to a server and get a stream of entries which it will then:

- First, use for a distribution reconstruction (mean/STD)
- Second, after some time start detecting anomalies based on supplied STD anomaly coefficient (I suggest you pick it big enough for this experiment, so anomalies wouldn't happen too frequently)
- Third, all anomalies should be written into a database in PostgreSQL using ORM

If Louise is right, these anomalies could be the key to a first contact with the aliens. But it is also a pretty direct approach for cases when you need to detect anomalies on a stream of data, which Go can be efficiently used for.

<h2 id="chapter-viii" >Глава VIII</h2>
<h3 id="reading">Reading</h3>

[Normal distribution](https://en.wikipedia.org/wiki/Normal_distribution) <br>
[68-95-99.7 rule](https://en.wikipedia.org/wiki/68%E2%80%9395%E2%80%9399.7_rule) <br>
[SQL Injections](https://en.wikipedia.org/wiki/SQL_injection) <br>
[go-pg](https://github.com/go-pg/pg) <br>
[GORM](https://gorm.io/index.html) <br>



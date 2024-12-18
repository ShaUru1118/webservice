WEBSERVICE

WebService - проект, позволяющий с помощью POST запросов решать не сложные математические выражения.

---

**КАК ПОЛЬЗОВАТЬСЯ?**

Для начала установите язык программирования GO (https://go.dev/dl/).
Далее скачайте архив с кодом из последнего релиза.
Распакуйте в удобную для вас папку, откойте ее и запустите _командную строку_ из этой папки.

Используя команду

    go run .\main\webservice.go

вы запустите сервер, и он сможет обрабатыввать ваши запросы.
После этого вы сможете, используя curl из любой _командной строки_ (кроме той, где вы запустили сервер), обращаться к нему.

В ответ вы получите результат работы:

1. Ответ на ваше математическое выражение ({"result":"6"}).
2. Ошибку:
     *  **"Internal server error : r.Body.Read(data)"** - не смог прочитать Body в r *http.Request.
     *  **"Internal server error : json.Unmarshal(body, &datago)"** - не смог преобразовать структуру json в структуру go.
     *  **"Expression is not valid : simplecalc.Calc(expretion)"** - не смог посчитать ответ, так как выражение написано неправильно.

Чтобы отправить запрос, воспользуйтесь командой

    curl -X POST -H 'Content-Type:application/json' -d '\{\"expression\":\"*YOUR_EXPRESSION*\"\}' http://localhost:8080/api/v1/calculate

где ***YOUR_EXPRESSION*** - ваше выражение

---

ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ.
  1. Команда

          curl -X POST -H 'Content-Type:application/json' -d '\{\"expression\":\"2+2\*2\"\}' http://localhost:8080/api/v1/calculate
     вернет {"result":"6"}, т. к. выражение задано правильно и других ошибок нет.

  3. Команда

          curl -X POST -H 'Content-Type:application/json' -d '\{\"expression\":\"+2\*2\"\}' http://localhost:8080/api/v1/calculate
     вернет {"error":"Expression is not valid : simplecalc.Calc(expretion)"}, т. к. выражение составлено не правильно.

  5. Команда

          curl -X POST -H 'Content-Type:application/json' -d '{"a":"a"}' http://localhost:8080/api/v1/calculate
     вернет {"error":"Internal server error : json.Unmarshal"}, т. к. не передаётся "expression", т. е. программа не может найти ключ к выражению.




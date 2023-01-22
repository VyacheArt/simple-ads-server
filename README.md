# Простой рекламный сервер на GoLang

Посмотреть видео по проекту можете тут https://youtu.be/27WPASOQs2w

В этом репозитории находится простейший рекламный сервер на Go. Несмотря на его простоту, он использует стек, который используется на самом настоящем высоконагруженном рекламном сервере: fasthttp, GeoIP от MaxMind и [mssola/user_agent](https://github.com/mssola/user_agent).

## Обновления!

* Совсем недавно вышло видео про __[реализацию высоконагруженной статистики](https://www.youtube.com/watch?v=eAmblLikTgo)__

  * [Ветка со статистикой на MySQL](https://github.com/VyacheArt/simple-ads-server/tree/mysql-stats)
  * [Ветка со статистикой на ClickHouse](https://github.com/VyacheslavGoryunov/simple-ads-server/tree/clickhouse-stats)

* А в ролике про __[Prometheus метрики в GoLang](https://youtu.be/6pQQw-qEoCo)__ я рассказываю про метрики на примере этого рекламного сервера
  * [Ветка с метриками на Prometheus](https://github.com/VyacheArt/simple-ads-server/tree/prometheus)

## Настройка

Не забудьте перед запуском положить GeoLite2-Country.mmdb в корень проекта. В остальном всё должно сразу работать :) 

## Стек

* Быстрый HTTP сервер с использованием [valyala/fasthttp](https://github.com/valyala/fasthttp)
* Парсинг User-Agent через [mssola/user_agent](https://github.com/mssola/user_agent)
* Взаимодействие с GeoIP через [oschwald/geoip2-golang](https://github.com/oschwald/geoip2-golang)


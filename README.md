# Go by Practice
-------------------------------
## 不积跬步无以至千里，Go练手实战
-------------------------------

* Weathet-Cli(命令行天气)
```
- 命令行效果
$ go run weather-cli.go
查询的城市：北京
┌────────────────┬─────────┬────────────────┬────────────────┬────────┬─────┬───────────┬────────┬────────┬──────────────────────────────────────┐
│ Date           │ Sunrise │ High           │ Low            │ Sunset │ Aqi │ Fx        │ Fl     │ Type   │ Notice                               │
├────────────────┼─────────┼────────────────┼────────────────┼────────┼─────┼───────────┼────────┼────────┼──────────────────────────────────────┤
│ 28日星期一     │ 04:51   │ 高温 30.0℃     │ 低温 16.0℃     │ 19:33  │ 51  │ 西北风    │ 4-5级  │ 晴     │ 愿你拥有比阳光明媚的心情             │
│ 29日星期二     │ 04:50   │ 高温 30.0℃     │ 低温 16.0℃     │ 19:34  │ 49  │ 西风      │ <3级   │ 多云   │ 阴晴之间，谨防紫外线侵扰             │
│ 30日星期三     │ 04:50   │ 高温 33.0℃     │ 低温 19.0℃     │ 19:34  │ 39  │ 西南风    │ <3级   │ 晴     │ 愿你拥有比阳光明媚的心情             │
│ 31日星期四     │ 04:49   │ 高温 34.0℃     │ 低温 20.0℃     │ 19:35  │ 41  │ 西南风    │ <3级   │ 晴     │ 愿你拥有比阳光明媚的心情             │
│ 01日星期五     │ 04:49   │ 高温 35.0℃     │ 低温 20.0℃     │ 19:36  │ 74  │ 西南风    │ <3级   │ 晴     │ 愿你拥有比阳光明媚的心情             │
└────────────────┴─────────┴────────────────┴────────────────┴────────┴─────┴───────────┴────────┴────────┴──────────────────────────────────────
```
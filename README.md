# Go by Practice
-------------------------------
## 不积跬步无以至千里，Go练手实战
-------------------------------

* Weather-Cli(命令行天气)
    原作者 [@yangwenmai](https://github.com/yangwenmai/wg)    
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

* Table(控制台输出表格)
    原作者 [@modood](https://github.com/modood/table)
    ！！！ 完整版请参照原版 ！！！
- [Go反射详解](http://www.cnblogs.com/golove/p/5909541.html)
```
*测试代码直接写在源码里面了*
-效果展示
$ go run table.go
┌───────────┬──────────┬──────────────────┐
│ Name      │ Sigil    │ Motto            │
├───────────┼──────────┼─────────────────┤┼
│ Stark     │ direwolf │ Winter is coming │
│ Targaryen │ dragon   │ Fire and Blood   │
│ Lannister │ lion     │ Hear Me Roar     │
└───────────┴──────────┴──────────────────┘
```

* Emoji(控制台输出表情)
    原作者 [@kyokomi](https://github.com/kyokomi/emoji)
```
- 效果展示(!!!直接用gitbash测试，windows不支持!!!)
$ go test
🍺   ビール!!!
PASS
ok      github.com/BeanWei/Go-by-Practice/Emoji 0.078s
```

* 时间解析Golang版
```
- 效果展示
    "1分钟前",
    "54分钟前",
    "1小时前",
    "13小时前",
    "昨天",
    "前天",
    "06-25",
    "12-20"
  /*-----------------------------*/
  2018/07/01 18:57:37 2018-07-01 18:56:37
  2018/07/01 18:57:37 2018-07-01 18:03:37
  2018/07/01 18:57:37 2018-07-01 17:57:37
  2018/07/01 18:57:37 2018-07-01 05:57:37
  2018/07/01 18:57:37 2018-06-30 18:57:37
  2018/07/01 18:57:37 2018-06-29 18:57:37
  2018/07/01 18:57:37 2018-06-25 00:00:00
  2018/07/01 18:57:37 2018-12-20 00:00:00
```
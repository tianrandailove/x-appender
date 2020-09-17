# X-APPENDER

x-appender 是一款用golang编写的日志适配器，旨在解决某些应用程序无法持久化日志的问题；可以实现日志分级持久化、日志自动分割；目前该适配器只适配了golang编写的程序。

## 特点

- [x] 日志持久化
- [x] 日志分级输出
- [x] 日志自动切割
- [x] 日志自动清理

## 使用

将应用程序的标准输出或标准错误输出给x-appender即可，例如hyperledger-fabric peer:

```shell
peer node start 2>&1 | x-appender
```

## 配置

x-appender 所有配置都是用环境变量进行控制，具体如下：

- XAPPENDER_LOG_NAME_FORMAT 日志文件名格式，如：%**Y**-%m-%d
- XAPPENDER_LOG_ROOT_PATH 日志持久化目录，如：/var/logs
- XAPPENDER_LOG_MAX_AGE 日志最大保存时间，单位为天，如：30
- XAPPENDER_LOG_ROTATION 日志切割时间，单位为小时，如：24


# hlog

#### 介绍
logger


使用说明：
```golang

package main

import "github.com/sirupsen/logrus"

func Now() (l *logrus.Logger) {
    defer func() {
        l = logrus.StandardLogger()
    }()
    logrus.SetLevel(logrus.InfoLevel) // 设置日志级别
    return
}

```
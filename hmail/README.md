# 邮件收发

## 使用IMAP收取邮件

使用方法

```
mail,err = hmail.ImapFetchSince(auth,time)  // 收指定时间内的邮件
```

## 使用SMTP发送邮件

使用方法

```
err = hmail.SmtpSend(mail,auth)
```

<br>

腾讯企业邮箱

```
接收服务器：    
imap.exmail.qq.com(使用SSL，端口号993)   
发送服务器：   
smtp.exmail.qq.com(使用SSL，端口号465)  
```

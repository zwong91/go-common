<!-- package=passport.service.identify.v1 -->
- [/passport.service.identify.v1.Identify/GetCookieInfo](#passport.service.identify.v1.IdentifyGetCookieInfo)  CookieInfo identify info by cookie.
- [/passport.service.identify.v1.Identify/GetTokenInfo](#passport.service.identify.v1.IdentifyGetTokenInfo)  TokenInfo identify info by token.

## /passport.service.identify.v1.Identify/GetCookieInfo
### CookieInfo identify info by cookie.

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|cookie|否|string| all user cookie of HTTP request example: 'SESSDATA=DEE4597D836A5A9DA29DFC1AB1EFFDEB;sid=exampleSID'|

#### 响应

```javascript
{
    "code": 0,
    "message": "ok",
    "data": {
        //  用户是否登录
        "is_login": true,
        //  user mid
        "mid": 0,
        //  cookie csrf
        //  when token reqest this field is empty
        "csrfToken": "",
        //  expire time(unix timestamp)
        "expires": 0
    }
}
```


## /passport.service.identify.v1.Identify/GetTokenInfo
### TokenInfo identify info by token.

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|token|是|string| user access token|
|buvid|否|string| buvid|

#### 响应

```javascript
{
    "code": 0,
    "message": "ok",
    "data": {
        //  用户是否登录
        "is_login": true,
        //  user mid
        "mid": 0,
        //  cookie csrf
        //  when token reqest this field is empty
        "csrfToken": "",
        //  expire time(unix timestamp)
        "expires": 0
    }
}
```


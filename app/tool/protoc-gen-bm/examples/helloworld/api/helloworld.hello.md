<!-- package=helloworld.v1 -->
- [/hello](#hello)  Sends a greeting
- [/echo](#echo) 

## /hello
### Sends a greeting

#### 方法：GET

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|name|否|string||

#### 响应

```javascript
{
    "code": 0,
    "message": "ok",
    "data": {
        "message": ""
    }
}
```


## /echo
### 无标题

#### 方法：POST

#### 请求参数

|参数名|必选|类型|描述|
|:---|:---|:---|:---|
|content|否|string||

#### 响应

```javascript
{
    "code": 0,
    "message": "ok",
    "data": {
        "content": ""
    }
}
```


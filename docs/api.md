# API for Html2X

## Html to pdf

### `POST` https://saas.xinjingcai.com/wk/html-to-pdf

### 请求参数

> Request Header

  |标签|类型|必填|说明|
  |:---:|:---:|:---:|:---|
  |Content-Type|string|是|必须为 `application/json`|

> Request Body `JSON` `最外层结构为：Object`

  |参数名|类型|必填|说明|
  |:---:|:---:|:---:|:---|
  |html|string|是|html 的字符串|
  |args|object|是|配置参数，key:value 的格式，详见：[wkhtmltopdf usage](https://overnaive.github.io/Html2X/docs/wkhtmltopdf.txt)|

> 示例

```json
{
  "html": "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>这是一个HTML5的网页</title></head><body><p>Hello HTML5</p></body></html>",
  "args": {
    "page-size": "A4",
    "no-debug-javascript": "",
    "minimum-font-size": "20"
  }
}
```

### 返回参数

> Response Header

  |标签|类型|必填|说明|
  |:---:|:---:|:---:|:---|
  |Content-Type|string|是|`application/pdf`|

> Response Body

直接返回 pdf 的内容。

### 异常

通过 `HTTP Status Codes` 来区分请求成功与否，请查看[HTTP状态码列表](https://httpstatuses.com/) 。

## Html to img

### `POST` https://saas.xinjingcai.com/wk/html-to-img


### 请求参数

> Request Header

  |标签|类型|必填|说明|
  |:---:|:---:|:---:|:---|
  |Content-Type|string|是|必须为 `application/json`|

> Request Body `JSON` `最外层结构为：Object`

  |参数名|类型|必填|说明|
  |:---:|:---:|:---:|:---|
  |html|string|是|html 的字符串|
  |args|object|是|配置参数，key:value 的格式，详见：[wkhtmltoimage usage](https://overnaive.github.io/Html2X/docs/wkhtmltoimage.txt)|
  |ext|string|否|图片后缀，默认值为 `jpg`|

> 示例

```json
{
  "html": "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>这是一个HTML5的网页</title></head><body><p>Hello HTML5</p></body></html>",
  "args": {
    "crop-w": "200",
    "crop-h": "200",
    "no-debug-javascript": ""
  }
}
```

### 返回参数

> Response Header

  |标签|类型|必填|说明|
  |:---:|:---:|:---:|:---|
  |Content-Type|string|是|`application/jpeg`, `application/png`等|

> Response Body

直接返回图片的内容。

### 异常

通过 `HTTP Status Codes` 来区分请求成功与否，请查看[HTTP状态码列表](https://httpstatuses.com/) 。

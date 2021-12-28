### 简要描述
#### 备忘录增删改查

### 请求地址
|  方式  | 请求方式 | 请求URL |
| :-----| ----: | :----: |
| 增      | POST | http://localhost:8000/api/v1/todo |
| 删单个   | GET | http://localhost:8000/api/v1/todo/id |
| 删所有 | GET  |  http://localhost:8000/api/v1/todo/all |
| 删所有完成/代办 | GET | http://localhost:8000/api/v1/todo/status |
| 改单个完成/代办 | PUT | http://localhost:8000/api/v1/todo/one |
| 改所有完成/代办 | PUT | http://localhost:8000/api/v1/todo/all |
| 查所有 | GET | http://localhost:8000/api/v1/todo/all |
| 查完成/代办 | GET | http://localhost:8000/api/v1/todo/status |
| 关键词查 | GET |http://localhost:8000/api/v1/todo/content |

### 参数名
#### 增
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| nil  | nil | nil |

#### 删单个
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
|  id |  int  |  要删除的id |

#### 删所有
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| nil  | nil | nil |

#### 删所有代办/完成
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| status | bool | 要删除的状态 |

#### 改单个完成/代办
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
|  ID | int | 要更改的id |

#### 改所有完成/代办
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| status | bool | 要修改的状态 |

#### 查所有
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| page | int | 要查找的页码 |
/*每页最多20条*/

#### 查完成/代办
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| status | bool | 要查找的状态 |
| page | int | 查找的页码 |

#### 关键词查找
| 参数名  | 类型   | 说明 |
|  :----| -----: | :-----:|
| keyword | string | 关键词 |
| page | int | 要查找的页码 |

### 返回实例
#### 增
```{
    "data": {
        "item": {
            "content": "",
            "created_time": "2021-12-28T15:47:52.5241531+08:00",
            "end_time": 0,
            "id": 101,
            "start_time": "2021-12-28T15:47:52.5241531+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 1
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 删单个
```{
    "data": {
        "item": {
            "content": "",
            "created_time": "2021-12-28T14:49:53+08:00",
            "end_time": "2021-12-28T15:52:02.0170916+08:00",
            "id": 90,
            "start_time": "2021-12-28T14:49:53+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 1
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 删所有完成/代办
```{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T14:49:53+08:00",
            "end_Time": "2021-12-28T15:53:23.0925913+08:00",
            "id": 91,
            "start_Time": "2021-12-28T14:49:53+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 3
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T14:49:54+08:00",
            "end_Time": "2021-12-28T15:53:23.0970448+08:00",
            "id": 92,
            "start_Time": "2021-12-28T14:49:54+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 3
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "操你妈",
            "created_At": "2021-12-28T15:01:03+08:00",
            "end_Time": "2021-12-28T15:53:23.1018025+08:00",
            "id": 93,
            "start_Time": "2021-12-28T15:01:03+08:00",
            "status": true,
            "title": "1234",
            "view": 0
        },
        "total": 3
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 删所有

```
{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T14:49:50+08:00",
            "end_Time": "2021-12-28T15:54:21.6048375+08:00",
            "id": 83,
            "start_Time": "2021-12-28T15:08:05+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T14:49:50+08:00",
            "end_Time": "2021-12-28T15:54:21.6084046+08:00",
            "id": 84,
            "start_Time": "2021-12-28T14:49:50+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T14:49:52+08:00",
            "end_Time": "2021-12-28T15:54:21.6130945+08:00",
            "id": 88,
            "start_Time": "2021-12-28T14:49:52+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T14:49:52+08:00",
            "end_Time": "2021-12-28T15:54:21.6171024+08:00",
            "id": 89,
            "start_Time": "2021-12-28T14:49:52+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "操你妈",
            "created_At": "2021-12-28T15:01:08+08:00",
            "end_Time": "2021-12-28T15:54:21.6210979+08:00",
            "id": 97,
            "start_Time": "2021-12-28T15:01:08+08:00",
            "status": false,
            "title": "1234",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T15:06:18+08:00",
            "end_Time": "2021-12-28T15:54:21.6256957+08:00",
            "id": 98,
            "start_Time": "2021-12-28T15:08:05+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T15:06:19+08:00",
            "end_Time": "2021-12-28T15:54:21.6306941+08:00",
            "id": 99,
            "start_Time": "2021-12-28T15:06:19+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T15:06:20+08:00",
            "end_Time": "2021-12-28T15:54:21.6347404+08:00",
            "id": 100,
            "start_Time": "2021-12-28T15:06:20+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "items": {
            "content": "",
            "created_At": "2021-12-28T15:47:53+08:00",
            "end_Time": "2021-12-28T15:54:21.6392951+08:00",
            "id": 101,
            "start_Time": "2021-12-28T15:47:53+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 9
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 改单个
```
{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 103,
            "start_time": "2021-12-28T15:56:01.9749922+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 1
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 改全部代办/完成
```
{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:32+08:00",
            "end_time": 0,
            "id": 102,
            "start_time": "2021-12-28T15:57:10.1052521+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 2
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 103,
            "start_time": "2021-12-28T15:57:10.1088482+08:00",
            "status": false,
            "title": "",
            "view": 0
        },
        "total": 2
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
```
{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:32+08:00",
            "end_time": 0,
            "id": 102,
            "start_time": "2021-12-28T15:57:27.4056554+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 103,
            "start_time": "2021-12-28T15:57:27.4097321+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 104,
            "start_time": "2021-12-28T15:57:27.4135488+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 105,
            "start_time": "2021-12-28T15:57:27.4182581+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 查完成/代办
```
{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:32+08:00",
            "end_time": 0,
            "id": 102,
            "start_time": "2021-12-28T15:57:27+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 103,
            "start_time": "2021-12-28T15:57:27+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 104,
            "start_time": "2021-12-28T15:57:27+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "",
            "created_At": "2021-12-28T15:55:33+08:00",
            "end_time": 0,
            "id": 105,
            "start_time": "2021-12-28T15:57:27+08:00",
            "status": true,
            "title": "",
            "view": 0
        },
        "total": 4
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```
#### 关键字搜索
```
{
    "data": {
        "item": {
            "content": "圣诞快乐",
            "created_At": "2021-12-28T16:02:08+08:00",
            "end_time": 0,
            "id": 107,
            "start_time": "2021-12-28T16:02:08+08:00",
            "status": false,
            "title": "w2ol",
            "view": 0
        },
        "total": 3
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "元旦快乐",
            "created_At": "2021-12-28T16:02:16+08:00",
            "end_time": 0,
            "id": 108,
            "start_time": "2021-12-28T16:02:16+08:00",
            "status": false,
            "title": "w2ol",
            "view": 0
        },
        "total": 3
    },
    "error": "",
    "msg": "ok",
    "status": 200
}{
    "data": {
        "item": {
            "content": "新年快乐",
            "created_At": "2021-12-28T16:02:23+08:00",
            "end_time": 0,
            "id": 109,
            "start_time": "2021-12-28T16:02:23+08:00",
            "status": false,
            "title": "w2ol",
            "view": 0
        },
        "total": 3
    },
    "error": "",
    "msg": "ok",
    "status": 200
}
```

### 返回参数说明
|  参数名  | 类型 | 说明 |
| :-----| ----: | :----: |
| status | int  | 状态码 |
| data  | map | 数据 |
| id | int | ID |
| title | string | 标题 |
| content | string | 文本内容 |
| view | int | 访问次数 |
| status | int | 完成状态 |
| created_at | Time | 创建时间 |
| start_time | Time | 完成时间 |
| end_time | Time | 删除时间 |
| msg | string | 返回信息 |
| error | error | 错误 | 

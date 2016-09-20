# 通知系统

## 事件

### 事件列表

 * 网址: GET /event/list
 * 网址参数: 
   * `page`: 页码(可选，默认为1)
   * `size`: 每页数据量(可选，默认为10)
 * 响应内容:
```
{
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
    "data": {
        "list": [
            {
                "id": 1,
                "name": "add",
                "target": "magazine",
                "target_id": 1,
                "memo": "",
                "created": 0,
                "finished": 0
            }
        ],
        "pages": 1,
        "page": 1,
        "total": 1
    }
}
```

### 添加新事件

 * 网址: POST /event/add
 * POST参数: 
   * `Name`: 行为名称:add,del,edit
   * `Target`: 目标:magazine,book,article
   * `TargetId`: 目标id。比如Target为magazine时，TargetId则应该是杂志的id，依次类推
   * `Memo`: 说明

 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```

### 修改事件

 * 网址: POST /event/edit?id=1
 * 网址参数:
   * `id`: id为大于0的整数
 * POST参数: 
   * `Name`: 行为名称:add,del,edit
   * `Target`: 目标:magazine,book,article
   * `TargetId`: 目标id。比如Target为magazine时，TargetId则应该是杂志的id，依次类推
   * `Memo`: 说明

 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```

### 删除事件

 * 网址: GET /event/del?id=1
 * 网址参数:
   * `id`: id为大于0的整数
 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```
## 客户端（通知接收者）

### 客户端列表

 * 网址: GET /client/list
 * 网址参数: 
   * `page`: 页码(可选，默认为1)
   * `size`: 每页数据量(可选，默认为10)
 * 响应内容:
```
{
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
    "data": {
        "list": [
            {
                "id": 1,
                "name": "test",
                "recv_url": "",
                "type_magazine": 0,
                "type_book": 0,
                "type_article": 0,
                "event_add": 0,
                "event_edit": 0,
                "event_del": 0,
                "pre_hook": "",
                "hook_param": ""
            },
            ...
        ],
        "pages": 1,
        "page": 1,
        "total": 2
    }
}
```

### 添加客户端

 * 网址: POST /client/add
 * POST参数: 
   * `Name`: 客户端名称
   * `RecvUrl`: 接收通知的网址（支持占位符：{type},{event},{id}）
   * `TypeMagazine`: 是否(1/0)接收杂志通知
   * `TypeBook`: 是否(1/0)接收图书通知
   * `TypeArticle`: 是否(1/0)接收文章通知
   * `EventAdd`: 是否(1/0)接收新增通知
   * `EventEdit`: 是否(1/0)接收修改通知
   * `EventDel`: 是否(1/0)接收删除通知
   * `PreHook`: 通知之前的处理钩子
   * `HookParam`: 钩子参数。多个参数用半角逗号隔开

 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```

### 修改客户端

 * 网址: POST /client/edit?id=1
 * 网址参数:
   * `id`: id为大于0的整数
 * POST参数: 
   * `Name`: 客户端名称
   * `RecvUrl`: 接收通知的网址（支持占位符：{type},{event},{id}）
   * `TypeMagazine`: 是否(1/0)接收杂志通知
   * `TypeBook`: 是否(1/0)接收图书通知
   * `TypeArticle`: 是否(1/0)接收文章通知
   * `EventAdd`: 是否(1/0)接收新增通知
   * `EventEdit`: 是否(1/0)接收修改通知
   * `EventDel`: 是否(1/0)接收删除通知
   * `PreHook`: 通知之前的处理钩子
   * `HookParam`: 钩子参数。多个参数用半角逗号隔开

 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```

### 删除客户端

 * 网址: GET /client/del?id=1
 * 网址参数:
   * `id`: id为大于0的整数
 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```
## 通知

### 通知列表

 * 网址: GET /notice/list
 * 网址参数: 
   * `page`: 页码(可选，默认为1)
   * `size`: 每页数据量(可选，默认为10)
 * 响应内容:
```
{
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
    "data": {
        "list": [
            {
                "id": 1,
                "client_id": 1,
                "event_id": 1,
                "created": 0,
                "retry": 0,
                "finished": 0,
                "client": {
                    "id": 1,
                    "name": "test",
                    "recv_url": "",
                    "type_magazine": 0,
                    "type_book": 0,
                    "type_article": 0,
                    "event_add": 0,
                    "event_edit": 0,
                    "event_del": 0,
                    "pre_hook": "",
                    "hook_param": ""
                },
                "event": {
                    "id": 1,
                    "name": "add",
                    "target": "magazine",
                    "target_id": 1,
                    "memo": "",
                    "created": 0,
                    "finished": 0
                }
            },
            ...
        ],
        "pages": 1,
        "page": 1,
        "total": 1
    }
}
```

### 添加新通知

 * 网址: POST /notice/add
 * POST参数: 
   * `ClientId`: 客户端id
   * `EventId`: 事件id

 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```

### 修改通知

 * 网址: POST /notice/edit?id=1
 * 网址参数:
   * `id`: id为大于0的整数
 * POST参数: 
   * `ClientId`: 客户端id
   * `EventId`: 事件id

 * 响应内容:
 ```
 {
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
 }
 ```

### 删除通知

 * 网址: GET /notice/del?id=1
 * 网址参数:
   * `id`: id为大于0的整数
 * 响应内容:
```
{
    "status": 1,   //状态。成功为1，失败为0，未登录为-1
    "message": "", //提示信息
}
```

# 接口文档

## 用户管理

### 创建用户
```
POST /api/users
@request
{
    name: (string),    // 名称
    age: (int),        // 年龄
    address: (string), // 地址
}
@response
{
    code: (int),       // 响应状态
    message: (string), // 响应消息
    data: {
        id: (int),         // 用户ID
        name: (string),    // 名称
        age: (int),        // 年龄
        address: (string), // 地址
    },  // 响应数据
}
```

### 删除用户
```
DELETE /api/users
@request
{
    id: (int),         // 用户ID
}
@response
{
    code: (int),       // 响应状态
    message: (string), // 响应消息
    data: {},          // 响应数据
}
```

### 修改用户
```
PUT /api/users
@request
{
    id: (int),         // 用户ID
    name: (string),    // 名称
    age: (int),        // 年龄
    address: (string), // 地址
}
@response
{
    code: (int),       // 响应状态
    message: (string), // 响应消息
    data: {},          // 响应数据
}
```

### 用户列表
```
GET /api/users
@request
{
    page: (int),  // 页码
    limit: (int), // 每页数量
}
@response
{
    code: (int),       // 响应状态
    message: (string), // 响应消息
    data: {
        total: (int),  // 数据总数
        items: [
            {
                id: (int),         // 用户ID
                name: (string),    // 名称
                age: (int),        // 年龄
                address: (string), // 地址
            }
        ]       // 数据列表
    },          // 响应数据
}
```
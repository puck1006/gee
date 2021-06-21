# 实现一个简易的gin框架

## version1

1. 封装Engine结构体 结构体中有router属性来存放所有路由
2. 实现ServeHTTP 方法来拦截所有的请求
3. 实现GET 以及 POST请求

## version2

1. 封装Content结构体 来作为所有路由处理函数的执行上下文
2. 为context添加 
- PostForm
- Query
- Status
- SetHeader
- String
- JSON
- Data
- HTML
方法来强化context功能
3. router单独抽离出来,方便扩展




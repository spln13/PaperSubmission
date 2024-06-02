# Middleware层
### jwt
负责对前端传来对请求鉴权，判断请求是否满足权限要求，并根据请求token解析出用户身份

在routers中加入该中间件，会解析请求中cookie中的token字段，获得用户身份`user_id`并存储到gin.context中。
在controller层中对应handler函数使用`userID, ok := context.MustGet("user_id).(int64)`获取user_id，
因为context.MustGet()返回的是`interface{}`类型，所以需要使用`.(int64)`进行断言, ok为是否断言成功，因为jwt中代码写死了不会出错。
直接使用`userID, _ := context.MustGet("user_id).(int64)`即可。
### encryption
对带有password的请求中的password进行sha256加密，使密码在服务端不是明文形式存储

# Middleware层
### jwt
负责对前端传来对请求鉴权，判断请求是否满足权限要求，并根据请求token解析出用户身份

### encryption
对带有password的请求中的password进行sha256加密，使密码在服务端不是明文形式存储

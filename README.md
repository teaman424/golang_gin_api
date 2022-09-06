#  嘗試以golang + gin 建置restful api
* 以Swagger建立文件
    * github.com/swaggo/swag
* Create Docker File
* 使用JWT建置Access Token and Refresh Token
    * github.com/dgrijalva/jwt-go
* 使用Redis紀錄token，以完成revoke and refresh token功能
    * github.com/go-redis/redis/v9
* 加入 Mysql+gorm 紀錄User Info and Login
    * gorm.io/driver/mysql
	* gorm.io/gorm
* 建立MVC模式

### 環境參數
```
cp .env.example .env
vi .env //修改對應的參數
```
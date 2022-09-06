package jwt_auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// custom claims
type Claims struct {
	Account   string `json:"account"`
	Uuid      string
	TokenUuid string
	jwt.StandardClaims
}

type AuthToken struct {
	AccessToken  string `json:"accessToken"  example:"dkdke3klwlwkkf..."`
	AccessExp    int    `json:"accessExp" example:"600"`
	RefreshToken string `json:"refreshToken" example:"dkdke3klwlwkkf..."`
	RefreshExp   int    `json:"refreshExp" example:"86400"`
	TokenType    string `json:"tokenType" example:"Bearer"`
}

var rdb *redis.Client
var jwtAccessSecret []byte
var jwtRefreshSecret []byte

func Init() {
	//fmt.Println("jwt_auth init!!!!!!")
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtAccessSecret = []byte(os.Getenv("ACCESS_SECRET"))
	jwtRefreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

	rdb = NewClient(ctx)

}

func NewClient(ctx context.Context) *redis.Client { // 實體化redis.Client 並返回實體的位址
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		//Addr: "127.0.0.1:" + os.Getenv("REDIS_PORT"),
		//Password: "", // no password set
		//DB:       0,  // use default DB
	})

	_, err_redis := rdb.Ping(ctx).Result()
	if err_redis != nil {
		log.Fatal("Error rdb" + err_redis.Error())
	}
	return rdb
}

// validate JWT
func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not access token",
		})
		c.Abort()
		return
	}

	token := strings.Split(auth, "Bearer ")[1]
	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtAccessSecret, nil
	})
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		userid, err := rdb.Get(c, claims.TokenUuid+"-accuess").Result()
		if err != nil {
			fmt.Println("rdb not find  key!")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "access token error",
			})
			c.Abort()
			return
		}
		if userid != claims.Uuid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "access token error!",
			})
			c.Abort()
			return
		}
		c.Set("account", claims.Account)
		c.Set("uuid", claims.Uuid)
		c.Next()
	} else {
		c.Abort()
		return
	}
}

//create token
func CreateToken(c *gin.Context, id, account string) (*AuthToken, error) {

	now := time.Now()
	jwtId := account + strconv.FormatInt(now.Unix(), 10)
	authToken := AuthToken{}

	tokenUuid := uuid.New().String()
	// set claims and sign
	claims := Claims{
		Account:   account,
		Uuid:      id,
		TokenUuid: tokenUuid,
		StandardClaims: jwt.StandardClaims{
			Audience:  account,
			ExpiresAt: now.Add(600 * time.Second).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(1 * time.Second).Unix(),
			Subject:   account,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtAccessSecret)
	if err != nil {

		fmt.Println("here")
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": err.Error(),
		// })
		return &authToken, err
	}

	errAccess := rdb.Set(c, tokenUuid+"-accuess", id, 3*time.Minute).Err() // => SET key value 0 數字代表過期秒數，在這裡0為永不過期
	if errAccess != nil {
		panic(errAccess)
	}
	//create refresh token
	rfClaims := Claims{
		Account:   account,
		Uuid:      id,
		TokenUuid: tokenUuid,
		StandardClaims: jwt.StandardClaims{
			Audience:  account,
			ExpiresAt: now.Add(time.Hour * 24).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(1 * time.Second).Unix(),
			Subject:   account,
		},
	}
	refreshTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, rfClaims)

	refreshToken, err := refreshTokenClaims.SignedString(jwtRefreshSecret)
	if err != nil {
		return &authToken, err
	}

	errRefresh := rdb.Set(c, tokenUuid+"-refresh", id, 10*time.Minute).Err() // => SET key value 0 數字代表過期秒數，在這裡0為永不過期
	if errRefresh != nil {
		panic(errRefresh)
	}

	authToken.AccessToken = token
	authToken.RefreshToken = refreshToken
	authToken.AccessExp = 600
	authToken.RefreshExp = 86400
	authToken.TokenType = "Bearer"

	fmt.Println("tokenUUid : " + tokenUuid)
	// c.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })
	return &authToken, nil
}

func Refresh(c *gin.Context, refreshToken string) (*AuthToken, error) {
	//verify the token
	tokenClaims, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtRefreshSecret, nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "Refresh token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "Refresh token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "Refresh signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "Refresh token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "Refresh token is not yet valid before sometime"
			} else {
				message = "can not handle this refresh token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return nil, err
	}
	//is token valid?
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		fmt.Println("account:", claims.Account)
		authToken := AuthToken{}
		//create new accuss token
		now := time.Now()
		// set claims and sign
		claims := Claims{
			Account:   claims.Account,
			TokenUuid: claims.TokenUuid,
			Uuid:      claims.Uuid,
			StandardClaims: jwt.StandardClaims{
				Audience:  claims.Account,
				ExpiresAt: now.Add(600 * time.Second).Unix(),
				Id:        claims.Id,
				IssuedAt:  now.Unix(),
				Issuer:    "ginJWT",
				NotBefore: now.Add(1 * time.Second).Unix(),
				Subject:   claims.Account,
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtAccessSecret)
		if err != nil {

			fmt.Println("here")
			// c.JSON(http.StatusInternalServerError, gin.H{
			// 	"error": err.Error(),
			// })
			return &authToken, err
		}

		errAccess := rdb.Set(c, claims.TokenUuid+"-accuess", claims.Uuid, 3*time.Minute).Err() // => SET key value 0 數字代表過期秒數，在這裡0為永不過期
		if errAccess != nil {
			fmt.Println(errAccess)
		}

		authToken.AccessToken = token
		authToken.RefreshToken = refreshToken
		authToken.AccessExp = 600
		authToken.RefreshExp = 86400
		authToken.TokenType = "Bearer"
		return &authToken, nil

	} else {
		fmt.Println("error here")
		return nil, tokenClaims.Claims.Valid()
	}

}

func Revoke(c *gin.Context, accessToken string) bool {
	//verify the token
	tokenClaims, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtRefreshSecret, nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "Refresh token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "Refresh token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "Refresh signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "Refresh token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "Refresh token is not yet valid before sometime"
			} else {
				message = "can not handle this refresh token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return false
	}
	//is token valid?
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		fmt.Println("account:", claims.Account)

		errDelToken := rdb.Del(c, claims.TokenUuid+"-accuess", claims.TokenUuid+"-refresh").Err() // del access token and refresh token
		if errDelToken != nil {
			fmt.Println(errDelToken)
		}
		return true

	} else {
		fmt.Println("error here")
		return false
	}
}

package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/auth/internal"
)

func init() {
	internal.AddRoute(internal.Routes{
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/checkToken",
			Handler: CheckToken,
		},
	})
}

type RegisterRequest struct {
	Account    string `json:"account"`     // 用户名
	Password   string `json:"password"`    // 密码
	InviteCode string `json:"invite_code"` // 邀请码
}

func Register(c *gin.Context) {
	var param RegisterRequest
	err := c.BindJSON(&param)
	internal.CheckValue(c, internal.ParamParseError(err))
	if internal.GetConfig().Server.NeedInviteCode {
		internal.CheckValue(c, param.InviteCode != "", "param[invite_code] is null")
	}
	internal.CheckValue(c, param.Account != "", "param[account] is null")
	internal.CheckValue(c, param.Password != "", "param[password] is null")

	// 创建账号
	err = createAccount(param.Account, param.Password, param.InviteCode)
	internal.CheckValue(c, err)

	internal.SuccessJsonRsp(c, returnLoginResponse(param.Account))
}

func createAccount(account, password, inviteCode string) error {
	// 校验邀请码
	if internal.GetConfig().Server.NeedInviteCode {
		rows, err := internal.Db().Query("select count(1) from invite_code where `invite_code` = ?", inviteCode)
		if err != nil {
			return internal.DatabaseError(err)
		}
		defer rows.Close()

		rows.Next()

		var count int
		err = rows.Scan(&count)
		if err != nil {
			return internal.DatabaseError(err)
		}

		if count == 0 {
			return internal.ParamValidateError("validate invite code error")
		}
	}

	tx, err := internal.Db().Begin()
	if err != nil {
		return internal.DatabaseError(err)
	}

	// 验证该账号未注册
	{
		rows, err := internal.Db().Query("select count(1) from account where `account` = ?", account)
		if err != nil {
			return internal.DatabaseError(err)
		}
		defer rows.Close()

		rows.Next()

		var count int
		err = rows.Scan(&count)
		if err != nil {
			return internal.DatabaseError(err)
		}

		if count > 0 {
			return internal.ParamValidateError("account already exist")
		}
	}

	// 创建账号
	{
		_, err = tx.Exec("insert into account(`account`,`password`,`create_time`)values(?,?,?)", account, password, internal.NowUnix())
		if err != nil {
			tx.Rollback()
			return internal.DatabaseError(err)
		}
	}

	// 删除邀请码
	if internal.GetConfig().Server.NeedInviteCode {
		_, err = internal.Db().Exec("delete from invite_code where `invite_code` = ?", inviteCode)
		if err != nil {
			tx.Rollback()
			return internal.DatabaseError(err)
		}
	}

	tx.Commit()
	return nil
}

type LoginRequest struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

type LoginResponse struct {
	Account string `json:"account"` // 用户名
	Token   string `json:"token"`   // token
}

func Login(c *gin.Context) {
	var param LoginRequest
	err := c.BindJSON(&param)
	internal.CheckValue(c, internal.ParamParseError(err))
	internal.CheckValue(c, param.Account != "", "param[account] is null")
	internal.CheckValue(c, param.Password != "", "param[password] is null")

	// 验证账号
	err = validateAccount(param.Account, param.Password)
	internal.CheckValue(c, err)

	internal.SuccessJsonRsp(c, returnLoginResponse(param.Account))
}

func validateAccount(account, checkPass string) error {
	rows, err := internal.Db().Query("select password from account where `account` = ?", account)
	if err != nil {
		return internal.DatabaseError(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return internal.ParamValidateError("account not exist")
	}

	var password string
	err = rows.Scan(&password)
	if err != nil {
		return internal.DatabaseError(err)
	}

	if password != checkPass {
		return internal.ParamValidateError("account or password incorrect")
	}

	return nil
}

func returnLoginResponse(account string) LoginResponse {
	token := internal.CreateToken(internal.GetConfig().Server.Secret, account)
	return LoginResponse{
		Account: account,
		Token:   token,
	}
}

type CheckTokenRequest struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

func CheckToken(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		log.Println("authorization is null")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	var param CheckTokenRequest
	err := c.BindJSON(&param)
	if err != nil {
		log.Println("request body error: ", err)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	if param.Method == "" || param.Url == "" {
		log.Println("param is null: ", param.Method, param.Url)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// 用户 token 校验
	account, err := internal.ValidToken(internal.GetConfig().Server.Secret, authorization)
	if err != nil {
		log.Println("invalid authorization: ", err)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	// 权限校验
	var act string
	switch param.Method {
	case http.MethodGet:
		act = "read"
	case http.MethodPost, http.MethodDelete, http.MethodPut:
		act = "write"
	default:
		log.Println("unsupported method: ", param.Method)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	if ok := internal.CheckAccess(account, buildCastbinObj(param.Method, param.Url), act); !ok {
		log.Println("castbin validate failed")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func buildCastbinObj(method, url string) string {
	return fmt.Sprintf("%s[%s]", strings.ToUpper(method), url)
}

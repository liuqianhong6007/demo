package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/demo/auth/config"
	"github.com/liuqianhong6007/demo/auth/internal"
)

func init() {
	internal.AddRoute(internal.Routes{
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: Login,
		},
	})
}

type RegisterRequest struct {
	Account    string // 用户名
	Password   string // 密码
	InviteCode string // 邀请码
}

type RegisterResponse struct {
	Account string
}

func Register(c *gin.Context) {
	var param RegisterRequest
	err := c.BindJSON(&param)
	internal.CheckValue(c, internal.ParamParseError(err))
	if config.NeedInviteCode() {
		internal.CheckValue(c, param.InviteCode != "", "param[invite_code] is null")
	}
	internal.CheckValue(c, param.Account != "", "param[account] is null")
	internal.CheckValue(c, param.Password != "", "param[password] is null")

	// 创建账号
	err = createAccount(param.Account, param.Password, param.InviteCode)
	internal.CheckValue(c, err)

	internal.SuccessJsonRsp(c, RegisterResponse{
		Account: param.Account,
	})
}

func createAccount(account, password, inviteCode string) error {
	// 校验邀请码
	if config.NeedInviteCode() {
		rows, err := internal.Db().Query(`select count(1) from invite_code where 'invite_code' = ?`, inviteCode)
		if err != nil {
			return internal.DatabaseError(err)
		}
		rows.Close()

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
		rows, err := internal.Db().Query(`select count(1) from account where 'account' = ?`, account)
		if err != nil {
			return internal.DatabaseError(err)
		}
		rows.Close()

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
		_, err = tx.Exec(`insert into account('account','password','create_time')values(?,?,?)`, account, password, internal.NowUnix())
		if err != nil {
			tx.Rollback()
			return internal.DatabaseError(err)
		}
	}

	// 删除邀请码
	if config.NeedInviteCode() {
		_, err = internal.Db().Exec(`delete from invite_code where 'invite_code' = ?`, inviteCode)
		if err != nil {
			tx.Rollback()
			return internal.DatabaseError(err)
		}
	}

	tx.Commit()
	return nil
}

type LoginRequest struct {
	Account  string // 用户名
	Password string // 密码
}

type LoginResponse struct {
	Account string // 用户名
	Token   string // token
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

	internal.SuccessJsonRsp(c, returnLoginResponse(param.Account, param.Password))
}

func validateAccount(account, checkPass string) error {
	rows, err := internal.Db().Query(`select password from account where 'account' = ?`, account)
	if err != nil {
		return internal.DatabaseError(err)
	}
	rows.Close()

	rows.Next()

	var password string
	err = rows.Scan(&password)
	if err != nil {
		return internal.DatabaseError(err)
	}

	if password == "" {
		return internal.ParamValidateError("account not exist")
	}

	if password != checkPass {
		return internal.ParamValidateError("account or password incorrect")
	}

	return nil
}

func returnLoginResponse(account, password string) LoginResponse {
	return LoginResponse{
		Account: account,
		Token:   "",
	}
}

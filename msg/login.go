package msg

type C2S_TokenAuthorize struct{
	Token string
	Connect bool
}

const (
	S2C_Close_LoginRepeated   = 1 // 您的账号在其他设备上线，非本人操作请注意修改密码
	S2C_Close_InnerError      = 2 // 登录出错，请重新登录
	S2C_Close_TokenInvalid    = 3 // 登录状态失效，请重新登录
	S2C_Close_UnionIDInvalid  = 4 // 登录出错，微信ID无效
	S2C_Close_UsernameInvalid = 5 // 登录出错，用户名无效
	S2C_Close_SystemOff       = 6 // 系统升级维护中，请稍后重试
	S2C_Close_RoleBlack       = 7 // 账号已冻结，请联系客服微信 S2C_Close.WeChatNumber
	S2C_Close_IPChanged       = 8 // 登录IP发生变化，非本人操作请注意修改密码
)

type S2C_Close struct{
	Error int
	WeChatNumber string
}

type S2C_Authorize struct{

}
package error

import (
	"fmt"
)

type Code int64

const (
	Success                 Code = 0
	Unknown                 Code = -1
	OtherError              Code = -2
	ServerError             Code = -3
	MissParamError          Code = 200000
	LoginTypeError          Code = 200001
	MissUserParamError      Code = 200002
	GetUserInfoError        Code = 200003
	NotEnterpriseUserError  Code = 200004
	DevelopingError         Code = 200005
	LoginTimeoutError       Code = 200006
	ProgramError            Code = 200007
	UnknownApplicationError Code = 200008
	WrongParamError         Code = 200009
	OutOfRangeError         Code = 200010
	IllegalSignError        Code = 200011
	RequestFrequentlyError  Code = 200012
	//实名认证场景
	UserRelationError          Code = 200013
	UserCancelVerifyReverify   Code = 200014
	UserNoVerify               Code = 200015
	AesDecryptError            Code = 200016
	InfoNotSame                Code = 200017
	UserCancelVerify           Code = 200018
	VerifyOnceMobileNotSame    Code = 200019
	VerifyOrCancelReverifyLock Code = 200020
	VerifyIdcardAlready        Code = 200021
	//验证码
	SMSMobileWrong           Code = 200022
	SMSMobileLock            Code = 200023
	SMSExpired               Code = 200024
	SMSNotSame               Code = 200025
	SMSMobileInBlackList     Code = 200026
	SMSUnsuportedRegion      Code = 200027
	SMSPhoneNumberDailyLimit Code = 200028
	UsernameSignedUp         Code = 200100
	UsernameNotSignUp        Code = 200101
	PasswordIsInconsistent   Code = 200102
	UsernameOrPasswordWrong  Code = 200103
	UsernameLocked           Code = 200104
	VerifyIdcardContainsx    Code = 200029 //实名认证场景
	GetverifyInfoError       Code = 200030
	NoPriviliages            Code = 200031
	UserNoVerifyNoMobile     Code = 200032
	UserNotFound             Code = 200033
	UserNoVerifyNoInfo       Code = 200034
)

// ErrorMap 全局返回码
var ErrorMap = map[Code]string{
	Success:                 "成功",
	MissParamError:          "缺少参数",
	Unknown:                 "未知错误",
	LoginTypeError:          "登陆方式错误",
	MissUserParamError:      "缺少用户参数",
	UnknownApplicationError: "未知应用",
	GetUserInfoError:        "获取用户信息错误",
	NotEnterpriseUserError:  "抱歉，您未通过所选企业的认证",
	DevelopingError:         "暂未开放，敬请期待",
	LoginTimeoutError:       "登录超时，请稍后再试",
	ProgramError:            "程序内部错误",
	WrongParamError:         "参数错误",
	ServerError:             "服务器错误",
	OutOfRangeError:         "传入参数超过可支持范围",
	IllegalSignError:        "非法签名错误",
	RequestFrequentlyError:  "抱歉请求过于频繁",
	//实名场景code
	UserCancelVerifyReverify: "您已经实名，暂时不支持再次实名认证",
	AesDecryptError:          "aes解密失败",
	InfoNotSame:              "您已经实名，但本次输入信息与实名信息不一致", //第二次页面输入信息的不一致，不能过腾讯云接口
	//取消实名场景code
	UserRelationError:          "用户实名关系错误",
	UserNoVerify:               "用户未实名",
	UserCancelVerify:           "抱歉，您已取消实名",
	VerifyOnceMobileNotSame:    "抱歉，您已通过别的手机号完成实名", //这种类型可以过腾讯云接口，个人的手机副卡
	VerifyOrCancelReverifyLock: "操作过于频繁，已对此账号锁定操作，24小时后自动解锁",
	VerifyIdcardAlready:        "抱歉，该身份证已实名认证", //已经实名，不支持再次实名
	//验证码
	SMSMobileWrong:           "请输入正确的手机号码",
	SMSMobileLock:            "操作次数过多，请更换手机号或今日24小时后重试",
	SMSExpired:               "验证码已过期，请重新获取",
	SMSNotSame:               "验证码错误，请重试",
	SMSMobileInBlackList:     "抱歉，此手机号码被运营商列为黑名单",
	SMSUnsuportedRegion:      "抱歉，暂不支持该地区短信下发",
	SMSPhoneNumberDailyLimit: "抱歉，该手机号每日下发短信条数超过设定的上限",
	UsernameSignedUp:         "该用户名已注册",
	UsernameNotSignUp:        "该用户名未注册，请注册后重新登录",
	PasswordIsInconsistent:   "抱歉，两次密码输入不一致",
	UsernameOrPasswordWrong:  "抱歉，账号或密码输入错误，请确认后重试",
	UsernameLocked:           "抱歉，因账号或密码输入错误次数过多，账号已锁定，请明天重试",
	VerifyIdcardContainsx:    "抱歉，您身份证号末位的字母应为大写",
	GetverifyInfoError:       "获取元素配置错误",
	NoPriviliages:            "抱歉，权限不足",
	UserNoVerifyNoMobile:     "该用户未实名，无法获取手机号码",
	UserNotFound:             "抱歉，该用户不存在",
	UserNoVerifyNoInfo:       "该用户未实名，无法获取用户实名信息",
}

// GetMsg 获取 Code 所对应的错误信息
func (c Code) GetMsg(msg ...string) string {
	if len(msg) != 0 {
		return msg[0]
	}
	s, ok := ErrorMap[c]
	if ok {
		return s
	}
	return msg[0]
}

// String 获取 Code 所对应的错误信息
func (c Code) String(msg ...string) string {
	str, ok := ErrorMap[c]
	if !ok {
		return fmt.Sprint(msg)
	}
	return str
}

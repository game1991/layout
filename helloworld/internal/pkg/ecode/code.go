package ecode

import (
	// "git.xq5.com/office/survey-backend/utils/ecode/rpc"
	"errors"

	//"git.xq5.com/office/survey-backend/pkg/captcha"
	"helloworld/pkg/response"
)

// 负数的错误码直接弹框进行展示。如创建用户，用户参数缺失
// 正数错误码前端进行拦截处理，如果前端没有进行拦截那么不做任何提示。如创建用户信息，用户已经存在
// 业务错误码统一长度6位数，前3位是业务模块，后3位具体业务
const (
	// OK 请求成功。一般用于GET与POST请求
	OK = 0
	// Created 已创建。成功请求并创建了新的资源
	Created = -201
	// NoContent 无内容。服务器成功处理，但未返回内容。在未更新网页的情况下，可确保浏览器继续显示当前文档
	NoContent = -204
	// BadRequest 请求参数错误
	BadRequest = -400
	// Unauthorized 请求要求用户的身份认证
	Unauthorized = -401
	// NotFound 服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置"您所请求的资源无法找到"的个性页面
	NotFound = -404
	// InternalServerError 服务器内部错误，无法完成请求
	InternalServerError = -500
	// DatabaseOperationError 数据库访问错误
	DatabaseOperationError = -100500
	// ErrorCheckAdminFail 登录
	ErrorCheckAdminFail = -200000 // 用户名或密码错误
	// NotImplemented 尚未实现
	NotImplemented = -100444
	// SystemError 系统错误
	SystemError = -100520

	ErrGetPhoneNumber = -1 // 小程序获取手机号码失败
)

// 用户业务
const (
	ErrUserNotExist             = 1      // 用户不在问卷系统名单
	ErrQuestionnaireNotStart    = 2      // 当前问卷未到时间
	ErrQuestionnaireExpired     = 3      // 当前问卷已到结束时间
	ErrTaskStatusOrNotExist     = 4      // 当前任务状态不正确或者不存在
	ErrTaskNotPickCompany       = 5      // 当前任务没有选择企业
	ErrTaskNotPickQuestionnaire = 6      // 当前任务没有选择问卷
	ErrTaskHasExpired           = 7      // 当前任务已过期
	ErrAnswerTimerTooShort      = 8      // 当前填写时长太短
	ErrCaptchaCode              = 100000 // 验证码无效
	ErrCaptchaSendTooBusy       = 100001 // 验证码发送太频繁(指的是当日频次)
	ErrCaptchaCoolDown          = 100002 // 验证码发送cd冷却时间内

	ErrAnswerNotFound                  = 100003 // 问卷答案不存在
	ErrAdminUserNamePassword           = 100004 // 用户名或密码错误
	ErrAdminUserNotExist               = 100005 // 用户不存在
	ErrAdminUserNameExist              = 100006 // 用户名存在
	ErrAdminUserDeleteSelf             = 100007 // 用户不可以删除自己
	ErrAdminUserIsDisable              = 100008 // 用户已被禁用
	ErrAdminUserDisableSelf            = 100009 // 用户不能禁用自己
	ErrPhoneReg                        = 100010 // 错误的手机号格式
	ErrQuestionnaireNameHasExisted     = 100011 // 问卷名称已存在
	ErrInvitationNotFoundSMSStatusInit = 100012 // 受邀企业中没有短信待确定的对象
	ErrTaskInvitationHasExisted        = 100013 //当前任务中已有此号码，无法重复添加
	ErrInvitationSmsStatus             = 100014 // 当前任务没有符合短信状态的受邀企业

	ErrAdminInvitationImportText = 200000 // 管理员批量导入输入错误

	ErrDictDataNameIsExist = 110001 // 数据字典的数据名称已存在
	ErrDictOptionIsExist   = 110002 // 数据字典的选项名称已存在
	ErrUpdateNodeDataName  = 110003 // 仅有根节点可以修改数据名称
	ErrNodeParentIsNull    = 110004 // 插入子节点，父节点字段不能为空
	ErrDictIDNotExist      = 110005 // 数据字典不存在

	ErrReportNumberNotExist = 120001 // 选择得分报表编号不存在
	ErrStatisticsIsRunning  = 120002 // 得分统计正在计算中

	ErrQuestionnaireNotFound = 130001 // 问卷不存在
	ErrTaskNotFound          = 140001 // 任务不存在
)

var (
	// ErrBadRequest ErrBadRequest
	ErrBadRequest error
	// ErrUnauthorized ErrUnauthorized
	ErrUnauthorized error
	// ErrNotFound ErrRecordNotFound
	ErrNotFound error
	// ErrDatabaseOperationError ErrDatabaseOperationError
	ErrDatabaseOperationError error
	// ErrNotImplemented ErrNotImplemented
	ErrNotImplemented error
	// ErrSystemError ErrSystemError
	ErrSystemError error
)

func init() {
	ErrBadRequest = errors.New(GetMsg(BadRequest))
	ErrUnauthorized = errors.New(GetMsg(Unauthorized))
	ErrNotFound = errors.New(GetMsg(NotFound))
	ErrDatabaseOperationError = errors.New(GetMsg(DatabaseOperationError))
	ErrNotImplemented = errors.New(GetMsg(NotImplemented))
	ErrSystemError = errors.New(GetMsg(SystemError))
}

// WrapErrorWithCode 封装错误通过code判断
func WrapErrorWithCode(err error) response.Error {
	if respErr, ok := err.(response.Error); ok {
		return respErr
	}
	code := InternalServerError
	switch err.(type) {
	// case *rpc.Error:
	// 	e := err.(*goMicroErr.Error)
	// 	switch e.Code {

	// 	}
	case error:
		e := err.(error)
		switch e {
		//case captcha.ErrCaptchaNotCoolDown:
		//	code = ErrCaptchaCoolDown
		//case captcha.ErrCaptchaVerifyFailed:
		//	code = ErrCaptchaCode
		//case captcha.ErrCaptchaSendTooBusy:
		//	code = ErrCaptchaSendTooBusy
		//case ErrUnauthorized:
		//	code = Unauthorized
		}
	}
	return &response.InnerError{Status: code, Msg: err.Error()}
}

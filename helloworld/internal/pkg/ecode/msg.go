package ecode

import "helloworld/pkg/response"

// MsgFlags msg flags
var MsgFlags = map[int]string{
	OK:                  "请求成功",
	Created:             "成功请求并创建了新的资源",
	NoContent:           "无内容，服务器成功处理，但未返回内容",
	BadRequest:          "请求参数错误",
	Unauthorized:        "请求要求用户的身份认证",
	NotFound:            "未找到资源",
	InternalServerError: "服务器内部错误",

	// GetInfoFail:            "用户不存在",
	// CreateInfoFail:         "创建用户失败",
	// CheckExistReqNameEmpty: "检测是否存在接口入参Name为空",
	ErrorCheckAdminFail: "用户名或密码错误",

	DatabaseOperationError: "数据库访问错误",

	NotImplemented: "尚未实现",
	SystemError:    "系统错误",

	ErrUserNotExist:             "用户不在问卷系统名单",
	ErrQuestionnaireNotStart:    "当前问卷未到时间",
	ErrQuestionnaireExpired:     "当前问卷已到结束时间",
	ErrTaskStatusOrNotExist:     "当期任务状态不正确或者不存在",
	ErrTaskNotPickCompany:       "当前任务没有选择企业",
	ErrTaskNotPickQuestionnaire: "当前任务没有选择问卷",
	ErrTaskHasExpired:           "当前任务已过期",
	ErrCaptchaCode:              "验证码无效",
	ErrCaptchaSendTooBusy:       "验证码发送太频繁(指的是当日频次)",
	ErrCaptchaCoolDown:          "验证码发送cd冷却时间内",

	ErrAnswerNotFound:                  "答案不存在",
	ErrAnswerTimerTooShort:             " 当前填写时长太短",
	ErrAdminUserNamePassword:           "用户名或者密码错误",
	ErrAdminUserNotExist:               "用户不存在",
	ErrAdminUserNameExist:              "用户名存在",
	ErrAdminUserDeleteSelf:             "用户不可以删除自己",
	ErrAdminUserIsDisable:              "用户已被禁用",
	ErrAdminUserDisableSelf:            "用户不能禁用自己",
	ErrPhoneReg:                        "错误的手机号格式",
	ErrQuestionnaireNameHasExisted:     "问卷名称已存在",
	ErrInvitationNotFoundSMSStatusInit: "受邀企业中没有短信待确定的对象",
	ErrAdminInvitationImportText:       "管理员批量导入输入错误",
	ErrDictDataNameIsExist:             "数据字典的数据名称已存在",
	ErrDictOptionIsExist:               "数据字典的选项名称已存在",
	ErrUpdateNodeDataName:              "仅有根节点可以修改数据名称",
	ErrNodeParentIsNull:                "插入子节点，父节点字段不能为空",
	ErrDictIDNotExist:                  "数据字典不存在",
	ErrReportNumberNotExist:            "选择得分报表编号不存在",
	ErrStatisticsIsRunning:             "得分统计正在计算中",
	ErrQuestionnaireNotFound:           "问卷不存在",
	ErrTaskNotFound:                    "任务不存在",
	ErrTaskInvitationHasExisted:        "当前任务中已有此号码，无法重复添加",
	ErrInvitationSmsStatus:             "当前任务没有符合短信状态的受邀企业",
}

// GetMsg 根据代码获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[InternalServerError]
}

// Fail 返回错误信息
func Fail(code int, msg ...string) response.Error {
	var info string
	if len(msg) == 0 {
		info = GetMsg(code)
	} else {
		info = msg[0]
	}
	return &response.InnerError{
		Status: code,
		Msg:    info,
	}
}

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var msgFlags = map[int]string{
	1: "ok",
	2: "fail",
	3: "请求参数错误",

	101: "Token鉴权失败",
	102: "Token已超时",
	103: "Token生成失败",
	104: "Token错误",
	105: "该用户被禁止登录",
	106: "仅在开发环境可用",
	107: "查询范围最多为7天",

	301: "customer 不存在",

	10001: "已存在该标签名称",
	10002: "该标签不存在",
	10003: "该文章不存在",
	11004: "无法登录,请联系管理员",
	11005: "验证码错误",

	// 账户相关
	11001: "用户名或密码错误",
	11002: "旧密码错误",

	12001: "该字段已经在使用中",
	12002: "数据源类型与规则组类型不匹配",

	13001: "执行脚本失败",
	13002: "计算后字段值类型不正确",
	13003: "执行脚本失败",

	14004: "不是待审批状态",

	15001: "已操作, 待审核",
}

func ApiReturn(code int, data interface{}, c *gin.Context) {
	msg, ok := msgFlags[code]
	if !ok {
		msg = msgFlags[2]
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": msg,
		"data":    data,
	})
}

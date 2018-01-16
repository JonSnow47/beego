package filters

import (
	"github.com/astaxie/beego/context"

	"github.com/JonSnow47/beego/application/Fan/common"
)

func LoginFilter(ctx *context.Context) {
	if _, ok := MapFilter[ctx.Request.RequestURI]; !ok {
		userID := ctx.Input.CruSession.Get(common.SessionUserID)

		if userID == nil {
			ctx.Output.JSON(map[string]interface{}{common.RespKeyStatus: common.ErrLoginRequired}, false, false)
		}
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
)

func WebAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var cs = web_service.NewConsumerService()
		//var tokenString string
		//headerContentType := ctx.GetHeader("Content-type")
		//headerArr := strings.Split(headerContentType,";")
		//header := headerArr[0]
		//fmt.Printf("header = %v\n",header)
		//var jwtParam web_param.JwtParam
		//var avatarParam web_param.PersonalAvatarParams
		//if header == "multipart/form-data" {
		//	_ = ctx.ShouldBind(&avatarParam)
		//	ctx.Set("avatarParam",avatarParam)
		//	tokenString = avatarParam.Token
		//} else {
		//	_ = ctx.ShouldBindBodyWith(&jwtParam,binding.JSON)
		//	tokenString = jwtParam.Token
		//}
		//token,claims,parseErr := common.ParseToken(tokenString)
		//if parseErr != nil || !token.Valid {
		//	common.Failed(ctx,"需要登陆1")
		//	ctx.Abort()
		//	return
		//}
		//userId := claims.UserId
		//sessId := strconv.FormatInt(userId, 10)
		//idInterface := common.GetSess(ctx,"consumer_" + sessId)
		//if idInterface == nil {
		//	common.Failed(ctx,"需要登陆2")
		//	ctx.Abort()
		//	return
		//}
		//idStr := idInterface.(string)
		//id,err := strconv.ParseInt(idStr,10,64)
		//if err != nil {
		//	common.Failed(ctx,"获取用户ID失败")
		//	ctx.Abort()
		//	return
		//}
		//res,resErr := cs.GetConsumerById(id)
		//if resErr != nil || res.Id == 0 || res.IsDel == 1 {
		//	common.Failed(ctx,"需要注册")
		//	ctx.Abort()
		//	return
		//}
		//res.PassWord = ""
		//ctx.Set("webUser",res)
		//if jwtParam.OrderList != nil {
		//	ctx.Set("orderListData",jwtParam.OrderList)
		//}
		ctx.Next()
	}
}


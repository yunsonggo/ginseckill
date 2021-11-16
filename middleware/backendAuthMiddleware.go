package middleware

import (
	"github.com/gin-gonic/gin"
)

func BackendAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var tokenString string
		//var jwtParam web_param.JwtParam
		//headerContentType := ctx.GetHeader("Content-type")
		//headerArr := strings.Split(headerContentType,";")
		//header := headerArr[0]
		//fmt.Printf("header = %v\n",header)
		//var bannerUploadParam backend_param.ArticleBannerUploadParam
		//var coverUploadParam backend_param.ArticleCoverUploadParam
		//var videoUploadParam backend_param.ArticleVideoUploadParam
		//var linkIconUploadParam backend_param.LinkIconUploadParam
		//var websiteAvatarParam backend_param.WebsiteAvatarParam
		//var websiteLogoParam backend_param.WebsiteLogoParam
		//var websiteBannerParam backend_param.WebsiteBannerParam
		//var websiteWeboParam backend_param.WebsiteWeboParam
		//var websiteWecharParam backend_param.WebsiteWecharParam
		//var websiteApkParam backend_param.WebsiteApkParam
		//if header == "multipart/form-data" {
		//	err := ctx.ShouldBind(&bannerUploadParam)
		//	if err == nil {
		//		ctx.Set("bannerParam",bannerUploadParam)
		//		tokenString = bannerUploadParam.Token
		//	}
		//	err = ctx.ShouldBind(&coverUploadParam)
		//	if err == nil {
		//		ctx.Set("coverParam",coverUploadParam)
		//		tokenString = coverUploadParam.Token
		//	}
		//	err = ctx.ShouldBind(&videoUploadParam)
		//	if err == nil {
		//		ctx.Set("videoParam",videoUploadParam)
		//		tokenString = videoUploadParam.Token
		//	}
		//	err = ctx.ShouldBind(&linkIconUploadParam)
		//	if err == nil {
		//		ctx.Set("linkIconParam",linkIconUploadParam)
		//		tokenString = linkIconUploadParam.Token
		//	}
		//	err = ctx.ShouldBind(&websiteAvatarParam)
		//	if err == nil {
		//		ctx.Set("websiteAvatarParam",websiteAvatarParam)
		//		tokenString = websiteAvatarParam.Token
		//	}
		//	err = ctx.ShouldBind(&websiteLogoParam)
		//	if err == nil {
		//		ctx.Set("websiteLogoParam",websiteLogoParam)
		//		tokenString = websiteLogoParam.Token
		//	}
		//	err = ctx.ShouldBind(&websiteBannerParam)
		//	if err == nil {
		//		ctx.Set("websiteBannerParam",websiteBannerParam)
		//		tokenString = websiteBannerParam.Token
		//	}
		//	err = ctx.ShouldBind(&websiteWeboParam)
		//	if err == nil {
		//		ctx.Set("websiteWeboParam",websiteWeboParam)
		//		tokenString = websiteWeboParam.Token
		//	}
		//	err = ctx.ShouldBind(&websiteWecharParam)
		//	if err == nil {
		//		ctx.Set("websiteWecharParam",websiteWecharParam)
		//		tokenString = websiteWecharParam.Token
		//	}
		//	err = ctx.ShouldBind(&websiteApkParam)
		//	if err == nil {
		//		ctx.Set("websiteApkParam",websiteApkParam)
		//		tokenString = websiteApkParam.Token
		//	}
		//} else {
		//	_ = ctx.ShouldBindBodyWith(&jwtParam,binding.JSON)
		//	tokenString = jwtParam.Token
		//}
		//if tokenString == "" {
		//	tokenString = ctx.GetHeader("Authorization")
		//
		//}
		//token,claims,parseErr := common.ParseToken(tokenString)
		//if parseErr != nil || !token.Valid {
		//	common.Failed(ctx,"需要登陆1")
		//	ctx.Abort()
		//	return
		//}
		//userId := claims.UserId
		//sessId := strconv.FormatInt(userId, 10)
		//idInterface := common.GetSess(ctx,"manager_" + sessId)
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
		//var backendManagerServer = backend_service.NewManagerServer()
		//res,resErr := backendManagerServer.GetManagerById(id)
		//if resErr != nil || res.Id == 0  {
		//	common.Failed(ctx,"需要注册")
		//	ctx.Abort()
		//	return
		//}
		//ctx.Set("loginManager",res)
		ctx.Next()
	}
}


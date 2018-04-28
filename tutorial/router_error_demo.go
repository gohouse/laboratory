package api

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"runtime/debug"
	"vigo.gaodun.com/vigo_loger"
	"vigo.gaodun.com/vigo_tool"
)

type ApiRoute struct {
	method []string
	handle func(*fasthttp.RequestCtx)
}


func GetVerifyList() map[string]bool {
	verifyList := make(map[string]bool)
	verifyList["/"] = true
	verifyList["/api/v1/makeidentity"] = true
	verifyList["/headers"] = true
	verifyList["/status"] = true
	return verifyList
}

var verifyList = GetVerifyList()

type Decorator struct {
	RunFuc  func(*fasthttp.RequestCtx)
	PathStr string
}

// 判断 ip
func (d Decorator) Decorator(ctx *fasthttp.RequestCtx) {

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			SystemError(ctx)
		}
	}()

	d.RunFuc(ctx)
	return
}

var dct Decorator

func GetRouter() *fasthttprouter.Router {
	// add api route in here.
	bs := NewBaseServer()
	ApiRouteList := map[string]ApiRoute{
		"/": {[]string{"GET", "POST"}, handRoot},
		"/api/v1/makeidentity": {[]string{"GET", "POST"}, bs.SetHandler},
		"/status":              {[]string{"GET"}, GetStatus},
		//"/headers":                          {[]string{"GET"}, headerHandle},
		"/api/v1/getallusers":      {[]string{"GET"}, bs.GetAllUser},
		"/api/v1/getusersagent":    {[]string{"POST"}, bs.GetUsersByAgent},
		"/api/v1/getuserswork":     {[]string{"GET"}, bs.GetUsersByWorkNo},
		"/api/v1/getuserpartial":   {[]string{"POST"}, bs.GetUserByPartial},
		"/api/v1/checklogin":       {[]string{"GET"}, bs.CheckLogin},
		"/api/v1/login":            {[]string{"POST"}, bs.Login},
		"/api/v1/changepwd":        {[]string{"POST"}, bs.ChangePwd},
		"/api/v1/resetpwd":         {[]string{"POST"}, bs.ResetPassword},
		"/api/v1/edituserfromjson": {[]string{"POST"}, bs.EditUser},
		"/api/v1/destroy":          {[]string{"PUT"}, bs.Destroy},
		"/api/v1/getsession":       {[]string{"GET"}, bs.GetSession},
		"/api/v1/getdepartnode":    {[]string{"GET"}, bs.GetDepartmentNode},
		"/api/v1/getalldepartnode": {[]string{"GET"}, bs.GetAllDepartmentNode},
		"/api/v1/getdepartlevel":   {[]string{"GET"}, bs.GetDepartByLevel},
		"/api/v1/getdeparttree":    {[]string{"GET"}, bs.GetDepartNodeByDepartId},
		"/api/v1/getleader":        {[]string{"GET"}, bs.GetLeader},
		"/api/v1/getusers":         {[]string{"GET", "POST"}, bs.GetUserInfo},
		"/api/v1/getroledepart":    {[]string{"GET", "POST"}, bs.GetRoleDepartUser},
		"/api/v1/getappgav":        {[]string{"GET"}, bs.GetAppNavigationNode},
		"/api/v1/getuserbyerp":     {[]string{"GET", "POST"}, bs.GetUserByErp},
		"/api/v1/edituser":         {[]string{"POST"}, bs.EditUserById},
		"/api/v1/findmq":           {[]string{"POST"}, bs.FindAliMQ},
		"/api/v1/findfuncdep":      {[]string{"GET"}, bs.FindDepartByFuncId},
		"/api/v1/findbycampus":     {[]string{"GET"}, bs.FindDepartByCampus},
		"/api/v1/findallrole":      {[]string{"GET"}, bs.FindRoles},
		"/api/v1/findpermission":   {[]string{"GET"}, bs.FindPermission},
		"/api/v1/edsuser":          {[]string{"POST"}, bs.EdsAddUser},
		// cache
		"/api/v1/refreshcache": {[]string{"GET"}, bs.RefreshCache},
		"/api/v1/cache/depart": {[]string{"GET"}, bs.DelCacheDepart}, // 清除部门缓存
		"/api/v1/cache/role":   {[]string{"GET"}, bs.DelCacheRole},   // 清除角色缓存
	}

	router := &fasthttprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}

	for k, v := range ApiRouteList {
		for _, value := range v.method {
			dct.PathStr = k
			dct.RunFuc = v.handle
			router.Handle(value, k, dct.Decorator)
		}
	}

	return router
}

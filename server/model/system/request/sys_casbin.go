/*
 * @Author: aliyun4052411970 keane_bj@aliyun.com
 * @Date: 2024-11-19 21:14:52
 * @LastEditors: aliyun4052411970 keane_bj@aliyun.com
 * @LastEditTime: 2024-11-19 21:22:19
 * @FilePath: /gin-vue-admin/server/model/system/request/sys_casbin.go
 * @Description: 生成遇到问题，请再试试
 */
package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint         `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/menu/getMenu", Method: "POST"},
		{Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/changePassword", Method: "POST"},
		{Path: "/user/setUserAuthority", Method: "POST"},
		{Path: "/user/getUserInfo", Method: "GET"},
		// 学习增删改查
		{Path: "/user/admin_registerV", Method: "POST"},
		{Path: "/user/changePasswordV", Method: "POST"},
		{Path: "/user/setUserAuthorityV", Method: "POST"},
		{Path: "/user/setUserInfoV", Method: "PUT"},
		{Path: "/user/getUserInfoV", Method: "GET"},
		{Path: "/user/setSelfInfo", Method: "PUT"},
		{Path: "/fileUploadAndDownload/upload", Method: "POST"},
		{Path: "/sysDictionary/findSysDictionary", Method: "GET"},
	}
}

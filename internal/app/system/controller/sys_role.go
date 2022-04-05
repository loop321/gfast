/*
* @desc:角色管理
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:08
 */

package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Role = roleController{}

type roleController struct {
	BaseController
}

// List 角色列表
func (c *roleController) List(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error) {
	res, err = service.Role().GetRoleListSearch(ctx, req)
	return
}

// GetParams 获取角色表单参数
func (c *roleController) GetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error) {
	res = new(system.RoleGetParamsRes)
	res.Menu, err = service.Rule().GetMenuList(ctx)
	return
}

// Add 添加角色信息
func (c *roleController) Add(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error) {
	err = service.Role().AddRole(ctx, req)
	return
}

// Get 获取角色信息
func (c *roleController) Get(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error) {
	res = new(system.RoleGetRes)
	res.Role, err = service.Role().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.MenuIds, err = service.Role().GetFilteredNamedPolicy(ctx, req.Id)
	return
}

// Edit 修改角色信息
func (c *roleController) Edit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error) {
	err = service.Role().EditRole(ctx, req)
	return
}

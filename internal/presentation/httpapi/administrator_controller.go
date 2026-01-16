package httpapi

import (
	"api-service-template/internal/app"
	"api-service-template/internal/domain"
	"api-service-template/internal/option"
	"api-service-template/internal/presentation/httpapi/request"
	"context"
)

type administratorController struct {
	App *app.Application
}

func newAdministratorController(opt *option.Options) *administratorController {
	return &administratorController{
		App: app.NewApplication(opt.GetDB()),
	}
}

// Create 创建管理员
// @Router /api/administrators [post]
func (c *administratorController) Create(ctx context.Context, req *request.Administrator) (*apiResponse, error) {
	administrator := domain.Administrator{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
	}
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.Administrator.Create()
	if err := c.App.Commands.CreateAdministrator.Handle(ctx, &administrator); err != nil {
		return nil, err
	}

	return sendResponse(withData(administrator), withMsg(createOKMsg)), nil
}

// Delete 删除管理员
// @Router /api/administrators/:id [delete]
func (c *administratorController) Delete(ctx context.Context, req *request.Administrator) (*apiResponse, error) {
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.Administrator.Delete()
	if err := c.App.Commands.DeleteAdministrator.Handle(ctx, req.ID); err != nil {
		return nil, err
	}

	return sendResponse(withMsg(deleteOKMsg)), nil
}

// Update 修改管理员
// @Router /api/administrators/:id [put]
func (c *administratorController) Update(ctx context.Context, req *request.Administrator) (*apiResponse, error) {
	administrator := domain.Administrator{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
	}
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.Administrator.Update()
	if err := c.App.Commands.UpdateAdministrator.Handle(ctx, &administrator); err != nil {
		return nil, err
	}

	return sendResponse(withMsg(updateOKMsg)), nil
}

// Find 获取管理员
// @Router /api/administrators/:id [get]
func (c *administratorController) Find(ctx context.Context, req *request.Administrator) (*apiResponse, error) {
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.Administrator.Find()
	administrator, err := c.App.Commands.FindAdministrator.Handle(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return sendResponse(withData(administrator)), nil
}

// List 查找管理员
// @Router /api/administrators [get]
func (c *administratorController) List(ctx context.Context, req *request.Query) (*apiResponse, error) {
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.Administrator.List()
	administrators, total, err := c.App.Commands.ListAdministrator.Handle(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	return sendResponse(withData(dataFields{
		"items": administrators,
		"total": total,
	})), nil
}

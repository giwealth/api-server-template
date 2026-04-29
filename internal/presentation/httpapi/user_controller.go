package httpapi

import (
	"api-service-template/internal/app"
	"api-service-template/internal/domain"
	"api-service-template/internal/presentation/httpapi/request"
	"time"

	// "api-service-template/internal/presentation/httpapi/response"
	"context"
)

type userController struct {
	App *app.Application
}

func newUserController(application *app.Application) *userController {
	return &userController{
		App: application,
	}
}

// Create 创建用户
// @Router /api/users [post]
func (c *userController) Create(ctx context.Context, req *request.User) (*apiResponse, error) {
	user := domain.User{
		ID:      req.ID,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	}

	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.User.Create()
	if err := c.App.Commands.CreateUser.Handle(ctx, &user); err != nil {
		return nil, err
	}

	// return response.WithData(user).SetMessage(response.CreateOKMsg), nil
	return sendResponse(withData(user), withMsg(createOKMsg)), nil
}

// Delete 删除用户
// @Router /api/users [delete]
func (c *userController) Delete(ctx context.Context, req *request.User) (*apiResponse, error) {
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.User.Delete()
	if err := c.App.Repositories.User.Delete(ctx, req.ID); err != nil {
		return nil, err
	}

	return sendResponse(withMsg(deleteOKMsg)), nil
}

// Update 修改用户
// @Router /api/users [put]
func (c *userController) Update(ctx context.Context, req *request.User) (*apiResponse, error) {
	user := domain.User{
		ID:      req.ID,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	}
	// // 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.User.Update()
	if err := c.App.Commands.UpdateUser.Handle(ctx, &user); err != nil {
		return nil, err
	}

	return sendResponse(withMsg(updateOKMsg)), nil
}

// Find 获取用户
// @Router /api/users/detail [get]
func (c *userController) Detail(ctx context.Context, req *request.User) (*apiResponse, error) {
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.User.Find()
	user, err := c.App.Repositories.User.Find(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return sendResponse(withData(user)), nil
}

// List 查找用户
// @Router /api/users [get]
func (c *userController) List(ctx context.Context, req *request.Query) (*apiResponse, error) {
	// 此处如果没有其它业务逻辑, 只是操作数据库就直接使用 c.App.Repositories.User.List()
	users, total, err := c.App.Repositories.User.List(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	time.Sleep(10 * time.Second)
	return sendResponse(withData(dataFields{
		"items": users,
		"total": total,
	})), nil
}

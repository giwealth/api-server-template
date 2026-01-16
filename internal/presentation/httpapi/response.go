package httpapi

const (
	createOKMsg = "创建成功"
	updateOKMsg = "更新成功"
	deleteOKMsg = "删除成功"
)

type dataFields map[string]any

// apiResponse 接口统一下行数据结构
type apiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type apiResponseOption func(*apiResponse)

// WithData 设置数据
func withData(data any) apiResponseOption {
	return func(ar *apiResponse) {
		ar.Data = data
	}
}

func withMsg(msg string) apiResponseOption {
	return func(ar *apiResponse) {
		ar.Message = msg
	}
}

func withError(err apiError) apiResponseOption {
	return func(ar *apiResponse) {
		ar.Code = err.Code
		ar.Message = err.Message
		if data, ok := err.Data(); ok {
			ar.Data = data
		}
	}
}

func sendResponse(options ...apiResponseOption) *apiResponse {
	resp := &apiResponse{}
	for _, fn := range options {
		fn(resp)
	}
	if resp.Data == nil {
		resp.Data = struct{}{}
	}

	return resp
}

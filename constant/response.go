package constant

type BasicResponse struct {
	Status   string      `json:"status"`
	Code     string      `json:"code"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message,omitempty"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
}

type PageInfo struct {
	PageNumber       string `json:"page_number"`
	PageSize         string `json:"page_size"`
	TotalRecordCount string `json:"total_record_count"`
}

type MixedResponse struct {
	BasicResponse
	MixData struct {
		Success interface{} `json:"success"`
		Fail    interface{} `json:"fail"`
	} `json:"mix_data"`
}

func CreateSuccessResponse(_data interface{}) BasicResponse {
	return BasicResponse{
		Status: HTTP_STATUS_SUCCESS,
		Code:   HTTP_CODE_200_OK,
		Data:   _data,
	}
}

func CreateErrorResponse(_errCode, _errMsg string) BasicResponse {
	return BasicResponse{
		Status:  HTTP_STATUS_ERROR,
		Code:    _errCode,
		Message: _errMsg,
	}
}

type RecordCount struct {
	TotalRecord int64 `json:"total_record"`
}

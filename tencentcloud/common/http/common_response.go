package common

import "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"

type actionResult map[string]interface{}
type CommonResponse struct {
	*BaseResponse
	*actionResult
}

func NewCommonResponse() (response *CommonResponse) {
	response = &CommonResponse{
		BaseResponse: &BaseResponse{},
		actionResult: &actionResult{},
	}
	return
}

func (r *CommonResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, r.actionResult)
}

func (r *CommonResponse) GetBody() []byte {
	raw, _ := json.Marshal(r.actionResult)
	return raw
}

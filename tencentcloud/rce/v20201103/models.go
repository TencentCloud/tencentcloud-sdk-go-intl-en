// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201103

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeRiskAssessmentRequestParams struct {
	// Business parameters
	BizCryptoData *InputBizCryptoData `json:"BizCryptoData,omitnil,omitempty" name:"BizCryptoData"`
}

type DescribeRiskAssessmentRequest struct {
	*tchttp.BaseRequest
	
	// Business parameters
	BizCryptoData *InputBizCryptoData `json:"BizCryptoData,omitnil,omitempty" name:"BizCryptoData"`
}

func (r *DescribeRiskAssessmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskAssessmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizCryptoData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskAssessmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskAssessmentResponseParams struct {
	// Business output parameters.
	Data *OutputRiskAssessment `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskAssessmentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskAssessmentResponseParams `json:"Response"`
}

func (r *DescribeRiskAssessmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskAssessmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InputBizCryptoData struct {
	// Whether to authorize.
	IsAuthorized *string `json:"IsAuthorized,omitnil,omitempty" name:"IsAuthorized"`

	// Encryption type.
	CryptoType *string `json:"CryptoType,omitnil,omitempty" name:"CryptoType"`

	// Encrypted content.
	CryptoContent *string `json:"CryptoContent,omitnil,omitempty" name:"CryptoContent"`
}

type OutputRiskAssessment struct {
	// Return code.
	// Note: This field may return null, indicating that no valid value is found.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Returned information
	// 
	// Note: This field may return null, indicating that no valid value is found.
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// Business details.
	// Note: This field may return null, indicating that no valid value is found.
	Value *OutputRiskAssessmentInfo `json:"Value,omitnil,omitempty" name:"Value"`

	// Request ID.
	// Note: This field may return null, indicating that no valid value is found.
	UUid *string `json:"UUid,omitnil,omitempty" name:"UUid"`
}

type OutputRiskAssessmentExtraInfoPair struct {
	// Key.
	// Note: This field may return null, indicating that no valid value is found.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value.
	// Note: This field may return null, indicating that no valid value is found.
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type OutputRiskAssessmentInfo struct {
	// Risk value.
	// Note: This field may return null, indicating that no valid value is found.
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// Risk type.
	// Note: This field may return null, indicating that no valid value is found.
	RiskType []*int64 `json:"RiskType,omitnil,omitempty" name:"RiskType"`

	// Device identification
	// 
	// Note: This field may return null, indicating that no valid value is found.
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// Extended attributes.
	// Note: This field may return null, indicating that no valid value is found.
	ExtraInfo []*OutputRiskAssessmentExtraInfoPair `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// Token generation timestamp.
	// Note: This field may return null, indicating that no valid value is found.
	TokenTime *string `json:"TokenTime,omitnil,omitempty" name:"TokenTime"`
}
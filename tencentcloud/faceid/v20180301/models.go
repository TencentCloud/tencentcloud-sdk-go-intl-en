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

package v20180301

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type LivenessCompareRequest struct {
	*tchttp.BaseRequest

	// Base64 value of a photo for face comparison;
	// Base64-encoded image data is up to 3 MB. Only JPG and PNG formats are supported.
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// Base64 value of a video for liveness detection;
	// The size after Base64-encoding cannot exceed 5 MB. MP4, AVI, and FLV formats are supported.
	VideoBase64 *string `json:"VideoBase64,omitempty" name:"VideoBase64"`

	// Liveness detection type. Valid values: LIP/ACTION/SILENT.
	// LIP: numeric mode; ACTION: motion mode; SILENT: silent mode. You need to select a mode to input.
	LivenessType *string `json:"LivenessType,omitempty" name:"LivenessType"`

	// Input parameter for the numeric mode: numeric verification code (1234). An API needs to be called first to get a numeric verification code;
	// Input parameter for the motion mode: motion order (2,1 or 1,2). An API needs to be called first to get the motion order;
	// Input parameter for silent mode: empty.
	ValidateData *string `json:"ValidateData,omitempty" name:"ValidateData"`

	// This parameter does not need to be passed in for this API.
	Optional *string `json:"Optional,omitempty" name:"Optional"`
}

func (r *LivenessCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessCompareRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LivenessCompareResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The best screenshot of the video after successful verification. The photo is Base64-encoded and in JPG format.
		BestFrameBase64 *string `json:"BestFrameBase64,omitempty" name:"BestFrameBase64"`

		// Similarity. Value range: [0.00, 100.00]. As a recommendation, when the similarity is greater than or equal to 70, it can be determined that the two faces are of the same person. You can adjust the threshold according to your specific scenario (the FAR at the threshold of 70 is 0.1%, and FAR at the threshold of 80 is 0.01%).
		Sim *float64 `json:"Sim,omitempty" name:"Sim"`

		// Service error code. `Success` will be returned for success. For error information, please see the `FailedOperation` section in the error code list below.
		Result *string `json:"Result,omitempty" name:"Result"`

		// Service result description.
		Description *string `json:"Description,omitempty" name:"Description"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LivenessCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LivenessCompareResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

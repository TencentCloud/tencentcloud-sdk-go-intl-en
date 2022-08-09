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

package v20190823

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type TextToVoiceRequestParams struct {
	// The source text for synthesizing speech, which is encoded in UTF-8.
	// It can contain up to 150 Chinese characters (a full-width punctuation as a Chinese character) or 500 letters ( a half-width punctuation as a letter).
	Text *string `json:"Text,omitempty" name:"Text"`

	// The `SessionId` of a request, which will be returned as-is. We recommend that you pass characters like uuid to prevent repetition.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Volume range: [0, 10], corresponding to 11 volume levels. 0 is the default value, indicating the normal volume. There is no mute option.
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`


	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// Project ID, which defaults to 0 and can be customized.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Model type, with `1` for the default model.
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// Standard voices <li>10510000-zhixiaoyao (Chinese)</li><li>1001-zhiyu (Chinese)</li><li>1002-zhiling (Chinese)</li><li>1003-zhimei (Chinese)</li><li>1004-zhiyun (Chinese)</li><li>1005-zhili (Chinese)</li><li>1007-zhina (Chinese)</li><li>1008-zhiqi (Chinese)</li><li>1009-zhiyun (Chinese)</li><li>1010-zhihua (Chinese)</li><li>1017-zhirong (Chinese)</li><li>1018-zhijing (Chinese)</li><li>1050-WeJack (English)</li><li>1051-WeRose (English)</li>Premium voices<br>Premium voices have higher fidelity and more natural-sounding quality than standard voices. For price details, see [Purchase Guide](https://intl.cloud.tencent.com/document/product/1073/34112?from_cn_redirect=1).<br><li>100510000-zhixiaoyao (Chinese)</li><li>101001-zhiyu (Chinese)</li><li>101002-zhiling (Chinese)</li><li>101003-zhimei (Chinese)</li><li>101004-zhiyun (Chinese)</li><li>101005-zhili (Chinese)</li><li>101006-zhiyan (Chinese)</li><li>101007-zhina (Chinese)</li><li>101008-zhiqi (Chinese)</li><li>101009-zhiyun (Chinese)</li><li>101010-zhihua (Chinese)</li><li>101011-zhiyan (Chinese)</li><li>101012-zhidan (Chinese)</li><li>101013-zhihui (Chinese)</li><li>101014-zhining (Chinese)</li><li>101015-zhimeng (Chinese)</li><li>101016-zhitian (Chinese)</li><li>101017-zhirong (Chinese)</li><li>101018-zhijing (Chinese)</li><li>101019-zhitong (Cantonese)</li><li>101020-zhigang (Chinese)</li><li>101021-zhirui (Chinese)</li><li>101022-zhihong (Chinese)</li><li>101023-zhixuan (Chinese)</li><li>101024-zhihao (Chinese)</li><li>101025-zhiwei (Chinese)</li><li>101026-zhixi (Chinese)</li><li>101027-zhimei (Chinese)</li><li>101028-zhijie (Chinese)</li><li>101029-zhikai (Chinese)</li><li>101030-zhike (Chinese)</li><li>101031-zhikui (Chinese)</li><li>101032-zhifang (Chinese)</li><li>101033-zhibei (Chinese)</li><li>101034-zhilian (Chinese)</li><li>101035-zhiyi (Chinese)</li><li>101040-zhichuan (Sichuan dialect)</li><li>101050-WeJack (English)</li><li>101051-WeRose (English)</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// Primary language type: <li>1 - Chinese (default)</li><li>2 - English</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// Audio sample rate: <li>16000: 16k (default)</li><li>8000: 8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Format of returned audio. Valid values: WAV (default), MP3, and PCM.
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type TextToVoiceRequest struct {
	*tchttp.BaseRequest
	
	// The source text for synthesizing speech, which is encoded in UTF-8.
	// It can contain up to 150 Chinese characters (a full-width punctuation as a Chinese character) or 500 letters ( a half-width punctuation as a letter).
	Text *string `json:"Text,omitempty" name:"Text"`

	// The `SessionId` of a request, which will be returned as-is. We recommend that you pass characters like uuid to prevent repetition.
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// Volume range: [0, 10], corresponding to 11 volume levels. 0 is the default value, indicating the normal volume. There is no mute option.
	Volume *float64 `json:"Volume,omitempty" name:"Volume"`

	Speed *float64 `json:"Speed,omitempty" name:"Speed"`

	// Project ID, which defaults to 0 and can be customized.
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// Model type, with `1` for the default model.
	ModelType *int64 `json:"ModelType,omitempty" name:"ModelType"`

	// Standard voices <li>10510000-zhixiaoyao (Chinese)</li><li>1001-zhiyu (Chinese)</li><li>1002-zhiling (Chinese)</li><li>1003-zhimei (Chinese)</li><li>1004-zhiyun (Chinese)</li><li>1005-zhili (Chinese)</li><li>1007-zhina (Chinese)</li><li>1008-zhiqi (Chinese)</li><li>1009-zhiyun (Chinese)</li><li>1010-zhihua (Chinese)</li><li>1017-zhirong (Chinese)</li><li>1018-zhijing (Chinese)</li><li>1050-WeJack (English)</li><li>1051-WeRose (English)</li>Premium voices<br>Premium voices have higher fidelity and more natural-sounding quality than standard voices. For price details, see [Purchase Guide](https://intl.cloud.tencent.com/document/product/1073/34112?from_cn_redirect=1).<br><li>100510000-zhixiaoyao (Chinese)</li><li>101001-zhiyu (Chinese)</li><li>101002-zhiling (Chinese)</li><li>101003-zhimei (Chinese)</li><li>101004-zhiyun (Chinese)</li><li>101005-zhili (Chinese)</li><li>101006-zhiyan (Chinese)</li><li>101007-zhina (Chinese)</li><li>101008-zhiqi (Chinese)</li><li>101009-zhiyun (Chinese)</li><li>101010-zhihua (Chinese)</li><li>101011-zhiyan (Chinese)</li><li>101012-zhidan (Chinese)</li><li>101013-zhihui (Chinese)</li><li>101014-zhining (Chinese)</li><li>101015-zhimeng (Chinese)</li><li>101016-zhitian (Chinese)</li><li>101017-zhirong (Chinese)</li><li>101018-zhijing (Chinese)</li><li>101019-zhitong (Cantonese)</li><li>101020-zhigang (Chinese)</li><li>101021-zhirui (Chinese)</li><li>101022-zhihong (Chinese)</li><li>101023-zhixuan (Chinese)</li><li>101024-zhihao (Chinese)</li><li>101025-zhiwei (Chinese)</li><li>101026-zhixi (Chinese)</li><li>101027-zhimei (Chinese)</li><li>101028-zhijie (Chinese)</li><li>101029-zhikai (Chinese)</li><li>101030-zhike (Chinese)</li><li>101031-zhikui (Chinese)</li><li>101032-zhifang (Chinese)</li><li>101033-zhibei (Chinese)</li><li>101034-zhilian (Chinese)</li><li>101035-zhiyi (Chinese)</li><li>101040-zhichuan (Sichuan dialect)</li><li>101050-WeJack (English)</li><li>101051-WeRose (English)</li>
	VoiceType *int64 `json:"VoiceType,omitempty" name:"VoiceType"`

	// Primary language type: <li>1 - Chinese (default)</li><li>2 - English</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitempty" name:"PrimaryLanguage"`

	// Audio sample rate: <li>16000: 16k (default)</li><li>8000: 8k</li>
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// Format of returned audio. Valid values: WAV (default), MP3, and PCM.
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

func (r *TextToVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "SessionId")
	delete(f, "Volume")
	delete(f, "Speed")
	delete(f, "ProjectId")
	delete(f, "ModelType")
	delete(f, "VoiceType")
	delete(f, "PrimaryLanguage")
	delete(f, "SampleRate")
	delete(f, "Codec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToVoiceResponseParams struct {
	// Base64-encoded WAV/MP3 audio data
	Audio *string `json:"Audio,omitempty" name:"Audio"`

	// The `SessionId` of a request
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextToVoiceResponse struct {
	*tchttp.BaseResponse
	Response *TextToVoiceResponseParams `json:"Response"`
}

func (r *TextToVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
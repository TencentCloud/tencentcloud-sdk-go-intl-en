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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

type Subtitle struct {
	// The word in the text that is sent.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The start timestamp of the word in the synthesized audio data, in milliseconds.
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// The end timestamp of the word in the synthesized audio data, in milliseconds.
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// The start index of the character in the whole sentence, starting from 0.
	BeginIndex *int64 `json:"BeginIndex,omitnil,omitempty" name:"BeginIndex"`

	// The end index of the character in the whole sentence, starting from 0.
	EndIndex *int64 `json:"EndIndex,omitnil,omitempty" name:"EndIndex"`

	// The phonemes of the word.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Phoneme *string `json:"Phoneme,omitnil,omitempty" name:"Phoneme"`
}

// Predefined struct for user
type TextToVoiceRequestParams struct {
	// The source text for synthesizing speech, which is encoded in UTF-8.
	// It can contain up to 150 Chinese characters (a full-width punctuation as a Chinese character) or 500 letters ( a half-width punctuation as a letter).
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The `SessionId` of a request, which will be returned as-is. We recommend that you pass characters like uuid to prevent repetition.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Volume range: [0, 10], corresponding to 11 volume levels. 0 is the default value, indicating the normal volume. There is no mute option.
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Speed range: [-2, 6], corresponding to different speeds<li>-2 for 0.6 times</li><li>-1 for 0.8 times</li><li>0 for 1.0 time (default)</li><li>1 for 1.2 times</li><li>2 for 1.5 times</li><li>6 for 2.5 times</li>To set finer-grained speed levels, keep one decimal place, such as 0.5, 1.1, and 1.8.<br>
	Speed *float64 `json:"Speed,omitnil,omitempty" name:"Speed"`

	// Project ID, which defaults to 0 and can be customized.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Model type, with `1` for the default model.
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// Standard voices <li>10510000-zhixiaoyao (Chinese)</li><li>1001-zhiyu (Chinese)</li><li>1002-zhiling (Chinese)</li><li>1003-zhimei (Chinese)</li><li>1004-zhiyun (Chinese)</li><li>1005-zhili (Chinese)</li><li>1007-zhina (Chinese)</li><li>1008-zhiqi (Chinese)</li><li>1009-zhiyun (Chinese)</li><li>1010-zhihua (Chinese)</li><li>1017-zhirong (Chinese)</li><li>1018-zhijing (Chinese)</li><li>1050-WeJack (English)</li><li>1051-WeRose (English)</li>Premium voices<br>Premium voices have higher fidelity and more natural-sounding quality than standard voices. For price details, see [Purchase Guide](https://www.tencentcloud.com/document/product/1154/47874).<br><li>100510000-zhixiaoyao (Chinese)</li><li>101001-zhiyu (Chinese)</li><li>101002-zhiling (Chinese)</li><li>101003-zhimei (Chinese)</li><li>101004-zhiyun (Chinese)</li><li>101005-zhili (Chinese)</li><li>101006-zhiyan (Chinese)</li><li>101007-zhina (Chinese)</li><li>101008-zhiqi (Chinese)</li><li>101009-zhiyun (Chinese)</li><li>101010-zhihua (Chinese)</li><li>101011-zhiyan (Chinese)</li><li>101012-zhidan (Chinese)</li><li>101013-zhihui (Chinese)</li><li>101014-zhining (Chinese)</li><li>101015-zhimeng (Chinese)</li><li>101016-zhitian (Chinese)</li><li>101017-zhirong (Chinese)</li><li>101018-zhijing (Chinese)</li><li>101019-zhitong (Cantonese)</li><li>101020-zhigang (Chinese)</li><li>101021-zhirui (Chinese)</li><li>101022-zhihong (Chinese)</li><li>101023-zhixuan (Chinese)</li><li>101024-zhihao (Chinese)</li><li>101025-zhiwei (Chinese)</li><li>101026-zhixi (Chinese)</li><li>101027-zhimei (Chinese)</li><li>101028-zhijie (Chinese)</li><li>101029-zhikai (Chinese)</li><li>101030-zhike (Chinese)</li><li>101031-zhikui (Chinese)</li><li>101032-zhifang (Chinese)</li><li>101033-zhibei (Chinese)</li><li>101034-zhilian (Chinese)</li><li>101035-zhiyi (Chinese)</li><li>101040-zhichuan (Sichuan dialect)</li><li>101050-WeJack (English)</li><li>101051-WeRose (English)</li><li>101052-zhiwei (Chinese)</li>
	// <li>101053-zhifang (Chinese)</li>
	// <li>101054-zhiyou (Chinese)</li>
	// <li>101055-zhiyou (Chinese)</li>
	// <li>101056-zhilin (Northeastern Mandarin)</li>
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// Primary language type: <li>1 - Chinese (default)</li><li>2 - English</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitnil,omitempty" name:"PrimaryLanguage"`

	// Audio sample rate: <li>16000: 16k (default)</li><li>8000: 8k</li>
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Format of returned audio. Valid values: WAV (default), MP3, and PCM.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Whether to enable the timestamp feature. Default value: `false`.
	EnableSubtitle *bool `json:"EnableSubtitle,omitnil,omitempty" name:"EnableSubtitle"`

	// The threshold of speech segmentation sensibility, which can be `0` (default), `1`, or `2`. A larger value indicates fewer segments, and the model tends to only segment sentences based on punctuation marks. We recommend you not change this parameter to avoid adverse effect on speech synthesis.
	SegmentRate *uint64 `json:"SegmentRate,omitnil,omitempty" name:"SegmentRate"`
}

type TextToVoiceRequest struct {
	*tchttp.BaseRequest
	
	// The source text for synthesizing speech, which is encoded in UTF-8.
	// It can contain up to 150 Chinese characters (a full-width punctuation as a Chinese character) or 500 letters ( a half-width punctuation as a letter).
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// The `SessionId` of a request, which will be returned as-is. We recommend that you pass characters like uuid to prevent repetition.
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Volume range: [0, 10], corresponding to 11 volume levels. 0 is the default value, indicating the normal volume. There is no mute option.
	Volume *float64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// Speed range: [-2, 6], corresponding to different speeds<li>-2 for 0.6 times</li><li>-1 for 0.8 times</li><li>0 for 1.0 time (default)</li><li>1 for 1.2 times</li><li>2 for 1.5 times</li><li>6 for 2.5 times</li>To set finer-grained speed levels, keep one decimal place, such as 0.5, 1.1, and 1.8.<br>
	Speed *float64 `json:"Speed,omitnil,omitempty" name:"Speed"`

	// Project ID, which defaults to 0 and can be customized.
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Model type, with `1` for the default model.
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// Standard voices <li>10510000-zhixiaoyao (Chinese)</li><li>1001-zhiyu (Chinese)</li><li>1002-zhiling (Chinese)</li><li>1003-zhimei (Chinese)</li><li>1004-zhiyun (Chinese)</li><li>1005-zhili (Chinese)</li><li>1007-zhina (Chinese)</li><li>1008-zhiqi (Chinese)</li><li>1009-zhiyun (Chinese)</li><li>1010-zhihua (Chinese)</li><li>1017-zhirong (Chinese)</li><li>1018-zhijing (Chinese)</li><li>1050-WeJack (English)</li><li>1051-WeRose (English)</li>Premium voices<br>Premium voices have higher fidelity and more natural-sounding quality than standard voices. For price details, see [Purchase Guide](https://www.tencentcloud.com/document/product/1154/47874).<br><li>100510000-zhixiaoyao (Chinese)</li><li>101001-zhiyu (Chinese)</li><li>101002-zhiling (Chinese)</li><li>101003-zhimei (Chinese)</li><li>101004-zhiyun (Chinese)</li><li>101005-zhili (Chinese)</li><li>101006-zhiyan (Chinese)</li><li>101007-zhina (Chinese)</li><li>101008-zhiqi (Chinese)</li><li>101009-zhiyun (Chinese)</li><li>101010-zhihua (Chinese)</li><li>101011-zhiyan (Chinese)</li><li>101012-zhidan (Chinese)</li><li>101013-zhihui (Chinese)</li><li>101014-zhining (Chinese)</li><li>101015-zhimeng (Chinese)</li><li>101016-zhitian (Chinese)</li><li>101017-zhirong (Chinese)</li><li>101018-zhijing (Chinese)</li><li>101019-zhitong (Cantonese)</li><li>101020-zhigang (Chinese)</li><li>101021-zhirui (Chinese)</li><li>101022-zhihong (Chinese)</li><li>101023-zhixuan (Chinese)</li><li>101024-zhihao (Chinese)</li><li>101025-zhiwei (Chinese)</li><li>101026-zhixi (Chinese)</li><li>101027-zhimei (Chinese)</li><li>101028-zhijie (Chinese)</li><li>101029-zhikai (Chinese)</li><li>101030-zhike (Chinese)</li><li>101031-zhikui (Chinese)</li><li>101032-zhifang (Chinese)</li><li>101033-zhibei (Chinese)</li><li>101034-zhilian (Chinese)</li><li>101035-zhiyi (Chinese)</li><li>101040-zhichuan (Sichuan dialect)</li><li>101050-WeJack (English)</li><li>101051-WeRose (English)</li><li>101052-zhiwei (Chinese)</li>
	// <li>101053-zhifang (Chinese)</li>
	// <li>101054-zhiyou (Chinese)</li>
	// <li>101055-zhiyou (Chinese)</li>
	// <li>101056-zhilin (Northeastern Mandarin)</li>
	VoiceType *int64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// Primary language type: <li>1 - Chinese (default)</li><li>2 - English</li>
	PrimaryLanguage *int64 `json:"PrimaryLanguage,omitnil,omitempty" name:"PrimaryLanguage"`

	// Audio sample rate: <li>16000: 16k (default)</li><li>8000: 8k</li>
	SampleRate *uint64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// Format of returned audio. Valid values: WAV (default), MP3, and PCM.
	Codec *string `json:"Codec,omitnil,omitempty" name:"Codec"`

	// Whether to enable the timestamp feature. Default value: `false`.
	EnableSubtitle *bool `json:"EnableSubtitle,omitnil,omitempty" name:"EnableSubtitle"`

	// The threshold of speech segmentation sensibility, which can be `0` (default), `1`, or `2`. A larger value indicates fewer segments, and the model tends to only segment sentences based on punctuation marks. We recommend you not change this parameter to avoid adverse effect on speech synthesis.
	SegmentRate *uint64 `json:"SegmentRate,omitnil,omitempty" name:"SegmentRate"`
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
	delete(f, "EnableSubtitle")
	delete(f, "SegmentRate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToVoiceResponseParams struct {
	// Base64-encoded WAV/MP3 audio data
	Audio *string `json:"Audio,omitnil,omitempty" name:"Audio"`

	// The `SessionId` of a request
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Timestamp information. If the timestamp feature is not enabled, an empty array will be returned.
	Subtitles []*Subtitle `json:"Subtitles,omitnil,omitempty" name:"Subtitles"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
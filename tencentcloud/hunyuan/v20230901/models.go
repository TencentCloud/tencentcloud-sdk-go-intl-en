// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20230901

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type ChatTranslationsRequestParams struct {
	// Model name. optional values include hunyuan-translation.
	// Please read the introduction in [the product overview](https://www.tencentcloud.com/document/product/1284/75277) for model descriptions.
	// 
	// Note:
	// Different models have different pricing. according to [the purchase guide](https://www.tencentcloud.com/document/product/1284/77186), call as needed.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Streaming call switch.
	// Note:.
	// 1. it defaults to non-streaming (false) when no value is passed.
	// 2. for streaming calls, the results are incrementally returned via the SSE protocol (the return value is taken from Choices[n].Delta, and incremental data must be concatenated to obtain the complete result).
	// 3. for non-streaming calls:.
	// The calling method is the same as an ordinary HTTP request.
	// The API response is time-consuming. if needed, set it to true for reduced latency.
	// Only return the final result once (return value takes the value from Choices[n].Message).
	// 
	// Note:.
	// When making an SDK call, streaming and non-streaming calls require **different ways** to obtain the return value. refer to the comments or sample code in the SDK (in the examples/hunyuan/v20230901/ directory of each language SDK code repository).
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// Text to be translated.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Source language.
	// Supported language list:. 
	// Simplified chinese: zh, traditional chinese: zh-TR, cantonese: yue, english: en, french: fr, portuguese: pt, spanish: es, japanese: ja, turkish: TR, russian: ru, arabic: ar, korean: ko, thai: th, italian: it, german: de, vietnamese: vi, malay: ms, indonesian: id.
	// The following languages are supported only by the hunyuan-translation model:.
	// Filipino: fil, hindi: hi, polish: pl, czech: cs, dutch: nl, khmer: km, burmese: my, persian: fa, gujarati: gu, urdu: ur, telugu: te, marathi: mr, hebrew: he, bengali: bn, tamil: ta, ukrainian: uk, tibetan: bo, kazakh: kk, mongolian: mn, uyghur: ug.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Target language.
	// Supported language list:. 
	// Simplified chinese: zh, traditional chinese: zh-TR, cantonese: yue, english: en, french: fr, portuguese: pt, spanish: es, japanese: ja, turkish: TR, russian: ru, arabic: ar, korean: ko, thai: th, italian: it, german: de, vietnamese: vi, malay: ms, indonesian: id.
	// The following languages are supported only by the hunyuan-translation model:.
	// Filipino: fil, hindi: hi, polish: pl, czech: cs, dutch: nl, khmer: km, burmese: my, persian: fa, gujarati: gu, urdu: ur, telugu: te, marathi: mr, hebrew: he, bengali: bn, tamil: ta, ukrainian: uk, tibetan: bo, kazakh: kk, mongolian: mn, uyghur: ug.
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// Domain of the text to be translated, such as game plot.
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// Reference example, up to 10.
	References []*Reference `json:"References,omitnil,omitempty" name:"References"`
}

type ChatTranslationsRequest struct {
	*tchttp.BaseRequest
	
	// Model name. optional values include hunyuan-translation.
	// Please read the introduction in [the product overview](https://www.tencentcloud.com/document/product/1284/75277) for model descriptions.
	// 
	// Note:
	// Different models have different pricing. according to [the purchase guide](https://www.tencentcloud.com/document/product/1284/77186), call as needed.
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Streaming call switch.
	// Note:.
	// 1. it defaults to non-streaming (false) when no value is passed.
	// 2. for streaming calls, the results are incrementally returned via the SSE protocol (the return value is taken from Choices[n].Delta, and incremental data must be concatenated to obtain the complete result).
	// 3. for non-streaming calls:.
	// The calling method is the same as an ordinary HTTP request.
	// The API response is time-consuming. if needed, set it to true for reduced latency.
	// Only return the final result once (return value takes the value from Choices[n].Message).
	// 
	// Note:.
	// When making an SDK call, streaming and non-streaming calls require **different ways** to obtain the return value. refer to the comments or sample code in the SDK (in the examples/hunyuan/v20230901/ directory of each language SDK code repository).
	Stream *bool `json:"Stream,omitnil,omitempty" name:"Stream"`

	// Text to be translated.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Source language.
	// Supported language list:. 
	// Simplified chinese: zh, traditional chinese: zh-TR, cantonese: yue, english: en, french: fr, portuguese: pt, spanish: es, japanese: ja, turkish: TR, russian: ru, arabic: ar, korean: ko, thai: th, italian: it, german: de, vietnamese: vi, malay: ms, indonesian: id.
	// The following languages are supported only by the hunyuan-translation model:.
	// Filipino: fil, hindi: hi, polish: pl, czech: cs, dutch: nl, khmer: km, burmese: my, persian: fa, gujarati: gu, urdu: ur, telugu: te, marathi: mr, hebrew: he, bengali: bn, tamil: ta, ukrainian: uk, tibetan: bo, kazakh: kk, mongolian: mn, uyghur: ug.
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// Target language.
	// Supported language list:. 
	// Simplified chinese: zh, traditional chinese: zh-TR, cantonese: yue, english: en, french: fr, portuguese: pt, spanish: es, japanese: ja, turkish: TR, russian: ru, arabic: ar, korean: ko, thai: th, italian: it, german: de, vietnamese: vi, malay: ms, indonesian: id.
	// The following languages are supported only by the hunyuan-translation model:.
	// Filipino: fil, hindi: hi, polish: pl, czech: cs, dutch: nl, khmer: km, burmese: my, persian: fa, gujarati: gu, urdu: ur, telugu: te, marathi: mr, hebrew: he, bengali: bn, tamil: ta, ukrainian: uk, tibetan: bo, kazakh: kk, mongolian: mn, uyghur: ug.
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// Domain of the text to be translated, such as game plot.
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// Reference example, up to 10.
	References []*Reference `json:"References,omitnil,omitempty" name:"References"`
}

func (r *ChatTranslationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatTranslationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Stream")
	delete(f, "Text")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "Field")
	delete(f, "References")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatTranslationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatTranslationsResponseParams struct {
	// Request'S RequestId this time.
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Disclaimer.
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// Unix timestamp, in seconds.
	Created *int64 `json:"Created,omitnil,omitempty" name:"Created"`

	// Token statistical information.
	// Billing by Token quantity.
	Usage *Usage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// Reply content.
	Choices []*TranslationChoice `json:"Choices,omitnil,omitempty" name:"Choices"`

	// Error message.
	// If the service encounters an exception during streaming return, return this error.
	ErrorMsg *ErrorMsg `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem. As a streaming response API, when the request is successfully completed, the RequestId will be placed in the Header "X-TC-RequestId" of the HTTP response.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatTranslationsResponse struct {
	tchttp.BaseSSEResponse `json:"-"`
	Response *ChatTranslationsResponseParams `json:"Response"`
}

func (r *ChatTranslationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatTranslationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Convert3DFormatRequestParams struct {
	// 3D file url address. model file size not greater than 60 mb
	// Supports fbx, obj, and glb format 3d file input
	File3D *string `json:"File3D,omitnil,omitempty" name:"File3D"`

	// Returns the 3D file format. valid values: 
	// STL, USDZ, FBX, MP4, GIF
	// Recommended input models below 50W, may timeout when selecting USDZ, MP4, or GIF format
	// Example value: STL.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type Convert3DFormatRequest struct {
	*tchttp.BaseRequest
	
	// 3D file url address. model file size not greater than 60 mb
	// Supports fbx, obj, and glb format 3d file input
	File3D *string `json:"File3D,omitnil,omitempty" name:"File3D"`

	// Returns the 3D file format. valid values: 
	// STL, USDZ, FBX, MP4, GIF
	// Recommended input models below 50W, may timeout when selecting USDZ, MP4, or GIF format
	// Example value: STL.
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

func (r *Convert3DFormatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Convert3DFormatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File3D")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "Convert3DFormatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Convert3DFormatResponseParams struct {
	// 3D file address
	ResultFile3D *string `json:"ResultFile3D,omitnil,omitempty" name:"ResultFile3D"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type Convert3DFormatResponse struct {
	*tchttp.BaseResponse
	Response *Convert3DFormatResponseParams `json:"Response"`
}

func (r *Convert3DFormatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Convert3DFormatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Describe3DSmartTopologyJobRequestParams struct {
	// Task ID.	
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type Describe3DSmartTopologyJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID.	
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *Describe3DSmartTopologyJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Describe3DSmartTopologyJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "Describe3DSmartTopologyJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Describe3DSmartTopologyJobResponseParams struct {
	// Task status. WAIT: waiting; RUN: running; FAIL: failed; DONE: successful. example value: RUN.	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code.	
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Error message.	
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// File generation URL address, with a valid period of 1 day.	
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type Describe3DSmartTopologyJobResponse struct {
	*tchttp.BaseResponse
	Response *Describe3DSmartTopologyJobResponseParams `json:"Response"`
}

func (r *Describe3DSmartTopologyJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Describe3DSmartTopologyJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHunyuanTo3DUVJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type DescribeHunyuanTo3DUVJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *DescribeHunyuanTo3DUVJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHunyuanTo3DUVJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHunyuanTo3DUVJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHunyuanTo3DUVJobResponseParams struct {
	// Task status
	// WAIT: waiting;
	// RUN: running; 
	// FAIL: failed; 
	// DONE: successful;
	// Example value: RUN
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// File generation URL address, with a valid period of 1 day
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHunyuanTo3DUVJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHunyuanTo3DUVJobResponseParams `json:"Response"`
}

func (r *DescribeHunyuanTo3DUVJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHunyuanTo3DUVJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorMsg struct {
	// Error message.
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// Error code.
	// 4000 internal service error.
	// 4001 request model timeout.
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`
}

type File3D struct {
	// 3D file format. valid values: GIF, OBJ.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Specifies the file Url (valid for 24 hours).
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Preview image Url.
	PreviewImageUrl *string `json:"PreviewImageUrl,omitnil,omitempty" name:"PreviewImageUrl"`
}

type ImageInfo struct {
	// Image Url.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Image Base64.
	Base64 *string `json:"Base64,omitnil,omitempty" name:"Base64"`
}

type InputFile3D struct {
	// Specifies the file Url with a valid period of 24 hours.	
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Specifies the file format.	
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type PromptTokensDetails struct {
	// The number of cache tokens.
	CachedTokens *string `json:"CachedTokens,omitnil,omitempty" name:"CachedTokens"`
}

// Predefined struct for user
type QueryHunyuan3DPartJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuan3DPartJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuan3DPartJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuan3DPartJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuan3DPartJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuan3DPartJobResponseParams struct {
	// Task status:
	// WAIT: waiting
	// RUN: running
	// FAIL: failed
	// DONE: successful
	// Example value: RUN
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Generates the file URL with a valid period of 1 day
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuan3DPartJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuan3DPartJobResponseParams `json:"Response"`
}

func (r *QueryHunyuan3DPartJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuan3DPartJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DProJobRequestParams struct {
	// Task ID.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanTo3DProJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID.
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanTo3DProJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DProJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanTo3DProJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DProJobResponseParams struct {
	// Task status:
	// WAIT: waiting
	// RUN: running; FAIL: failed
	// DONE: successful
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Describes the generated 3d file array
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanTo3DProJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanTo3DProJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanTo3DProJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DProJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DRapidJobRequestParams struct {
	// Task ID.	
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanTo3DRapidJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID.	
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanTo3DRapidJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DRapidJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanTo3DRapidJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DRapidJobResponseParams struct {
	// Task status:
	// WAIT: waiting
	// RUN: running
	// FAIL: failed
	// DONE: successful	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Error message	
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Specifies the generated 3D file array
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanTo3DRapidJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanTo3DRapidJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanTo3DRapidJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DRapidJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DTextureEditJobRequestParams struct {
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryHunyuanTo3DTextureEditJobRequest struct {
	*tchttp.BaseRequest
	
	// Task ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryHunyuanTo3DTextureEditJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DTextureEditJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryHunyuanTo3DTextureEditJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryHunyuanTo3DTextureEditJobResponseParams struct {
	// Task status
	// WAIT: waiting; 
	// RUN: running; 
	// FAIL: failed; 
	// DONE: successful;
	// Example value: RUN	
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Error code
	ErrorCode *string `json:"ErrorCode,omitnil,omitempty" name:"ErrorCode"`

	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// Generate the file URL with a valid period of 1 day
	ResultFile3Ds []*File3D `json:"ResultFile3Ds,omitnil,omitempty" name:"ResultFile3Ds"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryHunyuanTo3DTextureEditJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryHunyuanTo3DTextureEditJobResponseParams `json:"Response"`
}

func (r *QueryHunyuanTo3DTextureEditJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryHunyuanTo3DTextureEditJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Reference struct {
	// Translate text type, enumerate "sentence" means sentence, "term" means terminology.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Original.
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// Translation.
	Translation *string `json:"Translation,omitnil,omitempty" name:"Translation"`
}

// Predefined struct for user
type Submit3DSmartTopologyJobRequestParams struct {
	// Source 3D file model link
	// Reference value:glb,obj
	// One of the file formats is required.
	// Url: file size cannot exceed 200MB.
	// 3D model requirements: complex models and topologized models do not support face reduction. recommended input is non-topologized high-poly models, such as those generated by hunyuan 3d. high applicability categories include hard surface, game character, prop, daily objects.
	File3D *InputFile3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// Polygon type, indicates the model surface is composed of grid components, defaults to triangle
	// Reference value:
	// triangle: triangle face
	// quadrilateral: triangular and quadrilateral mixed face
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`

	// Reduction level bit type
	// valid values: high, medium, low
	FaceLevel *string `json:"FaceLevel,omitnil,omitempty" name:"FaceLevel"`
}

type Submit3DSmartTopologyJobRequest struct {
	*tchttp.BaseRequest
	
	// Source 3D file model link
	// Reference value:glb,obj
	// One of the file formats is required.
	// Url: file size cannot exceed 200MB.
	// 3D model requirements: complex models and topologized models do not support face reduction. recommended input is non-topologized high-poly models, such as those generated by hunyuan 3d. high applicability categories include hard surface, game character, prop, daily objects.
	File3D *InputFile3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// Polygon type, indicates the model surface is composed of grid components, defaults to triangle
	// Reference value:
	// triangle: triangle face
	// quadrilateral: triangular and quadrilateral mixed face
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`

	// Reduction level bit type
	// valid values: high, medium, low
	FaceLevel *string `json:"FaceLevel,omitnil,omitempty" name:"FaceLevel"`
}

func (r *Submit3DSmartTopologyJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Submit3DSmartTopologyJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File3D")
	delete(f, "PolygonType")
	delete(f, "FaceLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "Submit3DSmartTopologyJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type Submit3DSmartTopologyJobResponseParams struct {
	// Task ID.	
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type Submit3DSmartTopologyJobResponse struct {
	*tchttp.BaseResponse
	Response *Submit3DSmartTopologyJobResponseParams `json:"Response"`
}

func (r *Submit3DSmartTopologyJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *Submit3DSmartTopologyJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuan3DPartJobRequestParams struct {
	// Recommends inputting 3D models generated by AIGC
	// Recommended file size not greater than 100MB
	// Face count not greater than 30,000
	// Only supports FBX format
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

type SubmitHunyuan3DPartJobRequest struct {
	*tchttp.BaseRequest
	
	// Recommends inputting 3D models generated by AIGC
	// Recommended file size not greater than 100MB
	// Face count not greater than 30,000
	// Only supports FBX format
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

func (r *SubmitHunyuan3DPartJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuan3DPartJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuan3DPartJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuan3DPartJobResponseParams struct {
	// Task id (valid period: 24 hours)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuan3DPartJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuan3DPartJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuan3DPartJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuan3DPartJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DProJobRequestParams struct {
	// Tencent HY 3D Global model version
	// Defaults to 3.0, with optional choices: 3.0, 3.1
	// When selecting version 3.1, the [LowPoly] and [Sketch] parameter is unavailable
	// Example value:3.0
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Generates 3D content, describes 3D content.
	// Supports up to 1024 utf-8 characters.
	// Text-To-3D. Specifies either ImageBase64/ImageUrl or Prompt is required. Prompt and ImageBase64/ImageUrl cannot coexist.
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// Enter the Base64 code of the image.
	// Size: specifies the unilateral resolution requirement, not less than 128 and not greater than 5000. size should not exceed 8m (after encoding, it increases by about 30%, recommend actual input image size no more than 6m).
	// Input image suggestion:
	// 1.Simple background (solid-color background) 
	// 2.No text or blended colors
	// 3.Single object
	// 4.The object occupies over 50% of the frame
	// Valid values: jpg, png, jpeg, webp.
	// Specifies either ImageBase64/ImageUrl or Prompt is required. Prompt and ImageBase64/ImageUrl cannot coexist.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Input image Url.
	// Size: specifies the unilateral resolution requirement, not less than 128 and not greater than 5000. size should not exceed 8m (after encoding, it increases by about 30%, recommend actual input image size no more than 6m).
	// Input image suggestion:
	// 1.Simple background (solid-color background) 
	// 2.No text or blended colors
	// 3.Single object
	// 4.The object occupies over 50% of the frame
	// Valid values: jpg, png, jpeg, webp.
	// Specifies either ImageBase64/ImageUrl or Prompt is required. Prompt and ImageBase64/ImageUrl cannot coexist.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Multi-Perspective model image. reference value for viewing angle:.
	// left: Left view;
	// right: Right view;
	// back: Rear view;
	// top: Top view (only supported in Model 3.1);
	// bottom: Bottom view (only supported in Model 3.1);
	// left_front: Left front 45 degree view (only supported in Model 3.1);
	// right_front: Right front 45 degree view (only supported in Model 3.1);
	// 
	// Each perspective is limited to one image.
	// Image size limit. the value must not exceed 8 mb after encoding.
	// Image resolution limitation: the unilateral resolution should be less than 5000 and greater than 128.
	// Supported image format: JPG or PNG
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// Specifies whether PBR material generation is enabled. default false
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// Specifies the face count for 3D model generation. default value is 500000.
	// Specifies the supported face count generation range. value range: 40000-1500000
	FaceCount *int64 `json:"FaceCount,omitnil,omitempty" name:"FaceCount"`

	// Generation task type. default: Normal. valid values:
	// Normal: generates a geometric model with textures
	// LowPoly: specifies the model generated after intelligent polygon reduction.
	// Geometry: specifies whether to generate a Geometry model without textures (white model). when this task is selected, the EnablePBR parameter does not take effect
	// Specifies the Sketch for the generative model, allowing input of a Sketch or line drawing. in this mode, both prompt and ImageUrl/ImageBase64 can be entered together
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// This parameter only takes effect when LowPoly mode is selected from GenerateType
	// 
	// Polygon type, indicates the number of sides in the model's surface grid, defaults to triangle. valid values:
	// triangle. specifies the triangular face
	// quadrilateral: mix quadrangle and triangle faces to generate
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`
}

type SubmitHunyuanTo3DProJobRequest struct {
	*tchttp.BaseRequest
	
	// Tencent HY 3D Global model version
	// Defaults to 3.0, with optional choices: 3.0, 3.1
	// When selecting version 3.1, the [LowPoly] and [Sketch] parameter is unavailable
	// Example value:3.0
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// Generates 3D content, describes 3D content.
	// Supports up to 1024 utf-8 characters.
	// Text-To-3D. Specifies either ImageBase64/ImageUrl or Prompt is required. Prompt and ImageBase64/ImageUrl cannot coexist.
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// Enter the Base64 code of the image.
	// Size: specifies the unilateral resolution requirement, not less than 128 and not greater than 5000. size should not exceed 8m (after encoding, it increases by about 30%, recommend actual input image size no more than 6m).
	// Input image suggestion:
	// 1.Simple background (solid-color background) 
	// 2.No text or blended colors
	// 3.Single object
	// 4.The object occupies over 50% of the frame
	// Valid values: jpg, png, jpeg, webp.
	// Specifies either ImageBase64/ImageUrl or Prompt is required. Prompt and ImageBase64/ImageUrl cannot coexist.
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Input image Url.
	// Size: specifies the unilateral resolution requirement, not less than 128 and not greater than 5000. size should not exceed 8m (after encoding, it increases by about 30%, recommend actual input image size no more than 6m).
	// Input image suggestion:
	// 1.Simple background (solid-color background) 
	// 2.No text or blended colors
	// 3.Single object
	// 4.The object occupies over 50% of the frame
	// Valid values: jpg, png, jpeg, webp.
	// Specifies either ImageBase64/ImageUrl or Prompt is required. Prompt and ImageBase64/ImageUrl cannot coexist.
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Multi-Perspective model image. reference value for viewing angle:.
	// left: Left view;
	// right: Right view;
	// back: Rear view;
	// top: Top view (only supported in Model 3.1);
	// bottom: Bottom view (only supported in Model 3.1);
	// left_front: Left front 45 degree view (only supported in Model 3.1);
	// right_front: Right front 45 degree view (only supported in Model 3.1);
	// 
	// Each perspective is limited to one image.
	// Image size limit. the value must not exceed 8 mb after encoding.
	// Image resolution limitation: the unilateral resolution should be less than 5000 and greater than 128.
	// Supported image format: JPG or PNG
	MultiViewImages []*ViewImage `json:"MultiViewImages,omitnil,omitempty" name:"MultiViewImages"`

	// Specifies whether PBR material generation is enabled. default false
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// Specifies the face count for 3D model generation. default value is 500000.
	// Specifies the supported face count generation range. value range: 40000-1500000
	FaceCount *int64 `json:"FaceCount,omitnil,omitempty" name:"FaceCount"`

	// Generation task type. default: Normal. valid values:
	// Normal: generates a geometric model with textures
	// LowPoly: specifies the model generated after intelligent polygon reduction.
	// Geometry: specifies whether to generate a Geometry model without textures (white model). when this task is selected, the EnablePBR parameter does not take effect
	// Specifies the Sketch for the generative model, allowing input of a Sketch or line drawing. in this mode, both prompt and ImageUrl/ImageBase64 can be entered together
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// This parameter only takes effect when LowPoly mode is selected from GenerateType
	// 
	// Polygon type, indicates the number of sides in the model's surface grid, defaults to triangle. valid values:
	// triangle. specifies the triangular face
	// quadrilateral: mix quadrangle and triangle faces to generate
	PolygonType *string `json:"PolygonType,omitnil,omitempty" name:"PolygonType"`
}

func (r *SubmitHunyuanTo3DProJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DProJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "Prompt")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "MultiViewImages")
	delete(f, "EnablePBR")
	delete(f, "FaceCount")
	delete(f, "GenerateType")
	delete(f, "PolygonType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DProJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DProJobResponseParams struct {
	// Task ID (valid period: 24 hours).
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DProJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DProJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DProJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DProJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DRapidJobRequestParams struct {
	// Text-To-3D, description of 3D content, forward Prompt content
	// Supports up to 200 utf-8 characters
	// Either ImageBase64, ImageUrl, or Prompt is required, and Prompt cannot coexist with ImageBase64/ImageUrl
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// Input image Base64 data
	// Size: unilateral resolution requirement not less than 128, not greater than 5000, size not greater than 6mb (after encoding, size increases by approximately 30%). format:
	// jpg, png, jpeg, webp
	// Imagebase64, imageurl, and Prompt are required, but Prompt and imagebase64/imageurl cannot coexist
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Input image Url size: 
	// Unilateral resolution requirement not less than 128, not greater than 5000. size not greater than 8mb
	// Format: jpg, png, jpeg, webp
	// Imagebase64, imageurl, and Prompt are required, and Prompt cannot coexist with imagebase64/imageurl	
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Specifies whether PBR material generation is enabled, false by default.	
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// Specifies whether to enable the single geometry generation option
	// When enabled, it generates a 3D model without textures (white model)
	// When enabled, the generative model does not support OBJ format
	// Default model file format is GLB
	EnableGeometry *bool `json:"EnableGeometry,omitnil,omitempty" name:"EnableGeometry"`
}

type SubmitHunyuanTo3DRapidJobRequest struct {
	*tchttp.BaseRequest
	
	// Text-To-3D, description of 3D content, forward Prompt content
	// Supports up to 200 utf-8 characters
	// Either ImageBase64, ImageUrl, or Prompt is required, and Prompt cannot coexist with ImageBase64/ImageUrl
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// Input image Base64 data
	// Size: unilateral resolution requirement not less than 128, not greater than 5000, size not greater than 6mb (after encoding, size increases by approximately 30%). format:
	// jpg, png, jpeg, webp
	// Imagebase64, imageurl, and Prompt are required, but Prompt and imagebase64/imageurl cannot coexist
	ImageBase64 *string `json:"ImageBase64,omitnil,omitempty" name:"ImageBase64"`

	// Input image Url size: 
	// Unilateral resolution requirement not less than 128, not greater than 5000. size not greater than 8mb
	// Format: jpg, png, jpeg, webp
	// Imagebase64, imageurl, and Prompt are required, and Prompt cannot coexist with imagebase64/imageurl	
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// Specifies whether PBR material generation is enabled, false by default.	
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`

	// Specifies whether to enable the single geometry generation option
	// When enabled, it generates a 3D model without textures (white model)
	// When enabled, the generative model does not support OBJ format
	// Default model file format is GLB
	EnableGeometry *bool `json:"EnableGeometry,omitnil,omitempty" name:"EnableGeometry"`
}

func (r *SubmitHunyuanTo3DRapidJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DRapidJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "ImageBase64")
	delete(f, "ImageUrl")
	delete(f, "EnablePBR")
	delete(f, "EnableGeometry")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DRapidJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DRapidJobResponseParams struct {
	// Task ID (valid period: 24 hours)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DRapidJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DRapidJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DRapidJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DRapidJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DTextureEditJobRequestParams struct {
	// File URL of the 3D model file that requires texture edit
	// Supported formats: FBX, OBJ, GLB
	// 3D model limit: less than 100000 faces
	File3D *InputFile3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// Reference image for 3D model texture editing: Base64 data and image Url
	// Either Base64 or Url must be provided. if both are provided, Url prevails
	// Image restrictions: unilateral resolution less than 4096 and greater than 128, converted into Base64 string less than 10MB
	// Format support: jpg, jpeg, png
	Image *ImageInfo `json:"Image,omitnil,omitempty" name:"Image"`

	// Describes texture editing, forward prompt content
	// Supports up to 1024 utf-8 characters
	// Either image or prompt is required, and prompt and image cannot coexist
	// Example value: a kitten
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// Whether to enable the PBR texture parameter, only supports enabling when entering the Prompt
	// Example value: true
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

type SubmitHunyuanTo3DTextureEditJobRequest struct {
	*tchttp.BaseRequest
	
	// File URL of the 3D model file that requires texture edit
	// Supported formats: FBX, OBJ, GLB
	// 3D model limit: less than 100000 faces
	File3D *InputFile3D `json:"File3D,omitnil,omitempty" name:"File3D"`

	// Reference image for 3D model texture editing: Base64 data and image Url
	// Either Base64 or Url must be provided. if both are provided, Url prevails
	// Image restrictions: unilateral resolution less than 4096 and greater than 128, converted into Base64 string less than 10MB
	// Format support: jpg, jpeg, png
	Image *ImageInfo `json:"Image,omitnil,omitempty" name:"Image"`

	// Describes texture editing, forward prompt content
	// Supports up to 1024 utf-8 characters
	// Either image or prompt is required, and prompt and image cannot coexist
	// Example value: a kitten
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// Whether to enable the PBR texture parameter, only supports enabling when entering the Prompt
	// Example value: true
	EnablePBR *bool `json:"EnablePBR,omitnil,omitempty" name:"EnablePBR"`
}

func (r *SubmitHunyuanTo3DTextureEditJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DTextureEditJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File3D")
	delete(f, "Image")
	delete(f, "Prompt")
	delete(f, "EnablePBR")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DTextureEditJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DTextureEditJobResponseParams struct {
	// Task id (valid period 24 hours)
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DTextureEditJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DTextureEditJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DTextureEditJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DTextureEditJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DUVJobRequestParams struct {
	// File URL of the 3D model requiring UV unfold
	// Supported formats: FBX, OBJ, GLB
	// 3D model limit: less than 30000 faces
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

type SubmitHunyuanTo3DUVJobRequest struct {
	*tchttp.BaseRequest
	
	// File URL of the 3D model requiring UV unfold
	// Supported formats: FBX, OBJ, GLB
	// 3D model limit: less than 30000 faces
	File *InputFile3D `json:"File,omitnil,omitempty" name:"File"`
}

func (r *SubmitHunyuanTo3DUVJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DUVJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "File")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitHunyuanTo3DUVJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitHunyuanTo3DUVJobResponseParams struct {
	// Task ID.	
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitHunyuanTo3DUVJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitHunyuanTo3DUVJobResponseParams `json:"Response"`
}

func (r *SubmitHunyuanTo3DUVJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitHunyuanTo3DUVJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TranslationChoice struct {
	// End flag, can be stop or sensitive.
	// stop means output ends normally.
	// sensitive only appears when streaming output review is enabled, indicating security review not passed.
	FinishReason *string `json:"FinishReason,omitnil,omitempty" name:"FinishReason"`

	// Index value, used when streaming.
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// Incremental return value used when streaming this field.
	Delta *TranslationDelta `json:"Delta,omitnil,omitempty" name:"Delta"`

	// Return value, used when non-streaming.
	Message *TranslationMessage `json:"Message,omitnil,omitempty" name:"Message"`
}

type TranslationDelta struct {
	// Role name.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Content details.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type TranslationMessage struct {
	// Role. valid values: system, user, assistant, tool.
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// Text content.
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type Usage struct {
	// Input Token quantity.
	PromptTokens *int64 `json:"PromptTokens,omitnil,omitempty" name:"PromptTokens"`

	// Output Token quantity.
	CompletionTokens *int64 `json:"CompletionTokens,omitnil,omitempty" name:"CompletionTokens"`

	// Total Token quantity.
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// Details of the input token.
	PromptTokensDetails *PromptTokensDetails `json:"PromptTokensDetails,omitnil,omitempty" name:"PromptTokensDetails"`
}

type ViewImage struct {
	// Specifies the viewing angle type.
	// Valid values: back, left, right.
	ViewType *string `json:"ViewType,omitnil,omitempty" name:"ViewType"`

	// Image Url.
	ViewImageUrl *string `json:"ViewImageUrl,omitnil,omitempty" name:"ViewImageUrl"`

	// base64 address of the image.
	ViewImageBase64 *string `json:"ViewImageBase64,omitnil,omitempty" name:"ViewImageBase64"`
}
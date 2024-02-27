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

package v20200819

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type SendEmailRequestParams struct {
	// Sender
	FromAddress *string `json:"FromAddress,omitnil,omitempty" name:"FromAddress"`

	// Recipient
	ToAddress *string `json:"ToAddress,omitnil,omitempty" name:"ToAddress"`

	// Email summary
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Sender name
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// Reply-to address
	ReplyAddress *string `json:"ReplyAddress,omitnil,omitempty" name:"ReplyAddress"`

	// The body of an HTML email
	HtmlContent *string `json:"HtmlContent,omitnil,omitempty" name:"HtmlContent"`

	// The body of a plain-text email
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`
}

type SendEmailRequest struct {
	*tchttp.BaseRequest
	
	// Sender
	FromAddress *string `json:"FromAddress,omitnil,omitempty" name:"FromAddress"`

	// Recipient
	ToAddress *string `json:"ToAddress,omitnil,omitempty" name:"ToAddress"`

	// Email summary
	Subject *string `json:"Subject,omitnil,omitempty" name:"Subject"`

	// Sender name
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// Reply-to address
	ReplyAddress *string `json:"ReplyAddress,omitnil,omitempty" name:"ReplyAddress"`

	// The body of an HTML email
	HtmlContent *string `json:"HtmlContent,omitnil,omitempty" name:"HtmlContent"`

	// The body of a plain-text email
	TextContent *string `json:"TextContent,omitnil,omitempty" name:"TextContent"`
}

func (r *SendEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromAddress")
	delete(f, "ToAddress")
	delete(f, "Subject")
	delete(f, "FromName")
	delete(f, "ReplyAddress")
	delete(f, "HtmlContent")
	delete(f, "TextContent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendEmailResponseParams struct {
	// The result of creating an email task
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendEmailResponse struct {
	*tchttp.BaseResponse
	Response *SendEmailResponseParams `json:"Response"`
}

func (r *SendEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendTemplatedEmailRequestParams struct {
	// Sender address.
	FromAddress *string `json:"FromAddress,omitnil,omitempty" name:"FromAddress"`

	// Recipient address. Up to 100 recipient addresses are supported. Multiple addresses should be separated by semicolons (;).
	ToAddress *string `json:"ToAddress,omitnil,omitempty" name:"ToAddress"`

	// The name of the template created in advance.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template variable value, which is a JSON string.
	TemplateValue *string `json:"TemplateValue,omitnil,omitempty" name:"TemplateValue"`

	// Sender name.
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// Reply-to address.
	ReplyAddress *string `json:"ReplyAddress,omitnil,omitempty" name:"ReplyAddress"`
}

type SendTemplatedEmailRequest struct {
	*tchttp.BaseRequest
	
	// Sender address.
	FromAddress *string `json:"FromAddress,omitnil,omitempty" name:"FromAddress"`

	// Recipient address. Up to 100 recipient addresses are supported. Multiple addresses should be separated by semicolons (;).
	ToAddress *string `json:"ToAddress,omitnil,omitempty" name:"ToAddress"`

	// The name of the template created in advance.
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// Template variable value, which is a JSON string.
	TemplateValue *string `json:"TemplateValue,omitnil,omitempty" name:"TemplateValue"`

	// Sender name.
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// Reply-to address.
	ReplyAddress *string `json:"ReplyAddress,omitnil,omitempty" name:"ReplyAddress"`
}

func (r *SendTemplatedEmailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTemplatedEmailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromAddress")
	delete(f, "ToAddress")
	delete(f, "TemplateName")
	delete(f, "TemplateValue")
	delete(f, "FromName")
	delete(f, "ReplyAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendTemplatedEmailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendTemplatedEmailResponseParams struct {
	// The result of creating a template email task
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendTemplatedEmailResponse struct {
	*tchttp.BaseResponse
	Response *SendTemplatedEmailResponseParams `json:"Response"`
}

func (r *SendTemplatedEmailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendTemplatedEmailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
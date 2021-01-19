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

package v20201002

import (
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2020-10-02"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateEmailAddressRequest() (request *CreateEmailAddressRequest) {
    request = &CreateEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailAddress")
    return
}

func NewCreateEmailAddressResponse() (response *CreateEmailAddressResponse) {
    response = &CreateEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// After the sender domain is verified, you need a sender address to send emails. For example, if your sender domain is mail.qcloud.com, your sender address can be service@mail.qcloud.com. If you want to display your name (such as "Tencent Cloud") in the inbox list of the recipients, the sender address should be in the format of `Tencent Cloud <email address>`. Please note that there must be a space between your name and the first angle bracket.
func (c *Client) CreateEmailAddress(request *CreateEmailAddressRequest) (response *CreateEmailAddressResponse, err error) {
    if request == nil {
        request = NewCreateEmailAddressRequest()
    }
    response = NewCreateEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailIdentityRequest() (request *CreateEmailIdentityRequest) {
    request = &CreateEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailIdentity")
    return
}

func NewCreateEmailIdentityResponse() (response *CreateEmailIdentityResponse) {
    response = &CreateEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a sender domain. Before you can send an email using Tencent Cloud SES, you must create a sender domain as your identity. It can be the domain of your website or mobile app. You must verify the domain to prove that you own it and authorize Tencent Cloud SES to use it to send emails.
func (c *Client) CreateEmailIdentity(request *CreateEmailIdentityRequest) (response *CreateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewCreateEmailIdentityRequest()
    }
    response = NewCreateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmailTemplateRequest() (request *CreateEmailTemplateRequest) {
    request = &CreateEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "CreateEmailTemplate")
    return
}

func NewCreateEmailTemplateResponse() (response *CreateEmailTemplateResponse) {
    response = &CreateEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to create a TEXT or HTML email template. To create an HTML template, ensure that it does not include external CSS files. You can use {{variable name}} to specify a variable in the template.
// Note: only an approved template can be used to send emails.
func (c *Client) CreateEmailTemplate(request *CreateEmailTemplateRequest) (response *CreateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewCreateEmailTemplateRequest()
    }
    response = NewCreateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlackListRequest() (request *DeleteBlackListRequest) {
    request = &DeleteBlackListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteBlackList")
    return
}

func NewDeleteBlackListResponse() (response *DeleteBlackListResponse) {
    response = &DeleteBlackListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to unblocklist email addresses. If you confirm that a blocklisted recipient address is valid and active, you can remove it from Tencent Cloud’s address blocklist database.
func (c *Client) DeleteBlackList(request *DeleteBlackListRequest) (response *DeleteBlackListResponse, err error) {
    if request == nil {
        request = NewDeleteBlackListRequest()
    }
    response = NewDeleteBlackListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailAddressRequest() (request *DeleteEmailAddressRequest) {
    request = &DeleteEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailAddress")
    return
}

func NewDeleteEmailAddressResponse() (response *DeleteEmailAddressResponse) {
    response = &DeleteEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a sender address.
func (c *Client) DeleteEmailAddress(request *DeleteEmailAddressRequest) (response *DeleteEmailAddressResponse, err error) {
    if request == nil {
        request = NewDeleteEmailAddressRequest()
    }
    response = NewDeleteEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailIdentityRequest() (request *DeleteEmailIdentityRequest) {
    request = &DeleteEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailIdentity")
    return
}

func NewDeleteEmailIdentityResponse() (response *DeleteEmailIdentityResponse) {
    response = &DeleteEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete a sender domain. After deleted, the sender domain can no longer be used to send emails.
func (c *Client) DeleteEmailIdentity(request *DeleteEmailIdentityRequest) (response *DeleteEmailIdentityResponse, err error) {
    if request == nil {
        request = NewDeleteEmailIdentityRequest()
    }
    response = NewDeleteEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEmailTemplateRequest() (request *DeleteEmailTemplateRequest) {
    request = &DeleteEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "DeleteEmailTemplate")
    return
}

func NewDeleteEmailTemplateResponse() (response *DeleteEmailTemplateResponse) {
    response = &DeleteEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to delete an email template.
func (c *Client) DeleteEmailTemplate(request *DeleteEmailTemplateRequest) (response *DeleteEmailTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteEmailTemplateRequest()
    }
    response = NewDeleteEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmailIdentityRequest() (request *GetEmailIdentityRequest) {
    request = &GetEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetEmailIdentity")
    return
}

func NewGetEmailIdentityResponse() (response *GetEmailIdentityResponse) {
    response = &GetEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the configuration details of a sender domain.
func (c *Client) GetEmailIdentity(request *GetEmailIdentityRequest) (response *GetEmailIdentityResponse, err error) {
    if request == nil {
        request = NewGetEmailIdentityRequest()
    }
    response = NewGetEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmailTemplateRequest() (request *GetEmailTemplateRequest) {
    request = &GetEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetEmailTemplate")
    return
}

func NewGetEmailTemplateResponse() (response *GetEmailTemplateResponse) {
    response = &GetEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the details of a template.
func (c *Client) GetEmailTemplate(request *GetEmailTemplateRequest) (response *GetEmailTemplateResponse, err error) {
    if request == nil {
        request = NewGetEmailTemplateRequest()
    }
    response = NewGetEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewGetStatisticsReportRequest() (request *GetStatisticsReportRequest) {
    request = &GetStatisticsReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "GetStatisticsReport")
    return
}

func NewGetStatisticsReportResponse() (response *GetStatisticsReportResponse) {
    response = &GetStatisticsReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the email sending statistics over a recent period, including data on sent emails, delivery success rate, open rate, bounce rate, and so on. The maximum time span is 14 days.
func (c *Client) GetStatisticsReport(request *GetStatisticsReportRequest) (response *GetStatisticsReportResponse, err error) {
    if request == nil {
        request = NewGetStatisticsReportRequest()
    }
    response = NewGetStatisticsReportResponse()
    err = c.Send(request, response)
    return
}

func NewListBlackEmailAddressRequest() (request *ListBlackEmailAddressRequest) {
    request = &ListBlackEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListBlackEmailAddress")
    return
}

func NewListBlackEmailAddressResponse() (response *ListBlackEmailAddressResponse) {
    response = &ListBlackEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// The API is used to get blocklisted addresses. In the case of a hard bounce, Tencent Cloud will blocklist the recipient address and do not allow any user to send emails to this address. If you confirm that this is a misjudgment, you can remove it from the blocklist.
func (c *Client) ListBlackEmailAddress(request *ListBlackEmailAddressRequest) (response *ListBlackEmailAddressResponse, err error) {
    if request == nil {
        request = NewListBlackEmailAddressRequest()
    }
    response = NewListBlackEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailAddressRequest() (request *ListEmailAddressRequest) {
    request = &ListEmailAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailAddress")
    return
}

func NewListEmailAddressResponse() (response *ListEmailAddressResponse) {
    response = &ListEmailAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of sender addresses.
func (c *Client) ListEmailAddress(request *ListEmailAddressRequest) (response *ListEmailAddressResponse, err error) {
    if request == nil {
        request = NewListEmailAddressRequest()
    }
    response = NewListEmailAddressResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailIdentitiesRequest() (request *ListEmailIdentitiesRequest) {
    request = &ListEmailIdentitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailIdentities")
    return
}

func NewListEmailIdentitiesResponse() (response *ListEmailIdentitiesResponse) {
    response = &ListEmailIdentitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of sender domains, including verified and unverified domains.
func (c *Client) ListEmailIdentities(request *ListEmailIdentitiesRequest) (response *ListEmailIdentitiesResponse, err error) {
    if request == nil {
        request = NewListEmailIdentitiesRequest()
    }
    response = NewListEmailIdentitiesResponse()
    err = c.Send(request, response)
    return
}

func NewListEmailTemplatesRequest() (request *ListEmailTemplatesRequest) {
    request = &ListEmailTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "ListEmailTemplates")
    return
}

func NewListEmailTemplatesResponse() (response *ListEmailTemplatesResponse) {
    response = &ListEmailTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to get the list of email templates.
func (c *Client) ListEmailTemplates(request *ListEmailTemplatesRequest) (response *ListEmailTemplatesResponse, err error) {
    if request == nil {
        request = NewListEmailTemplatesRequest()
    }
    response = NewListEmailTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewSendEmailRequest() (request *SendEmailRequest) {
    request = &SendEmailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "SendEmail")
    return
}

func NewSendEmailResponse() (response *SendEmailResponse) {
    response = &SendEmailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to send a TEXT or HTML email. By default, you can only send emails using a template. To send custom content, please contact your sales rep to enable this feature.
func (c *Client) SendEmail(request *SendEmailRequest) (response *SendEmailResponse, err error) {
    if request == nil {
        request = NewSendEmailRequest()
    }
    response = NewSendEmailResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailIdentityRequest() (request *UpdateEmailIdentityRequest) {
    request = &UpdateEmailIdentityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailIdentity")
    return
}

func NewUpdateEmailIdentityResponse() (response *UpdateEmailIdentityResponse) {
    response = &UpdateEmailIdentityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to verify whether your DNS configuration is correct.
func (c *Client) UpdateEmailIdentity(request *UpdateEmailIdentityRequest) (response *UpdateEmailIdentityResponse, err error) {
    if request == nil {
        request = NewUpdateEmailIdentityRequest()
    }
    response = NewUpdateEmailIdentityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateEmailTemplateRequest() (request *UpdateEmailTemplateRequest) {
    request = &UpdateEmailTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ses", APIVersion, "UpdateEmailTemplate")
    return
}

func NewUpdateEmailTemplateResponse() (response *UpdateEmailTemplateResponse) {
    response = &UpdateEmailTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// This API is used to update an email template. An updated template must be approved again before it can be used.
func (c *Client) UpdateEmailTemplate(request *UpdateEmailTemplateRequest) (response *UpdateEmailTemplateResponse, err error) {
    if request == nil {
        request = NewUpdateEmailTemplateRequest()
    }
    response = NewUpdateEmailTemplateResponse()
    err = c.Send(request, response)
    return
}

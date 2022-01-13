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

package v20191012

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2019-10-12"

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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewAddEcdnDomainRequest() (request *AddEcdnDomainRequest) {
    request = &AddEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "AddEcdnDomain")
    
    
    return
}

func NewAddEcdnDomainResponse() (response *AddEcdnDomainResponse) {
    response = &AddEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddEcdnDomain
// This API is used to create an acceleration domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNCAMTAGKEYNOTEXIST = "InvalidParameter.EcdnCamTagKeyNotExist"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNRESOURCEMANYTAGKEY = "InvalidParameter.EcdnResourceManyTagKey"
//  INVALIDPARAMETER_ECDNTAGKEYINVALID = "InvalidParameter.EcdnTagKeyInvalid"
//  INVALIDPARAMETER_ECDNTAGKEYNOTEXIST = "InvalidParameter.EcdnTagKeyNotExist"
//  INVALIDPARAMETER_ECDNTAGKEYTOOMANYVALUE = "InvalidParameter.EcdnTagKeyTooManyValue"
//  INVALIDPARAMETER_ECDNTAGVALUEINVALID = "InvalidParameter.EcdnTagValueInvalid"
//  INVALIDPARAMETER_ECDNUSERTOOMANYTAGKEY = "InvalidParameter.EcdnUserTooManyTagKey"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) AddEcdnDomain(request *AddEcdnDomainRequest) (response *AddEcdnDomainResponse, err error) {
    if request == nil {
        request = NewAddEcdnDomainRequest()
    }
    
    response = NewAddEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

// AddEcdnDomain
// This API is used to create an acceleration domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNCAMTAGKEYNOTEXIST = "InvalidParameter.EcdnCamTagKeyNotExist"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNRESOURCEMANYTAGKEY = "InvalidParameter.EcdnResourceManyTagKey"
//  INVALIDPARAMETER_ECDNTAGKEYINVALID = "InvalidParameter.EcdnTagKeyInvalid"
//  INVALIDPARAMETER_ECDNTAGKEYNOTEXIST = "InvalidParameter.EcdnTagKeyNotExist"
//  INVALIDPARAMETER_ECDNTAGKEYTOOMANYVALUE = "InvalidParameter.EcdnTagKeyTooManyValue"
//  INVALIDPARAMETER_ECDNTAGVALUEINVALID = "InvalidParameter.EcdnTagValueInvalid"
//  INVALIDPARAMETER_ECDNUSERTOOMANYTAGKEY = "InvalidParameter.EcdnUserTooManyTagKey"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) AddEcdnDomainWithContext(ctx context.Context, request *AddEcdnDomainRequest) (response *AddEcdnDomainResponse, err error) {
    if request == nil {
        request = NewAddEcdnDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewAddEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEcdnDomainRequest() (request *DeleteEcdnDomainRequest) {
    request = &DeleteEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DeleteEcdnDomain")
    
    
    return
}

func NewDeleteEcdnDomainResponse() (response *DeleteEcdnDomainResponse) {
    response = &DeleteEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteEcdnDomain
// This API is used to delete a specified acceleration domain name. The acceleration domain name to be deleted must be in disabled status.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DeleteEcdnDomain(request *DeleteEcdnDomainRequest) (response *DeleteEcdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteEcdnDomainRequest()
    }
    
    response = NewDeleteEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

// DeleteEcdnDomain
// This API is used to delete a specified acceleration domain name. The acceleration domain name to be deleted must be in disabled status.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DeleteEcdnDomainWithContext(ctx context.Context, request *DeleteEcdnDomainRequest) (response *DeleteEcdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteEcdnDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomains
// This API is used to query the basic information of a CDN domain name, including the project ID, status, business type, creation time, update time, etc.
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomains
// This API is used to query the basic information of a CDN domain name, including the project ID, status, business type, creation time, update time, etc.
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomainsWithContext(ctx context.Context, request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsConfigRequest() (request *DescribeDomainsConfigRequest) {
    request = &DescribeDomainsConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeDomainsConfig")
    
    
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDomainsConfig
// This API is used to query the detailed configuration information of a CDN acceleration domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeDomainsConfig
// This API is used to query the detailed configuration information of a CDN acceleration domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomainsConfigWithContext(ctx context.Context, request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    if request == nil {
        request = NewDescribeDomainsConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeDomainsConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainLogsRequest() (request *DescribeEcdnDomainLogsRequest) {
    request = &DescribeEcdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainLogs")
    
    
    return
}

func NewDescribeEcdnDomainLogsResponse() (response *DescribeEcdnDomainLogsResponse) {
    response = &DescribeEcdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEcdnDomainLogs
// This API is used to query the access log download link of a domain name.
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
func (c *Client) DescribeEcdnDomainLogs(request *DescribeEcdnDomainLogsRequest) (response *DescribeEcdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainLogsRequest()
    }
    
    response = NewDescribeEcdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

// DescribeEcdnDomainLogs
// This API is used to query the access log download link of a domain name.
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
func (c *Client) DescribeEcdnDomainLogsWithContext(ctx context.Context, request *DescribeEcdnDomainLogsRequest) (response *DescribeEcdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainLogsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEcdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnDomainStatisticsRequest() (request *DescribeEcdnDomainStatisticsRequest) {
    request = &DescribeEcdnDomainStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnDomainStatistics")
    
    
    return
}

func NewDescribeEcdnDomainStatisticsResponse() (response *DescribeEcdnDomainStatisticsResponse) {
    response = &DescribeEcdnDomainStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEcdnDomainStatistics
// This API is used to query the statistical metrics of domain name access within a specified time period.
//
// error code that may be returned:
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
func (c *Client) DescribeEcdnDomainStatistics(request *DescribeEcdnDomainStatisticsRequest) (response *DescribeEcdnDomainStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainStatisticsRequest()
    }
    
    response = NewDescribeEcdnDomainStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeEcdnDomainStatistics
// This API is used to query the statistical metrics of domain name access within a specified time period.
//
// error code that may be returned:
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
func (c *Client) DescribeEcdnDomainStatisticsWithContext(ctx context.Context, request *DescribeEcdnDomainStatisticsRequest) (response *DescribeEcdnDomainStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnDomainStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEcdnDomainStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEcdnStatisticsRequest() (request *DescribeEcdnStatisticsRequest) {
    request = &DescribeEcdnStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeEcdnStatistics")
    
    
    return
}

func NewDescribeEcdnStatisticsResponse() (response *DescribeEcdnStatisticsResponse) {
    response = &DescribeEcdnStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeEcdnStatistics
// This API is used to query ECDN real-time access monitoring data and supports the following metrics:
//
// 
//
// + Traffic (in bytes)
//
// + Bandwidth (in bps)
//
// + Number of requests
//
// + Number of 2xx status codes and details of status codes starting with 2
//
// + Number of 3xx status codes and details of status codes starting with 3
//
// + Number of 4xx status codes and details of status codes starting with 4
//
// + Number of 5xx status codes and details of status codes starting with 5
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINVALIDPARAMINTERVAL = "InvalidParameter.EcdnInvalidParamInterval"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
func (c *Client) DescribeEcdnStatistics(request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnStatisticsRequest()
    }
    
    response = NewDescribeEcdnStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeEcdnStatistics
// This API is used to query ECDN real-time access monitoring data and supports the following metrics:
//
// 
//
// + Traffic (in bytes)
//
// + Bandwidth (in bps)
//
// + Number of requests
//
// + Number of 2xx status codes and details of status codes starting with 2
//
// + Number of 3xx status codes and details of status codes starting with 3
//
// + Number of 4xx status codes and details of status codes starting with 4
//
// + Number of 5xx status codes and details of status codes starting with 5
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINVALIDPARAMINTERVAL = "InvalidParameter.EcdnInvalidParamInterval"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNSTATINVALIDDATE = "InvalidParameter.EcdnStatInvalidDate"
//  INVALIDPARAMETER_ECDNSTATINVALIDMETRIC = "InvalidParameter.EcdnStatInvalidMetric"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.CdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_DOMAINNOPERMISSION = "UnauthorizedOperation.DomainNoPermission"
//  UNAUTHORIZEDOPERATION_DOMAINSNOPERMISSION = "UnauthorizedOperation.DomainsNoPermission"
//  UNAUTHORIZEDOPERATION_ECDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.EcdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNNODOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnNoDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTNOPERMISSION = "UnauthorizedOperation.ProjectNoPermission"
//  UNAUTHORIZEDOPERATION_PROJECTSNOPERMISSION = "UnauthorizedOperation.ProjectsNoPermission"
func (c *Client) DescribeEcdnStatisticsWithContext(ctx context.Context, request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeEcdnStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIpStatusRequest() (request *DescribeIpStatusRequest) {
    request = &DescribeIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribeIpStatus")
    
    
    return
}

func NewDescribeIpStatusResponse() (response *DescribeIpStatusResponse) {
    response = &DescribeIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIpStatus
// This API is used to query the detailed node information of the acceleration platform to which the domain name is connected.
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    
    response = NewDescribeIpStatusResponse()
    err = c.Send(request, response)
    return
}

// DescribeIpStatus
// This API is used to query the detailed node information of the acceleration platform to which the domain name is connected.
//
// error code that may be returned:
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNHOSTNOTEXISTS = "ResourceNotFound.EcdnHostNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) DescribeIpStatusWithContext(ctx context.Context, request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeIpStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeQuotaRequest() (request *DescribePurgeQuotaRequest) {
    request = &DescribePurgeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeQuota")
    
    
    return
}

func NewDescribePurgeQuotaResponse() (response *DescribePurgeQuotaResponse) {
    response = &DescribePurgeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeQuota
// This API is used to query the usage quota of the purge API.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeQuota(request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePurgeQuotaRequest()
    }
    
    response = NewDescribePurgeQuotaResponse()
    err = c.Send(request, response)
    return
}

// DescribePurgeQuota
// This API is used to query the usage quota of the purge API.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeQuotaWithContext(ctx context.Context, request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePurgeQuotaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePurgeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeTasksRequest() (request *DescribePurgeTasksRequest) {
    request = &DescribePurgeTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePurgeTasks
// This API is used to query the submission history of purge tasks and their execution progress.
//
// error code that may be returned:
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribePurgeTasks
// This API is used to query the submission history of purge tasks and their execution progress.
//
// error code that may be returned:
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePurgeTasksResponse()
    err = c.Send(request, response)
    return
}

func NewPurgePathCacheRequest() (request *PurgePathCacheRequest) {
    request = &PurgePathCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgePathCache")
    
    
    return
}

func NewPurgePathCacheResponse() (response *PurgePathCacheResponse) {
    response = &PurgePathCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgePathCache
// This API is used to purge cache directories in batches. One purge task ID will be returned for each submission.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"
//  LIMITEXCEEDED_ECDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgePathExceedBatchLimit"
//  LIMITEXCEEDED_ECDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgePathExceedDayLimit"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgePathCache(request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    if request == nil {
        request = NewPurgePathCacheRequest()
    }
    
    response = NewPurgePathCacheResponse()
    err = c.Send(request, response)
    return
}

// PurgePathCache
// This API is used to purge cache directories in batches. One purge task ID will be returned for each submission.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"
//  LIMITEXCEEDED_ECDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgePathExceedBatchLimit"
//  LIMITEXCEEDED_ECDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgePathExceedDayLimit"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgePathCacheWithContext(ctx context.Context, request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    if request == nil {
        request = NewPurgePathCacheRequest()
    }
    request.SetContext(ctx)
    
    response = NewPurgePathCacheResponse()
    err = c.Send(request, response)
    return
}

func NewPurgeUrlsCacheRequest() (request *PurgeUrlsCacheRequest) {
    request = &PurgeUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "PurgeUrlsCache")
    
    
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PurgeUrlsCache
// This API is used to batch purge URLs. One purge task ID will be returned for each submission.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

// PurgeUrlsCache
// This API is used to batch purge URLs. One purge task ID will be returned for each submission.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNDOMAININVALIDSTATUS = "InvalidParameter.EcdnDomainInvalidStatus"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  INVALIDPARAMETER_ECDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.EcdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_ECDNURLEXCEEDLENGTH = "InvalidParameter.EcdnUrlExceedLength"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.EcdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_ECDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.EcdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgeUrlsCacheWithContext(ctx context.Context, request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    request.SetContext(ctx)
    
    response = NewPurgeUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewStartEcdnDomainRequest() (request *StartEcdnDomainRequest) {
    request = &StartEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "StartEcdnDomain")
    
    
    return
}

func NewStartEcdnDomainResponse() (response *StartEcdnDomainResponse) {
    response = &StartEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartEcdnDomain
// This API is used to enable an acceleration domain name. The domain name to be enabled must be in deactivated status.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StartEcdnDomain(request *StartEcdnDomainRequest) (response *StartEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStartEcdnDomainRequest()
    }
    
    response = NewStartEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

// StartEcdnDomain
// This API is used to enable an acceleration domain name. The domain name to be enabled must be in deactivated status.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StartEcdnDomainWithContext(ctx context.Context, request *StartEcdnDomainRequest) (response *StartEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStartEcdnDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewStartEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopEcdnDomainRequest() (request *StopEcdnDomainRequest) {
    request = &StopEcdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "StopEcdnDomain")
    
    
    return
}

func NewStopEcdnDomainResponse() (response *StopEcdnDomainResponse) {
    response = &StopEcdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopEcdnDomain
// This API is used to disable an acceleration domain name. The domain name to be disabled must be in enabled or deploying status.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StopEcdnDomain(request *StopEcdnDomainRequest) (response *StopEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStopEcdnDomainRequest()
    }
    
    response = NewStopEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

// StopEcdnDomain
// This API is used to disable an acceleration domain name. The domain name to be disabled must be in enabled or deploying status.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StopEcdnDomainWithContext(ctx context.Context, request *StopEcdnDomainRequest) (response *StopEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStopEcdnDomainRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopEcdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDomainConfigRequest() (request *UpdateDomainConfigRequest) {
    request = &UpdateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ecdn", APIVersion, "UpdateDomainConfig")
    
    
    return
}

func NewUpdateDomainConfigResponse() (response *UpdateDomainConfigResponse) {
    response = &UpdateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDomainConfig
// This API is used to update the configuration information of an ECDN acceleration domain name.
//
// Note: if you need to update a complex configuration item, you must pass in all attributes of the entire object, and the default values will be used for the attributes that are not passed in. You are recommended to get the configuration attribute through the query API first and then directly modify and pass it to this API. Due to the special nature of the certificate for HTTPS configuration, you do not need to pass in the certificate and key fields during the update.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNINVALIDPARAMAREA = "InvalidParameter.EcdnInvalidParamArea"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

// UpdateDomainConfig
// This API is used to update the configuration information of an ECDN acceleration domain name.
//
// Note: if you need to update a complex configuration item, you must pass in all attributes of the entire object, and the default values will be used for the attributes that are not passed in. You are recommended to get the configuration attribute through the query API first and then directly modify and pass it to this API. Due to the special nature of the certificate for HTTPS configuration, you do not need to pass in the certificate and key fields during the update.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNINVALIDPARAMAREA = "InvalidParameter.EcdnInvalidParamArea"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNPROJECTNOTEXISTS = "ResourceNotFound.EcdnProjectNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) UpdateDomainConfigWithContext(ctx context.Context, request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

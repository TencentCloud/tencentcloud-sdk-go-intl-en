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
    "errors"
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41123?from_cn_redirect=1">corresponding CDN API</a>.
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
//  LIMITEXCEEDED_ECDNUSERTOOMANYDOMAINS = "LimitExceeded.EcdnUserTooManyDomains"
//  RESOURCEINUSE_ECDNDOMAINEXISTS = "ResourceInUse.EcdnDomainExists"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.EcdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_ECDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.EcdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) AddEcdnDomain(request *AddEcdnDomainRequest) (response *AddEcdnDomainResponse, err error) {
    return c.AddEcdnDomainWithContext(context.Background(), request)
}

// AddEcdnDomain
// This API is used to create an acceleration domain name.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41123?from_cn_redirect=1">corresponding CDN API</a>.
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
//  LIMITEXCEEDED_ECDNUSERTOOMANYDOMAINS = "LimitExceeded.EcdnUserTooManyDomains"
//  RESOURCEINUSE_ECDNDOMAINEXISTS = "ResourceInUse.EcdnDomainExists"
//  RESOURCEINUSE_ECDNOPINPROGRESS = "ResourceInUse.EcdnOpInProgress"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.EcdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_ECDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.EcdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) AddEcdnDomainWithContext(ctx context.Context, request *AddEcdnDomainRequest) (response *AddEcdnDomainResponse, err error) {
    if request == nil {
        request = NewAddEcdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddEcdnDomain require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/570/42471?from_cn_redirect=1">corresponding CDN API</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"
//  RESOURCEUNAVAILABLE_ECDNDOMAINISNOTOFFLINE = "ResourceUnavailable.EcdnDomainIsNotOffline"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DeleteEcdnDomain(request *DeleteEcdnDomainRequest) (response *DeleteEcdnDomainResponse, err error) {
    return c.DeleteEcdnDomainWithContext(context.Background(), request)
}

// DeleteEcdnDomain
// This API is used to delete a specified acceleration domain name. The acceleration domain name to be deleted must be in disabled status.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/570/42471?from_cn_redirect=1">corresponding CDN API</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"
//  RESOURCEUNAVAILABLE_ECDNDOMAINISNOTOFFLINE = "ResourceUnavailable.EcdnDomainIsNotOffline"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DeleteEcdnDomainWithContext(ctx context.Context, request *DeleteEcdnDomainRequest) (response *DeleteEcdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteEcdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteEcdnDomain require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41118?from_cn_redirect=1">corresponding CDN API</a>.
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
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// This API is used to query the basic information of a CDN domain name, including the project ID, status, business type, creation time, update time, etc.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41118?from_cn_redirect=1">corresponding CDN API</a>.
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomains require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41117?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"
//  INVALIDPARAMETER_ECDNCONFIGINVALIDCACHE = "InvalidParameter.EcdnConfigInvalidCache"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  LIMITEXCEEDED_ECDNDOMAINOPTOOOFTEN = "LimitExceeded.EcdnDomainOpTooOften"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    return c.DescribeDomainsConfigWithContext(context.Background(), request)
}

// DescribeDomainsConfig
// This API is used to query the detailed configuration information of a CDN acceleration domain name.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41117?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNCONFIGERROR = "InternalError.EcdnConfigError"
//  INTERNALERROR_ECDNDBERROR = "InternalError.EcdnDbError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNCERTNOCERTINFO = "InvalidParameter.EcdnCertNoCertInfo"
//  INVALIDPARAMETER_ECDNCONFIGINVALIDCACHE = "InvalidParameter.EcdnConfigInvalidCache"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDomainsConfig require credential")
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
    return c.DescribeEcdnDomainLogsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEcdnDomainLogs require credential")
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
    return c.DescribeEcdnDomainStatisticsWithContext(context.Background(), request)
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEcdnDomainStatistics require credential")
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
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
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
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeEcdnStatistics(request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    return c.DescribeEcdnStatisticsWithContext(context.Background(), request)
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
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
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
//  UNAUTHORIZEDOPERATION_UNKNOWN = "UnauthorizedOperation.Unknown"
func (c *Client) DescribeEcdnStatisticsWithContext(ctx context.Context, request *DescribeEcdnStatisticsRequest) (response *DescribeEcdnStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEcdnStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEcdnStatistics require credential")
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
// This API is used to query ECDN node IPs. This API is only available to beta users. Please submit a ticket to use it.
//
// 
//
// If you need to add the node IPs to your origin allowlist, keep querying the updating the IPs regularly to ensure the success of origin forwarding. 
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
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    return c.DescribeIpStatusWithContext(context.Background(), request)
}

// DescribeIpStatus
// This API is used to query ECDN node IPs. This API is only available to beta users. Please submit a ticket to use it.
//
// 
//
// If you need to add the node IPs to your origin allowlist, keep querying the updating the IPs regularly to ensure the success of origin forwarding. 
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
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNHOSTUNAUTHORIZED = "UnauthorizedOperation.EcdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) DescribeIpStatusWithContext(ctx context.Context, request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    if request == nil {
        request = NewDescribeIpStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpStatus require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41956?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeQuota(request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    return c.DescribePurgeQuotaWithContext(context.Background(), request)
}

// DescribePurgeQuota
// This API is used to query the usage quota of the purge API.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41956?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  FAILEDOPERATION_ECDNCONFIGERROR = "FailedOperation.EcdnConfigError"
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeQuotaWithContext(ctx context.Context, request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePurgeQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeQuota require credential")
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
// This API is used to query the submission record and progress of purge tasks.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/37873?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// This API is used to query the submission record and progress of purge tasks.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/37873?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  INTERNALERROR_ECDNSYSTEMERROR = "InternalError.EcdnSystemError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  INVALIDPARAMETER_ECDNPARAMERROR = "InvalidParameter.EcdnParamError"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCENOTFOUND_ECDNUSERNOTEXISTS = "ResourceNotFound.EcdnUserNotExists"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) DescribePurgeTasksWithContext(ctx context.Context, request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    if request == nil {
        request = NewDescribePurgeTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePurgeTasks require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/570/42475?from_cn_redirect=1">corresponding CDN API</a>.
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
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgePathCache(request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    return c.PurgePathCacheWithContext(context.Background(), request)
}

// PurgePathCache
// This API is used to purge cache directories in batches. One purge task ID will be returned for each submission.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/570/42475?from_cn_redirect=1">corresponding CDN API</a>.
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
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgePathCacheWithContext(ctx context.Context, request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    if request == nil {
        request = NewPurgePathCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PurgePathCache require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/37870?from_cn_redirect=1">corresponding CDN API</a>.
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
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    return c.PurgeUrlsCacheWithContext(context.Background(), request)
}

// PurgeUrlsCache
// This API is used to batch purge URLs. One purge task ID will be returned for each submission.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/37870?from_cn_redirect=1">corresponding CDN API</a>.
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
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) PurgeUrlsCacheWithContext(ctx context.Context, request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPurgeUrlsCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PurgeUrlsCache require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/product/228/41121?from_cn_redirect=1">corresponding CDN API</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"
//  RESOURCEUNAVAILABLE_ECDNDOMAINISNOTOFFLINE = "ResourceUnavailable.EcdnDomainIsNotOffline"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StartEcdnDomain(request *StartEcdnDomainRequest) (response *StartEcdnDomainResponse, err error) {
    return c.StartEcdnDomainWithContext(context.Background(), request)
}

// StartEcdnDomain
// This API is used to enable an acceleration domain name. The domain name to be enabled must be in deactivated status.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/product/228/41121?from_cn_redirect=1">corresponding CDN API</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"
//  RESOURCEUNAVAILABLE_ECDNDOMAINISNOTOFFLINE = "ResourceUnavailable.EcdnDomainIsNotOffline"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StartEcdnDomainWithContext(ctx context.Context, request *StartEcdnDomainRequest) (response *StartEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStartEcdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartEcdnDomain require credential")
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
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/product/228/41120?from_cn_redirect=1">corresponding CDN API</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISNOTONLINE = "ResourceUnavailable.EcdnDomainIsNotOnline"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StopEcdnDomain(request *StopEcdnDomainRequest) (response *StopEcdnDomainResponse, err error) {
    return c.StopEcdnDomainWithContext(context.Background(), request)
}

// StopEcdnDomain
// This API is used to disable an acceleration domain name. The domain name to be disabled must be in enabled or deploying status.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/product/228/41120?from_cn_redirect=1">corresponding CDN API</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISNOTONLINE = "ResourceUnavailable.EcdnDomainIsNotOnline"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
func (c *Client) StopEcdnDomainWithContext(ctx context.Context, request *StopEcdnDomainRequest) (response *StopEcdnDomainResponse, err error) {
    if request == nil {
        request = NewStopEcdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopEcdnDomain require credential")
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
// Note: if you need to update complex configuration items, you must pass all the attributes of the entire object. The default value will be used for attributes that are not passed. We recommend calling the querying API to obtain the configuration attributes first. You can then modify and pass the attributes to the API. The certificate and key fields do not need to be passed for HTTPS configuration.
//
// 
//
// >?  If your application has been migrated to Tencent Cloud CDN, you can use <a href="https://intl.cloud.tencent.com/document/product/228/41116?from_cn_redirect=1">CDN APIs</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    return c.UpdateDomainConfigWithContext(context.Background(), request)
}

// UpdateDomainConfig
// This API is used to update the configuration information of an ECDN acceleration domain name.
//
// Note: if you need to update complex configuration items, you must pass all the attributes of the entire object. The default value will be used for attributes that are not passed. We recommend calling the querying API to obtain the configuration attributes first. You can then modify and pass the attributes to the API. The certificate and key fields do not need to be passed for HTTPS configuration.
//
// 
//
// >?  If your application has been migrated to Tencent Cloud CDN, you can use <a href="https://intl.cloud.tencent.com/document/product/228/41116?from_cn_redirect=1">CDN APIs</a>.
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
//  RESOURCEUNAVAILABLE_ECDNDOMAINISLOCKED = "ResourceUnavailable.EcdnDomainIsLocked"
//  UNAUTHORIZEDOPERATION_ECDNCAMUNAUTHORIZED = "UnauthorizedOperation.EcdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNDOMAINUNAUTHORIZED = "UnauthorizedOperation.EcdnDomainUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_ECDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.EcdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_ECDNUSERISSUSPENDED = "UnauthorizedOperation.EcdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNUSERNOWHITELIST = "UnauthorizedOperation.EcdnUserNoWhitelist"
func (c *Client) UpdateDomainConfigWithContext(ctx context.Context, request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    if request == nil {
        request = NewUpdateDomainConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDomainConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDomainConfigResponse()
    err = c.Send(request, response)
    return
}

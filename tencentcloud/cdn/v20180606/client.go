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

package v20180606

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-06-06"

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


func NewAddCLSTopicDomainsRequest() (request *AddCLSTopicDomainsRequest) {
    request = &AddCLSTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "AddCLSTopicDomains")
    
    
    return
}

func NewAddCLSTopicDomainsResponse() (response *AddCLSTopicDomainsResponse) {
    response = &AddCLSTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCLSTopicDomains
// This API is used to add one or more domains to a specified log topic.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
func (c *Client) AddCLSTopicDomains(request *AddCLSTopicDomainsRequest) (response *AddCLSTopicDomainsResponse, err error) {
    return c.AddCLSTopicDomainsWithContext(context.Background(), request)
}

// AddCLSTopicDomains
// This API is used to add one or more domains to a specified log topic.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
func (c *Client) AddCLSTopicDomainsWithContext(ctx context.Context, request *AddCLSTopicDomainsRequest) (response *AddCLSTopicDomainsResponse, err error) {
    if request == nil {
        request = NewAddCLSTopicDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCLSTopicDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCLSTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewAddCdnDomainRequest() (request *AddCdnDomainRequest) {
    request = &AddCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "AddCdnDomain")
    
    
    return
}

func NewAddCdnDomainResponse() (response *AddCdnDomainResponse) {
    response = &AddCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddCdnDomain
// This API is used to add a CDN acceleration domain name. Up to 100 domain names can be added per minute.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CAMRESOURCEBELONGTODIFFERENTUSER = "InvalidParameter.CamResourceBelongToDifferentUser"
//  INVALIDPARAMETER_CAMRESOURCESIXSTAGEERROR = "InvalidParameter.CamResourceSixStageError"
//  INVALIDPARAMETER_CAMTAGKEYALREADYATTACHED = "InvalidParameter.CamTagKeyAlreadyAttached"
//  INVALIDPARAMETER_CAMTAGKEYILLEGAL = "InvalidParameter.CamTagKeyIllegal"
//  INVALIDPARAMETER_CAMTAGKEYNOTEXIST = "InvalidParameter.CamTagKeyNotExist"
//  INVALIDPARAMETER_CAMTAGVALUEILLEGAL = "InvalidParameter.CamTagValueIllegal"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNCONFIGINVALIDTAG = "InvalidParameter.CdnConfigInvalidTag"
//  INVALIDPARAMETER_CDNCONFIGTAGREQUIRED = "InvalidParameter.CdnConfigTagRequired"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNHOSTISCOSDEFAULTACCESS = "InvalidParameter.CdnHostIsCosDefaultAccess"
//  INVALIDPARAMETER_CDNHOSTTOOLONGHOST = "InvalidParameter.CdnHostTooLongHost"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CAMRESOURCEARRAYTOOLONG = "LimitExceeded.CamResourceArrayTooLong"
//  LIMITEXCEEDED_CAMRESOURCETOOMANYTAGKEY = "LimitExceeded.CamResourceTooManyTagKey"
//  LIMITEXCEEDED_CAMTAGKEYTOOLONG = "LimitExceeded.CamTagKeyTooLong"
//  LIMITEXCEEDED_CAMTAGKEYTOOMANYTAGVALUE = "LimitExceeded.CamTagKeyTooManyTagValue"
//  LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"
//  LIMITEXCEEDED_CAMUSERTOOMANYTAGKEY = "LimitExceeded.CamUserTooManyTagKey"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINDSA = "ResourceUnavailable.CdnHostExistsInDsa"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINTCB = "ResourceUnavailable.CdnHostExistsInTcb"
//  RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  RESOURCEUNAVAILABLE_HOSTEXISTINVOD = "ResourceUnavailable.HostExistInVod"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) AddCdnDomain(request *AddCdnDomainRequest) (response *AddCdnDomainResponse, err error) {
    return c.AddCdnDomainWithContext(context.Background(), request)
}

// AddCdnDomain
// This API is used to add a CDN acceleration domain name. Up to 100 domain names can be added per minute.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CAMRESOURCEBELONGTODIFFERENTUSER = "InvalidParameter.CamResourceBelongToDifferentUser"
//  INVALIDPARAMETER_CAMRESOURCESIXSTAGEERROR = "InvalidParameter.CamResourceSixStageError"
//  INVALIDPARAMETER_CAMTAGKEYALREADYATTACHED = "InvalidParameter.CamTagKeyAlreadyAttached"
//  INVALIDPARAMETER_CAMTAGKEYILLEGAL = "InvalidParameter.CamTagKeyIllegal"
//  INVALIDPARAMETER_CAMTAGKEYNOTEXIST = "InvalidParameter.CamTagKeyNotExist"
//  INVALIDPARAMETER_CAMTAGVALUEILLEGAL = "InvalidParameter.CamTagValueIllegal"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNCONFIGINVALIDTAG = "InvalidParameter.CdnConfigInvalidTag"
//  INVALIDPARAMETER_CDNCONFIGTAGREQUIRED = "InvalidParameter.CdnConfigTagRequired"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNHOSTISCOSDEFAULTACCESS = "InvalidParameter.CdnHostIsCosDefaultAccess"
//  INVALIDPARAMETER_CDNHOSTTOOLONGHOST = "InvalidParameter.CdnHostTooLongHost"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CAMRESOURCEARRAYTOOLONG = "LimitExceeded.CamResourceArrayTooLong"
//  LIMITEXCEEDED_CAMRESOURCETOOMANYTAGKEY = "LimitExceeded.CamResourceTooManyTagKey"
//  LIMITEXCEEDED_CAMTAGKEYTOOLONG = "LimitExceeded.CamTagKeyTooLong"
//  LIMITEXCEEDED_CAMTAGKEYTOOMANYTAGVALUE = "LimitExceeded.CamTagKeyTooManyTagValue"
//  LIMITEXCEEDED_CAMTAGQUOTAEXCEEDLIMIT = "LimitExceeded.CamTagQuotaExceedLimit"
//  LIMITEXCEEDED_CAMUSERTOOMANYTAGKEY = "LimitExceeded.CamUserTooManyTagKey"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNUSERTOOMANYHOSTS = "LimitExceeded.CdnUserTooManyHosts"
//  RESOURCEINUSE_CDNCONFLICTHOSTEXISTS = "ResourceInUse.CdnConflictHostExists"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CAMTAGKEYNOTEXIST = "ResourceNotFound.CamTagKeyNotExist"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINDSA = "ResourceUnavailable.CdnHostExistsInDsa"
//  RESOURCEUNAVAILABLE_CDNHOSTEXISTSINTCB = "ResourceUnavailable.CdnHostExistsInTcb"
//  RESOURCEUNAVAILABLE_CDNHOSTISMALICIOUS = "ResourceUnavailable.CdnHostIsMalicious"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  RESOURCEUNAVAILABLE_HOSTEXISTINVOD = "ResourceUnavailable.HostExistInVod"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNDOMAINRECORDNOTVERIFIED = "UnauthorizedOperation.CdnDomainRecordNotVerified"
//  UNAUTHORIZEDOPERATION_CDNHOSTEXISTSININTERNAL = "UnauthorizedOperation.CdnHostExistsInInternal"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTISOWNEDBYOTHER = "UnauthorizedOperation.CdnHostIsOwnedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) AddCdnDomainWithContext(ctx context.Context, request *AddCdnDomainRequest) (response *AddCdnDomainResponse, err error) {
    if request == nil {
        request = NewAddCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateClsLogTopicRequest() (request *CreateClsLogTopicRequest) {
    request = &CreateClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateClsLogTopic")
    
    
    return
}

func NewCreateClsLogTopicResponse() (response *CreateClsLogTopicResponse) {
    response = &CreateClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateClsLogTopic
// This API is used to create a log topic. Up to 10 log topics can be created under one logset.
//
// error code that may be returned:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateClsLogTopic(request *CreateClsLogTopicRequest) (response *CreateClsLogTopicResponse, err error) {
    return c.CreateClsLogTopicWithContext(context.Background(), request)
}

// CreateClsLogTopic
// This API is used to create a log topic. Up to 10 log topics can be created under one logset.
//
// error code that may be returned:
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNAMEINVALID = "InvalidParameter.CdnClsTopicNameInvalid"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) CreateClsLogTopicWithContext(ctx context.Context, request *CreateClsLogTopicRequest) (response *CreateClsLogTopicResponse, err error) {
    if request == nil {
        request = NewCreateClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateScdnFailedLogTaskRequest() (request *CreateScdnFailedLogTaskRequest) {
    request = &CreateScdnFailedLogTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "CreateScdnFailedLogTask")
    
    
    return
}

func NewCreateScdnFailedLogTaskResponse() (response *CreateScdnFailedLogTaskResponse) {
    response = &CreateScdnFailedLogTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateScdnFailedLogTask
// This API is used to recreate a failed event log task.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"
//  INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"
//  INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"
//  LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateScdnFailedLogTask(request *CreateScdnFailedLogTaskRequest) (response *CreateScdnFailedLogTaskResponse, err error) {
    return c.CreateScdnFailedLogTaskWithContext(context.Background(), request)
}

// CreateScdnFailedLogTask
// This API is used to recreate a failed event log task.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_SCDNLOGTASKEXPIRED = "InvalidParameter.ScdnLogTaskExpired"
//  INVALIDPARAMETER_SCDNLOGTASKNOTFOUNDORNOTFAIL = "InvalidParameter.ScdnLogTaskNotFoundOrNotFail"
//  INVALIDPARAMETER_SCDNLOGTASKTIMERANGEINVALID = "InvalidParameter.ScdnLogTaskTimeRangeInvalid"
//  LIMITEXCEEDED_SCDNLOGTASKEXCEEDDAYLIMIT = "LimitExceeded.ScdnLogTaskExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
func (c *Client) CreateScdnFailedLogTaskWithContext(ctx context.Context, request *CreateScdnFailedLogTaskRequest) (response *CreateScdnFailedLogTaskResponse, err error) {
    if request == nil {
        request = NewCreateScdnFailedLogTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateScdnFailedLogTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateScdnFailedLogTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCdnDomainRequest() (request *DeleteCdnDomainRequest) {
    request = &DeleteCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteCdnDomain")
    
    
    return
}

func NewDeleteCdnDomainResponse() (response *DeleteCdnDomainResponse) {
    response = &DeleteCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCdnDomain
// This API is used to delete a specified acceleration domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) (response *DeleteCdnDomainResponse, err error) {
    return c.DeleteCdnDomainWithContext(context.Background(), request)
}

// DeleteCdnDomain
// This API is used to delete a specified acceleration domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DeleteCdnDomainWithContext(ctx context.Context, request *DeleteCdnDomainRequest) (response *DeleteCdnDomainResponse, err error) {
    if request == nil {
        request = NewDeleteCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteClsLogTopicRequest() (request *DeleteClsLogTopicRequest) {
    request = &DeleteClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DeleteClsLogTopic")
    
    
    return
}

func NewDeleteClsLogTopicResponse() (response *DeleteClsLogTopicResponse) {
    response = &DeleteClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteClsLogTopic
// This API is used to delete a log topic. Note: when a log topic is deleted, all logs of the domain names bound to it will no longer be published to the topic, and the logs previously published to the topic will be deleted. This action will take effect within 5–15 minutes.
//
// error code that may be returned:
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DeleteClsLogTopic(request *DeleteClsLogTopicRequest) (response *DeleteClsLogTopicResponse, err error) {
    return c.DeleteClsLogTopicWithContext(context.Background(), request)
}

// DeleteClsLogTopic
// This API is used to delete a log topic. Note: when a log topic is deleted, all logs of the domain names bound to it will no longer be published to the topic, and the logs previously published to the topic will be deleted. This action will take effect within 5–15 minutes.
//
// error code that may be returned:
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DeleteClsLogTopicWithContext(ctx context.Context, request *DeleteClsLogTopicRequest) (response *DeleteClsLogTopicResponse, err error) {
    if request == nil {
        request = NewDeleteClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingDataRequest() (request *DescribeBillingDataRequest) {
    request = &DescribeBillingDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeBillingData")
    
    
    return
}

func NewDescribeBillingDataResponse() (response *DescribeBillingDataResponse) {
    response = &DescribeBillingDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingData
// This API is used to query billing data details.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeBillingData(request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    return c.DescribeBillingDataWithContext(context.Background(), request)
}

// DescribeBillingData
// This API is used to query billing data details.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeBillingDataWithContext(ctx context.Context, request *DescribeBillingDataRequest) (response *DescribeBillingDataResponse, err error) {
    if request == nil {
        request = NewDescribeBillingDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnDataRequest() (request *DescribeCdnDataRequest) {
    request = &DescribeCdnDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnData")
    
    
    return
}

func NewDescribeCdnDataResponse() (response *DescribeCdnDataResponse) {
    response = &DescribeCdnDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCdnData
// This API is used to query CDN real-time access monitoring data and supports the following metrics:
//
// 
//
// + Traffic (in bytes)
//
// + Bandwidth (in bps)
//
// + Number of requests
//
// + Number of hit requests
//
// + Request hit rate (in %)
//
// + Hit traffic (in bytes)
//
// + Traffic hit rate (in %)
//
// + Aggregate list of 2xx status codes and the details of status codes starting with 2 (in entries)
//
// + Aggregate list of 3xx status codes and the details of status codes starting with 3 (in entries)
//
// + Aggregate list of 4xx status codes and the details of status codes starting with 4 (in entries)
//
// + Aggregate list of 5xx status codes and the details of status codes starting with 5 (in entries)
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_CSRFERROR = "UnauthorizedOperation.CsrfError"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeCdnData(request *DescribeCdnDataRequest) (response *DescribeCdnDataResponse, err error) {
    return c.DescribeCdnDataWithContext(context.Background(), request)
}

// DescribeCdnData
// This API is used to query CDN real-time access monitoring data and supports the following metrics:
//
// 
//
// + Traffic (in bytes)
//
// + Bandwidth (in bps)
//
// + Number of requests
//
// + Number of hit requests
//
// + Request hit rate (in %)
//
// + Hit traffic (in bytes)
//
// + Traffic hit rate (in %)
//
// + Aggregate list of 2xx status codes and the details of status codes starting with 2 (in entries)
//
// + Aggregate list of 3xx status codes and the details of status codes starting with 3 (in entries)
//
// + Aggregate list of 4xx status codes and the details of status codes starting with 4 (in entries)
//
// + Aggregate list of 5xx status codes and the details of status codes starting with 5 (in entries)
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_COSTDATASYSTEMERROR = "InternalError.CostDataSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_CSRFERROR = "UnauthorizedOperation.CsrfError"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeCdnDataWithContext(ctx context.Context, request *DescribeCdnDataRequest) (response *DescribeCdnDataResponse, err error) {
    if request == nil {
        request = NewDescribeCdnDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnDomainLogsRequest() (request *DescribeCdnDomainLogsRequest) {
    request = &DescribeCdnDomainLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnDomainLogs")
    
    
    return
}

func NewDescribeCdnDomainLogsResponse() (response *DescribeCdnDomainLogsResponse) {
    response = &DescribeCdnDomainLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCdnDomainLogs
// This API is used to query the download link of an access log. You can use this API for access logs in the last 30 days either within or outside Mainland China.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
    return c.DescribeCdnDomainLogsWithContext(context.Background(), request)
}

// DescribeCdnDomainLogs
// This API is used to query the download link of an access log. You can use this API for access logs in the last 30 days either within or outside Mainland China.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCdnDomainLogsWithContext(ctx context.Context, request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
    if request == nil {
        request = NewDescribeCdnDomainLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnDomainLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnDomainLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnIpRequest() (request *DescribeCdnIpRequest) {
    request = &DescribeCdnIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnIp")
    
    
    return
}

func NewDescribeCdnIpResponse() (response *DescribeCdnIpResponse) {
    response = &DescribeCdnIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCdnIp
// This API is used to query CDN IP ownership.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeCdnIp(request *DescribeCdnIpRequest) (response *DescribeCdnIpResponse, err error) {
    return c.DescribeCdnIpWithContext(context.Background(), request)
}

// DescribeCdnIp
// This API is used to query CDN IP ownership.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeCdnIpWithContext(ctx context.Context, request *DescribeCdnIpRequest) (response *DescribeCdnIpResponse, err error) {
    if request == nil {
        request = NewDescribeCdnIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCdnOriginIpRequest() (request *DescribeCdnOriginIpRequest) {
    request = &DescribeCdnOriginIpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCdnOriginIp")
    
    
    return
}

func NewDescribeCdnOriginIpResponse() (response *DescribeCdnOriginIpResponse) {
    response = &DescribeCdnOriginIpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCdnOriginIp
// This API is used to query the IP information of CDN intermediate nodes. Note: this API will be deactivated soon and no longer be maintained. Please call `DescribeIpStatus` instead.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
func (c *Client) DescribeCdnOriginIp(request *DescribeCdnOriginIpRequest) (response *DescribeCdnOriginIpResponse, err error) {
    return c.DescribeCdnOriginIpWithContext(context.Background(), request)
}

// DescribeCdnOriginIp
// This API is used to query the IP information of CDN intermediate nodes. Note: this API will be deactivated soon and no longer be maintained. Please call `DescribeIpStatus` instead.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNCALLINGQUERYIPTOOOFTEN = "LimitExceeded.CdnCallingQueryIpTooOften"
//  LIMITEXCEEDED_CDNQUERYIPBATCHTOOMANY = "LimitExceeded.CdnQueryIpBatchTooMany"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
func (c *Client) DescribeCdnOriginIpWithContext(ctx context.Context, request *DescribeCdnOriginIpRequest) (response *DescribeCdnOriginIpResponse, err error) {
    if request == nil {
        request = NewDescribeCdnOriginIpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCdnOriginIp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCdnOriginIpResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCertDomainsRequest() (request *DescribeCertDomainsRequest) {
    request = &DescribeCertDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeCertDomains")
    
    
    return
}

func NewDescribeCertDomainsResponse() (response *DescribeCertDomainsResponse) {
    response = &DescribeCertDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCertDomains
// This API is used to verify a SSL certificate and obtain its domain names.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNCERTNOTPEM = "InvalidParameter.CdnCertNotPem"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCertDomains(request *DescribeCertDomainsRequest) (response *DescribeCertDomainsResponse, err error) {
    return c.DescribeCertDomainsWithContext(context.Background(), request)
}

// DescribeCertDomains
// This API is used to verify a SSL certificate and obtain its domain names.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNCERTNOTPEM = "InvalidParameter.CdnCertNotPem"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNHOSTISUSEDBYOTHER = "UnauthorizedOperation.CdnHostIsUsedByOther"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_CLSNOTALLOWED = "UnsupportedOperation.ClsNotAllowed"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeCertDomainsWithContext(ctx context.Context, request *DescribeCertDomainsRequest) (response *DescribeCertDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeCertDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCertDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCertDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
    request = &DescribeDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDomains")
    
    
    return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
    response = &DescribeDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomains
// This API is used to query the basic configuration information of CDN acceleration domain names (inside and outside mainland China), including the project ID, service status, service type, creation time, and update time, etc.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
    return c.DescribeDomainsWithContext(context.Background(), request)
}

// DescribeDomains
// This API is used to query the basic configuration information of CDN acceleration domain names (inside and outside mainland China), including the project ID, service status, service type, creation time, and update time, etc.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeDomainsConfig")
    
    
    return
}

func NewDescribeDomainsConfigResponse() (response *DescribeDomainsConfigResponse) {
    response = &DescribeDomainsConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDomainsConfig
// This API is used to query the complete configuration information of CDN acceleration domain names (inside and outside mainland China).
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeDomainsConfig(request *DescribeDomainsConfigRequest) (response *DescribeDomainsConfigResponse, err error) {
    return c.DescribeDomainsConfigWithContext(context.Background(), request)
}

// DescribeDomainsConfig
// This API is used to query the complete configuration information of CDN acceleration domain names (inside and outside mainland China).
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTNOCERTINFO = "InvalidParameter.CdnCertNoCertInfo"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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

func NewDescribeIpStatusRequest() (request *DescribeIpStatusRequest) {
    request = &DescribeIpStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeIpStatus")
    
    
    return
}

func NewDescribeIpStatusResponse() (response *DescribeIpStatusResponse) {
    response = &DescribeIpStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIpStatus
// This API is used to query the IP details of edge nodes (available soon) and intermediate nodes. Note that there is a certain delay in data availability.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41954?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (response *DescribeIpStatusResponse, err error) {
    return c.DescribeIpStatusWithContext(context.Background(), request)
}

// DescribeIpStatus
// This API is used to query the IP details of edge nodes (available soon) and intermediate nodes. Note that there is a certain delay in data availability.
//
// 
//
// >? If you have migrated your ECDN service to CDN, you can use the <a href="https://intl.cloud.tencent.com/document/api/228/41954?from_cn_redirect=1">corresponding CDN API</a>.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCONFIGINVALIDHOST = "InvalidParameter.CdnConfigInvalidHost"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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

func NewDescribeIpVisitRequest() (request *DescribeIpVisitRequest) {
    request = &DescribeIpVisitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeIpVisit")
    
    
    return
}

func NewDescribeIpVisitResponse() (response *DescribeIpVisitResponse) {
    response = &DescribeIpVisitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIpVisit
// This API (DescribeIpVisit) is used to query the number of users who remain active for 5 minutes and the detailed number of daily active users.
//
// 
//
// + Number of users who remain active for 5 minutes: Collects deduplicated statistics based on client IP addresses in the log with the 5-minute granularity.
//
// + Number of daily active users: Collects deduplicated statistics based on client IP addresses in the log with the 1-day granularity.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpVisit(request *DescribeIpVisitRequest) (response *DescribeIpVisitResponse, err error) {
    return c.DescribeIpVisitWithContext(context.Background(), request)
}

// DescribeIpVisit
// This API (DescribeIpVisit) is used to query the number of users who remain active for 5 minutes and the detailed number of daily active users.
//
// 
//
// + Number of users who remain active for 5 minutes: Collects deduplicated statistics based on client IP addresses in the log with the 5-minute granularity.
//
// + Number of daily active users: Collects deduplicated statistics based on client IP addresses in the log with the 1-day granularity.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMINTERVAL = "InvalidParameter.CdnInvalidParamInterval"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeIpVisitWithContext(ctx context.Context, request *DescribeIpVisitRequest) (response *DescribeIpVisitResponse, err error) {
    if request == nil {
        request = NewDescribeIpVisitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIpVisit require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIpVisitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMapInfoRequest() (request *DescribeMapInfoRequest) {
    request = &DescribeMapInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeMapInfo")
    
    
    return
}

func NewDescribeMapInfoResponse() (response *DescribeMapInfoResponse) {
    response = &DescribeMapInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMapInfo
// This API (DescribeMapInfo) is used to query the IDs of districts or ISPs.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeMapInfo(request *DescribeMapInfoRequest) (response *DescribeMapInfoResponse, err error) {
    return c.DescribeMapInfoWithContext(context.Background(), request)
}

// DescribeMapInfo
// This API (DescribeMapInfo) is used to query the IDs of districts or ISPs.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeMapInfoWithContext(ctx context.Context, request *DescribeMapInfoRequest) (response *DescribeMapInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMapInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMapInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMapInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOriginDataRequest() (request *DescribeOriginDataRequest) {
    request = &DescribeOriginDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeOriginData")
    
    
    return
}

func NewDescribeOriginDataResponse() (response *DescribeOriginDataResponse) {
    response = &DescribeOriginDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOriginData
// This API is used to query CDN real-time origin-pull monitoring data and supports the following metrics:
//
// 
//
// + Origin-pull traffic (in bytes)
//
// + Origin-pull bandwidth (in bps)
//
// + Number of origin-pull requests
//
// + Number of failed origin-pull requests
//
// + Origin-pull failure rate (in % with two decimal digits)
//
// + Aggregate list of 2xx origin-pull status codes and the details of origin-pull status codes starting with 2 (in entries)
//
// + Aggregate list of 3xx origin-pull status codes and the details of origin-pull status codes starting with 3 (in entries)
//
// + Aggregate list of 4xx origin-pull status codes and the details of origin-pull status codes starting with 4 (in entries)
//
// + Aggregate list of 5xx origin-pull status codes and the details of origin-pull status codes starting with 5 (in entries)
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeOriginData(request *DescribeOriginDataRequest) (response *DescribeOriginDataResponse, err error) {
    return c.DescribeOriginDataWithContext(context.Background(), request)
}

// DescribeOriginData
// This API is used to query CDN real-time origin-pull monitoring data and supports the following metrics:
//
// 
//
// + Origin-pull traffic (in bytes)
//
// + Origin-pull bandwidth (in bps)
//
// + Number of origin-pull requests
//
// + Number of failed origin-pull requests
//
// + Origin-pull failure rate (in % with two decimal digits)
//
// + Aggregate list of 2xx origin-pull status codes and the details of origin-pull status codes starting with 2 (in entries)
//
// + Aggregate list of 3xx origin-pull status codes and the details of origin-pull status codes starting with 3 (in entries)
//
// + Aggregate list of 4xx origin-pull status codes and the details of origin-pull status codes starting with 4 (in entries)
//
// + Aggregate list of 5xx origin-pull status codes and the details of origin-pull status codes starting with 5 (in entries)
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMMETRIC = "InvalidParameter.CdnInvalidParamMetric"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CDNUSERTOOMANYHOSTS = "ResourceNotFound.CdnUserTooManyHosts"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeOriginDataWithContext(ctx context.Context, request *DescribeOriginDataRequest) (response *DescribeOriginDataResponse, err error) {
    if request == nil {
        request = NewDescribeOriginDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOriginData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOriginDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePayTypeRequest() (request *DescribePayTypeRequest) {
    request = &DescribePayTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePayType")
    
    
    return
}

func NewDescribePayTypeResponse() (response *DescribePayTypeResponse) {
    response = &DescribePayTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePayType
// This API (DescribePayType) is used to query billing information of the current account, such as billing mode and billing cycle.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePayType(request *DescribePayTypeRequest) (response *DescribePayTypeResponse, err error) {
    return c.DescribePayTypeWithContext(context.Background(), request)
}

// DescribePayType
// This API (DescribePayType) is used to query billing information of the current account, such as billing mode and billing cycle.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_ECDNINTERFACEERROR = "InvalidParameter.EcdnInterfaceError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePayTypeWithContext(ctx context.Context, request *DescribePayTypeRequest) (response *DescribePayTypeResponse, err error) {
    if request == nil {
        request = NewDescribePayTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePayType require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePayTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePurgeQuotaRequest() (request *DescribePurgeQuotaRequest) {
    request = &DescribePurgeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePurgeQuota")
    
    
    return
}

func NewDescribePurgeQuotaResponse() (response *DescribePurgeQuotaResponse) {
    response = &DescribePurgeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePurgeQuota
// This API is used to query the purge usage quota and daily available usage for an account.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePurgeQuota(request *DescribePurgeQuotaRequest) (response *DescribePurgeQuotaResponse, err error) {
    return c.DescribePurgeQuotaWithContext(context.Background(), request)
}

// DescribePurgeQuota
// This API is used to query the purge usage quota and daily available usage for an account.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePurgeTasks")
    
    
    return
}

func NewDescribePurgeTasksResponse() (response *DescribePurgeTasksResponse) {
    response = &DescribePurgeTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePurgeTasks
// This API is used to query the record and progress of URL or directory purge tasks submitted via the `PurgePathCache` or `PurgeUrlsCache` APIs.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePurgeTasks(request *DescribePurgeTasksRequest) (response *DescribePurgeTasksResponse, err error) {
    return c.DescribePurgeTasksWithContext(context.Background(), request)
}

// DescribePurgeTasks
// This API is used to query the record and progress of URL or directory purge tasks submitted via the `PurgePathCache` or `PurgeUrlsCache` APIs.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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

func NewDescribePushQuotaRequest() (request *DescribePushQuotaRequest) {
    request = &DescribePushQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePushQuota")
    
    
    return
}

func NewDescribePushQuotaResponse() (response *DescribePushQuotaResponse) {
    response = &DescribePushQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePushQuota
// This API is used to query the prefetch quota and daily available usage.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushQuota(request *DescribePushQuotaRequest) (response *DescribePushQuotaResponse, err error) {
    return c.DescribePushQuotaWithContext(context.Background(), request)
}

// DescribePushQuota
// This API is used to query the prefetch quota and daily available usage.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushQuotaWithContext(ctx context.Context, request *DescribePushQuotaRequest) (response *DescribePushQuotaResponse, err error) {
    if request == nil {
        request = NewDescribePushQuotaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePushQuota require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePushQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePushTasksRequest() (request *DescribePushTasksRequest) {
    request = &DescribePushTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribePushTasks")
    
    
    return
}

func NewDescribePushTasksResponse() (response *DescribePushTasksResponse) {
    response = &DescribePushTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePushTasks
// This API is used to query the submission record and progress of prefetch tasks.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushTasks(request *DescribePushTasksRequest) (response *DescribePushTasksResponse, err error) {
    return c.DescribePushTasksWithContext(context.Background(), request)
}

// DescribePushTasks
// This API is used to query the submission record and progress of prefetch tasks.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribePushTasksWithContext(ctx context.Context, request *DescribePushTasksRequest) (response *DescribePushTasksResponse, err error) {
    if request == nil {
        request = NewDescribePushTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePushTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePushTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReportDataRequest() (request *DescribeReportDataRequest) {
    request = &DescribeReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeReportData")
    
    
    return
}

func NewDescribeReportDataResponse() (response *DescribeReportDataResponse) {
    response = &DescribeReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReportData
// This API is used to query the daily/weekly/monthly report data at domain name/project levels.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeReportData(request *DescribeReportDataRequest) (response *DescribeReportDataResponse, err error) {
    return c.DescribeReportDataWithContext(context.Background(), request)
}

// DescribeReportData
// This API is used to query the daily/weekly/monthly report data at domain name/project levels.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) DescribeReportDataWithContext(ctx context.Context, request *DescribeReportDataRequest) (response *DescribeReportDataResponse, err error) {
    if request == nil {
        request = NewDescribeReportDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReportData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReportDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUrlViolationsRequest() (request *DescribeUrlViolationsRequest) {
    request = &DescribeUrlViolationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DescribeUrlViolations")
    
    
    return
}

func NewDescribeUrlViolationsResponse() (response *DescribeUrlViolationsResponse) {
    response = &DescribeUrlViolationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUrlViolations
// This API is used to query the list of domain name URLs containing regulation-violating content scanned and detected by the CDN system, and the current status of the URLs.
//
// It corresponds to the **Pornography Detection** page on the CDN Console.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeUrlViolations(request *DescribeUrlViolationsRequest) (response *DescribeUrlViolationsResponse, err error) {
    return c.DescribeUrlViolationsWithContext(context.Background(), request)
}

// DescribeUrlViolations
// This API is used to query the list of domain name URLs containing regulation-violating content scanned and detected by the CDN system, and the current status of the URLs.
//
// It corresponds to the **Pornography Detection** page on the CDN Console.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DescribeUrlViolationsWithContext(ctx context.Context, request *DescribeUrlViolationsRequest) (response *DescribeUrlViolationsResponse, err error) {
    if request == nil {
        request = NewDescribeUrlViolationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUrlViolations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUrlViolationsResponse()
    err = c.Send(request, response)
    return
}

func NewDisableCachesRequest() (request *DisableCachesRequest) {
    request = &DisableCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DisableCaches")
    
    
    return
}

func NewDisableCachesResponse() (response *DisableCachesResponse) {
    response = &DisableCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableCaches
// This API is used to block access to a specific URL on CDN. When a URL is blocked, error 403 will be returned for requests from the Chinese mainland. URL blocking is not permanent. Note that this API is only available to beta users now. 
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableCaches(request *DisableCachesRequest) (response *DisableCachesResponse, err error) {
    return c.DisableCachesWithContext(context.Background(), request)
}

// DisableCaches
// This API is used to block access to a specific URL on CDN. When a URL is blocked, error 403 will be returned for requests from the Chinese mainland. URL blocking is not permanent. Note that this API is only available to beta users now. 
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableCachesWithContext(ctx context.Context, request *DisableCachesRequest) (response *DisableCachesResponse, err error) {
    if request == nil {
        request = NewDisableCachesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableCaches require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableCachesResponse()
    err = c.Send(request, response)
    return
}

func NewDisableClsLogTopicRequest() (request *DisableClsLogTopicRequest) {
    request = &DisableClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "DisableClsLogTopic")
    
    
    return
}

func NewDisableClsLogTopicResponse() (response *DisableClsLogTopicResponse) {
    response = &DisableClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableClsLogTopic
// This API is used to stop publishing to a log topic. Note: after a log topic is disabled, all logs of the domain names bound to it will no longer be published to the topic, and the logs that have already been published will be retained. This action will take effect within 5–15 minutes.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableClsLogTopic(request *DisableClsLogTopicRequest) (response *DisableClsLogTopicResponse, err error) {
    return c.DisableClsLogTopicWithContext(context.Background(), request)
}

// DisableClsLogTopic
// This API is used to stop publishing to a log topic. Note: after a log topic is disabled, all logs of the domain names bound to it will no longer be published to the topic, and the logs that have already been published will be retained. This action will take effect within 5–15 minutes.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) DisableClsLogTopicWithContext(ctx context.Context, request *DisableClsLogTopicRequest) (response *DisableClsLogTopicResponse, err error) {
    if request == nil {
        request = NewDisableClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewEnableCachesRequest() (request *EnableCachesRequest) {
    request = &EnableCachesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "EnableCaches")
    
    
    return
}

func NewEnableCachesResponse() (response *EnableCachesResponse) {
    response = &EnableCachesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableCaches
// This API (EnableCaches) is used to unblock manually blocked URLs. After a URL is successfully unblocked, it takes about 5 to 10 minutes to take effect across the entire network. (This API is during beta test and not fully available now.)
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableCaches(request *EnableCachesRequest) (response *EnableCachesResponse, err error) {
    return c.EnableCachesWithContext(context.Background(), request)
}

// EnableCaches
// This API (EnableCaches) is used to unblock manually blocked URLs. After a URL is successfully unblocked, it takes about 5 to 10 minutes to take effect across the entire network. (This API is during beta test and not fully available now.)
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableCachesWithContext(ctx context.Context, request *EnableCachesRequest) (response *EnableCachesResponse, err error) {
    if request == nil {
        request = NewEnableCachesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableCaches require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableCachesResponse()
    err = c.Send(request, response)
    return
}

func NewEnableClsLogTopicRequest() (request *EnableClsLogTopicRequest) {
    request = &EnableClsLogTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "EnableClsLogTopic")
    
    
    return
}

func NewEnableClsLogTopicResponse() (response *EnableClsLogTopicResponse) {
    response = &EnableClsLogTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// EnableClsLogTopic
// This API is used to start publishing to a log topic. Note: after a log topic is enabled, all logs of the domain names bound to the topic will be published to it. This action will take effect within 5–15 minutes.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableClsLogTopic(request *EnableClsLogTopicRequest) (response *EnableClsLogTopicResponse, err error) {
    return c.EnableClsLogTopicWithContext(context.Background(), request)
}

// EnableClsLogTopic
// This API is used to start publishing to a log topic. Note: after a log topic is enabled, all logs of the domain names bound to the topic will be published to it. This action will take effect within 5–15 minutes.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) EnableClsLogTopicWithContext(ctx context.Context, request *EnableClsLogTopicRequest) (response *EnableClsLogTopicResponse, err error) {
    if request == nil {
        request = NewEnableClsLogTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnableClsLogTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewEnableClsLogTopicResponse()
    err = c.Send(request, response)
    return
}

func NewGetDisableRecordsRequest() (request *GetDisableRecordsRequest) {
    request = &GetDisableRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "GetDisableRecords")
    
    
    return
}

func NewGetDisableRecordsResponse() (response *GetDisableRecordsResponse) {
    response = &GetDisableRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDisableRecords
// This API is used to query the resource blocking history and the current URL status. (This API is in beta test and not generally available yet.)
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) GetDisableRecords(request *GetDisableRecordsRequest) (response *GetDisableRecordsResponse, err error) {
    return c.GetDisableRecordsWithContext(context.Background(), request)
}

// GetDisableRecords
// This API is used to query the resource blocking history and the current URL status. (This API is in beta test and not generally available yet.)
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) GetDisableRecordsWithContext(ctx context.Context, request *GetDisableRecordsRequest) (response *GetDisableRecordsResponse, err error) {
    if request == nil {
        request = NewGetDisableRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDisableRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDisableRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewListClsLogTopicsRequest() (request *ListClsLogTopicsRequest) {
    request = &ListClsLogTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListClsLogTopics")
    
    
    return
}

func NewListClsLogTopicsResponse() (response *ListClsLogTopicsResponse) {
    response = &ListClsLogTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListClsLogTopics
// This API is used to display the list of log topics. Note: a logset can contain up to 10 log topics.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsLogTopics(request *ListClsLogTopicsRequest) (response *ListClsLogTopicsResponse, err error) {
    return c.ListClsLogTopicsWithContext(context.Background(), request)
}

// ListClsLogTopics
// This API is used to display the list of log topics. Note: a logset can contain up to 10 log topics.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSSERVICENOTACTIVATED = "UnauthorizedOperation.ClsServiceNotActivated"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsLogTopicsWithContext(ctx context.Context, request *ListClsLogTopicsRequest) (response *ListClsLogTopicsResponse, err error) {
    if request == nil {
        request = NewListClsLogTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListClsLogTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewListClsLogTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewListClsTopicDomainsRequest() (request *ListClsTopicDomainsRequest) {
    request = &ListClsTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListClsTopicDomains")
    
    
    return
}

func NewListClsTopicDomainsResponse() (response *ListClsTopicDomainsResponse) {
    response = &ListClsTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListClsTopicDomains
// This API is used to get the list of domain names bound to a log topic.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsTopicDomains(request *ListClsTopicDomainsRequest) (response *ListClsTopicDomainsResponse, err error) {
    return c.ListClsTopicDomainsWithContext(context.Background(), request)
}

// ListClsTopicDomains
// This API is used to get the list of domain names bound to a log topic.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListClsTopicDomainsWithContext(ctx context.Context, request *ListClsTopicDomainsRequest) (response *ListClsTopicDomainsResponse, err error) {
    if request == nil {
        request = NewListClsTopicDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListClsTopicDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewListClsTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewListTopDataRequest() (request *ListTopDataRequest) {
    request = &ListTopDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ListTopData")
    
    
    return
}

func NewListTopDataResponse() (response *ListTopDataResponse) {
    response = &ListTopDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListTopData
// This API is used to list data sorted the following ways by using different combinations of the Metric and Filter input parameters:
//
// 
//
// + It sorts access URLs by total traffic and total requests, and returns the top 1,000 URLs in descending order.
//
// + It sorts client districts by total traffic and total requests, and returns the list of districts in descending order.
//
// + It sorts client ISPs by total traffic and total requests, and returns the list of ISPs in descending order.
//
// + It sorts domain names by total traffic, peak bandwidth, total requests, average hit rate, and 2XX/3XX/4XX/5XX status codes, and returns the list of domain names in descending order.
//
// + It sorts domain names by total origin-pull traffic, peak origin-pull bandwidth, total origin-pull requests, average origin-pull failure rate, and 2XX/3XX/4XX/5XX origin-pull status codes, and returns the list of domain names in descending order.
//
// 
//
// Note: only data from the last 90 days will be queried.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListTopData(request *ListTopDataRequest) (response *ListTopDataResponse, err error) {
    return c.ListTopDataWithContext(context.Background(), request)
}

// ListTopData
// This API is used to list data sorted the following ways by using different combinations of the Metric and Filter input parameters:
//
// 
//
// + It sorts access URLs by total traffic and total requests, and returns the top 1,000 URLs in descending order.
//
// + It sorts client districts by total traffic and total requests, and returns the list of districts in descending order.
//
// + It sorts client ISPs by total traffic and total requests, and returns the list of ISPs in descending order.
//
// + It sorts domain names by total traffic, peak bandwidth, total requests, average hit rate, and 2XX/3XX/4XX/5XX status codes, and returns the list of domain names in descending order.
//
// + It sorts domain names by total origin-pull traffic, peak origin-pull bandwidth, total origin-pull requests, average origin-pull failure rate, and 2XX/3XX/4XX/5XX origin-pull status codes, and returns the list of domain names in descending order.
//
// 
//
// Note: only data from the last 90 days will be queried.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_DATASYSTEMERROR = "InternalError.DataSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_TAGSYSTEMERROR = "InternalError.TagSystemError"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDFILTER = "InvalidParameter.CdnStatInvalidFilter"
//  INVALIDPARAMETER_CDNSTATINVALIDMETRIC = "InvalidParameter.CdnStatInvalidMetric"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNPROJECTNOTEXISTS = "ResourceNotFound.CdnProjectNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNACCOUNTUNAUTHORIZED = "UnauthorizedOperation.CdnAccountUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ListTopDataWithContext(ctx context.Context, request *ListTopDataRequest) (response *ListTopDataResponse, err error) {
    if request == nil {
        request = NewListTopDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListTopData require credential")
    }

    request.SetContext(ctx)
    
    response = NewListTopDataResponse()
    err = c.Send(request, response)
    return
}

func NewManageClsTopicDomainsRequest() (request *ManageClsTopicDomainsRequest) {
    request = &ManageClsTopicDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ManageClsTopicDomains")
    
    
    return
}

func NewManageClsTopicDomainsResponse() (response *ManageClsTopicDomainsResponse) {
    response = &ManageClsTopicDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ManageClsTopicDomains
// This API is used to manage the list of domain names bound to a log topic.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ManageClsTopicDomains(request *ManageClsTopicDomainsRequest) (response *ManageClsTopicDomainsResponse, err error) {
    return c.ManageClsTopicDomainsWithContext(context.Background(), request)
}

// ManageClsTopicDomains
// This API is used to manage the list of domain names bound to a log topic.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ManageClsTopicDomainsWithContext(ctx context.Context, request *ManageClsTopicDomainsRequest) (response *ManageClsTopicDomainsResponse, err error) {
    if request == nil {
        request = NewManageClsTopicDomainsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ManageClsTopicDomains require credential")
    }

    request.SetContext(ctx)
    
    response = NewManageClsTopicDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainConfigRequest() (request *ModifyDomainConfigRequest) {
    request = &ModifyDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "ModifyDomainConfig")
    
    
    return
}

func NewModifyDomainConfigResponse() (response *ModifyDomainConfigResponse) {
    response = &ModifyDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDomainConfig
// This API is used to modify the configuration of a CDN acceleration domain name in a finer manner than `UpdateDomainConfig`.
//
// Notes:
//
// In `Route`, separate values by dots (.). The last value is called a leaf node. For non-leaf nodes, keep the configuration unchanged.
//
// The Value field is serialized to a JSON string {key:value}, where **key** is fixed to `update` and **value** is used to specify the value of the configuration parameter. To specify configurations with complex types, see https://intl.cloud.tencent.com/document/product/228/41116.?from_cn_redirect=1
//
// The input parameters of this API are not reported to CloudAudit as it may contain sensitive data, such as keys and secrets.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"
//  INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ModifyDomainConfig(request *ModifyDomainConfigRequest) (response *ModifyDomainConfigResponse, err error) {
    return c.ModifyDomainConfigWithContext(context.Background(), request)
}

// ModifyDomainConfig
// This API is used to modify the configuration of a CDN acceleration domain name in a finer manner than `UpdateDomainConfig`.
//
// Notes:
//
// In `Route`, separate values by dots (.). The last value is called a leaf node. For non-leaf nodes, keep the configuration unchanged.
//
// The Value field is serialized to a JSON string {key:value}, where **key** is fixed to `update` and **value** is used to specify the value of the configuration parameter. To specify configurations with complex types, see https://intl.cloud.tencent.com/document/product/228/41116.?from_cn_redirect=1
//
// The input parameters of this API are not reported to CloudAudit as it may contain sensitive data, such as keys and secrets.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"
//  INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) ModifyDomainConfigWithContext(ctx context.Context, request *ModifyDomainConfigRequest) (response *ModifyDomainConfigResponse, err error) {
    if request == nil {
        request = NewModifyDomainConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDomainConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDomainConfigResponse()
    err = c.Send(request, response)
    return
}

func NewPurgePathCacheRequest() (request *PurgePathCacheRequest) {
    request = &PurgePathCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "PurgePathCache")
    
    
    return
}

func NewPurgePathCacheResponse() (response *PurgePathCacheResponse) {
    response = &PurgePathCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PurgePathCache
// This API is used to submit multiple directory purge tasks, which are carried out according to the acceleration region of the domain names.
//
// By default, a maximum of 100 directories can be purged per day for acceleration regions either within or outside the Chinese mainland, and up to 500 tasks can be submitted at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgePathExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgePathExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PurgePathCache(request *PurgePathCacheRequest) (response *PurgePathCacheResponse, err error) {
    return c.PurgePathCacheWithContext(context.Background(), request)
}

// PurgePathCache
// This API is used to submit multiple directory purge tasks, which are carried out according to the acceleration region of the domain names.
//
// By default, a maximum of 100 directories can be purged per day for acceleration regions either within or outside the Chinese mainland, and up to 500 tasks can be submitted at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgePathExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEPATHEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgePathExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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
    
    request.Init().WithApiInfo("cdn", APIVersion, "PurgeUrlsCache")
    
    
    return
}

func NewPurgeUrlsCacheResponse() (response *PurgeUrlsCacheResponse) {
    response = &PurgeUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PurgeUrlsCache
// This API is used to submit multiple URL purge tasks, which are carried out according to the current acceleration region of the domain names in the URLs.
//
// By default, a maximum of 10,000 URLs can be purged per day for acceleration regions either within or outside Mainland China, and up to 1,000 tasks can be submitted at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PurgeUrlsCache(request *PurgeUrlsCacheRequest) (response *PurgeUrlsCacheResponse, err error) {
    return c.PurgeUrlsCacheWithContext(context.Background(), request)
}

// PurgeUrlsCache
// This API is used to submit multiple URL purge tasks, which are carried out according to the current acceleration region of the domain names in the URLs.
//
// By default, a maximum of 10,000 URLs can be purged per day for acceleration regions either within or outside Mainland China, and up to 1,000 tasks can be submitted at a time.
//
// error code that may be returned:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYPARAMERROR = "InternalError.CdnQueryParamError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPURGEWILDCARDNOTALLOWED = "InvalidParameter.CdnPurgeWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNURLEXCEEDLENGTH = "InvalidParameter.CdnUrlExceedLength"
//  LIMITEXCEEDED_CDNPURGEEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeExceedDayLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDBATCHLIMIT = "LimitExceeded.CdnPurgeUrlExceedBatchLimit"
//  LIMITEXCEEDED_CDNPURGEURLEXCEEDDAYLIMIT = "LimitExceeded.CdnPurgeUrlExceedDayLimit"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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

func NewPushUrlsCacheRequest() (request *PushUrlsCacheRequest) {
    request = &PushUrlsCacheRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "PushUrlsCache")
    
    
    return
}

func NewPushUrlsCacheResponse() (response *PushUrlsCacheResponse) {
    response = &PushUrlsCacheResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PushUrlsCache
// This API is used to cache specified URL resources to CDN nodes. You can specify acceleration regions for the prefetch.
//
// By default, a maximum of 1000 URLs can be prefetched per day for regions either within or outside the Chinese mainland, and up to 500 tasks can be submitted at a time. Note that resources prefetched outside the Chinese mainland will be cached to CDN nodes outside the Chinese mainland and the traffic generated will incur costs.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPUSHWILDCARDNOTALLOWED = "InvalidParameter.CdnPushWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPushExceedBatchLimit"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PushUrlsCache(request *PushUrlsCacheRequest) (response *PushUrlsCacheResponse, err error) {
    return c.PushUrlsCacheWithContext(context.Background(), request)
}

// PushUrlsCache
// This API is used to cache specified URL resources to CDN nodes. You can specify acceleration regions for the prefetch.
//
// By default, a maximum of 1000 URLs can be prefetched per day for regions either within or outside the Chinese mainland, and up to 500 tasks can be submitted at a time. Note that resources prefetched outside the Chinese mainland will be cached to CDN nodes outside the Chinese mainland and the traffic generated will incur costs.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INTERNALERROR_ROUTEERROR = "InternalError.RouteError"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDPARAM = "InvalidParameter.CdnHostInvalidParam"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNPUSHWILDCARDNOTALLOWED = "InvalidParameter.CdnPushWildcardNotAllowed"
//  INVALIDPARAMETER_CDNSTATINVALIDDATE = "InvalidParameter.CdnStatInvalidDate"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  LIMITEXCEEDED_CDNPUSHEXCEEDBATCHLIMIT = "LimitExceeded.CdnPushExceedBatchLimit"
//  LIMITEXCEEDED_CDNPUSHEXCEEDDAYLIMIT = "LimitExceeded.CdnPushExceedDayLimit"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTONLINE = "ResourceUnavailable.CdnHostIsNotOnline"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_ECDNMIGRATEDCDN = "UnauthorizedOperation.EcdnMigratedCdn"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) PushUrlsCacheWithContext(ctx context.Context, request *PushUrlsCacheRequest) (response *PushUrlsCacheResponse, err error) {
    if request == nil {
        request = NewPushUrlsCacheRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushUrlsCache require credential")
    }

    request.SetContext(ctx)
    
    response = NewPushUrlsCacheResponse()
    err = c.Send(request, response)
    return
}

func NewSearchClsLogRequest() (request *SearchClsLogRequest) {
    request = &SearchClsLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "SearchClsLog")
    
    
    return
}

func NewSearchClsLogResponse() (response *SearchClsLogResponse) {
    response = &SearchClsLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchClsLog
// This API is used to search for CLS logs. Search filters can be set to today, 24 hours (one of the last 7 days), and the last 7 days.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) SearchClsLog(request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    return c.SearchClsLogWithContext(context.Background(), request)
}

// SearchClsLog
// This API is used to search for CLS logs. Search filters can be set to today, 24 hours (one of the last 7 days), and the last 7 days.
//
// error code that may be returned:
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_CLSINTERNALERROR = "InternalError.ClsInternalError"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_CDNCLSDUPLICATETOPIC = "InvalidParameter.CdnClsDuplicateTopic"
//  INVALIDPARAMETER_CDNCLSTOPICNOTEXIST = "InvalidParameter.CdnClsTopicNotExist"
//  INVALIDPARAMETER_CLSINDEXCONFLICT = "InvalidParameter.ClsIndexConflict"
//  INVALIDPARAMETER_CLSINDEXRULEEMPTY = "InvalidParameter.ClsIndexRuleEmpty"
//  INVALIDPARAMETER_CLSINVALIDCONTENT = "InvalidParameter.ClsInvalidContent"
//  INVALIDPARAMETER_CLSINVALIDCONTENTTYPE = "InvalidParameter.ClsInvalidContentType"
//  INVALIDPARAMETER_CLSINVALIDPARAM = "InvalidParameter.ClsInvalidParam"
//  INVALIDPARAMETER_CLSLOGSETCONFLICT = "InvalidParameter.ClsLogsetConflict"
//  INVALIDPARAMETER_CLSLOGSETEMPTY = "InvalidParameter.ClsLogsetEmpty"
//  INVALIDPARAMETER_CLSLOGSETNOTEMPTY = "InvalidParameter.ClsLogsetNotEmpty"
//  INVALIDPARAMETER_CLSMISSINGAUTHORIZATION = "InvalidParameter.ClsMissingAuthorization"
//  INVALIDPARAMETER_CLSMISSINGCONTENT = "InvalidParameter.ClsMissingContent"
//  INVALIDPARAMETER_CLSSYNTAXERROR = "InvalidParameter.ClsSyntaxError"
//  INVALIDPARAMETER_CLSTOPICCLOSED = "InvalidParameter.ClsTopicClosed"
//  INVALIDPARAMETER_CLSTOPICCONFLICT = "InvalidParameter.ClsTopicConflict"
//  LIMITEXCEEDED_CDNCLSTOOMANYTOPICS = "LimitExceeded.CdnClsTooManyTopics"
//  LIMITEXCEEDED_CLSLOGSIZEEXCEED = "LimitExceeded.ClsLogSizeExceed"
//  LIMITEXCEEDED_CLSLOGSETEXCEED = "LimitExceeded.ClsLogsetExceed"
//  LIMITEXCEEDED_CLSTOPICEXCEED = "LimitExceeded.ClsTopicExceed"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_CLSINDEXNOTEXIST = "ResourceNotFound.ClsIndexNotExist"
//  RESOURCENOTFOUND_CLSLOGSETNOTEXIST = "ResourceNotFound.ClsLogsetNotExist"
//  RESOURCENOTFOUND_CLSTOPICNOTEXIST = "ResourceNotFound.ClsTopicNotExist"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNCLSNOTREGISTERED = "UnauthorizedOperation.CdnClsNotRegistered"
//  UNAUTHORIZEDOPERATION_CLSINVALIDAUTHORIZATION = "UnauthorizedOperation.ClsInvalidAuthorization"
//  UNAUTHORIZEDOPERATION_CLSUNAUTHORIZED = "UnauthorizedOperation.ClsUnauthorized"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) SearchClsLogWithContext(ctx context.Context, request *SearchClsLogRequest) (response *SearchClsLogResponse, err error) {
    if request == nil {
        request = NewSearchClsLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchClsLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchClsLogResponse()
    err = c.Send(request, response)
    return
}

func NewStartCdnDomainRequest() (request *StartCdnDomainRequest) {
    request = &StartCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "StartCdnDomain")
    
    
    return
}

func NewStartCdnDomainResponse() (response *StartCdnDomainResponse) {
    response = &StartCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCdnDomain
// This API is used to enable the acceleration service for a disabled domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERINVALIDCREDENTIAL = "UnauthorizedOperation.CdnUserInvalidCredential"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StartCdnDomain(request *StartCdnDomainRequest) (response *StartCdnDomainResponse, err error) {
    return c.StartCdnDomainWithContext(context.Background(), request)
}

// StartCdnDomain
// This API is used to enable the acceleration service for a disabled domain name.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERINVALIDCREDENTIAL = "UnauthorizedOperation.CdnUserInvalidCredential"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StartCdnDomainWithContext(ctx context.Context, request *StartCdnDomainRequest) (response *StartCdnDomainResponse, err error) {
    if request == nil {
        request = NewStartCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewStopCdnDomainRequest() (request *StopCdnDomainRequest) {
    request = &StopCdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "StopCdnDomain")
    
    
    return
}

func NewStopCdnDomainResponse() (response *StopCdnDomainResponse) {
    response = &StopCdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCdnDomain
// This API is used to suspend the acceleration service for a domain name.
//
// Note: after the acceleration service has been suspended, requests to the cache node will return a 404 error. In order to avoid impact to your business, please move the domain name to another service before suspending the acceleration service.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StopCdnDomain(request *StopCdnDomainRequest) (response *StopCdnDomainResponse, err error) {
    return c.StopCdnDomainWithContext(context.Background(), request)
}

// StopCdnDomain
// This API is used to suspend the acceleration service for a domain name.
//
// Note: after the acceleration service has been suspended, requests to the cache node will return a 404 error. In order to avoid impact to your business, please move the domain name to another service before suspending the acceleration service.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTHASSPECIALCONFIG = "InvalidParameter.CdnHostHasSpecialConfig"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
//  UNSUPPORTEDOPERATION_OPNOAUTH = "UnsupportedOperation.OpNoAuth"
func (c *Client) StopCdnDomainWithContext(ctx context.Context, request *StopCdnDomainRequest) (response *StopCdnDomainResponse, err error) {
    if request == nil {
        request = NewStopCdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCdnDomainResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDomainConfigRequest() (request *UpdateDomainConfigRequest) {
    request = &UpdateDomainConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateDomainConfig")
    
    
    return
}

func NewUpdateDomainConfigResponse() (response *UpdateDomainConfigResponse) {
    response = &UpdateDomainConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateDomainConfig
// This API is used to modify the configuration of CDN acceleration domain names.
//
// Note: To update complex configuration items, all attributes of the object must be specified, or the default values are used. We recommend calling the querying API to get attributes before modifying and passing them to this API.
//
// The input parameters of this API are not reported to CloudAudit as it may contain sensitive data, such as keys and secrets.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"
//  INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  OPERATIONDENIED_SHARECACHEAREADNSNOTMATCH = "OperationDenied.ShareCacheAreaDnsNotMatch"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdateDomainConfig(request *UpdateDomainConfigRequest) (response *UpdateDomainConfigResponse, err error) {
    return c.UpdateDomainConfigWithContext(context.Background(), request)
}

// UpdateDomainConfig
// This API is used to modify the configuration of CDN acceleration domain names.
//
// Note: To update complex configuration items, all attributes of the object must be specified, or the default values are used. We recommend calling the querying API to get attributes before modifying and passing them to this API.
//
// The input parameters of this API are not reported to CloudAudit as it may contain sensitive data, such as keys and secrets.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CAMSYSTEMERROR = "InternalError.CamSystemError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNQUERYSYSTEMERROR = "InternalError.CdnQuerySystemError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_ERROR = "InternalError.Error"
//  INTERNALERROR_INVALIDERRORCODE = "InternalError.InvalidErrorCode"
//  INTERNALERROR_PROXYSERVER = "InternalError.ProxyServer"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_BANDLIMITREQUIREDMAINLAND = "InvalidParameter.BandLimitRequiredMainland"
//  INVALIDPARAMETER_BANDWIDTHALERTCOUNTERMEASURECONFLICTORIGINTYPE = "InvalidParameter.BandwidthAlertCounterMeasureConflictOriginType"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNCERTINFONOTFOUND = "InvalidParameter.CdnCertInfoNotFound"
//  INVALIDPARAMETER_CDNCONFIGINVALIDCACHE = "InvalidParameter.CdnConfigInvalidCache"
//  INVALIDPARAMETER_CDNHOSTINVALIDMIDDLECONFIG = "InvalidParameter.CdnHostInvalidMiddleConfig"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNINVALIDPARAMSERVICETYPE = "InvalidParameter.CdnInvalidParamServiceType"
//  INVALIDPARAMETER_CDNKEYRULESEXCLUDECUSTOMREQUIRESFULLLEGO = "InvalidParameter.CdnKeyRulesExcludeCustomRequiresFullLego"
//  INVALIDPARAMETER_CDNKEYRULESINVALIDQUERYSTRINGVALUE = "InvalidParameter.CdnKeyRulesInvalidQueryStringValue"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  INVALIDPARAMETER_CDNSTATINVALIDPROJECTID = "InvalidParameter.CdnStatInvalidProjectId"
//  INVALIDPARAMETER_CDNSTATTOOMANYDOMAINS = "InvalidParameter.CdnStatTooManyDomains"
//  INVALIDPARAMETER_PATHREGEXTOOMANYSUBPATTERN = "InvalidParameter.PathRegexTooManySubPattern"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPLATFORM = "InvalidParameter.RemoteAuthInvalidPlatform"
//  INVALIDPARAMETER_REMOTEAUTHINVALIDPROTOCOL = "InvalidParameter.RemoteAuthInvalidProtocol"
//  LIMITEXCEEDED_CDNCONFIGTOOMANYCACHERULES = "LimitExceeded.CdnConfigTooManyCacheRules"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  OPERATIONDENIED_SHARECACHEAREADNSNOTMATCH = "OperationDenied.ShareCacheAreaDnsNotMatch"
//  RESOURCEINUSE_CDNHOSTEXISTS = "ResourceInUse.CdnHostExists"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCENOTFOUND_ECDNDOMAINNOTEXISTS = "ResourceNotFound.EcdnDomainNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINMAINLAND = "ResourceUnavailable.CdnHostBelongsToOthersInMainland"
//  RESOURCEUNAVAILABLE_CDNHOSTBELONGSTOOTHERSINOVERSEAS = "ResourceUnavailable.CdnHostBelongsToOthersInOverseas"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_CDNHOSTISNOTOFFLINE = "ResourceUnavailable.CdnHostIsNotOffline"
//  RESOURCEUNAVAILABLE_CDNHOSTNOICP = "ResourceUnavailable.CdnHostNoIcp"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTINICPBLACKLIST = "UnauthorizedOperation.CdnHostInIcpBlacklist"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNPROJECTUNAUTHORIZED = "UnauthorizedOperation.CdnProjectUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNTAGUNAUTHORIZED = "UnauthorizedOperation.CdnTagUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHFAIL = "UnauthorizedOperation.CdnUserAuthFail"
//  UNAUTHORIZEDOPERATION_CDNUSERAUTHWAIT = "UnauthorizedOperation.CdnUserAuthWait"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
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

func NewUpdatePayTypeRequest() (request *UpdatePayTypeRequest) {
    request = &UpdatePayTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdatePayType")
    
    
    return
}

func NewUpdatePayTypeResponse() (response *UpdatePayTypeResponse) {
    response = &UpdatePayTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePayType
// This API is used to modify the billing mode of an account. At present, the billing mode of accounts on a monthly billing cycle and sub-accounts cannot be modified.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdatePayType(request *UpdatePayTypeRequest) (response *UpdatePayTypeResponse, err error) {
    return c.UpdatePayTypeWithContext(context.Background(), request)
}

// UpdatePayType
// This API is used to modify the billing mode of an account. At present, the billing mode of accounts on a monthly billing cycle and sub-accounts cannot be modified.
//
// error code that may be returned:
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISISOLATED = "UnauthorizedOperation.CdnUserIsIsolated"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_OPNOAUTH = "UnauthorizedOperation.OpNoAuth"
//  UNAUTHORIZEDOPERATION_OPERATIONTOOOFTEN = "UnauthorizedOperation.OperationTooOften"
func (c *Client) UpdatePayTypeWithContext(ctx context.Context, request *UpdatePayTypeRequest) (response *UpdatePayTypeResponse, err error) {
    if request == nil {
        request = NewUpdatePayTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePayType require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePayTypeResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScdnDomainRequest() (request *UpdateScdnDomainRequest) {
    request = &UpdateScdnDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdn", APIVersion, "UpdateScdnDomain")
    
    
    return
}

func NewUpdateScdnDomainResponse() (response *UpdateScdnDomainResponse) {
    response = &UpdateScdnDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateScdnDomain
// This API is used to modify security configurations of SCDN acceleration domain names.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INTERNALERROR_SCDNUSERSUSPEND = "InternalError.ScdnUserSuspend"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) UpdateScdnDomain(request *UpdateScdnDomainRequest) (response *UpdateScdnDomainResponse, err error) {
    return c.UpdateScdnDomainWithContext(context.Background(), request)
}

// UpdateScdnDomain
// This API is used to modify security configurations of SCDN acceleration domain names.
//
// error code that may be returned:
//  FAILEDOPERATION_CDNCONFIGERROR = "FailedOperation.CdnConfigError"
//  INTERNALERROR_CDNCONFIGERROR = "InternalError.CdnConfigError"
//  INTERNALERROR_CDNDBERROR = "InternalError.CdnDbError"
//  INTERNALERROR_CDNSYSTEMERROR = "InternalError.CdnSystemError"
//  INTERNALERROR_SCDNUSERNOPACKAGE = "InternalError.ScdnUserNoPackage"
//  INTERNALERROR_SCDNUSERSUSPEND = "InternalError.ScdnUserSuspend"
//  INVALIDPARAMETER_ACCESSPORTOPENEDHTTPS = "InvalidParameter.AccessPortOpenedHttps"
//  INVALIDPARAMETER_CDNSTATUSINVALIDDOMAIN = "InvalidParameter.CDNStatusInvalidDomain"
//  INVALIDPARAMETER_CDNHOSTINTERNALHOST = "InvalidParameter.CdnHostInternalHost"
//  INVALIDPARAMETER_CDNHOSTINVALIDSTATUS = "InvalidParameter.CdnHostInvalidStatus"
//  INVALIDPARAMETER_CDNINTERFACEERROR = "InvalidParameter.CdnInterfaceError"
//  INVALIDPARAMETER_CDNPARAMERROR = "InvalidParameter.CdnParamError"
//  LIMITEXCEEDED_CDNHOSTOPTOOOFTEN = "LimitExceeded.CdnHostOpTooOften"
//  RESOURCEINUSE_CDNOPINPROGRESS = "ResourceInUse.CdnOpInProgress"
//  RESOURCENOTFOUND_CDNHOSTNOTEXISTS = "ResourceNotFound.CdnHostNotExists"
//  RESOURCENOTFOUND_CDNUSERNOTEXISTS = "ResourceNotFound.CdnUserNotExists"
//  RESOURCEUNAVAILABLE_CDNHOSTISLOCKED = "ResourceUnavailable.CdnHostIsLocked"
//  RESOURCEUNAVAILABLE_SCDNUSERNOPACKAGE = "ResourceUnavailable.ScdnUserNoPackage"
//  UNAUTHORIZEDOPERATION_CDNCAMUNAUTHORIZED = "UnauthorizedOperation.CdnCamUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNHOSTUNAUTHORIZED = "UnauthorizedOperation.CdnHostUnauthorized"
//  UNAUTHORIZEDOPERATION_CDNUSERISSUSPENDED = "UnauthorizedOperation.CdnUserIsSuspended"
//  UNAUTHORIZEDOPERATION_CDNUSERNOWHITELIST = "UnauthorizedOperation.CdnUserNoWhitelist"
func (c *Client) UpdateScdnDomainWithContext(ctx context.Context, request *UpdateScdnDomainRequest) (response *UpdateScdnDomainResponse, err error) {
    if request == nil {
        request = NewUpdateScdnDomainRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateScdnDomain require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateScdnDomainResponse()
    err = c.Send(request, response)
    return
}

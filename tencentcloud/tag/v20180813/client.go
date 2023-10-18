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

package v20180813

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-08-13"

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


func NewAddResourceTagRequest() (request *AddResourceTagRequest) {
    request = &AddResourceTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "AddResourceTag")
    
    
    return
}

func NewAddResourceTagResponse() (response *AddResourceTagResponse) {
    response = &AddResourceTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddResourceTag
// This API is used to associate resources with tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCEINUSE_TAGKEYATTACHED = "ResourceInUse.TagKeyAttached"
func (c *Client) AddResourceTag(request *AddResourceTagRequest) (response *AddResourceTagResponse, err error) {
    return c.AddResourceTagWithContext(context.Background(), request)
}

// AddResourceTag
// This API is used to associate resources with tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCEINUSE_TAGKEYATTACHED = "ResourceInUse.TagKeyAttached"
func (c *Client) AddResourceTagWithContext(ctx context.Context, request *AddResourceTagRequest) (response *AddResourceTagResponse, err error) {
    if request == nil {
        request = NewAddResourceTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddResourceTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddResourceTagResponse()
    err = c.Send(request, response)
    return
}

func NewAttachResourcesTagRequest() (request *AttachResourcesTagRequest) {
    request = &AttachResourcesTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "AttachResourcesTag")
    
    
    return
}

func NewAttachResourcesTagResponse() (response *AttachResourcesTagResponse) {
    response = &AttachResourcesTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AttachResourcesTag
// This API is used to associate a tag with multiple resources.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE_TAGKEYATTACHED = "ResourceInUse.TagKeyAttached"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) AttachResourcesTag(request *AttachResourcesTagRequest) (response *AttachResourcesTagResponse, err error) {
    return c.AttachResourcesTagWithContext(context.Background(), request)
}

// AttachResourcesTag
// This API is used to associate a tag with multiple resources.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE_TAGKEYATTACHED = "ResourceInUse.TagKeyAttached"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) AttachResourcesTagWithContext(ctx context.Context, request *AttachResourcesTagRequest) (response *AttachResourcesTagResponse, err error) {
    if request == nil {
        request = NewAttachResourcesTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachResourcesTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachResourcesTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagRequest() (request *CreateTagRequest) {
    request = &CreateTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "CreateTag")
    
    
    return
}

func NewCreateTagResponse() (response *CreateTagResponse) {
    response = &CreateTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTag
// This API is used to create a tag key and tag value pair.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCEINUSE_TAGDUPLICATE = "ResourceInUse.TagDuplicate"
func (c *Client) CreateTag(request *CreateTagRequest) (response *CreateTagResponse, err error) {
    return c.CreateTagWithContext(context.Background(), request)
}

// CreateTag
// This API is used to create a tag key and tag value pair.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCEINUSE_TAGDUPLICATE = "ResourceInUse.TagDuplicate"
func (c *Client) CreateTagWithContext(ctx context.Context, request *CreateTagRequest) (response *CreateTagResponse, err error) {
    if request == nil {
        request = NewCreateTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTagsRequest() (request *CreateTagsRequest) {
    request = &CreateTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "CreateTags")
    
    
    return
}

func NewCreateTagsResponse() (response *CreateTagsResponse) {
    response = &CreateTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTags
// This API is used to create multiple tag key-value pairs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCEINUSE_TAGDUPLICATE = "ResourceInUse.TagDuplicate"
func (c *Client) CreateTags(request *CreateTagsRequest) (response *CreateTagsResponse, err error) {
    return c.CreateTagsWithContext(context.Background(), request)
}

// CreateTags
// This API is used to create multiple tag key-value pairs.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCEINUSE_TAGDUPLICATE = "ResourceInUse.TagDuplicate"
func (c *Client) CreateTagsWithContext(ctx context.Context, request *CreateTagsRequest) (response *CreateTagsResponse, err error) {
    if request == nil {
        request = NewCreateTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceTagRequest() (request *DeleteResourceTagRequest) {
    request = &DeleteResourceTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DeleteResourceTag")
    
    
    return
}

func NewDeleteResourceTagResponse() (response *DeleteResourceTagResponse) {
    response = &DeleteResourceTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteResourceTag
// This API is used to unassociate tags and resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
func (c *Client) DeleteResourceTag(request *DeleteResourceTagRequest) (response *DeleteResourceTagResponse, err error) {
    return c.DeleteResourceTagWithContext(context.Background(), request)
}

// DeleteResourceTag
// This API is used to unassociate tags and resources.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
func (c *Client) DeleteResourceTagWithContext(ctx context.Context, request *DeleteResourceTagRequest) (response *DeleteResourceTagResponse, err error) {
    if request == nil {
        request = NewDeleteResourceTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteResourceTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagRequest() (request *DeleteTagRequest) {
    request = &DeleteTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DeleteTag")
    
    
    return
}

func NewDeleteTagResponse() (response *DeleteTagResponse) {
    response = &DeleteTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTag
// This API is used to delete a tag key and tag value pair.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGATTACHEDQUOTA = "FailedOperation.TagAttachedQuota"
//  FAILEDOPERATION_TAGATTACHEDRESOURCE = "FailedOperation.TagAttachedResource"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) DeleteTag(request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
    return c.DeleteTagWithContext(context.Background(), request)
}

// DeleteTag
// This API is used to delete a tag key and tag value pair.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGATTACHEDQUOTA = "FailedOperation.TagAttachedQuota"
//  FAILEDOPERATION_TAGATTACHEDRESOURCE = "FailedOperation.TagAttachedResource"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) DeleteTagWithContext(ctx context.Context, request *DeleteTagRequest) (response *DeleteTagResponse, err error) {
    if request == nil {
        request = NewDeleteTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTagResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTagsRequest() (request *DeleteTagsRequest) {
    request = &DeleteTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DeleteTags")
    
    
    return
}

func NewDeleteTagsResponse() (response *DeleteTagsResponse) {
    response = &DeleteTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTags
// This API is used to delete tag keys and tag values in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGATTACHEDQUOTA = "FailedOperation.TagAttachedQuota"
//  FAILEDOPERATION_TAGATTACHEDRESOURCE = "FailedOperation.TagAttachedResource"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) DeleteTags(request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
    return c.DeleteTagsWithContext(context.Background(), request)
}

// DeleteTags
// This API is used to delete tag keys and tag values in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_TAGATTACHEDQUOTA = "FailedOperation.TagAttachedQuota"
//  FAILEDOPERATION_TAGATTACHEDRESOURCE = "FailedOperation.TagAttachedResource"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) DeleteTagsWithContext(ctx context.Context, request *DeleteTagsRequest) (response *DeleteTagsResponse, err error) {
    if request == nil {
        request = NewDeleteTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProjectsRequest() (request *DescribeProjectsRequest) {
    request = &DescribeProjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeProjects")
    
    
    return
}

func NewDescribeProjectsResponse() (response *DescribeProjectsResponse) {
    response = &DescribeProjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProjects
// This API is used to get project lists.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProjects(request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    return c.DescribeProjectsWithContext(context.Background(), request)
}

// DescribeProjects
// This API is used to get project lists.
//
// error code that may be returned:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeProjectsWithContext(ctx context.Context, request *DescribeProjectsRequest) (response *DescribeProjectsResponse, err error) {
    if request == nil {
        request = NewDescribeProjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProjectsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsRequest() (request *DescribeResourceTagsRequest) {
    request = &DescribeResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourceTags")
    
    
    return
}

func NewDescribeResourceTagsResponse() (response *DescribeResourceTagsResponse) {
    response = &DescribeResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceTags
// This API is used to query the tags associated with a resource.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceTags(request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
    return c.DescribeResourceTagsWithContext(context.Background(), request)
}

// DescribeResourceTags
// This API is used to query the tags associated with a resource.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceTagsWithContext(ctx context.Context, request *DescribeResourceTagsRequest) (response *DescribeResourceTagsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsByResourceIdsRequest() (request *DescribeResourceTagsByResourceIdsRequest) {
    request = &DescribeResourceTagsByResourceIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourceTagsByResourceIds")
    
    
    return
}

func NewDescribeResourceTagsByResourceIdsResponse() (response *DescribeResourceTagsByResourceIdsResponse) {
    response = &DescribeResourceTagsByResourceIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceTagsByResourceIds
// This API is used to query the tag key-value pairs associated with an existing resource.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourceTagsByResourceIds(request *DescribeResourceTagsByResourceIdsRequest) (response *DescribeResourceTagsByResourceIdsResponse, err error) {
    return c.DescribeResourceTagsByResourceIdsWithContext(context.Background(), request)
}

// DescribeResourceTagsByResourceIds
// This API is used to query the tag key-value pairs associated with an existing resource.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourceTagsByResourceIdsWithContext(ctx context.Context, request *DescribeResourceTagsByResourceIdsRequest) (response *DescribeResourceTagsByResourceIdsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsByResourceIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTagsByResourceIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTagsByResourceIdsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsByResourceIdsSeqRequest() (request *DescribeResourceTagsByResourceIdsSeqRequest) {
    request = &DescribeResourceTagsByResourceIdsSeqRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourceTagsByResourceIdsSeq")
    
    
    return
}

func NewDescribeResourceTagsByResourceIdsSeqResponse() (response *DescribeResourceTagsByResourceIdsSeqResponse) {
    response = &DescribeResourceTagsByResourceIdsSeqResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceTagsByResourceIdsSeq
// This API is used to view the tags associated with a resource in sequence.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeResourceTagsByResourceIdsSeq(request *DescribeResourceTagsByResourceIdsSeqRequest) (response *DescribeResourceTagsByResourceIdsSeqResponse, err error) {
    return c.DescribeResourceTagsByResourceIdsSeqWithContext(context.Background(), request)
}

// DescribeResourceTagsByResourceIdsSeq
// This API is used to view the tags associated with a resource in sequence.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeResourceTagsByResourceIdsSeqWithContext(ctx context.Context, request *DescribeResourceTagsByResourceIdsSeqRequest) (response *DescribeResourceTagsByResourceIdsSeqResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsByResourceIdsSeqRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTagsByResourceIdsSeq require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTagsByResourceIdsSeqResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceTagsByTagKeysRequest() (request *DescribeResourceTagsByTagKeysRequest) {
    request = &DescribeResourceTagsByTagKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourceTagsByTagKeys")
    
    
    return
}

func NewDescribeResourceTagsByTagKeysResponse() (response *DescribeResourceTagsByTagKeysResponse) {
    response = &DescribeResourceTagsByTagKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourceTagsByTagKeys
// This API is used to get resource tags based on tag keys.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
func (c *Client) DescribeResourceTagsByTagKeys(request *DescribeResourceTagsByTagKeysRequest) (response *DescribeResourceTagsByTagKeysResponse, err error) {
    return c.DescribeResourceTagsByTagKeysWithContext(context.Background(), request)
}

// DescribeResourceTagsByTagKeys
// This API is used to get resource tags based on tag keys.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
func (c *Client) DescribeResourceTagsByTagKeysWithContext(ctx context.Context, request *DescribeResourceTagsByTagKeysRequest) (response *DescribeResourceTagsByTagKeysResponse, err error) {
    if request == nil {
        request = NewDescribeResourceTagsByTagKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceTagsByTagKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceTagsByTagKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByTagsRequest() (request *DescribeResourcesByTagsRequest) {
    request = &DescribeResourcesByTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourcesByTags")
    
    
    return
}

func NewDescribeResourcesByTagsResponse() (response *DescribeResourcesByTagsResponse) {
    response = &DescribeResourcesByTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcesByTags
// This API is used to query resources by tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_TAGFILTERS = "InvalidParameterValue.TagFilters"
//  INVALIDPARAMETERVALUE_TAGFILTERSLENGTHEXCEEDED = "InvalidParameterValue.TagFiltersLengthExceeded"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourcesByTags(request *DescribeResourcesByTagsRequest) (response *DescribeResourcesByTagsResponse, err error) {
    return c.DescribeResourcesByTagsWithContext(context.Background(), request)
}

// DescribeResourcesByTags
// This API is used to query resources by tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_TAGFILTERS = "InvalidParameterValue.TagFilters"
//  INVALIDPARAMETERVALUE_TAGFILTERSLENGTHEXCEEDED = "InvalidParameterValue.TagFiltersLengthExceeded"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeResourcesByTagsWithContext(ctx context.Context, request *DescribeResourcesByTagsRequest) (response *DescribeResourcesByTagsResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcesByTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesByTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesByTagsUnionRequest() (request *DescribeResourcesByTagsUnionRequest) {
    request = &DescribeResourcesByTagsUnionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeResourcesByTagsUnion")
    
    
    return
}

func NewDescribeResourcesByTagsUnionResponse() (response *DescribeResourcesByTagsUnionResponse) {
    response = &DescribeResourcesByTagsUnionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResourcesByTagsUnion
// This API is used to query resource list by tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_TAGFILTERS = "InvalidParameterValue.TagFilters"
//  INVALIDPARAMETERVALUE_TAGFILTERSLENGTHEXCEEDED = "InvalidParameterValue.TagFiltersLengthExceeded"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeResourcesByTagsUnion(request *DescribeResourcesByTagsUnionRequest) (response *DescribeResourcesByTagsUnionResponse, err error) {
    return c.DescribeResourcesByTagsUnionWithContext(context.Background(), request)
}

// DescribeResourcesByTagsUnion
// This API is used to query resource list by tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_TAGFILTERS = "InvalidParameterValue.TagFilters"
//  INVALIDPARAMETERVALUE_TAGFILTERSLENGTHEXCEEDED = "InvalidParameterValue.TagFiltersLengthExceeded"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeResourcesByTagsUnionWithContext(ctx context.Context, request *DescribeResourcesByTagsUnionRequest) (response *DescribeResourcesByTagsUnionResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesByTagsUnionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourcesByTagsUnion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourcesByTagsUnionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagKeysRequest() (request *DescribeTagKeysRequest) {
    request = &DescribeTagKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTagKeys")
    
    
    return
}

func NewDescribeTagKeysResponse() (response *DescribeTagKeysResponse) {
    response = &DescribeTagKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagKeys
// This API is used to query tag keys in an existing tag list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagKeys(request *DescribeTagKeysRequest) (response *DescribeTagKeysResponse, err error) {
    return c.DescribeTagKeysWithContext(context.Background(), request)
}

// DescribeTagKeys
// This API is used to query tag keys in an existing tag list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagKeysWithContext(ctx context.Context, request *DescribeTagKeysRequest) (response *DescribeTagKeysResponse, err error) {
    if request == nil {
        request = NewDescribeTagKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
    request = &DescribeTagValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTagValues")
    
    
    return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
    response = &DescribeTagValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagValues
// This API is used to query tag values in an existing tag list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    return c.DescribeTagValuesWithContext(context.Background(), request)
}

// DescribeTagValues
// This API is used to query tag values in an existing tag list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagValuesWithContext(ctx context.Context, request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
    if request == nil {
        request = NewDescribeTagValuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagValuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagValuesSeqRequest() (request *DescribeTagValuesSeqRequest) {
    request = &DescribeTagValuesSeqRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTagValuesSeq")
    
    
    return
}

func NewDescribeTagValuesSeqResponse() (response *DescribeTagValuesSeqResponse) {
    response = &DescribeTagValuesSeqResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagValuesSeq
// This API is used to query tag values in a created tag list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagValuesSeq(request *DescribeTagValuesSeqRequest) (response *DescribeTagValuesSeqResponse, err error) {
    return c.DescribeTagValuesSeqWithContext(context.Background(), request)
}

// DescribeTagValuesSeq
// This API is used to query tag values in a created tag list.
//
// error code that may be returned:
//  INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagValuesSeqWithContext(ctx context.Context, request *DescribeTagValuesSeqRequest) (response *DescribeTagValuesSeqResponse, err error) {
    if request == nil {
        request = NewDescribeTagValuesSeqRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagValuesSeq require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagValuesSeqResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsRequest() (request *DescribeTagsRequest) {
    request = &DescribeTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTags")
    
    
    return
}

func NewDescribeTagsResponse() (response *DescribeTagsResponse) {
    response = &DescribeTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTags
// This API is used to query existing tag lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTags(request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    return c.DescribeTagsWithContext(context.Background(), request)
}

// DescribeTags
// This API is used to query existing tag lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagsWithContext(ctx context.Context, request *DescribeTagsRequest) (response *DescribeTagsResponse, err error) {
    if request == nil {
        request = NewDescribeTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTagsSeqRequest() (request *DescribeTagsSeqRequest) {
    request = &DescribeTagsSeqRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DescribeTagsSeq")
    
    
    return
}

func NewDescribeTagsSeqResponse() (response *DescribeTagsSeqResponse) {
    response = &DescribeTagsSeqResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTagsSeq
// This API is used to query the created tag lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagsSeq(request *DescribeTagsSeqRequest) (response *DescribeTagsSeqResponse, err error) {
    return c.DescribeTagsSeqWithContext(context.Background(), request)
}

// DescribeTagsSeq
// This API is used to query the created tag lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE_OFFSETINVALID = "InvalidParameterValue.OffsetInvalid"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
func (c *Client) DescribeTagsSeqWithContext(ctx context.Context, request *DescribeTagsSeqRequest) (response *DescribeTagsSeqResponse, err error) {
    if request == nil {
        request = NewDescribeTagsSeqRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTagsSeq require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTagsSeqResponse()
    err = c.Send(request, response)
    return
}

func NewDetachResourcesTagRequest() (request *DetachResourcesTagRequest) {
    request = &DetachResourcesTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "DetachResourcesTag")
    
    
    return
}

func NewDetachResourcesTagResponse() (response *DetachResourcesTagResponse) {
    response = &DetachResourcesTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DetachResourcesTag
// This API is used to unbind a tag from multiple resources.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
func (c *Client) DetachResourcesTag(request *DetachResourcesTagRequest) (response *DetachResourcesTagResponse, err error) {
    return c.DetachResourcesTagWithContext(context.Background(), request)
}

// DetachResourcesTag
// This API is used to unbind a tag from multiple resources.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
func (c *Client) DetachResourcesTagWithContext(ctx context.Context, request *DetachResourcesTagRequest) (response *DetachResourcesTagResponse, err error) {
    if request == nil {
        request = NewDetachResourcesTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachResourcesTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachResourcesTagResponse()
    err = c.Send(request, response)
    return
}

func NewGetResourcesRequest() (request *GetResourcesRequest) {
    request = &GetResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "GetResources")
    
    
    return
}

func NewGetResourcesResponse() (response *GetResourcesResponse) {
    response = &GetResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetResources
// This API is used to query the list of resources associated with a tag.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
func (c *Client) GetResources(request *GetResourcesRequest) (response *GetResourcesResponse, err error) {
    return c.GetResourcesWithContext(context.Background(), request)
}

// GetResources
// This API is used to query the list of resources associated with a tag.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
func (c *Client) GetResourcesWithContext(ctx context.Context, request *GetResourcesRequest) (response *GetResourcesResponse, err error) {
    if request == nil {
        request = NewGetResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTagKeysRequest() (request *GetTagKeysRequest) {
    request = &GetTagKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "GetTagKeys")
    
    
    return
}

func NewGetTagKeysResponse() (response *GetTagKeysResponse) {
    response = &GetTagKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTagKeys
// This API is used to query the list of tag keys.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetTagKeys(request *GetTagKeysRequest) (response *GetTagKeysResponse, err error) {
    return c.GetTagKeysWithContext(context.Background(), request)
}

// GetTagKeys
// This API is used to query the list of tag keys.
//
// error code that may be returned:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetTagKeysWithContext(ctx context.Context, request *GetTagKeysRequest) (response *GetTagKeysResponse, err error) {
    if request == nil {
        request = NewGetTagKeysRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTagKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTagKeysResponse()
    err = c.Send(request, response)
    return
}

func NewGetTagValuesRequest() (request *GetTagValuesRequest) {
    request = &GetTagValuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "GetTagValues")
    
    
    return
}

func NewGetTagValuesResponse() (response *GetTagValuesResponse) {
    response = &GetTagValuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTagValues
// This API is used to query tag values in the list of created tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTagValues(request *GetTagValuesRequest) (response *GetTagValuesResponse, err error) {
    return c.GetTagValuesWithContext(context.Background(), request)
}

// GetTagValues
// This API is used to query tag values in the list of created tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETERVALUE_RESERVEDTAGKEY = "InvalidParameterValue.ReservedTagKey"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetTagValuesWithContext(ctx context.Context, request *GetTagValuesRequest) (response *GetTagValuesResponse, err error) {
    if request == nil {
        request = NewGetTagValuesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTagValues require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTagValuesResponse()
    err = c.Send(request, response)
    return
}

func NewGetTagsRequest() (request *GetTagsRequest) {
    request = &GetTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "GetTags")
    
    
    return
}

func NewGetTagsResponse() (response *GetTagsResponse) {
    response = &GetTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTags
// This API is used to get the list of created tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetTags(request *GetTagsRequest) (response *GetTagsResponse, err error) {
    return c.GetTagsWithContext(context.Background(), request)
}

// GetTags
// This API is used to get the list of created tags.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PAGINATIONTOKENINVALID = "InvalidParameter.PaginationTokenInvalid"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetTagsWithContext(ctx context.Context, request *GetTagsRequest) (response *GetTagsResponse, err error) {
    if request == nil {
        request = NewGetTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourceTagsRequest() (request *ModifyResourceTagsRequest) {
    request = &ModifyResourceTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "ModifyResourceTags")
    
    
    return
}

func NewModifyResourceTagsResponse() (response *ModifyResourceTagsResponse) {
    response = &ModifyResourceTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourceTags
// This API is used to modify all tags associated with a resource.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_DELETETAGSPARAMERROR = "InvalidParameterValue.DeleteTagsParamError"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
func (c *Client) ModifyResourceTags(request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    return c.ModifyResourceTagsWithContext(context.Background(), request)
}

// ModifyResourceTags
// This API is used to modify all tags associated with a resource.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_DELETETAGSPARAMERROR = "InvalidParameterValue.DeleteTagsParamError"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  LIMITEXCEEDED_RESOURCEATTACHEDTAGS = "LimitExceeded.ResourceAttachedTags"
//  LIMITEXCEEDED_TAGKEY = "LimitExceeded.TagKey"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
func (c *Client) ModifyResourceTagsWithContext(ctx context.Context, request *ModifyResourceTagsRequest) (response *ModifyResourceTagsResponse, err error) {
    if request == nil {
        request = NewModifyResourceTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourceTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourceTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyResourcesTagValueRequest() (request *ModifyResourcesTagValueRequest) {
    request = &ModifyResourcesTagValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "ModifyResourcesTagValue")
    
    
    return
}

func NewModifyResourcesTagValueResponse() (response *ModifyResourcesTagValueResponse) {
    response = &ModifyResourcesTagValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyResourcesTagValue
// This API is used to modify the tag value corresponding to a tag key associated with multiple resources.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) ModifyResourcesTagValue(request *ModifyResourcesTagValueRequest) (response *ModifyResourcesTagValueResponse, err error) {
    return c.ModifyResourcesTagValueWithContext(context.Background(), request)
}

// ModifyResourcesTagValue
// This API is used to modify the tag value corresponding to a tag key associated with multiple resources.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCEAPPIDNOTSAME = "FailedOperation.ResourceAppIdNotSame"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_REGIONINVALID = "InvalidParameterValue.RegionInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_RESOURCEIDINVALID = "InvalidParameterValue.ResourceIdInvalid"
//  INVALIDPARAMETERVALUE_RESOURCEPREFIXINVALID = "InvalidParameterValue.ResourcePrefixInvalid"
//  INVALIDPARAMETERVALUE_SERVICETYPEINVALID = "InvalidParameterValue.ServiceTypeInvalid"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  INVALIDPARAMETERVALUE_UININVALID = "InvalidParameterValue.UinInvalid"
//  LIMITEXCEEDED_QUOTA = "LimitExceeded.Quota"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
//  RESOURCENOTFOUND_TAGNONEXIST = "ResourceNotFound.TagNonExist"
func (c *Client) ModifyResourcesTagValueWithContext(ctx context.Context, request *ModifyResourcesTagValueRequest) (response *ModifyResourcesTagValueResponse, err error) {
    if request == nil {
        request = NewModifyResourcesTagValueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyResourcesTagValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyResourcesTagValueResponse()
    err = c.Send(request, response)
    return
}

func NewTagResourcesRequest() (request *TagResourcesRequest) {
    request = &TagResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "TagResources")
    
    
    return
}

func NewTagResourcesResponse() (response *TagResourcesResponse) {
    response = &TagResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TagResources
// This API is used to create and bind a tag uniformly to multiple specified resources of multiple Tencent Cloud services.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
func (c *Client) TagResources(request *TagResourcesRequest) (response *TagResourcesResponse, err error) {
    return c.TagResourcesWithContext(context.Background(), request)
}

// TagResources
// This API is used to create and bind a tag uniformly to multiple specified resources of multiple Tencent Cloud services.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUEEMPTY = "InvalidParameterValue.TagValueEmpty"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
func (c *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest) (response *TagResourcesResponse, err error) {
    if request == nil {
        request = NewTagResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TagResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewTagResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewUnTagResourcesRequest() (request *UnTagResourcesRequest) {
    request = &UnTagResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "UnTagResources")
    
    
    return
}

func NewUnTagResourcesResponse() (response *UnTagResourcesResponse) {
    response = &UnTagResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnTagResources
// This API is used to unbind a tag uniformly from multiple specified resources of multiple Tencent Cloud services.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"
func (c *Client) UnTagResources(request *UnTagResourcesRequest) (response *UnTagResourcesResponse, err error) {
    return c.UnTagResourcesWithContext(context.Background(), request)
}

// UnTagResources
// This API is used to unbind a tag uniformly from multiple specified resources of multiple Tencent Cloud services.
//
// error code that may be returned:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESOURCETAGPROCESSING = "FailedOperation.ResourceTagProcessing"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_TAG = "InvalidParameter.Tag"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGKEYDUPLICATE = "InvalidParameterValue.TagKeyDuplicate"
//  INVALIDPARAMETERVALUE_TAGKEYEMPTY = "InvalidParameterValue.TagKeyEmpty"
//  INVALIDPARAMETERVALUE_TAGKEYLENGTHEXCEEDED = "InvalidParameterValue.TagKeyLengthExceeded"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  LIMITEXCEEDED_RESOURCENUMPERREQUEST = "LimitExceeded.ResourceNumPerRequest"
//  LIMITEXCEEDED_TAGNUMPERREQUEST = "LimitExceeded.TagNumPerRequest"
func (c *Client) UnTagResourcesWithContext(ctx context.Context, request *UnTagResourcesRequest) (response *UnTagResourcesResponse, err error) {
    if request == nil {
        request = NewUnTagResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnTagResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnTagResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateResourceTagValueRequest() (request *UpdateResourceTagValueRequest) {
    request = &UpdateResourceTagValueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tag", APIVersion, "UpdateResourceTagValue")
    
    
    return
}

func NewUpdateResourceTagValueResponse() (response *UpdateResourceTagValueResponse) {
    response = &UpdateResourceTagValueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateResourceTagValue
// This API is used to modify the values of tags associated with a resource (the tag key does not change).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
func (c *Client) UpdateResourceTagValue(request *UpdateResourceTagValueRequest) (response *UpdateResourceTagValueResponse, err error) {
    return c.UpdateResourceTagValueWithContext(context.Background(), request)
}

// UpdateResourceTagValue
// This API is used to modify the values of tags associated with a resource (the tag key does not change).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETERVALUE_RESOURCEDESCRIPTIONERROR = "InvalidParameterValue.ResourceDescriptionError"
//  INVALIDPARAMETERVALUE_TAGKEYCHARACTERILLEGAL = "InvalidParameterValue.TagKeyCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUECHARACTERILLEGAL = "InvalidParameterValue.TagValueCharacterIllegal"
//  INVALIDPARAMETERVALUE_TAGVALUELENGTHEXCEEDED = "InvalidParameterValue.TagValueLengthExceeded"
//  LIMITEXCEEDED_TAGVALUE = "LimitExceeded.TagValue"
//  RESOURCENOTFOUND_ATTACHEDTAGKEYNOTFOUND = "ResourceNotFound.AttachedTagKeyNotFound"
func (c *Client) UpdateResourceTagValueWithContext(ctx context.Context, request *UpdateResourceTagValueRequest) (response *UpdateResourceTagValueResponse, err error) {
    if request == nil {
        request = NewUpdateResourceTagValueRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateResourceTagValue require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateResourceTagValueResponse()
    err = c.Send(request, response)
    return
}

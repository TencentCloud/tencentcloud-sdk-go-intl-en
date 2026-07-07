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

package v20260322

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2026-03-22"

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


func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
    request = &CreateApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateApiKey")
    
    
    return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
    response = &CreateApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApiKey
// Create an API key.
//
// 
//
// Create a new API key. Upon successful creation, return the API Key ID. Specify the platform kind, binding method, and initial state.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    return c.CreateApiKeyWithContext(context.Background(), request)
}

// CreateApiKey
// Create an API key.
//
// 
//
// Create a new API key. Upon successful creation, return the API Key ID. Specify the platform kind, binding method, and initial state.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateApiKeyWithContext(ctx context.Context, request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
    if request == nil {
        request = NewCreateApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlossaryRequest() (request *CreateGlossaryRequest) {
    request = &CreateGlossaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateGlossary")
    
    
    return
}

func NewCreateGlossaryResponse() (response *CreateGlossaryResponse) {
    response = &CreateGlossaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlossary
// Create a Termbase.
//
// 
//
// Create a new Termbase in this application for custom definition source to target language terminology mapping. Return the Termbase ID upon success, which can be used to carry out other management operations on terminology entries.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossary(request *CreateGlossaryRequest) (response *CreateGlossaryResponse, err error) {
    return c.CreateGlossaryWithContext(context.Background(), request)
}

// CreateGlossary
// Create a Termbase.
//
// 
//
// Create a new Termbase in this application for custom definition source to target language terminology mapping. Return the Termbase ID upon success, which can be used to carry out other management operations on terminology entries.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossaryWithContext(ctx context.Context, request *CreateGlossaryRequest) (response *CreateGlossaryResponse, err error) {
    if request == nil {
        request = NewCreateGlossaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateGlossary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlossary require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlossaryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGlossaryEntriesRequest() (request *CreateGlossaryEntriesRequest) {
    request = &CreateGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateGlossaryEntries")
    
    
    return
}

func NewCreateGlossaryEntriesResponse() (response *CreateGlossaryEntriesResponse) {
    response = &CreateGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateGlossaryEntries
// Create terminology entries in batches.
//
// 
//
// Create terminology entries in batches under the designated Termbase. You can create up to 100 entries at a time.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossaryEntries(request *CreateGlossaryEntriesRequest) (response *CreateGlossaryEntriesResponse, err error) {
    return c.CreateGlossaryEntriesWithContext(context.Background(), request)
}

// CreateGlossaryEntries
// Create terminology entries in batches.
//
// 
//
// Create terminology entries in batches under the designated Termbase. You can create up to 100 entries at a time.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateGlossaryEntriesWithContext(ctx context.Context, request *CreateGlossaryEntriesRequest) (response *CreateGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewCreateGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
    request = &DeleteApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteApiKey")
    
    
    return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
    response = &DeleteApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApiKey
// This API is used to delete specified api keys and clean up associated model binding relationships.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    return c.DeleteApiKeyWithContext(context.Background(), request)
}

// DeleteApiKey
// This API is used to delete specified api keys and clean up associated model binding relationships.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteApiKeyWithContext(ctx context.Context, request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlossaryRequest() (request *DeleteGlossaryRequest) {
    request = &DeleteGlossaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteGlossary")
    
    
    return
}

func NewDeleteGlossaryResponse() (response *DeleteGlossaryResponse) {
    response = &DeleteGlossaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlossary
// Delete a termbase.
//
// 
//
// This API is used to delete specified Termbase and ALL terminology entries under it. The deletion is idempotent and returns a successful result for non-existing Termbase. After calling the API, if the corresponding Termbase cannot be found via DescribeGlossaries, it indicates successful deletion.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossary(request *DeleteGlossaryRequest) (response *DeleteGlossaryResponse, err error) {
    return c.DeleteGlossaryWithContext(context.Background(), request)
}

// DeleteGlossary
// Delete a termbase.
//
// 
//
// This API is used to delete specified Termbase and ALL terminology entries under it. The deletion is idempotent and returns a successful result for non-existing Termbase. After calling the API, if the corresponding Termbase cannot be found via DescribeGlossaries, it indicates successful deletion.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossaryWithContext(ctx context.Context, request *DeleteGlossaryRequest) (response *DeleteGlossaryResponse, err error) {
    if request == nil {
        request = NewDeleteGlossaryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteGlossary")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlossary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlossaryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGlossaryEntriesRequest() (request *DeleteGlossaryEntriesRequest) {
    request = &DeleteGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteGlossaryEntries")
    
    
    return
}

func NewDeleteGlossaryEntriesResponse() (response *DeleteGlossaryEntriesResponse) {
    response = &DeleteGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteGlossaryEntries
// Delete terminology entries in batches.
//
// 
//
// Delete terminology entries in batches under the specified Termbase. You can delete up to 200 entries at a time. If the Termbase is nonexistent or NOT_IN this application, it returns a ResourceNotFound error.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossaryEntries(request *DeleteGlossaryEntriesRequest) (response *DeleteGlossaryEntriesResponse, err error) {
    return c.DeleteGlossaryEntriesWithContext(context.Background(), request)
}

// DeleteGlossaryEntries
// Delete terminology entries in batches.
//
// 
//
// Delete terminology entries in batches under the specified Termbase. You can delete up to 200 entries at a time. If the Termbase is nonexistent or NOT_IN this application, it returns a ResourceNotFound error.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteGlossaryEntriesWithContext(ctx context.Context, request *DeleteGlossaryEntriesRequest) (response *DeleteGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewDeleteGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlossariesRequest() (request *DescribeGlossariesRequest) {
    request = &DescribeGlossariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeGlossaries")
    
    
    return
}

func NewDescribeGlossariesResponse() (response *DescribeGlossariesResponse) {
    response = &DescribeGlossariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlossaries
// Query the terminology repository list.
//
// 
//
// Query the Termbase list under this application. Support paginate, filter, and sort.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossaries(request *DescribeGlossariesRequest) (response *DescribeGlossariesResponse, err error) {
    return c.DescribeGlossariesWithContext(context.Background(), request)
}

// DescribeGlossaries
// Query the terminology repository list.
//
// 
//
// Query the Termbase list under this application. Support paginate, filter, and sort.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossariesWithContext(ctx context.Context, request *DescribeGlossariesRequest) (response *DescribeGlossariesResponse, err error) {
    if request == nil {
        request = NewDescribeGlossariesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeGlossaries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlossaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlossariesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGlossaryEntriesRequest() (request *DescribeGlossaryEntriesRequest) {
    request = &DescribeGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeGlossaryEntries")
    
    
    return
}

func NewDescribeGlossaryEntriesResponse() (response *DescribeGlossaryEntriesResponse) {
    response = &DescribeGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeGlossaryEntries
// Query the terminology entry list.
//
// 
//
// Query specified entries in a Termbase. Support pagination.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossaryEntries(request *DescribeGlossaryEntriesRequest) (response *DescribeGlossaryEntriesResponse, err error) {
    return c.DescribeGlossaryEntriesWithContext(context.Background(), request)
}

// DescribeGlossaryEntries
// Query the terminology entry list.
//
// 
//
// Query specified entries in a Termbase. Support pagination.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeGlossaryEntriesWithContext(ctx context.Context, request *DescribeGlossaryEntriesRequest) (response *DescribeGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewDescribeGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiKeyInfoRequest() (request *ModifyApiKeyInfoRequest) {
    request = &ModifyApiKeyInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyApiKeyInfo")
    
    
    return
}

func NewModifyApiKeyInfoResponse() (response *ModifyApiKeyInfoResponse) {
    response = &ModifyApiKeyInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiKeyInfo
// Refresh API key information.
//
// 
//
// This API is used to update the remark information, IP allowlist and Token quota of an API key (recommended to use QuotaDesired parameter for quota modification). Passing no optional parameters means no modification.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyInfo(request *ModifyApiKeyInfoRequest) (response *ModifyApiKeyInfoResponse, err error) {
    return c.ModifyApiKeyInfoWithContext(context.Background(), request)
}

// ModifyApiKeyInfo
// Refresh API key information.
//
// 
//
// This API is used to update the remark information, IP allowlist and Token quota of an API key (recommended to use QuotaDesired parameter for quota modification). Passing no optional parameters means no modification.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyInfoWithContext(ctx context.Context, request *ModifyApiKeyInfoRequest) (response *ModifyApiKeyInfoResponse, err error) {
    if request == nil {
        request = NewModifyApiKeyInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyApiKeyInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiKeyInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiKeyInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApiKeyStatusRequest() (request *ModifyApiKeyStatusRequest) {
    request = &ModifyApiKeyStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyApiKeyStatus")
    
    
    return
}

func NewModifyApiKeyStatusResponse() (response *ModifyApiKeyStatusResponse) {
    response = &ModifyApiKeyStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApiKeyStatus
// This API is used to enable or disable the status of an api key.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyStatus(request *ModifyApiKeyStatusRequest) (response *ModifyApiKeyStatusResponse, err error) {
    return c.ModifyApiKeyStatusWithContext(context.Background(), request)
}

// ModifyApiKeyStatus
// This API is used to enable or disable the status of an api key.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyApiKeyStatusWithContext(ctx context.Context, request *ModifyApiKeyStatusRequest) (response *ModifyApiKeyStatusResponse, err error) {
    if request == nil {
        request = NewModifyApiKeyStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyApiKeyStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApiKeyStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApiKeyStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGlossaryEntriesRequest() (request *ModifyGlossaryEntriesRequest) {
    request = &ModifyGlossaryEntriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyGlossaryEntries")
    
    
    return
}

func NewModifyGlossaryEntriesResponse() (response *ModifyGlossaryEntriesResponse) {
    response = &ModifyGlossaryEntriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyGlossaryEntries
// Batch modify terminology entries.
//
// 
//
// This API is used to batch modify terminology entries in a designated Termbase. You can modify up to 200 entries at a time.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyGlossaryEntries(request *ModifyGlossaryEntriesRequest) (response *ModifyGlossaryEntriesResponse, err error) {
    return c.ModifyGlossaryEntriesWithContext(context.Background(), request)
}

// ModifyGlossaryEntries
// Batch modify terminology entries.
//
// 
//
// This API is used to batch modify terminology entries in a designated Termbase. You can modify up to 200 entries at a time.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_GLOSSARYNOTFOUND = "ResourceNotFound.GlossaryNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyGlossaryEntriesWithContext(ctx context.Context, request *ModifyGlossaryEntriesRequest) (response *ModifyGlossaryEntriesResponse, err error) {
    if request == nil {
        request = NewModifyGlossaryEntriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyGlossaryEntries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGlossaryEntries require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGlossaryEntriesResponse()
    err = c.Send(request, response)
    return
}

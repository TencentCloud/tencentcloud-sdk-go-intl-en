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

func NewCreateTokenPlanApiKeysRequest() (request *CreateTokenPlanApiKeysRequest) {
    request = &CreateTokenPlanApiKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateTokenPlanApiKeys")
    
    
    return
}

func NewCreateTokenPlanApiKeysResponse() (response *CreateTokenPlanApiKeysResponse) {
    response = &CreateTokenPlanApiKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTokenPlanApiKeys
// Batch create TokenPlan API Keys.
//
// 
//
// Import a name prefix and quantity to automatically generate names in the `{Api Key Name}-{serial number}` format (for example, aaa-1, aaa-2). Duplicate names are allowed. Partial success is supported for up to 100 entries.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanApiKeys(request *CreateTokenPlanApiKeysRequest) (response *CreateTokenPlanApiKeysResponse, err error) {
    return c.CreateTokenPlanApiKeysWithContext(context.Background(), request)
}

// CreateTokenPlanApiKeys
// Batch create TokenPlan API Keys.
//
// 
//
// Import a name prefix and quantity to automatically generate names in the `{Api Key Name}-{serial number}` format (for example, aaa-1, aaa-2). Duplicate names are allowed. Partial success is supported for up to 100 entries.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanApiKeysWithContext(ctx context.Context, request *CreateTokenPlanApiKeysRequest) (response *CreateTokenPlanApiKeysResponse, err error) {
    if request == nil {
        request = NewCreateTokenPlanApiKeysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateTokenPlanApiKeys")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTokenPlanApiKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTokenPlanApiKeysResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTokenPlanTeamOrderAndBuyRequest() (request *CreateTokenPlanTeamOrderAndBuyRequest) {
    request = &CreateTokenPlanTeamOrderAndBuyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "CreateTokenPlanTeamOrderAndBuy")
    
    
    return
}

func NewCreateTokenPlanTeamOrderAndBuyResponse() (response *CreateTokenPlanTeamOrderAndBuyResponse) {
    response = &CreateTokenPlanTeamOrderAndBuyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateTokenPlanTeamOrderAndBuy
// Purchase a package (This API is also used to reactivate and renew expired packages. The teamId of the expired package is required. After the renewal is successful, the total cycle count of the package will include historical cycles. The actual effective cycle of the package is determined by the effective time and expiration time.)
//
// 
//
// Initiate an order for a TokenPlan package and complete payment. Return the order ID and associated sub-orders and resource information upon success.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanTeamOrderAndBuy(request *CreateTokenPlanTeamOrderAndBuyRequest) (response *CreateTokenPlanTeamOrderAndBuyResponse, err error) {
    return c.CreateTokenPlanTeamOrderAndBuyWithContext(context.Background(), request)
}

// CreateTokenPlanTeamOrderAndBuy
// Purchase a package (This API is also used to reactivate and renew expired packages. The teamId of the expired package is required. After the renewal is successful, the total cycle count of the package will include historical cycles. The actual effective cycle of the package is determined by the effective time and expiration time.)
//
// 
//
// Initiate an order for a TokenPlan package and complete payment. Return the order ID and associated sub-orders and resource information upon success.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) CreateTokenPlanTeamOrderAndBuyWithContext(ctx context.Context, request *CreateTokenPlanTeamOrderAndBuyRequest) (response *CreateTokenPlanTeamOrderAndBuyResponse, err error) {
    if request == nil {
        request = NewCreateTokenPlanTeamOrderAndBuyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "CreateTokenPlanTeamOrderAndBuy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTokenPlanTeamOrderAndBuy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTokenPlanTeamOrderAndBuyResponse()
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

func NewDeleteTokenPlanApiKeyRequest() (request *DeleteTokenPlanApiKeyRequest) {
    request = &DeleteTokenPlanApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DeleteTokenPlanApiKey")
    
    
    return
}

func NewDeleteTokenPlanApiKeyResponse() (response *DeleteTokenPlanApiKeyResponse) {
    response = &DeleteTokenPlanApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteTokenPlanApiKey
// Delete the Token Plan API key.
//
// 
//
// Simultaneously delete the limit center sub-limit package and notify the Notification Gateway to purge cache.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteTokenPlanApiKey(request *DeleteTokenPlanApiKeyRequest) (response *DeleteTokenPlanApiKeyResponse, err error) {
    return c.DeleteTokenPlanApiKeyWithContext(context.Background(), request)
}

// DeleteTokenPlanApiKey
// Delete the Token Plan API key.
//
// 
//
// Simultaneously delete the limit center sub-limit package and notify the Notification Gateway to purge cache.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DeleteTokenPlanApiKeyWithContext(ctx context.Context, request *DeleteTokenPlanApiKeyRequest) (response *DeleteTokenPlanApiKeyResponse, err error) {
    if request == nil {
        request = NewDeleteTokenPlanApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DeleteTokenPlanApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTokenPlanApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTokenPlanApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyRequest() (request *DescribeApiKeyRequest) {
    request = &DescribeApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeApiKey")
    
    
    return
}

func NewDescribeApiKeyResponse() (response *DescribeApiKeyResponse) {
    response = &DescribeApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKey
// This API is used to query API Key details based on API Key ID or key value, and return the plaintext key. At least one of ApiKeyId and ApiKey must be input, with priority given to ApiKeyId.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKey(request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    return c.DescribeApiKeyWithContext(context.Background(), request)
}

// DescribeApiKey
// This API is used to query API Key details based on API Key ID or key value, and return the plaintext key. At least one of ApiKeyId and ApiKey must be input, with priority given to ApiKeyId.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKeyWithContext(ctx context.Context, request *DescribeApiKeyRequest) (response *DescribeApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApiKeyListRequest() (request *DescribeApiKeyListRequest) {
    request = &DescribeApiKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeApiKeyList")
    
    
    return
}

func NewDescribeApiKeyListResponse() (response *DescribeApiKeyListResponse) {
    response = &DescribeApiKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApiKeyList
// Query API key list.
//
// 
//
// Query the API key list of the current user with key values in masking display. Support pagination, filtering, and sorting.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKeyList(request *DescribeApiKeyListRequest) (response *DescribeApiKeyListResponse, err error) {
    return c.DescribeApiKeyListWithContext(context.Background(), request)
}

// DescribeApiKeyList
// Query API key list.
//
// 
//
// Query the API key list of the current user with key values in masking display. Support pagination, filtering, and sorting.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeApiKeyListWithContext(ctx context.Context, request *DescribeApiKeyListRequest) (response *DescribeApiKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeApiKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeApiKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApiKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApiKeyListResponse()
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

func NewDescribeTokenPlanRequest() (request *DescribeTokenPlanRequest) {
    request = &DescribeTokenPlanRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlan")
    
    
    return
}

func NewDescribeTokenPlanResponse() (response *DescribeTokenPlanResponse) {
    response = &DescribeTokenPlanResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlan
// Query the TokenPlan package details.
//
// 
//
// Return the package basic info and the remaining quota of the package.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlan(request *DescribeTokenPlanRequest) (response *DescribeTokenPlanResponse, err error) {
    return c.DescribeTokenPlanWithContext(context.Background(), request)
}

// DescribeTokenPlan
// Query the TokenPlan package details.
//
// 
//
// Return the package basic info and the remaining quota of the package.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanWithContext(ctx context.Context, request *DescribeTokenPlanRequest) (response *DescribeTokenPlanResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlan")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlan require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeyRequest() (request *DescribeTokenPlanApiKeyRequest) {
    request = &DescribeTokenPlanApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKey")
    
    
    return
}

func NewDescribeTokenPlanApiKeyResponse() (response *DescribeTokenPlanApiKeyResponse) {
    response = &DescribeTokenPlanApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKey
// Query TokenPlan APIKey details.
//
// 
//
// Return the complete APIKey information (including the plaintext key) and the remaining quota of the sub-quota package.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKey(request *DescribeTokenPlanApiKeyRequest) (response *DescribeTokenPlanApiKeyResponse, err error) {
    return c.DescribeTokenPlanApiKeyWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKey
// Query TokenPlan APIKey details.
//
// 
//
// Return the complete APIKey information (including the plaintext key) and the remaining quota of the sub-quota package.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyWithContext(ctx context.Context, request *DescribeTokenPlanApiKeyRequest) (response *DescribeTokenPlanApiKeyResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeyListRequest() (request *DescribeTokenPlanApiKeyListRequest) {
    request = &DescribeTokenPlanApiKeyListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKeyList")
    
    
    return
}

func NewDescribeTokenPlanApiKeyListResponse() (response *DescribeTokenPlanApiKeyListResponse) {
    response = &DescribeTokenPlanApiKeyListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKeyList
// Query the list of Token Plan API keys.
//
// 
//
// Returns the API key list under a specified package. Keys are masked. Root accounts can view all keys, while sub-accounts can only view keys created by themselves.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyList(request *DescribeTokenPlanApiKeyListRequest) (response *DescribeTokenPlanApiKeyListResponse, err error) {
    return c.DescribeTokenPlanApiKeyListWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKeyList
// Query the list of Token Plan API keys.
//
// 
//
// Returns the API key list under a specified package. Keys are masked. Root accounts can view all keys, while sub-accounts can only view keys created by themselves.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyListWithContext(ctx context.Context, request *DescribeTokenPlanApiKeyListRequest) (response *DescribeTokenPlanApiKeyListResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeyListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKeyList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKeyList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeyListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeySecretRequest() (request *DescribeTokenPlanApiKeySecretRequest) {
    request = &DescribeTokenPlanApiKeySecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKeySecret")
    
    
    return
}

func NewDescribeTokenPlanApiKeySecretResponse() (response *DescribeTokenPlanApiKeySecretResponse) {
    response = &DescribeTokenPlanApiKeySecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKeySecret
// Query the TokenPlan APIKey (plaintext).
//
// 
//
// Return the plaintext key value of the designated APIKey. Keep it safe.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeySecret(request *DescribeTokenPlanApiKeySecretRequest) (response *DescribeTokenPlanApiKeySecretResponse, err error) {
    return c.DescribeTokenPlanApiKeySecretWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKeySecret
// Query the TokenPlan APIKey (plaintext).
//
// 
//
// Return the plaintext key value of the designated APIKey. Keep it safe.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeySecretWithContext(ctx context.Context, request *DescribeTokenPlanApiKeySecretRequest) (response *DescribeTokenPlanApiKeySecretResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeySecretRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKeySecret")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKeySecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeySecretResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanApiKeyUsageDetailRequest() (request *DescribeTokenPlanApiKeyUsageDetailRequest) {
    request = &DescribeTokenPlanApiKeyUsageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanApiKeyUsageDetail")
    
    
    return
}

func NewDescribeTokenPlanApiKeyUsageDetailResponse() (response *DescribeTokenPlanApiKeyUsageDetailResponse) {
    response = &DescribeTokenPlanApiKeyUsageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanApiKeyUsageDetail
// Query the Token Plan APIKey call detail.
//
// 
//
// This API is used to query call details under a package from CLS log service, filter by team_id, and support cursor-based pagination.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyUsageDetail(request *DescribeTokenPlanApiKeyUsageDetailRequest) (response *DescribeTokenPlanApiKeyUsageDetailResponse, err error) {
    return c.DescribeTokenPlanApiKeyUsageDetailWithContext(context.Background(), request)
}

// DescribeTokenPlanApiKeyUsageDetail
// Query the Token Plan APIKey call detail.
//
// 
//
// This API is used to query call details under a package from CLS log service, filter by team_id, and support cursor-based pagination.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeTokenPlanApiKeyUsageDetailWithContext(ctx context.Context, request *DescribeTokenPlanApiKeyUsageDetailRequest) (response *DescribeTokenPlanApiKeyUsageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanApiKeyUsageDetailRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanApiKeyUsageDetail")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanApiKeyUsageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanApiKeyUsageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenPlanListRequest() (request *DescribeTokenPlanListRequest) {
    request = &DescribeTokenPlanListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeTokenPlanList")
    
    
    return
}

func NewDescribeTokenPlanListResponse() (response *DescribeTokenPlanListResponse) {
    response = &DescribeTokenPlanListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenPlanList
// Query the list of Token Plan package options.
//
// 
//
// Supports pagination, filtering, and sorting. Root accounts can view all packages, while sub-accounts can only view packages created by themselves. Returned results include the main limit package details associated with each package in the limit center.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeTokenPlanList(request *DescribeTokenPlanListRequest) (response *DescribeTokenPlanListResponse, err error) {
    return c.DescribeTokenPlanListWithContext(context.Background(), request)
}

// DescribeTokenPlanList
// Query the list of Token Plan package options.
//
// 
//
// Supports pagination, filtering, and sorting. Root accounts can view all packages, while sub-accounts can only view packages created by themselves. Returned results include the main limit package details associated with each package in the limit center.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeTokenPlanListWithContext(ctx context.Context, request *DescribeTokenPlanListRequest) (response *DescribeTokenPlanListResponse, err error) {
    if request == nil {
        request = NewDescribeTokenPlanListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeTokenPlanList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenPlanList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenPlanListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageRankListRequest() (request *DescribeUsageRankListRequest) {
    request = &DescribeUsageRankListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "DescribeUsageRankList")
    
    
    return
}

func NewDescribeUsageRankListResponse() (response *DescribeUsageRankListResponse) {
    response = &DescribeUsageRankListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageRankList
// Query the usage ranking list.
//
// 
//
// Metric family (MetricType)
//
// - `tokens` (default): Token usage statistics. Supports Dimension = apikey / endpoint / model.
//
// Metrics returned: TotalToken (total) / InputTotalToken (input) / OutputTotalToken (output) / CacheTotalToken (read cache).
//
// - `search`: [To be launched] Online search usage statistics. Supports Dimension = apikey / endpoint / model.
//
// Returns metrics: SearchRequestCount (search request count)/SearchCount (search engine call count).
//
// 
//
// content
//
// -The MetricType field is used to switch metric families. The response echoes back MetricType and MetricKeys.
//
// -TotalStats: The aggregated value of all objects over the entire time window.
//
// -PageStats: The aggregated value of objects on the current page.
//
// - TopList: A list of objects sorted by MetricKeys[0] in descending order, including the aggregated value over the entire period and point-in-time curves.
//
// error code that may be returned:
//  INTERNALERROR_BARADERROR = "InternalError.BaradError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERIODEXCEEDSSPAN = "InvalidParameter.PeriodExceedsSpan"
//  INVALIDPARAMETER_PERIODTOOFINEFORDATA = "InvalidParameter.PeriodTooFineForData"
//  INVALIDPARAMETER_TOOMANYOBJECTS = "InvalidParameter.TooManyObjects"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeUsageRankList(request *DescribeUsageRankListRequest) (response *DescribeUsageRankListResponse, err error) {
    return c.DescribeUsageRankListWithContext(context.Background(), request)
}

// DescribeUsageRankList
// Query the usage ranking list.
//
// 
//
// Metric family (MetricType)
//
// - `tokens` (default): Token usage statistics. Supports Dimension = apikey / endpoint / model.
//
// Metrics returned: TotalToken (total) / InputTotalToken (input) / OutputTotalToken (output) / CacheTotalToken (read cache).
//
// - `search`: [To be launched] Online search usage statistics. Supports Dimension = apikey / endpoint / model.
//
// Returns metrics: SearchRequestCount (search request count)/SearchCount (search engine call count).
//
// 
//
// content
//
// -The MetricType field is used to switch metric families. The response echoes back MetricType and MetricKeys.
//
// -TotalStats: The aggregated value of all objects over the entire time window.
//
// -PageStats: The aggregated value of objects on the current page.
//
// - TopList: A list of objects sorted by MetricKeys[0] in descending order, including the aggregated value over the entire period and point-in-time curves.
//
// error code that may be returned:
//  INTERNALERROR_BARADERROR = "InternalError.BaradError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_PERIODEXCEEDSSPAN = "InvalidParameter.PeriodExceedsSpan"
//  INVALIDPARAMETER_PERIODTOOFINEFORDATA = "InvalidParameter.PeriodTooFineForData"
//  INVALIDPARAMETER_TOOMANYOBJECTS = "InvalidParameter.TooManyObjects"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) DescribeUsageRankListWithContext(ctx context.Context, request *DescribeUsageRankListRequest) (response *DescribeUsageRankListResponse, err error) {
    if request == nil {
        request = NewDescribeUsageRankListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "DescribeUsageRankList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageRankList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageRankListResponse()
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

func NewModifyTokenPlanApiKeyRequest() (request *ModifyTokenPlanApiKeyRequest) {
    request = &ModifyTokenPlanApiKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyTokenPlanApiKey")
    
    
    return
}

func NewModifyTokenPlanApiKeyResponse() (response *ModifyTokenPlanApiKeyResponse) {
    response = &ModifyTokenPlanApiKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTokenPlanApiKey
// Modify the Token Plan APIKey configuration (field that the gateway focuses on).
//
// 
//
// After modification, automatically notify the gateway to update the cache and sync the limit center.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKey(request *ModifyTokenPlanApiKeyRequest) (response *ModifyTokenPlanApiKeyResponse, err error) {
    return c.ModifyTokenPlanApiKeyWithContext(context.Background(), request)
}

// ModifyTokenPlanApiKey
// Modify the Token Plan APIKey configuration (field that the gateway focuses on).
//
// 
//
// After modification, automatically notify the gateway to update the cache and sync the limit center.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKeyWithContext(ctx context.Context, request *ModifyTokenPlanApiKeyRequest) (response *ModifyTokenPlanApiKeyResponse, err error) {
    if request == nil {
        request = NewModifyTokenPlanApiKeyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyTokenPlanApiKey")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTokenPlanApiKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTokenPlanApiKeyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTokenPlanApiKeySecretRequest() (request *ModifyTokenPlanApiKeySecretRequest) {
    request = &ModifyTokenPlanApiKeySecretRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "ModifyTokenPlanApiKeySecret")
    
    
    return
}

func NewModifyTokenPlanApiKeySecretResponse() (response *ModifyTokenPlanApiKeySecretResponse) {
    response = &ModifyTokenPlanApiKeySecretResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyTokenPlanApiKeySecret
// Reset the TokenPlan API Key.
//
// 
//
// Regenerate the key value. The key version increments and the old key expires immediately. The API Key ID remains unchanged. After resetting, the new key can be queried through DescribeTokenPlanApiKeySecret.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKeySecret(request *ModifyTokenPlanApiKeySecretRequest) (response *ModifyTokenPlanApiKeySecretResponse, err error) {
    return c.ModifyTokenPlanApiKeySecretWithContext(context.Background(), request)
}

// ModifyTokenPlanApiKeySecret
// Reset the TokenPlan API Key.
//
// 
//
// Regenerate the key value. The key version increments and the old key expires immediately. The API Key ID remains unchanged. After resetting, the new key can be queried through DescribeTokenPlanApiKeySecret.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) ModifyTokenPlanApiKeySecretWithContext(ctx context.Context, request *ModifyTokenPlanApiKeySecretRequest) (response *ModifyTokenPlanApiKeySecretResponse, err error) {
    if request == nil {
        request = NewModifyTokenPlanApiKeySecretRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "ModifyTokenPlanApiKeySecret")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTokenPlanApiKeySecret require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTokenPlanApiKeySecretResponse()
    err = c.Send(request, response)
    return
}

func NewRenewTokenPlanTeamOrderRequest() (request *RenewTokenPlanTeamOrderRequest) {
    request = &RenewTokenPlanTeamOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "RenewTokenPlanTeamOrder")
    
    
    return
}

func NewRenewTokenPlanTeamOrderResponse() (response *RenewTokenPlanTeamOrderResponse) {
    response = &RenewTokenPlanTeamOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewTokenPlanTeamOrder
// Renew a package.
//
// 
//
// Initiate a renewal order for an existing Token Plan package and complete payment. Return the order ID and associated sub-orders and resource information upon success.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RenewTokenPlanTeamOrder(request *RenewTokenPlanTeamOrderRequest) (response *RenewTokenPlanTeamOrderResponse, err error) {
    return c.RenewTokenPlanTeamOrderWithContext(context.Background(), request)
}

// RenewTokenPlanTeamOrder
// Renew a package.
//
// 
//
// Initiate a renewal order for an existing Token Plan package and complete payment. Return the order ID and associated sub-orders and resource information upon success.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) RenewTokenPlanTeamOrderWithContext(ctx context.Context, request *RenewTokenPlanTeamOrderRequest) (response *RenewTokenPlanTeamOrderResponse, err error) {
    if request == nil {
        request = NewRenewTokenPlanTeamOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "RenewTokenPlanTeamOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewTokenPlanTeamOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewTokenPlanTeamOrderResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeTokenPlanTeamOrderRequest() (request *UpgradeTokenPlanTeamOrderRequest) {
    request = &UpgradeTokenPlanTeamOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tokenhub", APIVersion, "UpgradeTokenPlanTeamOrder")
    
    
    return
}

func NewUpgradeTokenPlanTeamOrderResponse() (response *UpgradeTokenPlanTeamOrderResponse) {
    response = &UpgradeTokenPlanTeamOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpgradeTokenPlanTeamOrder
// Upgrade the package.
//
// 
//
// Initiate an upgrade order for an existing Token Plan package and complete payment to expand point or token limits. Return the order ID and associated sub-orders and resource information upon success. The new limit must be greater than the current limit.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpgradeTokenPlanTeamOrder(request *UpgradeTokenPlanTeamOrderRequest) (response *UpgradeTokenPlanTeamOrderResponse, err error) {
    return c.UpgradeTokenPlanTeamOrderWithContext(context.Background(), request)
}

// UpgradeTokenPlanTeamOrder
// Upgrade the package.
//
// 
//
// Initiate an upgrade order for an existing Token Plan package and complete payment to expand point or token limits. Return the order ID and associated sub-orders and resource information upon success. The new limit must be greater than the current limit.
//
// error code that may be returned:
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  MISSINGPARAMETER_MISSINGPARAMETER = "MissingParameter.MissingParameter"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATION = "UnauthorizedOperation.UnauthorizedOperation"
func (c *Client) UpgradeTokenPlanTeamOrderWithContext(ctx context.Context, request *UpgradeTokenPlanTeamOrderRequest) (response *UpgradeTokenPlanTeamOrderResponse, err error) {
    if request == nil {
        request = NewUpgradeTokenPlanTeamOrderRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tokenhub", APIVersion, "UpgradeTokenPlanTeamOrder")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeTokenPlanTeamOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeTokenPlanTeamOrderResponse()
    err = c.Send(request, response)
    return
}

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

package v20180416

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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


func NewDescribeFabricBlockRequest() (request *DescribeFabricBlockRequest) {
    request = &DescribeFabricBlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "DescribeFabricBlock")
    
    
    return
}

func NewDescribeFabricBlockResponse() (response *DescribeFabricBlockResponse) {
    response = &DescribeFabricBlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFabricBlock
// This API is used to retrieve the detailed information of a block in Fabric.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICBLOCKNOEXIST = "FailedOperation.FabricBlockNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricBlock(request *DescribeFabricBlockRequest) (response *DescribeFabricBlockResponse, err error) {
    return c.DescribeFabricBlockWithContext(context.Background(), request)
}

// DescribeFabricBlock
// This API is used to retrieve the detailed information of a block in Fabric.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICBLOCKNOEXIST = "FailedOperation.FabricBlockNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricBlockWithContext(ctx context.Context, request *DescribeFabricBlockRequest) (response *DescribeFabricBlockResponse, err error) {
    if request == nil {
        request = NewDescribeFabricBlockRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFabricBlock require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFabricBlockResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFabricTransactionRequest() (request *DescribeFabricTransactionRequest) {
    request = &DescribeFabricTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "DescribeFabricTransaction")
    
    
    return
}

func NewDescribeFabricTransactionResponse() (response *DescribeFabricTransactionResponse) {
    response = &DescribeFabricTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFabricTransaction
// This API is used to obtain detailed information of Fabric transactions.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICTRANSACTIONNOEXIST = "FailedOperation.FabricTransactionNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricTransaction(request *DescribeFabricTransactionRequest) (response *DescribeFabricTransactionResponse, err error) {
    return c.DescribeFabricTransactionWithContext(context.Background(), request)
}

// DescribeFabricTransaction
// This API is used to obtain detailed information of Fabric transactions.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICTRANSACTIONNOEXIST = "FailedOperation.FabricTransactionNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) DescribeFabricTransactionWithContext(ctx context.Context, request *DescribeFabricTransactionRequest) (response *DescribeFabricTransactionResponse, err error) {
    if request == nil {
        request = NewDescribeFabricTransactionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFabricTransaction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFabricTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeFabricChaincodeRequest() (request *InvokeFabricChaincodeRequest) {
    request = &InvokeFabricChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "InvokeFabricChaincode")
    
    
    return
}

func NewInvokeFabricChaincodeResponse() (response *InvokeFabricChaincodeResponse) {
    response = &InvokeFabricChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// InvokeFabricChaincode
// This API is used to call a Fabric user contract to execute a transaction.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICCHAINCODEINVOKEFAILED = "FailedOperation.FabricChaincodeInvokeFailed"
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) InvokeFabricChaincode(request *InvokeFabricChaincodeRequest) (response *InvokeFabricChaincodeResponse, err error) {
    return c.InvokeFabricChaincodeWithContext(context.Background(), request)
}

// InvokeFabricChaincode
// This API is used to call a Fabric user contract to execute a transaction.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICCHAINCODEINVOKEFAILED = "FailedOperation.FabricChaincodeInvokeFailed"
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) InvokeFabricChaincodeWithContext(ctx context.Context, request *InvokeFabricChaincodeRequest) (response *InvokeFabricChaincodeResponse, err error) {
    if request == nil {
        request = NewInvokeFabricChaincodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InvokeFabricChaincode require credential")
    }

    request.SetContext(ctx)
    
    response = NewInvokeFabricChaincodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFabricChaincodeRequest() (request *QueryFabricChaincodeRequest) {
    request = &QueryFabricChaincodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbaas", APIVersion, "QueryFabricChaincode")
    
    
    return
}

func NewQueryFabricChaincodeResponse() (response *QueryFabricChaincodeResponse) {
    response = &QueryFabricChaincodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryFabricChaincode
// This API is used to make a user contract call in Fabric for querying.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_FABRICCHAINCODEQUERYFAILED = "FailedOperation.FabricChaincodeQueryFailed"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) QueryFabricChaincode(request *QueryFabricChaincodeRequest) (response *QueryFabricChaincodeResponse, err error) {
    return c.QueryFabricChaincodeWithContext(context.Background(), request)
}

// QueryFabricChaincode
// This API is used to make a user contract call in Fabric for querying.
//
// error code that may be returned:
//  FAILEDOPERATION_FABRICCHAINCODENOEXIST = "FailedOperation.FabricChaincodeNoExist"
//  FAILEDOPERATION_FABRICCHAINCODEQUERYFAILED = "FailedOperation.FabricChaincodeQueryFailed"
//  FAILEDOPERATION_INVALIDAUTH = "FailedOperation.InvalidAuth"
//  INTERNALERROR_SERVERERROR = "InternalError.ServerError"
func (c *Client) QueryFabricChaincodeWithContext(ctx context.Context, request *QueryFabricChaincodeRequest) (response *QueryFabricChaincodeResponse, err error) {
    if request == nil {
        request = NewQueryFabricChaincodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryFabricChaincode require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryFabricChaincodeResponse()
    err = c.Send(request, response)
    return
}

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

package v20231130

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2023-11-30"

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


func NewCreateCorpRequest() (request *CreateCorpRequest) {
    request = &CreateCorpRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateCorp")
    
    
    return
}

func NewCreateCorpResponse() (response *CreateCorpResponse) {
    response = &CreateCorpResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCorp
// This API is used to create enterprises.
func (c *Client) CreateCorp(request *CreateCorpRequest) (response *CreateCorpResponse, err error) {
    return c.CreateCorpWithContext(context.Background(), request)
}

// CreateCorp
// This API is used to create enterprises.
func (c *Client) CreateCorpWithContext(ctx context.Context, request *CreateCorpRequest) (response *CreateCorpResponse, err error) {
    if request == nil {
        request = NewCreateCorpRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCorp require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCorpResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReconstructDocumentFlowRequest() (request *CreateReconstructDocumentFlowRequest) {
    request = &CreateReconstructDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateReconstructDocumentFlow")
    
    
    return
}

func NewCreateReconstructDocumentFlowResponse() (response *CreateReconstructDocumentFlowResponse) {
    response = &CreateReconstructDocumentFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReconstructDocumentFlow
// This API is used to initiate requests for this asynchronous API, for initiating document parsing tasks.
//
// Document parsing supports converting images or PDF files into Markdown format files, and can parse content elements including tables, formulas, images, headings, paragraphs, headers, and footers, and intelligently convert the content into reading order.
//
// During the trial period, the QPS limit for a single account is only 1. If you need to access officially, please contact our R&D team.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlow(request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    return c.CreateReconstructDocumentFlowWithContext(context.Background(), request)
}

// CreateReconstructDocumentFlow
// This API is used to initiate requests for this asynchronous API, for initiating document parsing tasks.
//
// Document parsing supports converting images or PDF files into Markdown format files, and can parse content elements including tables, formulas, images, headings, paragraphs, headers, and footers, and intelligently convert the content into reading order.
//
// During the trial period, the QPS limit for a single account is only 1. If you need to access officially, please contact our R&D team.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FILEDECODEFAILED = "FailedOperation.FileDecodeFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWFILETYPEERROR = "FailedOperation.UnKnowFileTypeError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlowWithContext(ctx context.Context, request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateReconstructDocumentFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReconstructDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReconstructDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseRequest() (request *CreateReleaseRequest) {
    request = &CreateReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "CreateRelease")
    
    
    return
}

func NewCreateReleaseResponse() (response *CreateReleaseResponse) {
    response = &CreateReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRelease
// This API is used to create a release.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    return c.CreateReleaseWithContext(context.Background(), request)
}

// CreateRelease
// This API is used to create a release.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateReleaseWithContext(ctx context.Context, request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
    if request == nil {
        request = NewCreateReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAppRequest() (request *DeleteAppRequest) {
    request = &DeleteAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteApp")
    
    
    return
}

func NewDeleteAppResponse() (response *DeleteAppResponse) {
    response = &DeleteAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteApp
// This API is used to delete applications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteApp(request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    return c.DeleteAppWithContext(context.Background(), request)
}

// DeleteApp
// This API is used to delete applications.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteAppWithContext(ctx context.Context, request *DeleteAppRequest) (response *DeleteAppResponse, err error) {
    if request == nil {
        request = NewDeleteAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAppResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDocRequest() (request *DeleteDocRequest) {
    request = &DeleteDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteDoc")
    
    
    return
}

func NewDeleteDocResponse() (response *DeleteDocResponse) {
    response = &DeleteDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDoc
// This API is used to delete documents.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDoc(request *DeleteDocRequest) (response *DeleteDocResponse, err error) {
    return c.DeleteDocWithContext(context.Background(), request)
}

// DeleteDoc
// This API is used to delete documents.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteDocWithContext(ctx context.Context, request *DeleteDocRequest) (response *DeleteDocResponse, err error) {
    if request == nil {
        request = NewDeleteDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDocResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQARequest() (request *DeleteQARequest) {
    request = &DeleteQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteQA")
    
    
    return
}

func NewDeleteQAResponse() (response *DeleteQAResponse) {
    response = &DeleteQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQA
// This API is used to delete Q&As.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQA(request *DeleteQARequest) (response *DeleteQAResponse, err error) {
    return c.DeleteQAWithContext(context.Background(), request)
}

// DeleteQA
// This API is used to delete Q&As.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQAWithContext(ctx context.Context, request *DeleteQARequest) (response *DeleteQAResponse, err error) {
    if request == nil {
        request = NewDeleteQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQAResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQACateRequest() (request *DeleteQACateRequest) {
    request = &DeleteQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteQACate")
    
    
    return
}

func NewDeleteQACateResponse() (response *DeleteQACateResponse) {
    response = &DeleteQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteQACate
// This API is used to delete categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQACate(request *DeleteQACateRequest) (response *DeleteQACateResponse, err error) {
    return c.DeleteQACateWithContext(context.Background(), request)
}

// DeleteQACate
// This API is used to delete categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteQACateWithContext(ctx context.Context, request *DeleteQACateRequest) (response *DeleteQACateResponse, err error) {
    if request == nil {
        request = NewDeleteQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteQACateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRejectedQuestionRequest() (request *DeleteRejectedQuestionRequest) {
    request = &DeleteRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DeleteRejectedQuestion")
    
    
    return
}

func NewDeleteRejectedQuestionResponse() (response *DeleteRejectedQuestionResponse) {
    response = &DeleteRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRejectedQuestion
// This API is used to delete rejected questions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRejectedQuestion(request *DeleteRejectedQuestionRequest) (response *DeleteRejectedQuestionResponse, err error) {
    return c.DeleteRejectedQuestionWithContext(context.Background(), request)
}

// DeleteRejectedQuestion
// This API is used to delete rejected questions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DeleteRejectedQuestionWithContext(ctx context.Context, request *DeleteRejectedQuestionRequest) (response *DeleteRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewDeleteRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAttributeLabelRequest() (request *DescribeAttributeLabelRequest) {
    request = &DescribeAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeAttributeLabel")
    
    
    return
}

func NewDescribeAttributeLabelResponse() (response *DescribeAttributeLabelResponse) {
    response = &DescribeAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAttributeLabel
// This API is used to query attribute label details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAttributeLabel(request *DescribeAttributeLabelRequest) (response *DescribeAttributeLabelResponse, err error) {
    return c.DescribeAttributeLabelWithContext(context.Background(), request)
}

// DescribeAttributeLabel
// This API is used to query attribute label details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeAttributeLabelWithContext(ctx context.Context, request *DescribeAttributeLabelRequest) (response *DescribeAttributeLabelResponse, err error) {
    if request == nil {
        request = NewDescribeAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallStatsGraphRequest() (request *DescribeCallStatsGraphRequest) {
    request = &DescribeCallStatsGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeCallStatsGraph")
    
    
    return
}

func NewDescribeCallStatsGraphResponse() (response *DescribeCallStatsGraphResponse) {
    response = &DescribeCallStatsGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallStatsGraph
// This API is used to show line chart of API calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCallStatsGraph(request *DescribeCallStatsGraphRequest) (response *DescribeCallStatsGraphResponse, err error) {
    return c.DescribeCallStatsGraphWithContext(context.Background(), request)
}

// DescribeCallStatsGraph
// This API is used to show line chart of API calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeCallStatsGraphWithContext(ctx context.Context, request *DescribeCallStatsGraphRequest) (response *DescribeCallStatsGraphResponse, err error) {
    if request == nil {
        request = NewDescribeCallStatsGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallStatsGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallStatsGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyUsageRequest() (request *DescribeConcurrencyUsageRequest) {
    request = &DescribeConcurrencyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeConcurrencyUsage")
    
    
    return
}

func NewDescribeConcurrencyUsageResponse() (response *DescribeConcurrencyUsageResponse) {
    response = &DescribeConcurrencyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyUsage
// This API is used to response to concurrent calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsage(request *DescribeConcurrencyUsageRequest) (response *DescribeConcurrencyUsageResponse, err error) {
    return c.DescribeConcurrencyUsageWithContext(context.Background(), request)
}

// DescribeConcurrencyUsage
// This API is used to response to concurrent calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageWithContext(ctx context.Context, request *DescribeConcurrencyUsageRequest) (response *DescribeConcurrencyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConcurrencyUsageGraphRequest() (request *DescribeConcurrencyUsageGraphRequest) {
    request = &DescribeConcurrencyUsageGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeConcurrencyUsageGraph")
    
    
    return
}

func NewDescribeConcurrencyUsageGraphResponse() (response *DescribeConcurrencyUsageGraphResponse) {
    response = &DescribeConcurrencyUsageGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConcurrencyUsageGraph
// This API is used to show line chart of concurrent calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageGraph(request *DescribeConcurrencyUsageGraphRequest) (response *DescribeConcurrencyUsageGraphResponse, err error) {
    return c.DescribeConcurrencyUsageGraphWithContext(context.Background(), request)
}

// DescribeConcurrencyUsageGraph
// This API is used to show line chart of concurrent calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeConcurrencyUsageGraphWithContext(ctx context.Context, request *DescribeConcurrencyUsageGraphRequest) (response *DescribeConcurrencyUsageGraphResponse, err error) {
    if request == nil {
        request = NewDescribeConcurrencyUsageGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConcurrencyUsageGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConcurrencyUsageGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDocRequest() (request *DescribeDocRequest) {
    request = &DescribeDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeDoc")
    
    
    return
}

func NewDescribeDocResponse() (response *DescribeDocResponse) {
    response = &DescribeDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDoc
// This API is used to query document details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDoc(request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    return c.DescribeDocWithContext(context.Background(), request)
}

// DescribeDoc
// This API is used to query document details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeDocWithContext(ctx context.Context, request *DescribeDocRequest) (response *DescribeDocResponse, err error) {
    if request == nil {
        request = NewDescribeDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDocResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeUsageRequest() (request *DescribeKnowledgeUsageRequest) {
    request = &DescribeKnowledgeUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeKnowledgeUsage")
    
    
    return
}

func NewDescribeKnowledgeUsageResponse() (response *DescribeKnowledgeUsageResponse) {
    response = &DescribeKnowledgeUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeUsage
// This API is used to query the knowledge library usage.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsage(request *DescribeKnowledgeUsageRequest) (response *DescribeKnowledgeUsageResponse, err error) {
    return c.DescribeKnowledgeUsageWithContext(context.Background(), request)
}

// DescribeKnowledgeUsage
// This API is used to query the knowledge library usage.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsageWithContext(ctx context.Context, request *DescribeKnowledgeUsageRequest) (response *DescribeKnowledgeUsageResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKnowledgeUsagePieGraphRequest() (request *DescribeKnowledgeUsagePieGraphRequest) {
    request = &DescribeKnowledgeUsagePieGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeKnowledgeUsagePieGraph")
    
    
    return
}

func NewDescribeKnowledgeUsagePieGraphResponse() (response *DescribeKnowledgeUsagePieGraphResponse) {
    response = &DescribeKnowledgeUsagePieGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeKnowledgeUsagePieGraph
// This API is used to query pie chart of the enterprise knowledge base capacity .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsagePieGraph(request *DescribeKnowledgeUsagePieGraphRequest) (response *DescribeKnowledgeUsagePieGraphResponse, err error) {
    return c.DescribeKnowledgeUsagePieGraphWithContext(context.Background(), request)
}

// DescribeKnowledgeUsagePieGraph
// This API is used to query pie chart of the enterprise knowledge base capacity .
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeKnowledgeUsagePieGraphWithContext(ctx context.Context, request *DescribeKnowledgeUsagePieGraphRequest) (response *DescribeKnowledgeUsagePieGraphResponse, err error) {
    if request == nil {
        request = NewDescribeKnowledgeUsagePieGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKnowledgeUsagePieGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKnowledgeUsagePieGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReleaseRequest() (request *DescribeReleaseRequest) {
    request = &DescribeReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRelease")
    
    
    return
}

func NewDescribeReleaseResponse() (response *DescribeReleaseResponse) {
    response = &DescribeReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRelease
// This API is used to query release details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRelease(request *DescribeReleaseRequest) (response *DescribeReleaseResponse, err error) {
    return c.DescribeReleaseWithContext(context.Background(), request)
}

// DescribeRelease
// This API is used to query release details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeReleaseWithContext(ctx context.Context, request *DescribeReleaseRequest) (response *DescribeReleaseResponse, err error) {
    if request == nil {
        request = NewDescribeReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRobotBizIDByAppKeyRequest() (request *DescribeRobotBizIDByAppKeyRequest) {
    request = &DescribeRobotBizIDByAppKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeRobotBizIDByAppKey")
    
    
    return
}

func NewDescribeRobotBizIDByAppKeyResponse() (response *DescribeRobotBizIDByAppKeyResponse) {
    response = &DescribeRobotBizIDByAppKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRobotBizIDByAppKey
// This API is used to get application business ID through appKey.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRobotBizIDByAppKey(request *DescribeRobotBizIDByAppKeyRequest) (response *DescribeRobotBizIDByAppKeyResponse, err error) {
    return c.DescribeRobotBizIDByAppKeyWithContext(context.Background(), request)
}

// DescribeRobotBizIDByAppKey
// This API is used to get application business ID through appKey.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeRobotBizIDByAppKeyWithContext(ctx context.Context, request *DescribeRobotBizIDByAppKeyRequest) (response *DescribeRobotBizIDByAppKeyResponse, err error) {
    if request == nil {
        request = NewDescribeRobotBizIDByAppKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRobotBizIDByAppKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRobotBizIDByAppKeyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSearchStatsGraphRequest() (request *DescribeSearchStatsGraphRequest) {
    request = &DescribeSearchStatsGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeSearchStatsGraph")
    
    
    return
}

func NewDescribeSearchStatsGraphResponse() (response *DescribeSearchStatsGraphResponse) {
    response = &DescribeSearchStatsGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSearchStatsGraph
// This API is used to query line chart of search service calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSearchStatsGraph(request *DescribeSearchStatsGraphRequest) (response *DescribeSearchStatsGraphResponse, err error) {
    return c.DescribeSearchStatsGraphWithContext(context.Background(), request)
}

// DescribeSearchStatsGraph
// This API is used to query line chart of search service calls.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeSearchStatsGraphWithContext(ctx context.Context, request *DescribeSearchStatsGraphRequest) (response *DescribeSearchStatsGraphResponse, err error) {
    if request == nil {
        request = NewDescribeSearchStatsGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSearchStatsGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSearchStatsGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenUsageRequest() (request *DescribeTokenUsageRequest) {
    request = &DescribeTokenUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeTokenUsage")
    
    
    return
}

func NewDescribeTokenUsageResponse() (response *DescribeTokenUsageResponse) {
    response = &DescribeTokenUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenUsage
// This API is used to query API call token details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsage(request *DescribeTokenUsageRequest) (response *DescribeTokenUsageResponse, err error) {
    return c.DescribeTokenUsageWithContext(context.Background(), request)
}

// DescribeTokenUsage
// This API is used to query API call token details.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageWithContext(ctx context.Context, request *DescribeTokenUsageRequest) (response *DescribeTokenUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTokenUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTokenUsageGraphRequest() (request *DescribeTokenUsageGraphRequest) {
    request = &DescribeTokenUsageGraphRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeTokenUsageGraph")
    
    
    return
}

func NewDescribeTokenUsageGraphResponse() (response *DescribeTokenUsageGraphResponse) {
    response = &DescribeTokenUsageGraphResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTokenUsageGraph
// This API is used to show API call token line chart.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageGraph(request *DescribeTokenUsageGraphRequest) (response *DescribeTokenUsageGraphResponse, err error) {
    return c.DescribeTokenUsageGraphWithContext(context.Background(), request)
}

// DescribeTokenUsageGraph
// This API is used to show API call token line chart.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeTokenUsageGraphWithContext(ctx context.Context, request *DescribeTokenUsageGraphRequest) (response *DescribeTokenUsageGraphResponse, err error) {
    if request == nil {
        request = NewDescribeTokenUsageGraphRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTokenUsageGraph require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTokenUsageGraphResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnsatisfiedReplyContextRequest() (request *DescribeUnsatisfiedReplyContextRequest) {
    request = &DescribeUnsatisfiedReplyContextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "DescribeUnsatisfiedReplyContext")
    
    
    return
}

func NewDescribeUnsatisfiedReplyContextResponse() (response *DescribeUnsatisfiedReplyContextResponse) {
    response = &DescribeUnsatisfiedReplyContextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnsatisfiedReplyContext
// This API is used to get the context of dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUnsatisfiedReplyContext(request *DescribeUnsatisfiedReplyContextRequest) (response *DescribeUnsatisfiedReplyContextResponse, err error) {
    return c.DescribeUnsatisfiedReplyContextWithContext(context.Background(), request)
}

// DescribeUnsatisfiedReplyContext
// This API is used to get the context of dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeUnsatisfiedReplyContextWithContext(ctx context.Context, request *DescribeUnsatisfiedReplyContextRequest) (response *DescribeUnsatisfiedReplyContextResponse, err error) {
    if request == nil {
        request = NewDescribeUnsatisfiedReplyContextRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnsatisfiedReplyContext require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnsatisfiedReplyContextResponse()
    err = c.Send(request, response)
    return
}

func NewExportQAListRequest() (request *ExportQAListRequest) {
    request = &ExportQAListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportQAList")
    
    
    return
}

func NewExportQAListResponse() (response *ExportQAListResponse) {
    response = &ExportQAListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportQAList
// This API is used to export Q&A list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportQAList(request *ExportQAListRequest) (response *ExportQAListResponse, err error) {
    return c.ExportQAListWithContext(context.Background(), request)
}

// ExportQAList
// This API is used to export Q&A list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportQAListWithContext(ctx context.Context, request *ExportQAListRequest) (response *ExportQAListResponse, err error) {
    if request == nil {
        request = NewExportQAListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportQAList require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportQAListResponse()
    err = c.Send(request, response)
    return
}

func NewExportUnsatisfiedReplyRequest() (request *ExportUnsatisfiedReplyRequest) {
    request = &ExportUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ExportUnsatisfiedReply")
    
    
    return
}

func NewExportUnsatisfiedReplyResponse() (response *ExportUnsatisfiedReplyResponse) {
    response = &ExportUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExportUnsatisfiedReply
// This API is used to export dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportUnsatisfiedReply(request *ExportUnsatisfiedReplyRequest) (response *ExportUnsatisfiedReplyResponse, err error) {
    return c.ExportUnsatisfiedReplyWithContext(context.Background(), request)
}

// ExportUnsatisfiedReply
// This API is used to export dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ExportUnsatisfiedReplyWithContext(ctx context.Context, request *ExportUnsatisfiedReplyRequest) (response *ExportUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewExportUnsatisfiedReplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExportUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewExportUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateQARequest() (request *GenerateQARequest) {
    request = &GenerateQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GenerateQA")
    
    
    return
}

func NewGenerateQAResponse() (response *GenerateQAResponse) {
    response = &GenerateQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateQA
// This API is used to generate Q%A from document.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenerateQA(request *GenerateQARequest) (response *GenerateQAResponse, err error) {
    return c.GenerateQAWithContext(context.Background(), request)
}

// GenerateQA
// This API is used to generate Q%A from document.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GenerateQAWithContext(ctx context.Context, request *GenerateQARequest) (response *GenerateQAResponse, err error) {
    if request == nil {
        request = NewGenerateQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateQAResponse()
    err = c.Send(request, response)
    return
}

func NewGetMsgRecordRequest() (request *GetMsgRecordRequest) {
    request = &GetMsgRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetMsgRecord")
    
    
    return
}

func NewGetMsgRecordResponse() (response *GetMsgRecordResponse) {
    response = &GetMsgRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMsgRecord
// This API is used to obtain chat history based on the session ID (only historical session data within the past 180 days will be retained).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetMsgRecord(request *GetMsgRecordRequest) (response *GetMsgRecordResponse, err error) {
    return c.GetMsgRecordWithContext(context.Background(), request)
}

// GetMsgRecord
// This API is used to obtain chat history based on the session ID (only historical session data within the past 180 days will be retained).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetMsgRecordWithContext(ctx context.Context, request *GetMsgRecordRequest) (response *GetMsgRecordResponse, err error) {
    if request == nil {
        request = NewGetMsgRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMsgRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMsgRecordResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskStatusRequest() (request *GetTaskStatusRequest) {
    request = &GetTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetTaskStatus")
    
    
    return
}

func NewGetTaskStatusResponse() (response *GetTaskStatusResponse) {
    response = &GetTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskStatus
// This API is used to get the task status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    return c.GetTaskStatusWithContext(context.Background(), request)
}

// GetTaskStatus
// This API is used to get the task status.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetTaskStatusWithContext(ctx context.Context, request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
    if request == nil {
        request = NewGetTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewGetWsTokenRequest() (request *GetWsTokenRequest) {
    request = &GetWsTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "GetWsToken")
    
    
    return
}

func NewGetWsTokenResponse() (response *GetWsTokenResponse) {
    response = &GetWsTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetWsToken
// This API is used to get ws token.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetWsToken(request *GetWsTokenRequest) (response *GetWsTokenResponse, err error) {
    return c.GetWsTokenWithContext(context.Background(), request)
}

// GetWsToken
// This API is used to get ws token.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) GetWsTokenWithContext(ctx context.Context, request *GetWsTokenRequest) (response *GetWsTokenResponse, err error) {
    if request == nil {
        request = NewGetWsTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetWsToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetWsTokenResponse()
    err = c.Send(request, response)
    return
}

func NewIgnoreUnsatisfiedReplyRequest() (request *IgnoreUnsatisfiedReplyRequest) {
    request = &IgnoreUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "IgnoreUnsatisfiedReply")
    
    
    return
}

func NewIgnoreUnsatisfiedReplyResponse() (response *IgnoreUnsatisfiedReplyResponse) {
    response = &IgnoreUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IgnoreUnsatisfiedReply
// This API is used to ignore dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IgnoreUnsatisfiedReply(request *IgnoreUnsatisfiedReplyRequest) (response *IgnoreUnsatisfiedReplyResponse, err error) {
    return c.IgnoreUnsatisfiedReplyWithContext(context.Background(), request)
}

// IgnoreUnsatisfiedReply
// This API is used to ignore dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) IgnoreUnsatisfiedReplyWithContext(ctx context.Context, request *IgnoreUnsatisfiedReplyRequest) (response *IgnoreUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewIgnoreUnsatisfiedReplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IgnoreUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewIgnoreUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewListAppRequest() (request *ListAppRequest) {
    request = &ListAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListApp")
    
    
    return
}

func NewListAppResponse() (response *ListAppResponse) {
    response = &ListAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListApp
// This API is used to get the application list under the corporate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListApp(request *ListAppRequest) (response *ListAppResponse, err error) {
    return c.ListAppWithContext(context.Background(), request)
}

// ListApp
// This API is used to get the application list under the corporate.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAppWithContext(ctx context.Context, request *ListAppRequest) (response *ListAppResponse, err error) {
    if request == nil {
        request = NewListAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAppResponse()
    err = c.Send(request, response)
    return
}

func NewListAttributeLabelRequest() (request *ListAttributeLabelRequest) {
    request = &ListAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListAttributeLabel")
    
    
    return
}

func NewListAttributeLabelResponse() (response *ListAttributeLabelResponse) {
    response = &ListAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListAttributeLabel
// This API is used to query attribute label lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAttributeLabel(request *ListAttributeLabelRequest) (response *ListAttributeLabelResponse, err error) {
    return c.ListAttributeLabelWithContext(context.Background(), request)
}

// ListAttributeLabel
// This API is used to query attribute label lists.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListAttributeLabelWithContext(ctx context.Context, request *ListAttributeLabelRequest) (response *ListAttributeLabelResponse, err error) {
    if request == nil {
        request = NewListAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewListAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewListQACateRequest() (request *ListQACateRequest) {
    request = &ListQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListQACate")
    
    
    return
}

func NewListQACateResponse() (response *ListQACateResponse) {
    response = &ListQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListQACate
// This API is used to get Q&A categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQACate(request *ListQACateRequest) (response *ListQACateResponse, err error) {
    return c.ListQACateWithContext(context.Background(), request)
}

// ListQACate
// This API is used to get Q&A categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListQACateWithContext(ctx context.Context, request *ListQACateRequest) (response *ListQACateResponse, err error) {
    if request == nil {
        request = NewListQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewListQACateResponse()
    err = c.Send(request, response)
    return
}

func NewListRejectedQuestionRequest() (request *ListRejectedQuestionRequest) {
    request = &ListRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListRejectedQuestion")
    
    
    return
}

func NewListRejectedQuestionResponse() (response *ListRejectedQuestionResponse) {
    response = &ListRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListRejectedQuestion
// This API is used to get rejected questions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestion(request *ListRejectedQuestionRequest) (response *ListRejectedQuestionResponse, err error) {
    return c.ListRejectedQuestionWithContext(context.Background(), request)
}

// ListRejectedQuestion
// This API is used to get rejected questions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListRejectedQuestionWithContext(ctx context.Context, request *ListRejectedQuestionRequest) (response *ListRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewListRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewListRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewListReleaseConfigPreviewRequest() (request *ListReleaseConfigPreviewRequest) {
    request = &ListReleaseConfigPreviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListReleaseConfigPreview")
    
    
    return
}

func NewListReleaseConfigPreviewResponse() (response *ListReleaseConfigPreviewResponse) {
    response = &ListReleaseConfigPreviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListReleaseConfigPreview
// This API is used to preview the release configuration items.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseConfigPreview(request *ListReleaseConfigPreviewRequest) (response *ListReleaseConfigPreviewResponse, err error) {
    return c.ListReleaseConfigPreviewWithContext(context.Background(), request)
}

// ListReleaseConfigPreview
// This API is used to preview the release configuration items.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListReleaseConfigPreviewWithContext(ctx context.Context, request *ListReleaseConfigPreviewRequest) (response *ListReleaseConfigPreviewResponse, err error) {
    if request == nil {
        request = NewListReleaseConfigPreviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListReleaseConfigPreview require credential")
    }

    request.SetContext(ctx)
    
    response = NewListReleaseConfigPreviewResponse()
    err = c.Send(request, response)
    return
}

func NewListSelectDocRequest() (request *ListSelectDocRequest) {
    request = &ListSelectDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListSelectDoc")
    
    
    return
}

func NewListSelectDocResponse() (response *ListSelectDocResponse) {
    response = &ListSelectDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListSelectDoc
// This API is used to get account information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSelectDoc(request *ListSelectDocRequest) (response *ListSelectDocResponse, err error) {
    return c.ListSelectDocWithContext(context.Background(), request)
}

// ListSelectDoc
// This API is used to get account information.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ListSelectDocWithContext(ctx context.Context, request *ListSelectDocRequest) (response *ListSelectDocResponse, err error) {
    if request == nil {
        request = NewListSelectDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListSelectDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewListSelectDocResponse()
    err = c.Send(request, response)
    return
}

func NewListUnsatisfiedReplyRequest() (request *ListUnsatisfiedReplyRequest) {
    request = &ListUnsatisfiedReplyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListUnsatisfiedReply")
    
    
    return
}

func NewListUnsatisfiedReplyResponse() (response *ListUnsatisfiedReplyResponse) {
    response = &ListUnsatisfiedReplyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUnsatisfiedReply
// This API is used to query a list of dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUnsatisfiedReply(request *ListUnsatisfiedReplyRequest) (response *ListUnsatisfiedReplyResponse, err error) {
    return c.ListUnsatisfiedReplyWithContext(context.Background(), request)
}

// ListUnsatisfiedReply
// This API is used to query a list of dissatisfied responses.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUnsatisfiedReplyWithContext(ctx context.Context, request *ListUnsatisfiedReplyRequest) (response *ListUnsatisfiedReplyResponse, err error) {
    if request == nil {
        request = NewListUnsatisfiedReplyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUnsatisfiedReply require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUnsatisfiedReplyResponse()
    err = c.Send(request, response)
    return
}

func NewListUsageCallDetailRequest() (request *ListUsageCallDetailRequest) {
    request = &ListUsageCallDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ListUsageCallDetail")
    
    
    return
}

func NewListUsageCallDetailResponse() (response *ListUsageCallDetailResponse) {
    response = &ListUsageCallDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ListUsageCallDetail
// This API is used to query single call details in a list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUsageCallDetail(request *ListUsageCallDetailRequest) (response *ListUsageCallDetailResponse, err error) {
    return c.ListUsageCallDetailWithContext(context.Background(), request)
}

// ListUsageCallDetail
// This API is used to query single call details in a list.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ListUsageCallDetailWithContext(ctx context.Context, request *ListUsageCallDetailRequest) (response *ListUsageCallDetailResponse, err error) {
    if request == nil {
        request = NewListUsageCallDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ListUsageCallDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewListUsageCallDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppRequest() (request *ModifyAppRequest) {
    request = &ModifyAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyApp")
    
    
    return
}

func NewModifyAppResponse() (response *ModifyAppResponse) {
    response = &ModifyAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApp
// This API is used to modify application request structure.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyApp(request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    return c.ModifyAppWithContext(context.Background(), request)
}

// ModifyApp
// This API is used to modify application request structure.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAppWithContext(ctx context.Context, request *ModifyAppRequest) (response *ModifyAppResponse, err error) {
    if request == nil {
        request = NewModifyAppRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApp require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAppResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAttributeLabelRequest() (request *ModifyAttributeLabelRequest) {
    request = &ModifyAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyAttributeLabel")
    
    
    return
}

func NewModifyAttributeLabelResponse() (response *ModifyAttributeLabelResponse) {
    response = &ModifyAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAttributeLabel
// This API is used to edit attribute labels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAttributeLabel(request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    return c.ModifyAttributeLabelWithContext(context.Background(), request)
}

// ModifyAttributeLabel
// This API is used to edit attribute labels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyAttributeLabelWithContext(ctx context.Context, request *ModifyAttributeLabelRequest) (response *ModifyAttributeLabelResponse, err error) {
    if request == nil {
        request = NewModifyAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocRequest() (request *ModifyDocRequest) {
    request = &ModifyDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDoc")
    
    
    return
}

func NewModifyDocResponse() (response *ModifyDocResponse) {
    response = &ModifyDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDoc
// This API is used to modify documents.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDoc(request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    return c.ModifyDocWithContext(context.Background(), request)
}

// ModifyDoc
// This API is used to modify documents.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyDocWithContext(ctx context.Context, request *ModifyDocRequest) (response *ModifyDocResponse, err error) {
    if request == nil {
        request = NewModifyDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocAttrRangeRequest() (request *ModifyDocAttrRangeRequest) {
    request = &ModifyDocAttrRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDocAttrRange")
    
    
    return
}

func NewModifyDocAttrRangeResponse() (response *ModifyDocAttrRangeResponse) {
    response = &ModifyDocAttrRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocAttrRange
// This API is used to modify applicable scope of the document in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocAttrRange(request *ModifyDocAttrRangeRequest) (response *ModifyDocAttrRangeResponse, err error) {
    return c.ModifyDocAttrRangeWithContext(context.Background(), request)
}

// ModifyDocAttrRange
// This API is used to modify applicable scope of the document in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocAttrRangeWithContext(ctx context.Context, request *ModifyDocAttrRangeRequest) (response *ModifyDocAttrRangeResponse, err error) {
    if request == nil {
        request = NewModifyDocAttrRangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocAttrRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocAttrRangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDocCateRequest() (request *ModifyDocCateRequest) {
    request = &ModifyDocCateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyDocCate")
    
    
    return
}

func NewModifyDocCateResponse() (response *ModifyDocCateResponse) {
    response = &ModifyDocCateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDocCate
// This API is used to modify Doc categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocCate(request *ModifyDocCateRequest) (response *ModifyDocCateResponse, err error) {
    return c.ModifyDocCateWithContext(context.Background(), request)
}

// ModifyDocCate
// This API is used to modify Doc categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyDocCateWithContext(ctx context.Context, request *ModifyDocCateRequest) (response *ModifyDocCateResponse, err error) {
    if request == nil {
        request = NewModifyDocCateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDocCate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDocCateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQARequest() (request *ModifyQARequest) {
    request = &ModifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQA")
    
    
    return
}

func NewModifyQAResponse() (response *ModifyQAResponse) {
    response = &ModifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQA
// This API is used to update Q&As.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQA(request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    return c.ModifyQAWithContext(context.Background(), request)
}

// ModifyQA
// This API is used to update Q&As.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAWithContext(ctx context.Context, request *ModifyQARequest) (response *ModifyQAResponse, err error) {
    if request == nil {
        request = NewModifyQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQAAttrRangeRequest() (request *ModifyQAAttrRangeRequest) {
    request = &ModifyQAAttrRangeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQAAttrRange")
    
    
    return
}

func NewModifyQAAttrRangeResponse() (response *ModifyQAAttrRangeResponse) {
    response = &ModifyQAAttrRangeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQAAttrRange
// This API is used to modify Q&A applicable scope in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAAttrRange(request *ModifyQAAttrRangeRequest) (response *ModifyQAAttrRangeResponse, err error) {
    return c.ModifyQAAttrRangeWithContext(context.Background(), request)
}

// ModifyQAAttrRange
// This API is used to modify Q&A applicable scope in batches.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQAAttrRangeWithContext(ctx context.Context, request *ModifyQAAttrRangeRequest) (response *ModifyQAAttrRangeResponse, err error) {
    if request == nil {
        request = NewModifyQAAttrRangeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQAAttrRange require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQAAttrRangeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyQACateRequest() (request *ModifyQACateRequest) {
    request = &ModifyQACateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyQACate")
    
    
    return
}

func NewModifyQACateResponse() (response *ModifyQACateResponse) {
    response = &ModifyQACateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyQACate
// This API is used to modify Q&A categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQACate(request *ModifyQACateRequest) (response *ModifyQACateResponse, err error) {
    return c.ModifyQACateWithContext(context.Background(), request)
}

// ModifyQACate
// This API is used to modify Q&A categories.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyQACateWithContext(ctx context.Context, request *ModifyQACateRequest) (response *ModifyQACateResponse, err error) {
    if request == nil {
        request = NewModifyQACateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyQACate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyQACateResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRejectedQuestionRequest() (request *ModifyRejectedQuestionRequest) {
    request = &ModifyRejectedQuestionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "ModifyRejectedQuestion")
    
    
    return
}

func NewModifyRejectedQuestionResponse() (response *ModifyRejectedQuestionResponse) {
    response = &ModifyRejectedQuestionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyRejectedQuestion
// This API is used to modify rejected questions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRejectedQuestion(request *ModifyRejectedQuestionRequest) (response *ModifyRejectedQuestionResponse, err error) {
    return c.ModifyRejectedQuestionWithContext(context.Background(), request)
}

// ModifyRejectedQuestion
// This API is used to modify rejected questions.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) ModifyRejectedQuestionWithContext(ctx context.Context, request *ModifyRejectedQuestionRequest) (response *ModifyRejectedQuestionResponse, err error) {
    if request == nil {
        request = NewModifyRejectedQuestionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyRejectedQuestion require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyRejectedQuestionResponse()
    err = c.Send(request, response)
    return
}

func NewRateMsgRecordRequest() (request *RateMsgRecordRequest) {
    request = &RateMsgRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RateMsgRecord")
    
    
    return
}

func NewRateMsgRecordResponse() (response *RateMsgRecordResponse) {
    response = &RateMsgRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RateMsgRecord
// This API is used to show messages for likes and dislikes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RateMsgRecord(request *RateMsgRecordRequest) (response *RateMsgRecordResponse, err error) {
    return c.RateMsgRecordWithContext(context.Background(), request)
}

// RateMsgRecord
// This API is used to show messages for likes and dislikes.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RateMsgRecordWithContext(ctx context.Context, request *RateMsgRecordRequest) (response *RateMsgRecordResponse, err error) {
    if request == nil {
        request = NewRateMsgRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RateMsgRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewRateMsgRecordResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDocAuditRequest() (request *RetryDocAuditRequest) {
    request = &RetryDocAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryDocAudit")
    
    
    return
}

func NewRetryDocAuditResponse() (response *RetryDocAuditResponse) {
    response = &RetryDocAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDocAudit
// This API is used to retry document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocAudit(request *RetryDocAuditRequest) (response *RetryDocAuditResponse, err error) {
    return c.RetryDocAuditWithContext(context.Background(), request)
}

// RetryDocAudit
// This API is used to retry document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocAuditWithContext(ctx context.Context, request *RetryDocAuditRequest) (response *RetryDocAuditResponse, err error) {
    if request == nil {
        request = NewRetryDocAuditRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDocAudit require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDocAuditResponse()
    err = c.Send(request, response)
    return
}

func NewRetryDocParseRequest() (request *RetryDocParseRequest) {
    request = &RetryDocParseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryDocParse")
    
    
    return
}

func NewRetryDocParseResponse() (response *RetryDocParseResponse) {
    response = &RetryDocParseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryDocParse
// This API is used to retry document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocParse(request *RetryDocParseRequest) (response *RetryDocParseResponse, err error) {
    return c.RetryDocParseWithContext(context.Background(), request)
}

// RetryDocParse
// This API is used to retry document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryDocParseWithContext(ctx context.Context, request *RetryDocParseRequest) (response *RetryDocParseResponse, err error) {
    if request == nil {
        request = NewRetryDocParseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryDocParse require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryDocParseResponse()
    err = c.Send(request, response)
    return
}

func NewRetryReleaseRequest() (request *RetryReleaseRequest) {
    request = &RetryReleaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "RetryRelease")
    
    
    return
}

func NewRetryReleaseResponse() (response *RetryReleaseResponse) {
    response = &RetryReleaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RetryRelease
// This API is used to retry after release suspension.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryRelease(request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    return c.RetryReleaseWithContext(context.Background(), request)
}

// RetryRelease
// This API is used to retry after release suspension.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RetryReleaseWithContext(ctx context.Context, request *RetryReleaseRequest) (response *RetryReleaseResponse, err error) {
    if request == nil {
        request = NewRetryReleaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryRelease require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryReleaseResponse()
    err = c.Send(request, response)
    return
}

func NewSaveDocRequest() (request *SaveDocRequest) {
    request = &SaveDocRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "SaveDoc")
    
    
    return
}

func NewSaveDocResponse() (response *SaveDocResponse) {
    response = &SaveDocResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SaveDoc
// This API is used to save a knowledge base document Q&As.
//
// Three steps to store a file in the knowledge library of the application: 
//
// 1. Obtain a temporary key. For more information, see [API Documentation](https://cloud.tencent.com/document/product/1759/105050). Different parameter combinations of the temporary key have different permissions. For more information, see [Tencent Cloud Agent Development Platform/TCADP COS Guide](https://cloud.tencent.com/document/product/1759/116238). 
//
// 2. Call the COS storage API provided by Tencent Cloud to store the file in the COS of TCADP. For details, see [COS SDK Overview](https://cloud.tencent.com/document/product/436/6474). Note that the temporary key method is used to operate COS. 
//
// 3. Call this API to store the basic information of the file in TCADP. 
//
// For the above steps, see [Documentation](https://cloud.tencent.com/document/product/1759/108903). At the end of the documentation, there is a [code demo](https://cloud.tencent.com/document/product/1759/108903#demo), which can be used as a reference.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveDoc(request *SaveDocRequest) (response *SaveDocResponse, err error) {
    return c.SaveDocWithContext(context.Background(), request)
}

// SaveDoc
// This API is used to save a knowledge base document Q&As.
//
// Three steps to store a file in the knowledge library of the application: 
//
// 1. Obtain a temporary key. For more information, see [API Documentation](https://cloud.tencent.com/document/product/1759/105050). Different parameter combinations of the temporary key have different permissions. For more information, see [Tencent Cloud Agent Development Platform/TCADP COS Guide](https://cloud.tencent.com/document/product/1759/116238). 
//
// 2. Call the COS storage API provided by Tencent Cloud to store the file in the COS of TCADP. For details, see [COS SDK Overview](https://cloud.tencent.com/document/product/436/6474). Note that the temporary key method is used to operate COS. 
//
// 3. Call this API to store the basic information of the file in TCADP. 
//
// For the above steps, see [Documentation](https://cloud.tencent.com/document/product/1759/108903). At the end of the documentation, there is a [code demo](https://cloud.tencent.com/document/product/1759/108903#demo), which can be used as a reference.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) SaveDocWithContext(ctx context.Context, request *SaveDocRequest) (response *SaveDocResponse, err error) {
    if request == nil {
        request = NewSaveDocRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveDoc require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveDocResponse()
    err = c.Send(request, response)
    return
}

func NewStopDocParseRequest() (request *StopDocParseRequest) {
    request = &StopDocParseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "StopDocParse")
    
    
    return
}

func NewStopDocParseResponse() (response *StopDocParseResponse) {
    response = &StopDocParseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopDocParse
// This API is used to terminate document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopDocParse(request *StopDocParseRequest) (response *StopDocParseResponse, err error) {
    return c.StopDocParseWithContext(context.Background(), request)
}

// StopDocParse
// This API is used to terminate document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) StopDocParseWithContext(ctx context.Context, request *StopDocParseRequest) (response *StopDocParseResponse, err error) {
    if request == nil {
        request = NewStopDocParseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopDocParse require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopDocParseResponse()
    err = c.Send(request, response)
    return
}

func NewUploadAttributeLabelRequest() (request *UploadAttributeLabelRequest) {
    request = &UploadAttributeLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "UploadAttributeLabel")
    
    
    return
}

func NewUploadAttributeLabelResponse() (response *UploadAttributeLabelResponse) {
    response = &UploadAttributeLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadAttributeLabel
// This API is used to upload and import attribute labels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadAttributeLabel(request *UploadAttributeLabelRequest) (response *UploadAttributeLabelResponse, err error) {
    return c.UploadAttributeLabelWithContext(context.Background(), request)
}

// UploadAttributeLabel
// This API is used to upload and import attribute labels.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) UploadAttributeLabelWithContext(ctx context.Context, request *UploadAttributeLabelRequest) (response *UploadAttributeLabelResponse, err error) {
    if request == nil {
        request = NewUploadAttributeLabelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadAttributeLabel require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadAttributeLabelResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyQARequest() (request *VerifyQARequest) {
    request = &VerifyQARequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lke", APIVersion, "VerifyQA")
    
    
    return
}

func NewVerifyQAResponse() (response *VerifyQAResponse) {
    response = &VerifyQAResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyQA
// This API is used to verify Q&A.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) VerifyQA(request *VerifyQARequest) (response *VerifyQAResponse, err error) {
    return c.VerifyQAWithContext(context.Background(), request)
}

// VerifyQA
// This API is used to verify Q&A.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) VerifyQAWithContext(ctx context.Context, request *VerifyQARequest) (response *VerifyQAResponse, err error) {
    if request == nil {
        request = NewVerifyQARequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyQA require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyQAResponse()
    err = c.Send(request, response)
    return
}

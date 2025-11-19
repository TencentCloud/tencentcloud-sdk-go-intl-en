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

package v20240522

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2024-05-22"

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


func NewCreateReconstructDocumentFlowRequest() (request *CreateReconstructDocumentFlowRequest) {
    request = &CreateReconstructDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateReconstructDocumentFlow")
    
    
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
// Document parsing supports converting images or PDF files into Markdown format files, and can parse content elements including tables, formulas, images, headings, paragraphs, headers, and footers, and intelligently convert the content into reading order. Please refer to the input parameter list below for specific supported file types.
//
// During the trial period, the QPS limit for a single account is only 1. If you need to access officially, please contact our R&D team.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlow(request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    return c.CreateReconstructDocumentFlowWithContext(context.Background(), request)
}

// CreateReconstructDocumentFlow
// This API is used to initiate requests for this asynchronous API, for initiating document parsing tasks.
//
// Document parsing supports converting images or PDF files into Markdown format files, and can parse content elements including tables, formulas, images, headings, paragraphs, headers, and footers, and intelligently convert the content into reading order. Please refer to the input parameter list below for specific supported file types.
//
// During the trial period, the QPS limit for a single account is only 1. If you need to access officially, please contact our R&D team.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CreateReconstructDocumentFlowWithContext(ctx context.Context, request *CreateReconstructDocumentFlowRequest) (response *CreateReconstructDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateReconstructDocumentFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "CreateReconstructDocumentFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReconstructDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReconstructDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSplitDocumentFlowRequest() (request *CreateSplitDocumentFlowRequest) {
    request = &CreateSplitDocumentFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "CreateSplitDocumentFlow")
    
    
    return
}

func NewCreateSplitDocumentFlowResponse() (response *CreateSplitDocumentFlowResponse) {
    response = &CreateSplitDocumentFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSplitDocumentFlow
// This API is used to create document segmentation tasks, support various file types, possess mllm capacity, and can analyze and deeply understand the information in charts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateSplitDocumentFlow(request *CreateSplitDocumentFlowRequest) (response *CreateSplitDocumentFlowResponse, err error) {
    return c.CreateSplitDocumentFlowWithContext(context.Background(), request)
}

// CreateSplitDocumentFlow
// This API is used to create document segmentation tasks, support various file types, possess mllm capacity, and can analyze and deeply understand the information in charts.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) CreateSplitDocumentFlowWithContext(ctx context.Context, request *CreateSplitDocumentFlowRequest) (response *CreateSplitDocumentFlowResponse, err error) {
    if request == nil {
        request = NewCreateSplitDocumentFlowRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "CreateSplitDocumentFlow")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSplitDocumentFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSplitDocumentFlowResponse()
    err = c.Send(request, response)
    return
}

func NewGetEmbeddingRequest() (request *GetEmbeddingRequest) {
    request = &GetEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetEmbedding")
    
    
    return
}

func NewGetEmbeddingResponse() (response *GetEmbeddingResponse) {
    response = &GetEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetEmbedding
// This API is used to call the text representation model to convert text into a vector represented by numbers, which can be used in scenarios such as text retrieval, information recommendation, and knowledge mining. There is a single-account call limit control for this API. If you need to increase the concurrency limit, please contact us (https://cloud.tencent.com/act/event/Online_service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) GetEmbedding(request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    return c.GetEmbeddingWithContext(context.Background(), request)
}

// GetEmbedding
// This API is used to call the text representation model to convert text into a vector represented by numbers, which can be used in scenarios such as text retrieval, information recommendation, and knowledge mining. There is a single-account call limit control for this API. If you need to increase the concurrency limit, please contact us (https://cloud.tencent.com/act/event/Online_service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) GetEmbeddingWithContext(ctx context.Context, request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
    if request == nil {
        request = NewGetEmbeddingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "GetEmbedding")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewGetReconstructDocumentResultRequest() (request *GetReconstructDocumentResultRequest) {
    request = &GetReconstructDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetReconstructDocumentResult")
    
    
    return
}

func NewGetReconstructDocumentResultResponse() (response *GetReconstructDocumentResultResponse) {
    response = &GetReconstructDocumentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetReconstructDocumentResult
// This is an asynchronous API for querying results, which is used to obtain the result of document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION_FILEPARSEERROR = "FailedOperation.FileParseError"
//  FAILEDOPERATION_FILEPARSETIMEOUT = "FailedOperation.FileParseTimeout"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetReconstructDocumentResult(request *GetReconstructDocumentResultRequest) (response *GetReconstructDocumentResultResponse, err error) {
    return c.GetReconstructDocumentResultWithContext(context.Background(), request)
}

// GetReconstructDocumentResult
// This is an asynchronous API for querying results, which is used to obtain the result of document parsing.
//
// error code that may be returned:
//  FAILEDOPERATION_FILEPARSEERROR = "FailedOperation.FileParseError"
//  FAILEDOPERATION_FILEPARSETIMEOUT = "FailedOperation.FileParseTimeout"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetReconstructDocumentResultWithContext(ctx context.Context, request *GetReconstructDocumentResultRequest) (response *GetReconstructDocumentResultResponse, err error) {
    if request == nil {
        request = NewGetReconstructDocumentResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "GetReconstructDocumentResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetReconstructDocumentResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetReconstructDocumentResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetSplitDocumentResultRequest() (request *GetSplitDocumentResultRequest) {
    request = &GetSplitDocumentResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "GetSplitDocumentResult")
    
    
    return
}

func NewGetSplitDocumentResultResponse() (response *GetSplitDocumentResultResponse) {
    response = &GetSplitDocumentResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetSplitDocumentResult
// This API is used to query the results of document splitting tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_FILEPARSEERROR = "FailedOperation.FileParseError"
//  FAILEDOPERATION_FILEPARSETIMEOUT = "FailedOperation.FileParseTimeout"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetSplitDocumentResult(request *GetSplitDocumentResultRequest) (response *GetSplitDocumentResultResponse, err error) {
    return c.GetSplitDocumentResultWithContext(context.Background(), request)
}

// GetSplitDocumentResult
// This API is used to query the results of document splitting tasks.
//
// error code that may be returned:
//  FAILEDOPERATION_FILEPARSEERROR = "FailedOperation.FileParseError"
//  FAILEDOPERATION_FILEPARSETIMEOUT = "FailedOperation.FileParseTimeout"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETER_FILEURLINVALID = "InvalidParameter.FileURLInvalid"
//  INVALIDPARAMETER_INVALIDFILEFORMAT = "InvalidParameter.InvalidFileFormat"
//  INVALIDPARAMETER_INVALIDFILETYPE = "InvalidParameter.InvalidFileType"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GetSplitDocumentResultWithContext(ctx context.Context, request *GetSplitDocumentResultRequest) (response *GetSplitDocumentResultResponse, err error) {
    if request == nil {
        request = NewGetSplitDocumentResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "GetSplitDocumentResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSplitDocumentResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetSplitDocumentResultResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRewriteRequest() (request *QueryRewriteRequest) {
    request = &QueryRewriteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "QueryRewrite")
    
    
    return
}

func NewQueryRewriteResponse() (response *QueryRewriteResponse) {
    response = &QueryRewriteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryRewrite
// QueryRewrite is mainly used in multi-round conversations for reference resolution and ellipsis completion. Using this API, you don't need to input prompt descriptions. A more accurate user query can be generated based on the conversation history. In terms of application scenarios, this API can be applied to various scenarios such as intelligent Q&A and conversational search.
//
// There is a call limit for single-account for this API. If you need to increase the concurrency limit, please contact us (https://cloud.tencent.com/act/event/Online_service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryRewrite(request *QueryRewriteRequest) (response *QueryRewriteResponse, err error) {
    return c.QueryRewriteWithContext(context.Background(), request)
}

// QueryRewrite
// QueryRewrite is mainly used in multi-round conversations for reference resolution and ellipsis completion. Using this API, you don't need to input prompt descriptions. A more accurate user query can be generated based on the conversation history. In terms of application scenarios, this API can be applied to various scenarios such as intelligent Q&A and conversational search.
//
// There is a call limit for single-account for this API. If you need to increase the concurrency limit, please contact us (https://cloud.tencent.com/act/event/Online_service).
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) QueryRewriteWithContext(ctx context.Context, request *QueryRewriteRequest) (response *QueryRewriteResponse, err error) {
    if request == nil {
        request = NewQueryRewriteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "QueryRewrite")
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryRewrite require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryRewriteResponse()
    err = c.Send(request, response)
    return
}

func NewReconstructDocumentSSERequest() (request *ReconstructDocumentSSERequest) {
    request = &ReconstructDocumentSSERequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "ReconstructDocumentSSE")
    
    
    return
}

func NewReconstructDocumentSSEResponse() (response *ReconstructDocumentSSEResponse) {
    response = &ReconstructDocumentSSEResponse{} 
    return

}

// ReconstructDocumentSSE
// This API is used for quasi-real-time document parsing, using HTTP SSE protocol for communication.
//
// error code that may be returned:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentSSE(request *ReconstructDocumentSSERequest) (response *ReconstructDocumentSSEResponse, err error) {
    return c.ReconstructDocumentSSEWithContext(context.Background(), request)
}

// ReconstructDocumentSSE
// This API is used for quasi-real-time document parsing, using HTTP SSE protocol for communication.
//
// error code that may be returned:
//  FAILEDOPERATION_NONSUPPORTPARSE = "FailedOperation.NonsupportParse"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UPLOADRESULTFILEFAILED = "FailedOperation.UploadResultFileFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_EXCEEDEDMAXPAGESERROR = "LimitExceeded.ExceededMaxPagesError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReconstructDocumentSSEWithContext(ctx context.Context, request *ReconstructDocumentSSERequest) (response *ReconstructDocumentSSEResponse, err error) {
    if request == nil {
        request = NewReconstructDocumentSSERequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "ReconstructDocumentSSE")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReconstructDocumentSSE require credential")
    }

    request.SetContext(ctx)
    
    response = NewReconstructDocumentSSEResponse()
    err = c.Send(request, response)
    return
}

func NewRunRerankRequest() (request *RunRerankRequest) {
    request = &RunRerankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("lkeap", APIVersion, "RunRerank")
    
    
    return
}

func NewRunRerankResponse() (response *RunRerankResponse) {
    response = &RunRerankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunRerank
// This API is used to reorder the results of multi-channel recall based on the rerank model of knowledge engine fine-tuning model technology, sort the segments according to the relevance between the query and the segment content from high to low score, and output the corresponding scoring results.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerank(request *RunRerankRequest) (response *RunRerankResponse, err error) {
    return c.RunRerankWithContext(context.Background(), request)
}

// RunRerank
// This API is used to reorder the results of multi-channel recall based on the rerank model of knowledge engine fine-tuning model technology, sort the segments according to the relevance between the query and the segment content from high to low score, and output the corresponding scoring results.
//
// error code that may be returned:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) RunRerankWithContext(ctx context.Context, request *RunRerankRequest) (response *RunRerankResponse, err error) {
    if request == nil {
        request = NewRunRerankRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "lkeap", APIVersion, "RunRerank")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunRerank require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunRerankResponse()
    err = c.Send(request, response)
    return
}

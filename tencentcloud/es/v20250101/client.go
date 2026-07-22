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

package v20250101

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2025-01-01"

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


func NewChunkDocumentRequest() (request *ChunkDocumentRequest) {
    request = &ChunkDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ChunkDocument")
    
    
    return
}

func NewChunkDocumentResponse() (response *ChunkDocumentResponse) {
    response = &ChunkDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChunkDocument
// Text segmentation is a technology that splits long text into short fragments for adapting to model input, improving processing efficiency, or information retrieval. It balances fragment length and semantic consistency, suitable for NLP and data analysis scenarios.
//
// This API is used to slice text based on delimiter rules. It has a single-account call limit. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
func (c *Client) ChunkDocument(request *ChunkDocumentRequest) (response *ChunkDocumentResponse, err error) {
    return c.ChunkDocumentWithContext(context.Background(), request)
}

// ChunkDocument
// Text segmentation is a technology that splits long text into short fragments for adapting to model input, improving processing efficiency, or information retrieval. It balances fragment length and semantic consistency, suitable for NLP and data analysis scenarios.
//
// This API is used to slice text based on delimiter rules. It has a single-account call limit. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
func (c *Client) ChunkDocumentWithContext(ctx context.Context, request *ChunkDocumentRequest) (response *ChunkDocumentResponse, err error) {
    if request == nil {
        request = NewChunkDocumentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "ChunkDocument")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChunkDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewChunkDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewChunkDocumentAsyncRequest() (request *ChunkDocumentAsyncRequest) {
    request = &ChunkDocumentAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ChunkDocumentAsync")
    
    
    return
}

func NewChunkDocumentAsyncResponse() (response *ChunkDocumentAsyncResponse) {
    response = &ChunkDocumentAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChunkDocumentAsync
// Text segmentation is a technology that splits long text into short clips for adapting to model input, improving processing efficiency, or information retrieval. It balances clip length and semantic consistency, suitable for NLP and data analysis scenarios.
//
// This API is an async API with a model dimensional call limit. Each model has a qps limit of 5. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChunkDocumentAsync(request *ChunkDocumentAsyncRequest) (response *ChunkDocumentAsyncResponse, err error) {
    return c.ChunkDocumentAsyncWithContext(context.Background(), request)
}

// ChunkDocumentAsync
// Text segmentation is a technology that splits long text into short clips for adapting to model input, improving processing efficiency, or information retrieval. It balances clip length and semantic consistency, suitable for NLP and data analysis scenarios.
//
// This API is an async API with a model dimensional call limit. Each model has a qps limit of 5. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ChunkDocumentAsyncWithContext(ctx context.Context, request *ChunkDocumentAsyncRequest) (response *ChunkDocumentAsyncResponse, err error) {
    if request == nil {
        request = NewChunkDocumentAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "ChunkDocumentAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChunkDocumentAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewChunkDocumentAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewGetDocumentChunkResultRequest() (request *GetDocumentChunkResultRequest) {
    request = &GetDocumentChunkResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetDocumentChunkResult")
    
    
    return
}

func NewGetDocumentChunkResultResponse() (response *GetDocumentChunkResultResponse) {
    response = &GetDocumentChunkResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDocumentChunkResult
// Retrieve document slices
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentChunkResult(request *GetDocumentChunkResultRequest) (response *GetDocumentChunkResultResponse, err error) {
    return c.GetDocumentChunkResultWithContext(context.Background(), request)
}

// GetDocumentChunkResult
// Retrieve document slices
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentChunkResultWithContext(ctx context.Context, request *GetDocumentChunkResultRequest) (response *GetDocumentChunkResultResponse, err error) {
    if request == nil {
        request = NewGetDocumentChunkResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "GetDocumentChunkResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDocumentChunkResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDocumentChunkResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetDocumentParseResultRequest() (request *GetDocumentParseResultRequest) {
    request = &GetDocumentParseResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetDocumentParseResult")
    
    
    return
}

func NewGetDocumentParseResultResponse() (response *GetDocumentParseResultResponse) {
    response = &GetDocumentParseResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetDocumentParseResult
// This API is used to retrieve the asynchronous processing result of document parsing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentParseResult(request *GetDocumentParseResultRequest) (response *GetDocumentParseResultResponse, err error) {
    return c.GetDocumentParseResultWithContext(context.Background(), request)
}

// GetDocumentParseResult
// This API is used to retrieve the asynchronous processing result of document parsing.
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetDocumentParseResultWithContext(ctx context.Context, request *GetDocumentParseResultRequest) (response *GetDocumentParseResultResponse, err error) {
    if request == nil {
        request = NewGetDocumentParseResultRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "GetDocumentParseResult")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetDocumentParseResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetDocumentParseResultResponse()
    err = c.Send(request, response)
    return
}

func NewGetMultiModalEmbeddingRequest() (request *GetMultiModalEmbeddingRequest) {
    request = &GetMultiModalEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetMultiModalEmbedding")
    
    
    return
}

func NewGetMultiModalEmbeddingResponse() (response *GetMultiModalEmbeddingResponse) {
    response = &GetMultiModalEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetMultiModalEmbedding
// Embedding is a technology that maps high-dimensional data to a low-dimensional space, usually used for converting unstructured data such as text, images, or audio into vector representation, making it easier to input into machine models for processing, and the distance between vectors can reflect the similarity between objects. 
//
// This API has a model dimensional call limit. Each model has a qps limit of 10. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMultiModalEmbedding(request *GetMultiModalEmbeddingRequest) (response *GetMultiModalEmbeddingResponse, err error) {
    return c.GetMultiModalEmbeddingWithContext(context.Background(), request)
}

// GetMultiModalEmbedding
// Embedding is a technology that maps high-dimensional data to a low-dimensional space, usually used for converting unstructured data such as text, images, or audio into vector representation, making it easier to input into machine models for processing, and the distance between vectors can reflect the similarity between objects. 
//
// This API has a model dimensional call limit. Each model has a qps limit of 10. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetMultiModalEmbeddingWithContext(ctx context.Context, request *GetMultiModalEmbeddingRequest) (response *GetMultiModalEmbeddingResponse, err error) {
    if request == nil {
        request = NewGetMultiModalEmbeddingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "GetMultiModalEmbedding")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetMultiModalEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetMultiModalEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewGetTextEmbeddingRequest() (request *GetTextEmbeddingRequest) {
    request = &GetTextEmbeddingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetTextEmbedding")
    
    
    return
}

func NewGetTextEmbeddingResponse() (response *GetTextEmbeddingResponse) {
    response = &GetTextEmbeddingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTextEmbedding
// Embedding is a technology that maps high-dimensional data to a low-dimensional space, usually used for converting unstructured data such as text, images, or audio into vector representation, making it easier to input into machine models for processing, and the distance between vectors can reflect the similarity between objects.
//
// This API has a model dimensional call limit. Each model has a qps limit of 20. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTextEmbedding(request *GetTextEmbeddingRequest) (response *GetTextEmbeddingResponse, err error) {
    return c.GetTextEmbeddingWithContext(context.Background(), request)
}

// GetTextEmbedding
// Embedding is a technology that maps high-dimensional data to a low-dimensional space, usually used for converting unstructured data such as text, images, or audio into vector representation, making it easier to input into machine models for processing, and the distance between vectors can reflect the similarity between objects.
//
// This API has a model dimensional call limit. Each model has a qps limit of 20. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetTextEmbeddingWithContext(ctx context.Context, request *GetTextEmbeddingRequest) (response *GetTextEmbeddingResponse, err error) {
    if request == nil {
        request = NewGetTextEmbeddingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "GetTextEmbedding")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTextEmbedding require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTextEmbeddingResponse()
    err = c.Send(request, response)
    return
}

func NewParseDocumentRequest() (request *ParseDocumentRequest) {
    request = &ParseDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ParseDocument")
    
    
    return
}

func NewParseDocumentResponse() (response *ParseDocumentResponse) {
    response = &ParseDocumentResponse{} 
    return

}

// ParseDocument
// This service can accurately convert various types of documents into a standard format to meet the requirements for building an enterprise knowledge base, migrating technical documentation, and structured storage for content platforms.
//
// This API has a model dimensional call limit. Each model has a qps limit of 5. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ParseDocument(request *ParseDocumentRequest) (response *ParseDocumentResponse, err error) {
    return c.ParseDocumentWithContext(context.Background(), request)
}

// ParseDocument
// This service can accurately convert various types of documents into a standard format to meet the requirements for building an enterprise knowledge base, migrating technical documentation, and structured storage for content platforms.
//
// This API has a model dimensional call limit. Each model has a qps limit of 5. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ParseDocumentWithContext(ctx context.Context, request *ParseDocumentRequest) (response *ParseDocumentResponse, err error) {
    if request == nil {
        request = NewParseDocumentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "ParseDocument")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseDocument require credential")
    }

    request.SetContext(ctx)
    request.SetRootDomain("ai." + tchttp.RootDomain)
    
    response = NewParseDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewParseDocumentAsyncRequest() (request *ParseDocumentAsyncRequest) {
    request = &ParseDocumentAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "ParseDocumentAsync")
    
    
    return
}

func NewParseDocumentAsyncResponse() (response *ParseDocumentAsyncResponse) {
    response = &ParseDocumentAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ParseDocumentAsync
// This service accurately converts various format documents into standard format, meeting requirements for Enterprise Knowledge Base construction, technical documentation migration, and structured storage for content platforms.
//
// This API is an async API with a model dimensional call limit. Each model has a qps limit of 5. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ParseDocumentAsync(request *ParseDocumentAsyncRequest) (response *ParseDocumentAsyncResponse, err error) {
    return c.ParseDocumentAsyncWithContext(context.Background(), request)
}

// ParseDocumentAsync
// This service accurately converts various format documents into standard format, meeting requirements for Enterprise Knowledge Base construction, technical documentation migration, and structured storage for content platforms.
//
// This API is an async API with a model dimensional call limit. Each model has a qps limit of 5. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) ParseDocumentAsyncWithContext(ctx context.Context, request *ParseDocumentAsyncRequest) (response *ParseDocumentAsyncResponse, err error) {
    if request == nil {
        request = NewParseDocumentAsyncRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "ParseDocumentAsync")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ParseDocumentAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewParseDocumentAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewRunRerankRequest() (request *RunRerankRequest) {
    request = &RunRerankRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RunRerank")
    
    
    return
}

func NewRunRerankResponse() (response *RunRerankResponse) {
    response = &RunRerankResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RunRerank
// Rearrangement refers to the process in RAG where, by assessing the relevance between documents and queries, the most relevant documents are placed at the front. This ensures that the language model prioritizes high-ranking context when generating responses, improving the accuracy and reliability of generated results. It can also be used for filtering to reduce large model costs.
//
// This API has a single-account call limit. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RunRerank(request *RunRerankRequest) (response *RunRerankResponse, err error) {
    return c.RunRerankWithContext(context.Background(), request)
}

// RunRerank
// Rearrangement refers to the process in RAG where, by assessing the relevance between documents and queries, the most relevant documents are placed at the front. This ensures that the language model prioritizes high-ranking context when generating responses, improving the accuracy and reliability of generated results. It can also be used for filtering to reduce large model costs.
//
// This API has a single-account call limit. If you need to increase the concurrent limit, please contact us (https://www.tencentcloud.com/act/event/Online_service?from_cn_redirect=1).
//
// error code that may be returned:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RunRerankWithContext(ctx context.Context, request *RunRerankRequest) (response *RunRerankResponse, err error) {
    if request == nil {
        request = NewRunRerankRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "es", APIVersion, "RunRerank")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunRerank require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunRerankResponse()
    err = c.Send(request, response)
    return
}

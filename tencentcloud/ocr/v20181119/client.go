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

package v20181119

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
)

const APIVersion = "2018-11-19"

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


func NewBankCardOCRRequest() (request *BankCardOCRRequest) {
    request = &BankCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "BankCardOCR")
    
    
    return
}

func NewBankCardOCRResponse() (response *BankCardOCRResponse) {
    response = &BankCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BankCardOCR
// This API is used to detect and recognize key fields such as the card number, bank information, and expiration date on mainstream bank cards in Mainland China.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_ILLEGALBANKCARDERROR = "FailedOperation.IllegalBankCardError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBANKCARDERROR = "FailedOperation.NoBankCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BankCardOCR(request *BankCardOCRRequest) (response *BankCardOCRResponse, err error) {
    return c.BankCardOCRWithContext(context.Background(), request)
}

// BankCardOCR
// This API is used to detect and recognize key fields such as the card number, bank information, and expiration date on mainstream bank cards in Mainland China.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_ILLEGALBANKCARDERROR = "FailedOperation.IllegalBankCardError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOBANKCARDERROR = "FailedOperation.NoBankCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) BankCardOCRWithContext(ctx context.Context, request *BankCardOCRRequest) (response *BankCardOCRResponse, err error) {
    if request == nil {
        request = NewBankCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BankCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewBankCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewExtractDocMultiRequest() (request *ExtractDocMultiRequest) {
    request = &ExtractDocMultiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "ExtractDocMulti")
    
    
    return
}

func NewExtractDocMultiResponse() (response *ExtractDocMultiResponse) {
    response = &ExtractDocMultiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ExtractDocMulti
// This API supports identifying and extracting field information in structured scenarios such as complex scenarios and multiple formats. Key scenarios include: finance, health care, transportation, travel, insurance. Click [experience now](https://ocrdemo.cloud.tencent.com/).
//
// 
//
// This API is used to set the alias SmartStructuralPro.
//
// 
//
// The default API request rate limit is 5 requests per second.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ExtractDocMulti(request *ExtractDocMultiRequest) (response *ExtractDocMultiResponse, err error) {
    return c.ExtractDocMultiWithContext(context.Background(), request)
}

// ExtractDocMulti
// This API supports identifying and extracting field information in structured scenarios such as complex scenarios and multiple formats. Key scenarios include: finance, health care, transportation, travel, insurance. Click [experience now](https://ocrdemo.cloud.tencent.com/).
//
// 
//
// This API is used to set the alias SmartStructuralPro.
//
// 
//
// The default API request rate limit is 5 requests per second.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ExtractDocMultiWithContext(ctx context.Context, request *ExtractDocMultiRequest) (response *ExtractDocMultiResponse, err error) {
    if request == nil {
        request = NewExtractDocMultiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ExtractDocMulti require credential")
    }

    request.SetContext(ctx)
    
    response = NewExtractDocMultiResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralAccurateOCRRequest() (request *GeneralAccurateOCRRequest) {
    request = &GeneralAccurateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralAccurateOCR")
    
    
    return
}

func NewGeneralAccurateOCRResponse() (response *GeneralAccurateOCRResponse) {
    response = &GeneralAccurateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GeneralAccurateOCR
// This API is used to detect and recognize characters in an image. It can recognize Chinese, English, Chinese-English, digits, and special symbols and return the text box positions and characters.
//
// 
//
// It is suitable for scenarios with a lot of characters in complex layouts and requiring high recognition accuracy, such as examination papers, online images, signboards, and legal documents.
//
// 
//
// Strengths: compared with general print recognition, it provides higher-precision character recognition services. Its accuracy and recall rate are higher in difficult scenarios such as a large number of characters, long strings of digits, small characters, blurry characters, and tilted text.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralAccurateOCR(request *GeneralAccurateOCRRequest) (response *GeneralAccurateOCRResponse, err error) {
    return c.GeneralAccurateOCRWithContext(context.Background(), request)
}

// GeneralAccurateOCR
// This API is used to detect and recognize characters in an image. It can recognize Chinese, English, Chinese-English, digits, and special symbols and return the text box positions and characters.
//
// 
//
// It is suitable for scenarios with a lot of characters in complex layouts and requiring high recognition accuracy, such as examination papers, online images, signboards, and legal documents.
//
// 
//
// Strengths: compared with general print recognition, it provides higher-precision character recognition services. Its accuracy and recall rate are higher in difficult scenarios such as a large number of characters, long strings of digits, small characters, blurry characters, and tilted text.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralAccurateOCRWithContext(ctx context.Context, request *GeneralAccurateOCRRequest) (response *GeneralAccurateOCRResponse, err error) {
    if request == nil {
        request = NewGeneralAccurateOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralAccurateOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralAccurateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralBasicOCRRequest() (request *GeneralBasicOCRRequest) {
    request = &GeneralBasicOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralBasicOCR")
    
    
    return
}

func NewGeneralBasicOCRResponse() (response *GeneralBasicOCRResponse) {
    response = &GeneralBasicOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GeneralBasicOCR
// This API is used to detect and recognize characters in an image in the following 20 languages: Chinese, English, Japanese, Korean, Spanish, French, German, Portuguese, Vietnamese, Malay, Russian, Italian, Dutch, Swedish, Finnish, Danish, Norwegian, Hungarian, Thai, and Arabic. Mixed characters in English and each supported language can be recognized together.
//
// 
//
// It can recognize printed text in paper documents, online images, ads, signboards, menus, video titles, profile photos, etc.
//
// 
//
// Strengths: it can automatically recognize the text language, return the text box coordinate information, and automatically rotate tilted text to the upright direction.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralBasicOCR(request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
    return c.GeneralBasicOCRWithContext(context.Background(), request)
}

// GeneralBasicOCR
// This API is used to detect and recognize characters in an image in the following 20 languages: Chinese, English, Japanese, Korean, Spanish, French, German, Portuguese, Vietnamese, Malay, Russian, Italian, Dutch, Swedish, Finnish, Danish, Norwegian, Hungarian, Thai, and Arabic. Mixed characters in English and each supported language can be recognized together.
//
// 
//
// It can recognize printed text in paper documents, online images, ads, signboards, menus, video titles, profile photos, etc.
//
// 
//
// Strengths: it can automatically recognize the text language, return the text box coordinate information, and automatically rotate tilted text to the upright direction.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_ENGINERECOGNIZETIMEOUT = "FailedOperation.EngineRecognizeTimeout"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GeneralBasicOCRWithContext(ctx context.Context, request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
    if request == nil {
        request = NewGeneralBasicOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GeneralBasicOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewGeneralBasicOCRResponse()
    err = c.Send(request, response)
    return
}

func NewHKIDCardOCRRequest() (request *HKIDCardOCRRequest) {
    request = &HKIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "HKIDCardOCR")
    
    
    return
}

func NewHKIDCardOCRResponse() (response *HKIDCardOCRResponse) {
    response = &HKIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// HKIDCardOCR
// This API is used to recognize key fields on the photo side of a Hong Kong (China) identity card, including name in Chinese, name in English, telecode for name, date of birth, gender, document symbol, date of the first issue, date of the last receipt, identity card number, and permanent residency attribute. 
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOHKIDCARD = "FailedOperation.NoHKIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HKIDCardOCR(request *HKIDCardOCRRequest) (response *HKIDCardOCRResponse, err error) {
    return c.HKIDCardOCRWithContext(context.Background(), request)
}

// HKIDCardOCR
// This API is used to recognize key fields on the photo side of a Hong Kong (China) identity card, including name in Chinese, name in English, telecode for name, date of birth, gender, document symbol, date of the first issue, date of the last receipt, identity card number, and permanent residency attribute. 
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOHKIDCARD = "FailedOperation.NoHKIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HKIDCardOCRWithContext(ctx context.Context, request *HKIDCardOCRRequest) (response *HKIDCardOCRResponse, err error) {
    if request == nil {
        request = NewHKIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HKIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewHKIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewHmtResidentPermitOCRRequest() (request *HmtResidentPermitOCRRequest) {
    request = &HmtResidentPermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "HmtResidentPermitOCR")
    
    
    return
}

func NewHmtResidentPermitOCRResponse() (response *HmtResidentPermitOCRResponse) {
    response = &HmtResidentPermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// HmtResidentPermitOCR
// This API is used to recognize key fields on the front and back sides of a residence permit for Hong Kong, Macao, or Taiwan residents, including name, gender, date of birth, address, ID number, issuing authority, validity period, number of issues, and permit number. It can be used for residence permit OCR in scenarios such as bank account opening and user registration.
//
// 
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HmtResidentPermitOCR(request *HmtResidentPermitOCRRequest) (response *HmtResidentPermitOCRResponse, err error) {
    return c.HmtResidentPermitOCRWithContext(context.Background(), request)
}

// HmtResidentPermitOCR
// This API is used to recognize key fields on the front and back sides of a residence permit for Hong Kong, Macao, or Taiwan residents, including name, gender, date of birth, address, ID number, issuing authority, validity period, number of issues, and permit number. It can be used for residence permit OCR in scenarios such as bank account opening and user registration.
//
// 
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HmtResidentPermitOCRWithContext(ctx context.Context, request *HmtResidentPermitOCRRequest) (response *HmtResidentPermitOCRResponse, err error) {
    if request == nil {
        request = NewHmtResidentPermitOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("HmtResidentPermitOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewHmtResidentPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewIDCardOCRRequest() (request *IDCardOCRRequest) {
    request = &IDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "IDCardOCR")
    
    
    return
}

func NewIDCardOCRResponse() (response *IDCardOCRResponse) {
    response = &IDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IDCardOCR
// This API is used to recognize all fields on the front and back sides of a second-generation resident identity card for the Chinese mainland: name, gender, ethnicity, date of birth, domicile, identification number, issuing authority, and validity period, with a recognition accuracy of over 99%.
//
// 
//
// In addition, this API supports multiple value-added capabilities to meet the needs of different scenarios. It can crop ID card photos and profile photos, and provide warnings for nine cases, as detailed below.
//
// 
//
// <table style="width:650px">
//
//       <thead>
//
//         <tr>
//
//        <th width="150">Capability</th>
//
//           <th width="500">Description</th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td rowspan="2">Cropping</td>
//
//           <td>Crops the ID card photo (by removing extra edges outside the ID card and automatically correcting the shooting angle).</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Crops the profile photo (by automatically cutting out the face area in the ID card).</td>
//
//         </tr>
//
//         <tr>
//
//           <td rowspan="9">Warning</td>
//
//           <td>Warns about invalid ID card validity periods.</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Warns about  incomplete ID card borders.</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Warns about photocopied images.</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Warns about spoofed images.</td>
//
//         </tr>
//
//           <tr>
//
//           <td>Warns about border and frame occlusions.</td>
//
//         </tr>
//
//          <tr>
//
//           <td>Warns about temporary ID cards.</td>
//
//         </tr>
//
//           <tr>
//
//           <td>Warns about photoshopped images.</td>
//
//         </tr>
//
//           <tr>
//
//           <td>Warns about blurry ID card images (blurriness can be determined based on the image quality score).</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_CARDSIDEERROR = "FailedOperation.CardSideError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"
//  FAILEDOPERATION_IDCARDTOOSMALL = "FailedOperation.IdCardTooSmall"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) IDCardOCR(request *IDCardOCRRequest) (response *IDCardOCRResponse, err error) {
    return c.IDCardOCRWithContext(context.Background(), request)
}

// IDCardOCR
// This API is used to recognize all fields on the front and back sides of a second-generation resident identity card for the Chinese mainland: name, gender, ethnicity, date of birth, domicile, identification number, issuing authority, and validity period, with a recognition accuracy of over 99%.
//
// 
//
// In addition, this API supports multiple value-added capabilities to meet the needs of different scenarios. It can crop ID card photos and profile photos, and provide warnings for nine cases, as detailed below.
//
// 
//
// <table style="width:650px">
//
//       <thead>
//
//         <tr>
//
//        <th width="150">Capability</th>
//
//           <th width="500">Description</th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td rowspan="2">Cropping</td>
//
//           <td>Crops the ID card photo (by removing extra edges outside the ID card and automatically correcting the shooting angle).</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Crops the profile photo (by automatically cutting out the face area in the ID card).</td>
//
//         </tr>
//
//         <tr>
//
//           <td rowspan="9">Warning</td>
//
//           <td>Warns about invalid ID card validity periods.</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Warns about  incomplete ID card borders.</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Warns about photocopied images.</td>
//
//         </tr>
//
//         <tr>
//
//           <td>Warns about spoofed images.</td>
//
//         </tr>
//
//           <tr>
//
//           <td>Warns about border and frame occlusions.</td>
//
//         </tr>
//
//          <tr>
//
//           <td>Warns about temporary ID cards.</td>
//
//         </tr>
//
//           <tr>
//
//           <td>Warns about photoshopped images.</td>
//
//         </tr>
//
//           <tr>
//
//           <td>Warns about blurry ID card images (blurriness can be determined based on the image quality score).</td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// 
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_CARDSIDEERROR = "FailedOperation.CardSideError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"
//  FAILEDOPERATION_IDCARDTOOSMALL = "FailedOperation.IdCardTooSmall"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) IDCardOCRWithContext(ctx context.Context, request *IDCardOCRRequest) (response *IDCardOCRResponse, err error) {
    if request == nil {
        request = NewIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewLicensePlateOCRRequest() (request *LicensePlateOCRRequest) {
    request = &LicensePlateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "LicensePlateOCR")
    
    
    return
}

func NewLicensePlateOCRResponse() (response *LicensePlateOCRResponse) {
    response = &LicensePlateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// LicensePlateOCR
// This API is used to recognize a license plate attached to a motor vehicle in the Chinese mainland and return the regional code, license plate number, and license plate color.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) LicensePlateOCR(request *LicensePlateOCRRequest) (response *LicensePlateOCRResponse, err error) {
    return c.LicensePlateOCRWithContext(context.Background(), request)
}

// LicensePlateOCR
// This API is used to recognize a license plate attached to a motor vehicle in the Chinese mainland and return the regional code, license plate number, and license plate color.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) LicensePlateOCRWithContext(ctx context.Context, request *LicensePlateOCRRequest) (response *LicensePlateOCRResponse, err error) {
    if request == nil {
        request = NewLicensePlateOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("LicensePlateOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewLicensePlateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMLIDCardOCRRequest() (request *MLIDCardOCRRequest) {
    request = &MLIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MLIDCardOCR")
    
    
    return
}

func NewMLIDCardOCRResponse() (response *MLIDCardOCRResponse) {
    response = &MLIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MLIDCardOCR
// This API is used to recognize a Malaysian identity card, including identity card number, name, gender, and address. It is also used to crop identity photos and give alarms for photographed or photocopied certificates.
//
// 
//
// This API is not fully available for the time being. For more information, contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOMASIDCARD = "FailedOperation.NoMASIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDCardOCR(request *MLIDCardOCRRequest) (response *MLIDCardOCRResponse, err error) {
    return c.MLIDCardOCRWithContext(context.Background(), request)
}

// MLIDCardOCR
// This API is used to recognize a Malaysian identity card, including identity card number, name, gender, and address. It is also used to crop identity photos and give alarms for photographed or photocopied certificates.
//
// 
//
// This API is not fully available for the time being. For more information, contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOMASIDCARD = "FailedOperation.NoMASIDCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDCardOCRWithContext(ctx context.Context, request *MLIDCardOCRRequest) (response *MLIDCardOCRResponse, err error) {
    if request == nil {
        request = NewMLIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MLIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMLIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMLIDPassportOCRRequest() (request *MLIDPassportOCRRequest) {
    request = &MLIDPassportOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MLIDPassportOCR")
    
    
    return
}

func NewMLIDPassportOCRResponse() (response *MLIDPassportOCRResponse) {
    response = &MLIDPassportOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MLIDPassportOCR
// This API is used to recognize a passport issued in Hong Kong/Macao/Taiwan (China) or other countries/regions. Recognizable fields include passport ID, name, date of birth, gender, expiration date, issuing country/region, and nationality. It has the features of cropping identity photos and alarming for photographed or photocopied documents.
//
// This interface supports regional scope: countries with machine-readable passports
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FIELDEXCEPTION = "FailedOperation.FieldException"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_INCONSISTENCYBETWEENMRZANDVRZ = "FailedOperation.InconsistencyBetweenMRZAndVRZ"
//  FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDPassportOCR(request *MLIDPassportOCRRequest) (response *MLIDPassportOCRResponse, err error) {
    return c.MLIDPassportOCRWithContext(context.Background(), request)
}

// MLIDPassportOCR
// This API is used to recognize a passport issued in Hong Kong/Macao/Taiwan (China) or other countries/regions. Recognizable fields include passport ID, name, date of birth, gender, expiration date, issuing country/region, and nationality. It has the features of cropping identity photos and alarming for photographed or photocopied documents.
//
// This interface supports regional scope: countries with machine-readable passports
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_FIELDEXCEPTION = "FailedOperation.FieldException"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_INCONSISTENCYBETWEENMRZANDVRZ = "FailedOperation.InconsistencyBetweenMRZAndVRZ"
//  FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDPassportOCRWithContext(ctx context.Context, request *MLIDPassportOCRRequest) (response *MLIDPassportOCRResponse, err error) {
    if request == nil {
        request = NewMLIDPassportOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MLIDPassportOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMLIDPassportOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMainlandPermitOCRRequest() (request *MainlandPermitOCRRequest) {
    request = &MainlandPermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "MainlandPermitOCR")
    
    
    return
}

func NewMainlandPermitOCRResponse() (response *MainlandPermitOCRResponse) {
    response = &MainlandPermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// MainlandPermitOCR
// This API is used to recognize all fields on the front of a mainland travel permit for Hong Kong, Macao, or Taiwan residents: name in Chinese, name in English, gender, date of birth, issuing authority, validity period, document number, place of issuance, number of issues, and document type.
//
// 
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MainlandPermitOCR(request *MainlandPermitOCRRequest) (response *MainlandPermitOCRResponse, err error) {
    return c.MainlandPermitOCRWithContext(context.Background(), request)
}

// MainlandPermitOCR
// This API is used to recognize all fields on the front of a mainland travel permit for Hong Kong, Macao, or Taiwan residents: name in Chinese, name in English, gender, date of birth, issuing authority, validity period, document number, place of issuance, number of issues, and document type.
//
// 
//
// A maximum of 20 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MainlandPermitOCRWithContext(ctx context.Context, request *MainlandPermitOCRRequest) (response *MainlandPermitOCRResponse, err error) {
    if request == nil {
        request = NewMainlandPermitOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("MainlandPermitOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewMainlandPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPermitOCRRequest() (request *PermitOCRRequest) {
    request = &PermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "PermitOCR")
    
    
    return
}

func NewPermitOCRResponse() (response *PermitOCRResponse) {
    response = &PermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PermitOCR
// This API is used to recognize the fields on an exit/entry permit (card) for traveling to and from Hong Kong, Macao, or Taiwan, including place of issuance, issuing authority, validity period, gender, date of birth, name in English, name in Chinese, and document number.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PermitOCR(request *PermitOCRRequest) (response *PermitOCRResponse, err error) {
    return c.PermitOCRWithContext(context.Background(), request)
}

// PermitOCR
// This API is used to recognize the fields on an exit/entry permit (card) for traveling to and from Hong Kong, Macao, or Taiwan, including place of issuance, issuing authority, validity period, gender, date of birth, name in English, name in Chinese, and document number.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) PermitOCRWithContext(ctx context.Context, request *PermitOCRRequest) (response *PermitOCRResponse, err error) {
    if request == nil {
        request = NewPermitOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PermitOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeBrazilDriverLicenseOCRRequest() (request *RecognizeBrazilDriverLicenseOCRRequest) {
    request = &RecognizeBrazilDriverLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeBrazilDriverLicenseOCR")
    
    
    return
}

func NewRecognizeBrazilDriverLicenseOCRResponse() (response *RecognizeBrazilDriverLicenseOCRResponse) {
    response = &RecognizeBrazilDriverLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeBrazilDriverLicenseOCR
// This interface supports identification of the front and back of Brazilian driver's license. The identification fields include name, driver's license category, number, validity period, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RecognizeBrazilDriverLicenseOCR(request *RecognizeBrazilDriverLicenseOCRRequest) (response *RecognizeBrazilDriverLicenseOCRResponse, err error) {
    return c.RecognizeBrazilDriverLicenseOCRWithContext(context.Background(), request)
}

// RecognizeBrazilDriverLicenseOCR
// This interface supports identification of the front and back of Brazilian driver's license. The identification fields include name, driver's license category, number, validity period, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RecognizeBrazilDriverLicenseOCRWithContext(ctx context.Context, request *RecognizeBrazilDriverLicenseOCRRequest) (response *RecognizeBrazilDriverLicenseOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeBrazilDriverLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeBrazilDriverLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeBrazilDriverLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeBrazilIDCardOCRRequest() (request *RecognizeBrazilIDCardOCRRequest) {
    request = &RecognizeBrazilIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeBrazilIDCardOCR")
    
    
    return
}

func NewRecognizeBrazilIDCardOCRResponse() (response *RecognizeBrazilIDCardOCRResponse) {
    response = &RecognizeBrazilIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeBrazilIDCardOCR
// This interface supports identification of the front and back of Brazilian ID license. The identification fields include name, driver's license category, number, validity period, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeBrazilIDCardOCR(request *RecognizeBrazilIDCardOCRRequest) (response *RecognizeBrazilIDCardOCRResponse, err error) {
    return c.RecognizeBrazilIDCardOCRWithContext(context.Background(), request)
}

// RecognizeBrazilIDCardOCR
// This interface supports identification of the front and back of Brazilian ID license. The identification fields include name, driver's license category, number, validity period, etc.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeBrazilIDCardOCRWithContext(ctx context.Context, request *RecognizeBrazilIDCardOCRRequest) (response *RecognizeBrazilIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeBrazilIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeBrazilIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeBrazilIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeBrazilRNEOCRRequest() (request *RecognizeBrazilRNEOCRRequest) {
    request = &RecognizeBrazilRNEOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeBrazilRNEOCR")
    
    
    return
}

func NewRecognizeBrazilRNEOCRResponse() (response *RecognizeBrazilRNEOCRResponse) {
    response = &RecognizeBrazilRNEOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeBrazilRNEOCR
// Brazil RNE document recognition Default interface request frequency limit: 5 times/second
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeBrazilRNEOCR(request *RecognizeBrazilRNEOCRRequest) (response *RecognizeBrazilRNEOCRResponse, err error) {
    return c.RecognizeBrazilRNEOCRWithContext(context.Background(), request)
}

// RecognizeBrazilRNEOCR
// Brazil RNE document recognition Default interface request frequency limit: 5 times/second
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeBrazilRNEOCRWithContext(ctx context.Context, request *RecognizeBrazilRNEOCRRequest) (response *RecognizeBrazilRNEOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeBrazilRNEOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeBrazilRNEOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeBrazilRNEOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeBrazilRNMOCRRequest() (request *RecognizeBrazilRNMOCRRequest) {
    request = &RecognizeBrazilRNMOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeBrazilRNMOCR")
    
    
    return
}

func NewRecognizeBrazilRNMOCRResponse() (response *RecognizeBrazilRNMOCRResponse) {
    response = &RecognizeBrazilRNMOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeBrazilRNMOCR
// This interface supports identification of the front and back of Brazilian RNM license. The default interface request frequency limit is 5 times per second.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeBrazilRNMOCR(request *RecognizeBrazilRNMOCRRequest) (response *RecognizeBrazilRNMOCRResponse, err error) {
    return c.RecognizeBrazilRNMOCRWithContext(context.Background(), request)
}

// RecognizeBrazilRNMOCR
// This interface supports identification of the front and back of Brazilian RNM license. The default interface request frequency limit is 5 times per second.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeBrazilRNMOCRWithContext(ctx context.Context, request *RecognizeBrazilRNMOCRRequest) (response *RecognizeBrazilRNMOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeBrazilRNMOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeBrazilRNMOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeBrazilRNMOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeGeneralInvoiceRequest() (request *RecognizeGeneralInvoiceRequest) {
    request = &RecognizeGeneralInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeGeneralInvoice")
    
    
    return
}

func NewRecognizeGeneralInvoiceResponse() (response *RecognizeGeneralInvoiceResponse) {
    response = &RecognizeGeneralInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeGeneralInvoice
// This API is used to recognize various types of invoices or tickets in an image or PDF file. You can also specify a type. 14 types of standard expense reimbursement invoices are supported, including value-added tax (VAT) invoice (special, general, roll, blockchain, and toll), fully digitalized electronic invoice (special and general), non-tax revenue invoice (general receipt and general payment voucher), quota invoice, general machine-printed invoice, car sales invoice (motor vehicle sales invoice and used car invoice), train ticket, taxi receipt, itinerary/receipt of e-ticket for air transportation, bus ticket, ship ticket, toll receipt, and medical invoice (inpatient and outpatient). This API can also be used for intelligent recognition of other types of invoices. To try now, click [here](https://intl.cloud.tencent.com/product/ocr?from_cn_redirect=1).
//
// 
//
// A maximum of 5 requests can be initiated per second for this API.
//
// 
//
// 
//
// The invoice/ticket subtype (SubType), subtype description (TypeDescription), and parent type (Type) can be returned, as described below:
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:200px">SubType</th>
//
//           <th style="width:200px">TypeDescription</th>
//
//           <th >Type</th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> VatSpecialInvoice</td>
//
//           <td> Special VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatCommonInvoice</td>
//
//           <td> General VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicCommonInvoice </td>
//
//           <td> Electronic general VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicSpecialInvoice </td>
//
//           <td> Electronic special VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicInvoiceBlockchain</td>
//
//           <td> Blockchain electronic invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicInvoiceToll</td>
//
//           <td> Electronic general VAT invoice (toll)</td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicSpecialInvoiceFull</td>
//
//           <td> Electronic invoice (special)</td>
//
//           <td> 16 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicInvoiceFull</td>
//
//           <td> Electronic invoice (general) </td>
//
//           <td> 16 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> MotorVehicleSaleInvoice </td>
//
//           <td> Motor vehicle sales invoice </td>
//
//           <td> 12 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> UsedCarPurchaseInvoice </td>
//
//           <td> Used car invoice </td>
//
//           <td> 12 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatInvoiceRoll </td>
//
//           <td> General VAT invoice (roll) </td>
//
//           <td> 11 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> TaxiTicket </td>
//
//           <td> Taxi receipt </td>
//
//           <td> 0 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> QuotaInvoice </td>
//
//           <td> Quota invoice </td>
//
//           <td> 1 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> TrainTicket </td>
//
//           <td> Train ticket </td>
//
//           <td> 2 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> AirTransport </td>
//
//           <td> Itinerary/Receipt of e-ticket for air transportation </td>
//
//           <td> 5 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> MachinePrintedInvoice </td>
//
//           <td> General machine-printed invoice </td>
//
//           <td> 8 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> BusInvoice </td>
//
//           <td> Bus ticket </td>
//
//           <td> 9 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> ShippingInvoice </td>
//
//           <td> Ship ticket </td>
//
//           <td> 10 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> NonTaxIncomeGeneralBill </td>
//
//           <td> General receipt for non-tax revenue </td>
//
//           <td> 15 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> NonTaxIncomeElectronicBill </td>
//
//           <td> General payment voucher for non-tax revenue (electronic) </td>
//
//           <td> 15 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> TollInvoice </td>
//
//           <td> Toll receipt </td>
//
//           <td> 13 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> OtherInvoice </td>
//
//           <td> Other </td>
//
//           <td> -1 </td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeGeneralInvoice(request *RecognizeGeneralInvoiceRequest) (response *RecognizeGeneralInvoiceResponse, err error) {
    return c.RecognizeGeneralInvoiceWithContext(context.Background(), request)
}

// RecognizeGeneralInvoice
// This API is used to recognize various types of invoices or tickets in an image or PDF file. You can also specify a type. 14 types of standard expense reimbursement invoices are supported, including value-added tax (VAT) invoice (special, general, roll, blockchain, and toll), fully digitalized electronic invoice (special and general), non-tax revenue invoice (general receipt and general payment voucher), quota invoice, general machine-printed invoice, car sales invoice (motor vehicle sales invoice and used car invoice), train ticket, taxi receipt, itinerary/receipt of e-ticket for air transportation, bus ticket, ship ticket, toll receipt, and medical invoice (inpatient and outpatient). This API can also be used for intelligent recognition of other types of invoices. To try now, click [here](https://intl.cloud.tencent.com/product/ocr?from_cn_redirect=1).
//
// 
//
// A maximum of 5 requests can be initiated per second for this API.
//
// 
//
// 
//
// The invoice/ticket subtype (SubType), subtype description (TypeDescription), and parent type (Type) can be returned, as described below:
//
// <table style="width:715px">
//
//       <thead>
//
//         <tr>
//
//           <th style="width:200px">SubType</th>
//
//           <th style="width:200px">TypeDescription</th>
//
//           <th >Type</th>
//
//         </tr>
//
//       </thead>
//
//       <tbody>
//
//         <tr>
//
//           <td> VatSpecialInvoice</td>
//
//           <td> Special VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatCommonInvoice</td>
//
//           <td> General VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicCommonInvoice </td>
//
//           <td> Electronic general VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicSpecialInvoice </td>
//
//           <td> Electronic special VAT invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicInvoiceBlockchain</td>
//
//           <td> Blockchain electronic invoice </td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicInvoiceToll</td>
//
//           <td> Electronic general VAT invoice (toll)</td>
//
//           <td> 3 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicSpecialInvoiceFull</td>
//
//           <td> Electronic invoice (special)</td>
//
//           <td> 16 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatElectronicInvoiceFull</td>
//
//           <td> Electronic invoice (general) </td>
//
//           <td> 16 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> MotorVehicleSaleInvoice </td>
//
//           <td> Motor vehicle sales invoice </td>
//
//           <td> 12 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> UsedCarPurchaseInvoice </td>
//
//           <td> Used car invoice </td>
//
//           <td> 12 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> VatInvoiceRoll </td>
//
//           <td> General VAT invoice (roll) </td>
//
//           <td> 11 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> TaxiTicket </td>
//
//           <td> Taxi receipt </td>
//
//           <td> 0 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> QuotaInvoice </td>
//
//           <td> Quota invoice </td>
//
//           <td> 1 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> TrainTicket </td>
//
//           <td> Train ticket </td>
//
//           <td> 2 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> AirTransport </td>
//
//           <td> Itinerary/Receipt of e-ticket for air transportation </td>
//
//           <td> 5 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> MachinePrintedInvoice </td>
//
//           <td> General machine-printed invoice </td>
//
//           <td> 8 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> BusInvoice </td>
//
//           <td> Bus ticket </td>
//
//           <td> 9 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> ShippingInvoice </td>
//
//           <td> Ship ticket </td>
//
//           <td> 10 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> NonTaxIncomeGeneralBill </td>
//
//           <td> General receipt for non-tax revenue </td>
//
//           <td> 15 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> NonTaxIncomeElectronicBill </td>
//
//           <td> General payment voucher for non-tax revenue (electronic) </td>
//
//           <td> 15 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> TollInvoice </td>
//
//           <td> Toll receipt </td>
//
//           <td> 13 </td>
//
//         </tr>
//
//         <tr>
//
//           <td> OtherInvoice </td>
//
//           <td> Other </td>
//
//           <td> -1 </td>
//
//         </tr>
//
//       </tbody>
//
//     </table>
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeGeneralInvoiceWithContext(ctx context.Context, request *RecognizeGeneralInvoiceRequest) (response *RecognizeGeneralInvoiceResponse, err error) {
    if request == nil {
        request = NewRecognizeGeneralInvoiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeGeneralInvoice require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeGeneralInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeIndonesiaIDCardOCRRequest() (request *RecognizeIndonesiaIDCardOCRRequest) {
    request = &RecognizeIndonesiaIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeIndonesiaIDCardOCR")
    
    
    return
}

func NewRecognizeIndonesiaIDCardOCRResponse() (response *RecognizeIndonesiaIDCardOCRResponse) {
    response = &RecognizeIndonesiaIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeIndonesiaIDCardOCR
// This API is used to recognize an Indonesian identity card.
//
// 
//
// The API request rate is limited to 20 requests/sec by default.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeIndonesiaIDCardOCR(request *RecognizeIndonesiaIDCardOCRRequest) (response *RecognizeIndonesiaIDCardOCRResponse, err error) {
    return c.RecognizeIndonesiaIDCardOCRWithContext(context.Background(), request)
}

// RecognizeIndonesiaIDCardOCR
// This API is used to recognize an Indonesian identity card.
//
// 
//
// The API request rate is limited to 20 requests/sec by default.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeIndonesiaIDCardOCRWithContext(ctx context.Context, request *RecognizeIndonesiaIDCardOCRRequest) (response *RecognizeIndonesiaIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeIndonesiaIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeIndonesiaIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeIndonesiaIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeKoreanDrivingLicenseOCRRequest() (request *RecognizeKoreanDrivingLicenseOCRRequest) {
    request = &RecognizeKoreanDrivingLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeKoreanDrivingLicenseOCR")
    
    
    return
}

func NewRecognizeKoreanDrivingLicenseOCRResponse() (response *RecognizeKoreanDrivingLicenseOCRResponse) {
    response = &RecognizeKoreanDrivingLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeKoreanDrivingLicenseOCR
// This API is used to recognize a South Korean driver's license.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeKoreanDrivingLicenseOCR(request *RecognizeKoreanDrivingLicenseOCRRequest) (response *RecognizeKoreanDrivingLicenseOCRResponse, err error) {
    return c.RecognizeKoreanDrivingLicenseOCRWithContext(context.Background(), request)
}

// RecognizeKoreanDrivingLicenseOCR
// This API is used to recognize a South Korean driver's license.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeKoreanDrivingLicenseOCRWithContext(ctx context.Context, request *RecognizeKoreanDrivingLicenseOCRRequest) (response *RecognizeKoreanDrivingLicenseOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeKoreanDrivingLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeKoreanDrivingLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeKoreanDrivingLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeKoreanIDCardOCRRequest() (request *RecognizeKoreanIDCardOCRRequest) {
    request = &RecognizeKoreanIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeKoreanIDCardOCR")
    
    
    return
}

func NewRecognizeKoreanIDCardOCRResponse() (response *RecognizeKoreanIDCardOCRResponse) {
    response = &RecognizeKoreanIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeKoreanIDCardOCR
// This API is used to recognize a South Korean ID card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeKoreanIDCardOCR(request *RecognizeKoreanIDCardOCRRequest) (response *RecognizeKoreanIDCardOCRResponse, err error) {
    return c.RecognizeKoreanIDCardOCRWithContext(context.Background(), request)
}

// RecognizeKoreanIDCardOCR
// This API is used to recognize a South Korean ID card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeKoreanIDCardOCRWithContext(ctx context.Context, request *RecognizeKoreanIDCardOCRRequest) (response *RecognizeKoreanIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeKoreanIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeKoreanIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeKoreanIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeMacaoIDCardOCRRequest() (request *RecognizeMacaoIDCardOCRRequest) {
    request = &RecognizeMacaoIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeMacaoIDCardOCR")
    
    
    return
}

func NewRecognizeMacaoIDCardOCRResponse() (response *RecognizeMacaoIDCardOCRResponse) {
    response = &RecognizeMacaoIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeMacaoIDCardOCR
// This API is used to recognize key fields on the photo side of a Macao (China) identity card, including name in Chinese, name in English, telecode for name, date of birth, gender, document symbol, date of the first issue, date of the last receipt, identity card number, and permanent residency attribute. 
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
func (c *Client) RecognizeMacaoIDCardOCR(request *RecognizeMacaoIDCardOCRRequest) (response *RecognizeMacaoIDCardOCRResponse, err error) {
    return c.RecognizeMacaoIDCardOCRWithContext(context.Background(), request)
}

// RecognizeMacaoIDCardOCR
// This API is used to recognize key fields on the photo side of a Macao (China) identity card, including name in Chinese, name in English, telecode for name, date of birth, gender, document symbol, date of the first issue, date of the last receipt, identity card number, and permanent residency attribute. 
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
func (c *Client) RecognizeMacaoIDCardOCRWithContext(ctx context.Context, request *RecognizeMacaoIDCardOCRRequest) (response *RecognizeMacaoIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeMacaoIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeMacaoIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeMacaoIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeMainlandIDCardOCRRequest() (request *RecognizeMainlandIDCardOCRRequest) {
    request = &RecognizeMainlandIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeMainlandIDCardOCR")
    
    
    return
}

func NewRecognizeMainlandIDCardOCRResponse() (response *RecognizeMainlandIDCardOCRResponse) {
    response = &RecognizeMainlandIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeMainlandIDCardOCR
// This interface supports the identification of all fields on the front and back of the second-generation ID card for mainland Chinese residents.Including name, gender, ethnicity, date of birth, address, citizen ID number, issuing authority, and validity period, the identification accuracy reaches more than 99%.In addition, this interface also supports a variety of value-added capabilities to meet the needs of different scenarios. Such as the cropping function of ID card photos and portrait photos, and also has 5 alarm functions.
//
// As shown in the table below. <table style="width:650px"> <thead> <tr> <th width="150">Value-added ability</th> <th width="500">Ability items</th> </tr> </thead> <tbody> <tr> <td rowspan="9">Alarm function</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>Alarm for occlusion in the ID card frame</td> </tr> <tr> <td>ID card reflective warning</td> </tr> <tr> <td>Blurry picture warning</td> </tr> </tbody> </table> Default interface request frequency limit: 20 times/second
//
// error code that may be returned:
//  FAILEDOPERATION_CARDSIDEERROR = "FailedOperation.CardSideError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"
//  FAILEDOPERATION_IDCARDTOOSMALL = "FailedOperation.IdCardTooSmall"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeMainlandIDCardOCR(request *RecognizeMainlandIDCardOCRRequest) (response *RecognizeMainlandIDCardOCRResponse, err error) {
    return c.RecognizeMainlandIDCardOCRWithContext(context.Background(), request)
}

// RecognizeMainlandIDCardOCR
// This interface supports the identification of all fields on the front and back of the second-generation ID card for mainland Chinese residents.Including name, gender, ethnicity, date of birth, address, citizen ID number, issuing authority, and validity period, the identification accuracy reaches more than 99%.In addition, this interface also supports a variety of value-added capabilities to meet the needs of different scenarios. Such as the cropping function of ID card photos and portrait photos, and also has 5 alarm functions.
//
// As shown in the table below. <table style="width:650px"> <thead> <tr> <th width="150">Value-added ability</th> <th width="500">Ability items</th> </tr> </thead> <tbody> <tr> <td rowspan="9">Alarm function</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>Alarm for occlusion in the ID card frame</td> </tr> <tr> <td>ID card reflective warning</td> </tr> <tr> <td>Blurry picture warning</td> </tr> </tbody> </table> Default interface request frequency limit: 20 times/second
//
// error code that may be returned:
//  FAILEDOPERATION_CARDSIDEERROR = "FailedOperation.CardSideError"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IDCARDINFOILLEGAL = "FailedOperation.IdCardInfoIllegal"
//  FAILEDOPERATION_IDCARDTOOSMALL = "FailedOperation.IdCardTooSmall"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOIDCARD = "FailedOperation.ImageNoIdCard"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_MULTICARDERROR = "FailedOperation.MultiCardError"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_CONFIGFORMATERROR = "InvalidParameter.ConfigFormatError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeMainlandIDCardOCRWithContext(ctx context.Context, request *RecognizeMainlandIDCardOCRRequest) (response *RecognizeMainlandIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeMainlandIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeMainlandIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeMainlandIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeMexicoVTIDRequest() (request *RecognizeMexicoVTIDRequest) {
    request = &RecognizeMexicoVTIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeMexicoVTID")
    
    
    return
}

func NewRecognizeMexicoVTIDResponse() (response *RecognizeMexicoVTIDResponse) {
    response = &RecognizeMexicoVTIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeMexicoVTID
// This interface supports identification of the front and back of Mexican Voter ID Card. The default interface request frequency limit is 5 times per second.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeMexicoVTID(request *RecognizeMexicoVTIDRequest) (response *RecognizeMexicoVTIDResponse, err error) {
    return c.RecognizeMexicoVTIDWithContext(context.Background(), request)
}

// RecognizeMexicoVTID
// This interface supports identification of the front and back of Mexican Voter ID Card. The default interface request frequency limit is 5 times per second.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
func (c *Client) RecognizeMexicoVTIDWithContext(ctx context.Context, request *RecognizeMexicoVTIDRequest) (response *RecognizeMexicoVTIDResponse, err error) {
    if request == nil {
        request = NewRecognizeMexicoVTIDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeMexicoVTID require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeMexicoVTIDResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesDrivingLicenseOCRRequest() (request *RecognizePhilippinesDrivingLicenseOCRRequest) {
    request = &RecognizePhilippinesDrivingLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesDrivingLicenseOCR")
    
    
    return
}

func NewRecognizePhilippinesDrivingLicenseOCRResponse() (response *RecognizePhilippinesDrivingLicenseOCRResponse) {
    response = &RecognizePhilippinesDrivingLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizePhilippinesDrivingLicenseOCR
// This API is used to recognize a Philippine driver's license.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesDrivingLicenseOCR(request *RecognizePhilippinesDrivingLicenseOCRRequest) (response *RecognizePhilippinesDrivingLicenseOCRResponse, err error) {
    return c.RecognizePhilippinesDrivingLicenseOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesDrivingLicenseOCR
// This API is used to recognize a Philippine driver's license.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesDrivingLicenseOCRWithContext(ctx context.Context, request *RecognizePhilippinesDrivingLicenseOCRRequest) (response *RecognizePhilippinesDrivingLicenseOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesDrivingLicenseOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesDrivingLicenseOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesDrivingLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesSssIDOCRRequest() (request *RecognizePhilippinesSssIDOCRRequest) {
    request = &RecognizePhilippinesSssIDOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesSssIDOCR")
    
    
    return
}

func NewRecognizePhilippinesSssIDOCRResponse() (response *RecognizePhilippinesSssIDOCRResponse) {
    response = &RecognizePhilippinesSssIDOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizePhilippinesSssIDOCR
// This API is used to recognize a Philippine SSSID/UMID card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesSssIDOCR(request *RecognizePhilippinesSssIDOCRRequest) (response *RecognizePhilippinesSssIDOCRResponse, err error) {
    return c.RecognizePhilippinesSssIDOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesSssIDOCR
// This API is used to recognize a Philippine SSSID/UMID card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesSssIDOCRWithContext(ctx context.Context, request *RecognizePhilippinesSssIDOCRRequest) (response *RecognizePhilippinesSssIDOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesSssIDOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesSssIDOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesSssIDOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesTinIDOCRRequest() (request *RecognizePhilippinesTinIDOCRRequest) {
    request = &RecognizePhilippinesTinIDOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesTinIDOCR")
    
    
    return
}

func NewRecognizePhilippinesTinIDOCRResponse() (response *RecognizePhilippinesTinIDOCRResponse) {
    response = &RecognizePhilippinesTinIDOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizePhilippinesTinIDOCR
// This API is used to recognize a Philippine TIN ID card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesTinIDOCR(request *RecognizePhilippinesTinIDOCRRequest) (response *RecognizePhilippinesTinIDOCRResponse, err error) {
    return c.RecognizePhilippinesTinIDOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesTinIDOCR
// This API is used to recognize a Philippine TIN ID card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesTinIDOCRWithContext(ctx context.Context, request *RecognizePhilippinesTinIDOCRRequest) (response *RecognizePhilippinesTinIDOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesTinIDOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesTinIDOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesTinIDOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesUMIDOCRRequest() (request *RecognizePhilippinesUMIDOCRRequest) {
    request = &RecognizePhilippinesUMIDOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesUMIDOCR")
    
    
    return
}

func NewRecognizePhilippinesUMIDOCRResponse() (response *RecognizePhilippinesUMIDOCRResponse) {
    response = &RecognizePhilippinesUMIDOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizePhilippinesUMIDOCR
// This API is used to recognize a Philippine Unified Multi-Purpose ID (UMID) card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesUMIDOCR(request *RecognizePhilippinesUMIDOCRRequest) (response *RecognizePhilippinesUMIDOCRResponse, err error) {
    return c.RecognizePhilippinesUMIDOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesUMIDOCR
// This API is used to recognize a Philippine Unified Multi-Purpose ID (UMID) card.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesUMIDOCRWithContext(ctx context.Context, request *RecognizePhilippinesUMIDOCRRequest) (response *RecognizePhilippinesUMIDOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesUMIDOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesUMIDOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesUMIDOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizePhilippinesVoteIDOCRRequest() (request *RecognizePhilippinesVoteIDOCRRequest) {
    request = &RecognizePhilippinesVoteIDOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizePhilippinesVoteIDOCR")
    
    
    return
}

func NewRecognizePhilippinesVoteIDOCRResponse() (response *RecognizePhilippinesVoteIDOCRResponse) {
    response = &RecognizePhilippinesVoteIDOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizePhilippinesVoteIDOCR
// This API is used to recognize a Philippine voters ID card. It can recognize fields such as first name, family name, date of birth, civil status, citizenship, address, precinct, and voter's identification number (VIN).
//
// 
//
// The API request rate is limited to 20 requests/sec by default.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesVoteIDOCR(request *RecognizePhilippinesVoteIDOCRRequest) (response *RecognizePhilippinesVoteIDOCRResponse, err error) {
    return c.RecognizePhilippinesVoteIDOCRWithContext(context.Background(), request)
}

// RecognizePhilippinesVoteIDOCR
// This API is used to recognize a Philippine voters ID card. It can recognize fields such as first name, family name, date of birth, civil status, citizenship, address, precinct, and voter's identification number (VIN).
//
// 
//
// The API request rate is limited to 20 requests/sec by default.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizePhilippinesVoteIDOCRWithContext(ctx context.Context, request *RecognizePhilippinesVoteIDOCRRequest) (response *RecognizePhilippinesVoteIDOCRResponse, err error) {
    if request == nil {
        request = NewRecognizePhilippinesVoteIDOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizePhilippinesVoteIDOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizePhilippinesVoteIDOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeSingaporeIDCardOCRRequest() (request *RecognizeSingaporeIDCardOCRRequest) {
    request = &RecognizeSingaporeIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeSingaporeIDCardOCR")
    
    
    return
}

func NewRecognizeSingaporeIDCardOCRResponse() (response *RecognizeSingaporeIDCardOCRResponse) {
    response = &RecognizeSingaporeIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeSingaporeIDCardOCR
// This interface supports the identification of all fields on the front side of ID card for Singapore residents.The identification accuracy reaches more than 99%.In addition, this interface also supports a variety of value-added capabilities to meet the needs of different scenarios. Such as the cropping function of ID card photos and portrait photos, and also has 5 alarm functions.
//
// As shown in the table below. <table style="width:650px"> <thead> <tr> <th width="150">Value-added ability</th> <th width="500">Ability items</th> </tr> </thead> <tbody> <tr> <td rowspan="9">Alarm function</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>Alarm for occlusion in the ID card frame</td> </tr> <tr> <td>ID card reflective warning</td> </tr> <tr> <td>Blurry picture warning</td> </tr> </tbody> </table> Default interface request frequency limit: 20 times/second
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeSingaporeIDCardOCR(request *RecognizeSingaporeIDCardOCRRequest) (response *RecognizeSingaporeIDCardOCRResponse, err error) {
    return c.RecognizeSingaporeIDCardOCRWithContext(context.Background(), request)
}

// RecognizeSingaporeIDCardOCR
// This interface supports the identification of all fields on the front side of ID card for Singapore residents.The identification accuracy reaches more than 99%.In addition, this interface also supports a variety of value-added capabilities to meet the needs of different scenarios. Such as the cropping function of ID card photos and portrait photos, and also has 5 alarm functions.
//
// As shown in the table below. <table style="width:650px"> <thead> <tr> <th width="150">Value-added ability</th> <th width="500">Ability items</th> </tr> </thead> <tbody> <tr> <td rowspan="9">Alarm function</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>ID card copy warning</td> </tr> <tr> <td>Alarm for occlusion in the ID card frame</td> </tr> <tr> <td>ID card reflective warning</td> </tr> <tr> <td>Blurry picture warning</td> </tr> </tbody> </table> Default interface request frequency limit: 20 times/second
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeSingaporeIDCardOCRWithContext(ctx context.Context, request *RecognizeSingaporeIDCardOCRRequest) (response *RecognizeSingaporeIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeSingaporeIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeSingaporeIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeSingaporeIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeTableAccurateOCRRequest() (request *RecognizeTableAccurateOCRRequest) {
    request = &RecognizeTableAccurateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeTableAccurateOCR")
    
    
    return
}

func NewRecognizeTableAccurateOCRResponse() (response *RecognizeTableAccurateOCRResponse) {
    response = &RecognizeTableAccurateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeTableAccurateOCR
// This API is used to recognize regular tables, borderless tables, or multi-tables in images or PDF files containing Chinese and English texts. It returns the text content of each cell, supports recognition of rotated table images, and can save the recognition results into an Excel document. It delivers higher recognition accuracy than that of table OCR v2 and applies to more scenarios. The recognition accuracy in difficult table scenarios, such as irregular tables and nested tables (borderless tables contained in bordered tables), is better than that of table OCR v2. To try it, click [here](https://intl.cloud.tencent.com/product/smart?from_cn_redirect=1-ocr).
//
// 
//
// A maximum of 2 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTableAccurateOCR(request *RecognizeTableAccurateOCRRequest) (response *RecognizeTableAccurateOCRResponse, err error) {
    return c.RecognizeTableAccurateOCRWithContext(context.Background(), request)
}

// RecognizeTableAccurateOCR
// This API is used to recognize regular tables, borderless tables, or multi-tables in images or PDF files containing Chinese and English texts. It returns the text content of each cell, supports recognition of rotated table images, and can save the recognition results into an Excel document. It delivers higher recognition accuracy than that of table OCR v2 and applies to more scenarios. The recognition accuracy in difficult table scenarios, such as irregular tables and nested tables (borderless tables contained in bordered tables), is better than that of table OCR v2. To try it, click [here](https://intl.cloud.tencent.com/product/smart?from_cn_redirect=1-ocr).
//
// 
//
// A maximum of 2 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGESIZETOOLARGE = "FailedOperation.ImageSizeTooLarge"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeTableAccurateOCRWithContext(ctx context.Context, request *RecognizeTableAccurateOCRRequest) (response *RecognizeTableAccurateOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeTableAccurateOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeTableAccurateOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeTableAccurateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeThaiIDCardOCRRequest() (request *RecognizeThaiIDCardOCRRequest) {
    request = &RecognizeThaiIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeThaiIDCardOCR")
    
    
    return
}

func NewRecognizeThaiIDCardOCRResponse() (response *RecognizeThaiIDCardOCRResponse) {
    response = &RecognizeThaiIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeThaiIDCardOCR
// This API is used to recognize the fields on a Thai identity card, including name in Thai, name in English, address, date of birth, identification number, date of issue, and date of expiry.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeThaiIDCardOCR(request *RecognizeThaiIDCardOCRRequest) (response *RecognizeThaiIDCardOCRResponse, err error) {
    return c.RecognizeThaiIDCardOCRWithContext(context.Background(), request)
}

// RecognizeThaiIDCardOCR
// This API is used to recognize the fields on a Thai identity card, including name in Thai, name in English, address, date of birth, identification number, date of issue, and date of expiry.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEBLUR = "FailedOperation.ImageBlur"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_WARNINGSERVICEFAILED = "FailedOperation.WarningServiceFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeThaiIDCardOCRWithContext(ctx context.Context, request *RecognizeThaiIDCardOCRRequest) (response *RecognizeThaiIDCardOCRResponse, err error) {
    if request == nil {
        request = NewRecognizeThaiIDCardOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeThaiIDCardOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeThaiIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeThaiPinkCardRequest() (request *RecognizeThaiPinkCardRequest) {
    request = &RecognizeThaiPinkCardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "RecognizeThaiPinkCard")
    
    
    return
}

func NewRecognizeThaiPinkCardResponse() (response *RecognizeThaiPinkCardResponse) {
    response = &RecognizeThaiPinkCardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecognizeThaiPinkCard
// This API is used to recognize the fields on a Thai identity card, including name in Thai, name in English, address, date of birth, identification number, date of issue, and date of expiry.
//
// Currently, this API is not generally available. For more information, please [contact your sales rep](https://intl.cloud.tencent.com/about/connect?from_cn_redirect=1).
//
// 
//
// A maximum of 5 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RecognizeThaiPinkCard(request *RecognizeThaiPinkCardRequest) (response *RecognizeThaiPinkCardResponse, err error) {
    return c.RecognizeThaiPinkCardWithContext(context.Background(), request)
}

// RecognizeThaiPinkCard
// This API is used to recognize the fields on a Thai identity card, including name in Thai, name in English, address, date of birth, identification number, date of issue, and date of expiry.
//
// Currently, this API is not generally available. For more information, please [contact your sales rep](https://intl.cloud.tencent.com/about/connect?from_cn_redirect=1).
//
// 
//
// A maximum of 5 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOSPECIFIEDCARD = "FailedOperation.ImageNoSpecifiedCard"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) RecognizeThaiPinkCardWithContext(ctx context.Context, request *RecognizeThaiPinkCardRequest) (response *RecognizeThaiPinkCardResponse, err error) {
    if request == nil {
        request = NewRecognizeThaiPinkCardRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeThaiPinkCard require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeThaiPinkCardResponse()
    err = c.Send(request, response)
    return
}

func NewSealOCRRequest() (request *SealOCRRequest) {
    request = &SealOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "SealOCR")
    
    
    return
}

func NewSealOCRResponse() (response *SealOCRResponse) {
    response = &SealOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SealOCR
// This API is used to recognize various types of seals, including invoice seals and finance seals. It is suitable for scenarios such as official document and invoice/ticket OCR.
//
// 
//
// A maximum of 5 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SealOCR(request *SealOCRRequest) (response *SealOCRResponse, err error) {
    return c.SealOCRWithContext(context.Background(), request)
}

// SealOCR
// This API is used to recognize various types of seals, including invoice seals and finance seals. It is suitable for scenarios such as official document and invoice/ticket OCR.
//
// 
//
// A maximum of 5 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SealOCRWithContext(ctx context.Context, request *SealOCRRequest) (response *SealOCRResponse, err error) {
    if request == nil {
        request = NewSealOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SealOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewSealOCRResponse()
    err = c.Send(request, response)
    return
}

func NewSmartStructuralOCRV2Request() (request *SmartStructuralOCRV2Request) {
    request = &SmartStructuralOCRV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "SmartStructuralOCRV2")
    
    
    return
}

func NewSmartStructuralOCRV2Response() (response *SmartStructuralOCRV2Response) {
    response = &SmartStructuralOCRV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SmartStructuralOCRV2
// This API is used to recognize fields from cards, documents, bills, forms, contracts, and other structured information. It is flexible and efficient to use, without any configuration required. This API is suitable for recognizing structured information.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SmartStructuralOCRV2(request *SmartStructuralOCRV2Request) (response *SmartStructuralOCRV2Response, err error) {
    return c.SmartStructuralOCRV2WithContext(context.Background(), request)
}

// SmartStructuralOCRV2
// This API is used to recognize fields from cards, documents, bills, forms, contracts, and other structured information. It is flexible and efficient to use, without any configuration required. This API is suitable for recognizing structured information.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SmartStructuralOCRV2WithContext(ctx context.Context, request *SmartStructuralOCRV2Request) (response *SmartStructuralOCRV2Response, err error) {
    if request == nil {
        request = NewSmartStructuralOCRV2Request()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmartStructuralOCRV2 require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmartStructuralOCRV2Response()
    err = c.Send(request, response)
    return
}

func NewSmartStructuralProRequest() (request *SmartStructuralProRequest) {
    request = &SmartStructuralProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "SmartStructuralPro")
    
    
    return
}

func NewSmartStructuralProResponse() (response *SmartStructuralProResponse) {
    response = &SmartStructuralProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SmartStructuralPro
// This API is used to recognize fields from cards, documents, bills, forms, contracts, and other structured information. It is flexible and efficient to use, without any configuration required. This API is suitable for recognizing structured information.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SmartStructuralPro(request *SmartStructuralProRequest) (response *SmartStructuralProResponse, err error) {
    return c.SmartStructuralProWithContext(context.Background(), request)
}

// SmartStructuralPro
// This API is used to recognize fields from cards, documents, bills, forms, contracts, and other structured information. It is flexible and efficient to use, without any configuration required. This API is suitable for recognizing structured information.
//
// 
//
// A maximum of 10 requests can be initiated per second for this API.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SmartStructuralProWithContext(ctx context.Context, request *SmartStructuralProRequest) (response *SmartStructuralProResponse, err error) {
    if request == nil {
        request = NewSmartStructuralProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SmartStructuralPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewSmartStructuralProResponse()
    err = c.Send(request, response)
    return
}

func NewTableOCRRequest() (request *TableOCRRequest) {
    request = &TableOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "TableOCR")
    
    
    return
}

func NewTableOCRResponse() (response *TableOCRResponse) {
    response = &TableOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TableOCR
// This API is used to detect and recognize Chinese and English forms in images. It can return the text content of each cell and save the recognition result as Excel.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TableOCR(request *TableOCRRequest) (response *TableOCRResponse, err error) {
    return c.TableOCRWithContext(context.Background(), request)
}

// TableOCR
// This API is used to detect and recognize Chinese and English forms in images. It can return the text content of each cell and save the recognition result as Excel.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTEXT = "FailedOperation.ImageNoText"
//  FAILEDOPERATION_LANGUAGENOTSUPPORT = "FailedOperation.LanguageNotSupport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TableOCRWithContext(ctx context.Context, request *TableOCRRequest) (response *TableOCRResponse, err error) {
    if request == nil {
        request = NewTableOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TableOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewTableOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVinOCRRequest() (request *VinOCRRequest) {
    request = &VinOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ocr", APIVersion, "VinOCR")
    
    
    return
}

func NewVinOCRResponse() (response *VinOCRResponse) {
    response = &VinOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VinOCR
// This API is used to recognize the vehicle identification number (VIN) in an image.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VinOCR(request *VinOCRRequest) (response *VinOCRResponse, err error) {
    return c.VinOCRWithContext(context.Background(), request)
}

// VinOCR
// This API is used to recognize the vehicle identification number (VIN) in an image.
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_RESOURCEPACKAGERUNOUT = "ResourceUnavailable.ResourcePackageRunOut"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) VinOCRWithContext(ctx context.Context, request *VinOCRRequest) (response *VinOCRResponse, err error) {
    if request == nil {
        request = NewVinOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VinOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewVinOCRResponse()
    err = c.Send(request, response)
    return
}

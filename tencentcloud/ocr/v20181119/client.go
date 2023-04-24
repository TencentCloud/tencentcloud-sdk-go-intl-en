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
// This API is used to recognize key fields on the photo side of a Hong Kong (China) identity card, including name in Chinese, name in English, telecode for name, date of birth, gender, document symbol, date of the first issue, date of the last receipt, identity card number, and permanent residency attribute. It can check for card authenticity and crop the identity photo.
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) HKIDCardOCR(request *HKIDCardOCRRequest) (response *HKIDCardOCRResponse, err error) {
    return c.HKIDCardOCRWithContext(context.Background(), request)
}

// HKIDCardOCR
// This API is used to recognize key fields on the photo side of a Hong Kong (China) identity card, including name in Chinese, name in English, telecode for name, date of birth, gender, document symbol, date of the first issue, date of the last receipt, identity card number, and permanent residency attribute. It can check for card authenticity and crop the identity photo.
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
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
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
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
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
//  INVALIDPARAMETER_ENGINEIMAGEDECODEFAILED = "InvalidParameter.EngineImageDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
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
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) MLIDPassportOCR(request *MLIDPassportOCRRequest) (response *MLIDPassportOCRResponse, err error) {
    return c.MLIDPassportOCRWithContext(context.Background(), request)
}

// MLIDPassportOCR
// This API is used to recognize a passport issued in Hong Kong/Macao/Taiwan (China) or other countries/regions. Recognizable fields include passport ID, name, date of birth, gender, expiration date, issuing country/region, and nationality. It has the features of cropping identity photos and alarming for photographed or photocopied documents.
//
// 
//
// This API is not fully available for the time being. For more information, please contact your [Tencent Cloud sales rep](https://intl.cloud.tencent.com/contact-sales).
//
// error code that may be returned:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_NOPASSPORT = "FailedOperation.NoPassport"
//  FAILEDOPERATION_OCRFAILED = "FailedOperation.OcrFailed"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
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

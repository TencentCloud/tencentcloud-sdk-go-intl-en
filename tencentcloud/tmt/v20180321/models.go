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

package v20180321

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type TextTranslateRequestParams struct {
	// The texts to be translated, which must be encoded in UTF-8 and can contain up to 2,000 characters in a request. For non-pure texts such as those with HTML tags, the translation may fail.
	SourceText *string `json:"SourceText,omitnil" name:"SourceText"`

	// Supported source languages:
	// auto: Automatic language detection
	// zh: Simplified Chinese
	// zh_TW: Traditional Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Bahasa Indonesian
	// th: Thai
	// ms: Malay
	// ar: Arabic
	// hi: Hindi
	Source *string `json:"Source,omitnil" name:"Source"`

	// Supported target languages for the above source languages:
	// 
	// <li> zh (Simplified Chinese): en (English), ja (Japanese), ko (Korean), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), pt (Portuguese), vi (Vietnamese), id (Bahasa Indonesian), th (Thai), and ms (Malay)</li>
	// <li> zh-TW (Traditional Chinese): en (English), ja (Japanese), ko (Korean), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), pt (Portuguese), vi (Vietnamese), id (Bahasa Indonesian), th (Thai), and ms (Malay)</li>
	// <li> en (English): zh (Simplified Chinese), ja (Japanese), ko (Korean), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), pt (Portuguese), vi (Vietnamese), id (Bahasa Indonesian), th (Thai), ms (Malay), ar (Arabic), and hi (Hindi)</li>
	// <li>ja (Japanese): zh (Simplified Chinese), en (English), and ko (Korean)</li>
	// <li>ko (Korean): zh (Simplified Chinese), en (English), and ja (Japanese)</li>
	// <li>fr (French): zh (Simplified Chinese), en (English), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>es (Spanish): zh (Simplified Chinese), en (English), fr (French), it (Italian), de (German), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>it (Italian): zh (Simplified Chinese), en (English), fr (French), es (Spanish), de (German), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>de (German): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>tr (Turkish): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), de (German), ru (Russian), and pt (Portuguese)</li>
	// <li>ru (Russian): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), and pt (Portuguese)</li>
	// <li>pt (Portuguese): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), and ru (Russian)</li>
	// <li>vi (Vietnamese): zh (Simplified Chinese) and en (English)</li
	// <li>id (Bahasa Indonesian): zh (Simplified Chinese) and en (English)</li
	// <li>th (Thai): zh (Simplified Chinese) and en (English)</li
	// <li>ms (Malay): zh (Simplified Chinese) and en (English)</li
	// <li>ar (Arabic): en (English)</li>
	// <li>hi (Hindi): en (English)</li
	Target *string `json:"Target,omitnil" name:"Target"`

	// The project ID, which can be obtained from **Console -> Account Center -> Project Management**. If no one is set, enter the default project ID `0`.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// The parameter used to mark the text content that needs to remain untranslated, such as special symbols and names of people and places. You can set only one word for this parameter in each request. Only nouns (like names of people and places) are supported, and verbs or phrases may cause poor translation outcomes.
	UntranslatedText *string `json:"UntranslatedText,omitnil" name:"UntranslatedText"`
}

type TextTranslateRequest struct {
	*tchttp.BaseRequest
	
	// The texts to be translated, which must be encoded in UTF-8 and can contain up to 2,000 characters in a request. For non-pure texts such as those with HTML tags, the translation may fail.
	SourceText *string `json:"SourceText,omitnil" name:"SourceText"`

	// Supported source languages:
	// auto: Automatic language detection
	// zh: Simplified Chinese
	// zh_TW: Traditional Chinese
	// en: English
	// ja: Japanese
	// ko: Korean
	// fr: French
	// es: Spanish
	// it: Italian
	// de: German
	// tr: Turkish
	// ru: Russian
	// pt: Portuguese
	// vi: Vietnamese
	// id: Bahasa Indonesian
	// th: Thai
	// ms: Malay
	// ar: Arabic
	// hi: Hindi
	Source *string `json:"Source,omitnil" name:"Source"`

	// Supported target languages for the above source languages:
	// 
	// <li> zh (Simplified Chinese): en (English), ja (Japanese), ko (Korean), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), pt (Portuguese), vi (Vietnamese), id (Bahasa Indonesian), th (Thai), and ms (Malay)</li>
	// <li> zh-TW (Traditional Chinese): en (English), ja (Japanese), ko (Korean), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), pt (Portuguese), vi (Vietnamese), id (Bahasa Indonesian), th (Thai), and ms (Malay)</li>
	// <li> en (English): zh (Simplified Chinese), ja (Japanese), ko (Korean), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), pt (Portuguese), vi (Vietnamese), id (Bahasa Indonesian), th (Thai), ms (Malay), ar (Arabic), and hi (Hindi)</li>
	// <li>ja (Japanese): zh (Simplified Chinese), en (English), and ko (Korean)</li>
	// <li>ko (Korean): zh (Simplified Chinese), en (English), and ja (Japanese)</li>
	// <li>fr (French): zh (Simplified Chinese), en (English), es (Spanish), it (Italian), de (German), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>es (Spanish): zh (Simplified Chinese), en (English), fr (French), it (Italian), de (German), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>it (Italian): zh (Simplified Chinese), en (English), fr (French), es (Spanish), de (German), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>de (German): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), tr (Turkish), ru (Russian), and pt (Portuguese)</li>
	// <li>tr (Turkish): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), de (German), ru (Russian), and pt (Portuguese)</li>
	// <li>ru (Russian): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), and pt (Portuguese)</li>
	// <li>pt (Portuguese): zh (Simplified Chinese), en (English), fr (French), es (Spanish), it (Italian), de (German), tr (Turkish), and ru (Russian)</li>
	// <li>vi (Vietnamese): zh (Simplified Chinese) and en (English)</li
	// <li>id (Bahasa Indonesian): zh (Simplified Chinese) and en (English)</li
	// <li>th (Thai): zh (Simplified Chinese) and en (English)</li
	// <li>ms (Malay): zh (Simplified Chinese) and en (English)</li
	// <li>ar (Arabic): en (English)</li>
	// <li>hi (Hindi): en (English)</li
	Target *string `json:"Target,omitnil" name:"Target"`

	// The project ID, which can be obtained from **Console -> Account Center -> Project Management**. If no one is set, enter the default project ID `0`.
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// The parameter used to mark the text content that needs to remain untranslated, such as special symbols and names of people and places. You can set only one word for this parameter in each request. Only nouns (like names of people and places) are supported, and verbs or phrases may cause poor translation outcomes.
	UntranslatedText *string `json:"UntranslatedText,omitnil" name:"UntranslatedText"`
}

func (r *TextTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceText")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "ProjectId")
	delete(f, "UntranslatedText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateResponseParams struct {
	// The translation outcome.
	TargetText *string `json:"TargetText,omitnil" name:"TargetText"`

	// The source language. See the request parameter `Source` for details.
	Source *string `json:"Source,omitnil" name:"Source"`

	// The target language. See the request parameter `Target` for details.
	Target *string `json:"Target,omitnil" name:"Target"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextTranslateResponse struct {
	*tchttp.BaseResponse
	Response *TextTranslateResponseParams `json:"Response"`
}

func (r *TextTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
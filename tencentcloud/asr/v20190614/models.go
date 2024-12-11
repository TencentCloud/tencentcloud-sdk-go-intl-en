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

package v20190614

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type CreateRecTaskRequestParams struct {
	// Engine model type.
	// Each recognition engine adopts a specific billing plan. Engines marked with "large model version" adopt the large model billing plan. For product billing instructions, [click here] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1).
	// 
	// 
	// Note: If you want to recognize telecommunication audio but find that a 16k engine is required, you can use a 16k engine as described below for recognition. However, ** the 16k engines are not trained for recognizing telecommunication audio. Therefore, the recognition results cannot be guaranteed. You need to check whether the recognition results can be used. **
	// 
	// Engines for general scenarios:
	// ** Note: Use 16k engines for scenarios other than telecommunication. **
	// ** 16k_zh_large: ** Engine (large model version) for Mandarin, Chinese dialects, and English. This engine supports recognizing audio in Chinese, English, and [various Chinese dialects] (https://intl.cloud.tencent.com/document/product/1093/35682?from_cn_redirect=1). It has a large number of parameters, enhanced performance, and greatly improved recognition accuracy for low-quality audio with loud noise, too much echo, low voice volume, or faint voices. [Click here] (https://console.cloud.tencent.com/asr/demonstrate) to compare the recognition performance of the 16k_zh engine and this one.
	// ** 16k_multi_lang: ** Engine (large model version) for multiple languages. This engine supports recognizing audio in English, Japanese, Korean, Arabic, Filipino, French, Hindi, Indonesian, Malay, Portuguese, Spanish, Thai, Turkish, Vietnamese, and German (sentence-level or paragraph-level).
	// ** 16k_zh-PY: ** Engine for Chinese, English, and Cantonese. The engine supports recognizing audio in Mandarin, English, and Cantonese at the same time.
	// ** 16k_ms: ** Engine for Malay.
	// ** 16k_id: ** Engine for Indonesian.
	// ** 16k_th: ** Engine for Thai.
	EngineModelType *string `json:"EngineModelType,omitnil,omitempty" name:"EngineModelType"`

	// Number of recognition channels.
	// 1: Mono. (16k engines only support mono. ** Do no t** set to stereo.)
	// 2: Stereo. (Stereo is supported only for 8k engines, and the two channels should correspond to the respective communication parties.)
	// 
	// Note:
	// 16k engines: Only support mono. ** ChannelNum should be set to 1 **.
	// 8k engines: Support both mono and stereo. ** It is recommended to set ChannelNum to 2 (indicating stereo) **. Stereo can physically distinguish speakers to avoid recognition mistakes caused by overlapping speech. It can provide the best speaker separation and recognition effects. Once stereo is set, the speakers are automatically separated. ** You do not need to enable the speaker separation feature **. You can use the default values for related parameters (** SpeakerDiarization and SpeakerNumber **). For speakerId in the returned ResultDetail, the value 0 represents the left channel, and the value 1 represents the right channel.
	ChannelNum *uint64 `json:"ChannelNum,omitnil,omitempty" name:"ChannelNum"`

	// Format of the returned recognition result.
	// 0: The basic recognition result (containing only valid voice timestamps but no word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail)).
	// 1: The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps and speech speed value but ** no punctuation **).
	// 2: The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **).
	// 3: The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **). The recognition results are segmented by punctuation. ** This format applies to subtitle scenarios **.
	// 4: ** [Value-added paid feature] ** The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **). The recognition results are segmented by NLP semantics. ** This format applies to scenarios such as meeting and court record transcription ** and is supported only for 8k_zh and 16k_zh engines.
	// 5: ** [Value-added paid feature] ** Basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **). The oral-to-written transcription result is also output, which has excluded modal particles and consecutive identical words, optimized expressions, and corrected speech mistakes. ** This format applies to scenarios of generating minutes for online and offline meetings** and is supported only for 8k_zh and 16k_zh engines.
	// 
	// Notes:
	// If this parameter is set to 4, make sure that a [semantics-based segmentation resource package] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#97ae4aa0-29a0-4066-9f07-ccaf8856a16b) is purchased for your account or that your account has enabled post-payment. ** If post-payment is enabled and this parameter is set to 4, [automatic billing] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#d912167d-ffd5-41a9-8b1c-2e89845a6852) will apply **.
	// If this parameter is set to 5, make sure that an [oral-to-written resource package] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#97ae4aa0-29a0-4066-9f07-ccaf8856a16b) is purchased for your account or that your account has enabled post-payment. ** If post-payment is enabled and this parameter is set to 5, [automatic billing] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#d912167d-ffd5-41a9-8b1c-2e89845a6852) will apply **.
	ResTextFormat *uint64 `json:"ResTextFormat,omitnil,omitempty" name:"ResTextFormat"`

	// Audio source.
	// 0: Audio URL.
	// 1: Local audio file (body of the POST request).
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Audio file Base64 code.
	// ** This parameter is required if SourceType is set to 1. Otherwise, it can be left blank. **
	// 
	// Note: The audio data size cannot exceed 5 MB.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Data length (before Base64 encoding).
	DataLen *uint64 `json:"DataLen,omitnil,omitempty" name:"DataLen"`

	// Audio URL. (The audio should be downloadable via a public network browser.)
	// ** This parameter is required if SourceType is set to 0. Otherwise, it can be left blank. **
	// 
	// Notes:
	// 1. Make sure that the total audio duration of recording files does not exceed 5 hours. Otherwise, recognition may fail.
	// 2. Pay attention to file download to avoid download failure.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Callback URL
	// 
	// User-defined service URL for receiving recognition results.
	// For the callback format and content, see [Callback Description] (https://intl.cloud.tencent.com/document/product/1093/52632?from_cn_redirect=1).
	// 
	// Notes:
	// 
	// - If you use the polling method to obtain recognition results, this parameter is not required.
	// - It is recommended to include your business ID and other information in the callback URL for handling business logic.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// Whether to enable speaker separation.
	// 0: Disable.
	// 1: Enable. (This value is supported only for the following engines: 8k_zh, 16k_zh, 16k_ms, 16k_en, 16k_id, 16k_zh_large, and 16k_zh_dialect. ChannelNum should be set to 1.)
	// The default value is 0.
	// 
	// Note:
	// If an 8k engine is used and ChannelNum is set to 2 (stereo), use the default values for corresponding parameters as stated in the ** ChannelNum ** parameter description.
	SpeakerDiarization *int64 `json:"SpeakerDiarization,omitnil,omitempty" name:"SpeakerDiarization"`

	// Number of speakers to be separated.
	// ** Speaker separation must be enabled. Otherwise, this parameter does not take effect. ** Value range: 0-10.
	// 0: Automatic separation. (Up to 20 speakers can be separated.)
	// 1-10: Specify the number of speakers.
	// The default value is 0.
	SpeakerNumber *int64 `json:"SpeakerNumber,omitnil,omitempty" name:"SpeakerNumber"`

	// This service is not available.
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// This service is not available.
	//
	// Deprecated: ReinforceHotword is deprecated.
	ReinforceHotword *int64 `json:"ReinforceHotword,omitnil,omitempty" name:"ReinforceHotword"`

	// This service is not available.
	CustomizationId *string `json:"CustomizationId,omitnil,omitempty" name:"CustomizationId"`

	// This service is not available.
	EmotionRecognition *int64 `json:"EmotionRecognition,omitnil,omitempty" name:"EmotionRecognition"`

	// Emotional energy value.
	// The value is the result of dividing the sound volume in dB by 10. Value range: [1,10]. The higher the value, the stronger the emotion.
	// 0: Disable.
	// 1: Enable.
	// The default value is 0.
	EmotionalEnergy *int64 `json:"EmotionalEnergy,omitnil,omitempty" name:"EmotionalEnergy"`

	// Intelligent conversion into Arabic numerals (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not convert, but directly output Chinese numerals.
	// 1: Intelligently convert into Arabic numerals based on the scenario.
	// 3: Enable conversion for math-related letters.
	// The default value is 1.
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// Dirty word filtering (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not filter out dirty words.
	// 1: Filter out dirty words.
	// 2: Replace dirty words with *.
	// The default value is 0.
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// Punctuation filtering (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not filter out punctuation.
	// 1: Filter out sentence-ending punctuation.
	// 2: Filter out all punctuation.
	// The default value is 0.
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// Modal particle filtering (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not filter out modal particles.
	// 1: Filter out specified modal particles.
	// 2: Filter out all modal particles.
	// The default value is 0.
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// The maximum number of characters per line (supported only for engines for recognizing audio in Mandarin). A punctuation mark is added if this limit is reached.
	// ** This parameter can control the maximum number of characters per line, which applies to subtitle generation scenarios. ** Value range: [6,40].
	// 0: Disable this feature.
	// The default value is 0.
	// 
	// Note: To enable this feature, ResTextFormat should be set to 3. The recognition result can be obtained from FinalSentence by parsing the list in the returned ResultDetail.
	SentenceMaxLength *int64 `json:"SentenceMaxLength,omitnil,omitempty" name:"SentenceMaxLength"`

	// Additional parameter. ** (This parameter is meaningless. Ignore it.) **
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Temporary term list. This parameter is used to improve the recognition accuracy.
	// 
	// - Restrictions for individual terms: The format is "term|weight". Each term can contain no more than 30 characters (or 10 Chinese characters. The weight can be in the range of [1-11] or 100. For example, "Tencent Cloud|5" or "ASR|11".
	// 
	// - Restrictions for the temporary term list: Multiple terms are separated by commas. The list can contain up to 128 terms. For example, "Tencent Cloud|10, Audio Recognition|5, ASR|11".
	// 
	// - Difference between hotword_id (term list) and hotword_list (temporary term list):
	// 
	//     - hotword_id: Term list. You need to create a term list in the console or by using the API first and obtain the term list ID.
	// 
	//     - hotword_list: Temporary term list. You can directly enter the ID of the temporary term list each time you initiate a request. The temporary term list is not retained on the cloud. This parameter applies to users with a massive number of terms.
	// 
	// Notes:
	// 
	// - If both hotword_id and hotword_list are specified, hotword_list will take effect first.
	// 
	// - When the weight of a term is set to 11, this term becomes a super term. It is recommended that the weight is set to 11 only for critical and necessary terms. The overall accuracy will be affected if the weight is set to 11 for too many terms.
	// 
	// - When the weight of a term is set to 100, the term enhancement feature is enabled to replace homophones of this term. (This feature is supported only for 8k_zh and 16k_zh engines.) For example, if you configure "mizhi 1|100", the recognized word "mizhi 2", which is a homophone of "mizhi 2", will be forcibly replaced with "mizhi 2". It is recommended that you enable this feature based on the actual needs. You can set the weight to 100 for only critical and necessary terms. The overall accuracy will be affected if the weight is set to 100 for too many terms.
	HotwordList *string `json:"HotwordList,omitnil,omitempty" name:"HotwordList"`

	// List of keyword IDs for recognition. This parameter is left blank by default, indicating that no keyword is recognized. You can enter up to 10 IDs.
	KeyWordLibIdList []*string `json:"KeyWordLibIdList,omitnil,omitempty" name:"KeyWordLibIdList"`
}

type CreateRecTaskRequest struct {
	*tchttp.BaseRequest
	
	// Engine model type.
	// Each recognition engine adopts a specific billing plan. Engines marked with "large model version" adopt the large model billing plan. For product billing instructions, [click here] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1).
	// 
	// 
	// Note: If you want to recognize telecommunication audio but find that a 16k engine is required, you can use a 16k engine as described below for recognition. However, ** the 16k engines are not trained for recognizing telecommunication audio. Therefore, the recognition results cannot be guaranteed. You need to check whether the recognition results can be used. **
	// 
	// Engines for general scenarios:
	// ** Note: Use 16k engines for scenarios other than telecommunication. **
	// ** 16k_zh_large: ** Engine (large model version) for Mandarin, Chinese dialects, and English. This engine supports recognizing audio in Chinese, English, and [various Chinese dialects] (https://intl.cloud.tencent.com/document/product/1093/35682?from_cn_redirect=1). It has a large number of parameters, enhanced performance, and greatly improved recognition accuracy for low-quality audio with loud noise, too much echo, low voice volume, or faint voices. [Click here] (https://console.cloud.tencent.com/asr/demonstrate) to compare the recognition performance of the 16k_zh engine and this one.
	// ** 16k_multi_lang: ** Engine (large model version) for multiple languages. This engine supports recognizing audio in English, Japanese, Korean, Arabic, Filipino, French, Hindi, Indonesian, Malay, Portuguese, Spanish, Thai, Turkish, Vietnamese, and German (sentence-level or paragraph-level).
	// ** 16k_zh-PY: ** Engine for Chinese, English, and Cantonese. The engine supports recognizing audio in Mandarin, English, and Cantonese at the same time.
	// ** 16k_ms: ** Engine for Malay.
	// ** 16k_id: ** Engine for Indonesian.
	// ** 16k_th: ** Engine for Thai.
	EngineModelType *string `json:"EngineModelType,omitnil,omitempty" name:"EngineModelType"`

	// Number of recognition channels.
	// 1: Mono. (16k engines only support mono. ** Do no t** set to stereo.)
	// 2: Stereo. (Stereo is supported only for 8k engines, and the two channels should correspond to the respective communication parties.)
	// 
	// Note:
	// 16k engines: Only support mono. ** ChannelNum should be set to 1 **.
	// 8k engines: Support both mono and stereo. ** It is recommended to set ChannelNum to 2 (indicating stereo) **. Stereo can physically distinguish speakers to avoid recognition mistakes caused by overlapping speech. It can provide the best speaker separation and recognition effects. Once stereo is set, the speakers are automatically separated. ** You do not need to enable the speaker separation feature **. You can use the default values for related parameters (** SpeakerDiarization and SpeakerNumber **). For speakerId in the returned ResultDetail, the value 0 represents the left channel, and the value 1 represents the right channel.
	ChannelNum *uint64 `json:"ChannelNum,omitnil,omitempty" name:"ChannelNum"`

	// Format of the returned recognition result.
	// 0: The basic recognition result (containing only valid voice timestamps but no word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail)).
	// 1: The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps and speech speed value but ** no punctuation **).
	// 2: The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **).
	// 3: The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **). The recognition results are segmented by punctuation. ** This format applies to subtitle scenarios **.
	// 4: ** [Value-added paid feature] ** The basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **). The recognition results are segmented by NLP semantics. ** This format applies to scenarios such as meeting and court record transcription ** and is supported only for 8k_zh and 16k_zh engines.
	// 5: ** [Value-added paid feature] ** Basic recognition result and word-level [detailed recognition result] (https://intl.cloud.tencent.com/document/api/1093/37824?from_cn_redirect=1#SentenceDetail) (containing word-level timestamps, speech speed value, and ** punctuation **). The oral-to-written transcription result is also output, which has excluded modal particles and consecutive identical words, optimized expressions, and corrected speech mistakes. ** This format applies to scenarios of generating minutes for online and offline meetings** and is supported only for 8k_zh and 16k_zh engines.
	// 
	// Notes:
	// If this parameter is set to 4, make sure that a [semantics-based segmentation resource package] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#97ae4aa0-29a0-4066-9f07-ccaf8856a16b) is purchased for your account or that your account has enabled post-payment. ** If post-payment is enabled and this parameter is set to 4, [automatic billing] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#d912167d-ffd5-41a9-8b1c-2e89845a6852) will apply **.
	// If this parameter is set to 5, make sure that an [oral-to-written resource package] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#97ae4aa0-29a0-4066-9f07-ccaf8856a16b) is purchased for your account or that your account has enabled post-payment. ** If post-payment is enabled and this parameter is set to 5, [automatic billing] (https://intl.cloud.tencent.com/document/product/1093/35686?from_cn_redirect=1#d912167d-ffd5-41a9-8b1c-2e89845a6852) will apply **.
	ResTextFormat *uint64 `json:"ResTextFormat,omitnil,omitempty" name:"ResTextFormat"`

	// Audio source.
	// 0: Audio URL.
	// 1: Local audio file (body of the POST request).
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// Audio file Base64 code.
	// ** This parameter is required if SourceType is set to 1. Otherwise, it can be left blank. **
	// 
	// Note: The audio data size cannot exceed 5 MB.
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// Data length (before Base64 encoding).
	DataLen *uint64 `json:"DataLen,omitnil,omitempty" name:"DataLen"`

	// Audio URL. (The audio should be downloadable via a public network browser.)
	// ** This parameter is required if SourceType is set to 0. Otherwise, it can be left blank. **
	// 
	// Notes:
	// 1. Make sure that the total audio duration of recording files does not exceed 5 hours. Otherwise, recognition may fail.
	// 2. Pay attention to file download to avoid download failure.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Callback URL
	// 
	// User-defined service URL for receiving recognition results.
	// For the callback format and content, see [Callback Description] (https://intl.cloud.tencent.com/document/product/1093/52632?from_cn_redirect=1).
	// 
	// Notes:
	// 
	// - If you use the polling method to obtain recognition results, this parameter is not required.
	// - It is recommended to include your business ID and other information in the callback URL for handling business logic.
	CallbackUrl *string `json:"CallbackUrl,omitnil,omitempty" name:"CallbackUrl"`

	// Whether to enable speaker separation.
	// 0: Disable.
	// 1: Enable. (This value is supported only for the following engines: 8k_zh, 16k_zh, 16k_ms, 16k_en, 16k_id, 16k_zh_large, and 16k_zh_dialect. ChannelNum should be set to 1.)
	// The default value is 0.
	// 
	// Note:
	// If an 8k engine is used and ChannelNum is set to 2 (stereo), use the default values for corresponding parameters as stated in the ** ChannelNum ** parameter description.
	SpeakerDiarization *int64 `json:"SpeakerDiarization,omitnil,omitempty" name:"SpeakerDiarization"`

	// Number of speakers to be separated.
	// ** Speaker separation must be enabled. Otherwise, this parameter does not take effect. ** Value range: 0-10.
	// 0: Automatic separation. (Up to 20 speakers can be separated.)
	// 1-10: Specify the number of speakers.
	// The default value is 0.
	SpeakerNumber *int64 `json:"SpeakerNumber,omitnil,omitempty" name:"SpeakerNumber"`

	// This service is not available.
	HotwordId *string `json:"HotwordId,omitnil,omitempty" name:"HotwordId"`

	// This service is not available.
	ReinforceHotword *int64 `json:"ReinforceHotword,omitnil,omitempty" name:"ReinforceHotword"`

	// This service is not available.
	CustomizationId *string `json:"CustomizationId,omitnil,omitempty" name:"CustomizationId"`

	// This service is not available.
	EmotionRecognition *int64 `json:"EmotionRecognition,omitnil,omitempty" name:"EmotionRecognition"`

	// Emotional energy value.
	// The value is the result of dividing the sound volume in dB by 10. Value range: [1,10]. The higher the value, the stronger the emotion.
	// 0: Disable.
	// 1: Enable.
	// The default value is 0.
	EmotionalEnergy *int64 `json:"EmotionalEnergy,omitnil,omitempty" name:"EmotionalEnergy"`

	// Intelligent conversion into Arabic numerals (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not convert, but directly output Chinese numerals.
	// 1: Intelligently convert into Arabic numerals based on the scenario.
	// 3: Enable conversion for math-related letters.
	// The default value is 1.
	ConvertNumMode *int64 `json:"ConvertNumMode,omitnil,omitempty" name:"ConvertNumMode"`

	// Dirty word filtering (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not filter out dirty words.
	// 1: Filter out dirty words.
	// 2: Replace dirty words with *.
	// The default value is 0.
	FilterDirty *int64 `json:"FilterDirty,omitnil,omitempty" name:"FilterDirty"`

	// Punctuation filtering (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not filter out punctuation.
	// 1: Filter out sentence-ending punctuation.
	// 2: Filter out all punctuation.
	// The default value is 0.
	FilterPunc *int64 `json:"FilterPunc,omitnil,omitempty" name:"FilterPunc"`

	// Modal particle filtering (supported only for engines for recognizing audio in Mandarin).
	// 0: Do not filter out modal particles.
	// 1: Filter out specified modal particles.
	// 2: Filter out all modal particles.
	// The default value is 0.
	FilterModal *int64 `json:"FilterModal,omitnil,omitempty" name:"FilterModal"`

	// The maximum number of characters per line (supported only for engines for recognizing audio in Mandarin). A punctuation mark is added if this limit is reached.
	// ** This parameter can control the maximum number of characters per line, which applies to subtitle generation scenarios. ** Value range: [6,40].
	// 0: Disable this feature.
	// The default value is 0.
	// 
	// Note: To enable this feature, ResTextFormat should be set to 3. The recognition result can be obtained from FinalSentence by parsing the list in the returned ResultDetail.
	SentenceMaxLength *int64 `json:"SentenceMaxLength,omitnil,omitempty" name:"SentenceMaxLength"`

	// Additional parameter. ** (This parameter is meaningless. Ignore it.) **
	Extra *string `json:"Extra,omitnil,omitempty" name:"Extra"`

	// Temporary term list. This parameter is used to improve the recognition accuracy.
	// 
	// - Restrictions for individual terms: The format is "term|weight". Each term can contain no more than 30 characters (or 10 Chinese characters. The weight can be in the range of [1-11] or 100. For example, "Tencent Cloud|5" or "ASR|11".
	// 
	// - Restrictions for the temporary term list: Multiple terms are separated by commas. The list can contain up to 128 terms. For example, "Tencent Cloud|10, Audio Recognition|5, ASR|11".
	// 
	// - Difference between hotword_id (term list) and hotword_list (temporary term list):
	// 
	//     - hotword_id: Term list. You need to create a term list in the console or by using the API first and obtain the term list ID.
	// 
	//     - hotword_list: Temporary term list. You can directly enter the ID of the temporary term list each time you initiate a request. The temporary term list is not retained on the cloud. This parameter applies to users with a massive number of terms.
	// 
	// Notes:
	// 
	// - If both hotword_id and hotword_list are specified, hotword_list will take effect first.
	// 
	// - When the weight of a term is set to 11, this term becomes a super term. It is recommended that the weight is set to 11 only for critical and necessary terms. The overall accuracy will be affected if the weight is set to 11 for too many terms.
	// 
	// - When the weight of a term is set to 100, the term enhancement feature is enabled to replace homophones of this term. (This feature is supported only for 8k_zh and 16k_zh engines.) For example, if you configure "mizhi 1|100", the recognized word "mizhi 2", which is a homophone of "mizhi 2", will be forcibly replaced with "mizhi 2". It is recommended that you enable this feature based on the actual needs. You can set the weight to 100 for only critical and necessary terms. The overall accuracy will be affected if the weight is set to 100 for too many terms.
	HotwordList *string `json:"HotwordList,omitnil,omitempty" name:"HotwordList"`

	// List of keyword IDs for recognition. This parameter is left blank by default, indicating that no keyword is recognized. You can enter up to 10 IDs.
	KeyWordLibIdList []*string `json:"KeyWordLibIdList,omitnil,omitempty" name:"KeyWordLibIdList"`
}

func (r *CreateRecTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineModelType")
	delete(f, "ChannelNum")
	delete(f, "ResTextFormat")
	delete(f, "SourceType")
	delete(f, "Data")
	delete(f, "DataLen")
	delete(f, "Url")
	delete(f, "CallbackUrl")
	delete(f, "SpeakerDiarization")
	delete(f, "SpeakerNumber")
	delete(f, "HotwordId")
	delete(f, "ReinforceHotword")
	delete(f, "CustomizationId")
	delete(f, "EmotionRecognition")
	delete(f, "EmotionalEnergy")
	delete(f, "ConvertNumMode")
	delete(f, "FilterDirty")
	delete(f, "FilterPunc")
	delete(f, "FilterModal")
	delete(f, "SentenceMaxLength")
	delete(f, "Extra")
	delete(f, "HotwordList")
	delete(f, "KeyWordLibIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRecTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRecTaskResponseParams struct {
	// Returned result of the recording recognition request, containing the task ID required for querying the result.
	// ** Note: The task ID is valid for 24 hours, and duplicate task IDs of different dates may exist. Do not use task ID as the unique ID in your business system. **
	Data *Task `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRecTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRecTaskResponseParams `json:"Response"`
}

func (r *CreateRecTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRecTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// Task ID obtained from the CreateRecTask API, which is used to obtain the task status and results.
	// ** Note: A task is valid for 24 hours. Do not query the results with the tasks that have existed for more than 24 hours. **
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// Task ID obtained from the CreateRecTask API, which is used to obtain the task status and results.
	// ** Note: A task is valid for 24 hours. Do not query the results with the tasks that have existed for more than 24 hours. **
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// Returned result of the recording recognition request.
	Data *TaskStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// The unique request ID, generated by the server, will be returned for every request (if the request fails to reach the server for other reasons, the request will not obtain a RequestId). RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyWordResult struct {
	// Keyword library ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyWordLibID *string `json:"KeyWordLibID,omitnil,omitempty" name:"KeyWordLibID"`

	// Keyword library name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyWordLibName *string `json:"KeyWordLibName,omitnil,omitempty" name:"KeyWordLibName"`

	// Matching keywords.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyWords []*string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`
}

type SentenceDetail struct {
	// Final recognition result of a sentence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	FinalSentence *string `json:"FinalSentence,omitnil,omitempty" name:"FinalSentence"`

	// Intermediate recognition result of a sentence. The sentence is split into multiple phrases by spaces.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SliceSentence *string `json:"SliceSentence,omitnil,omitempty" name:"SliceSentence"`

	// Oral-to-written transcription result. This parameter has a value only if the corresponding feature is enabled.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WrittenText *string `json:"WrittenText,omitnil,omitempty" name:"WrittenText"`

	// Start time of a sentence (ms).
	// Note: This field may return null, indicating that no valid values can be obtained.
	StartMs *int64 `json:"StartMs,omitnil,omitempty" name:"StartMs"`

	// End time of a sentence (ms).
	// Note: This field may return null, indicating that no valid values can be obtained.
	EndMs *int64 `json:"EndMs,omitnil,omitempty" name:"EndMs"`

	// Number of words in a sentence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	WordsNum *int64 `json:"WordsNum,omitnil,omitempty" name:"WordsNum"`

	// Word details of a sentence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Words []*SentenceWords `json:"Words,omitnil,omitempty" name:"Words"`

	// Speech speed of a sentence. Unit: Number of words per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpeechSpeed *float64 `json:"SpeechSpeed,omitnil,omitempty" name:"SpeechSpeed"`

	// Channel or speaker ID. (If speaker_diarization is specified or ChannelNum is set to 2 (stereo) in the request, speakers or channels can be distinguished.)
	// Different values represent different speakers in mono mode. For the speakerId values, 0 represents the left channel, and 1 represents the right channel in stereo mode if an 8k engine is used.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SpeakerId *int64 `json:"SpeakerId,omitnil,omitempty" name:"SpeakerId"`

	// Emotional energy value. This value is the result of dividing the sound volume in dB by 10. Value range: [1,10]. The higher the value, the stronger the emotion.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EmotionalEnergy *float64 `json:"EmotionalEnergy,omitnil,omitempty" name:"EmotionalEnergy"`

	// Silent duration between the current sentence and the last sentence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SilenceTime *int64 `json:"SilenceTime,omitnil,omitempty" name:"SilenceTime"`

	// Emotion type. (This parameter may be left blank in two scenarios: 1. No corresponding resource package exists; 2. The emotion is not recognized because it is not strong enough, which is related to the emotional energy.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	EmotionType []*string `json:"EmotionType,omitnil,omitempty" name:"EmotionType"`

	// List of recognized keywords.
	// Note: This field may return null, indicating that no valid values can be obtained.
	KeyWordResults []*KeyWordResult `json:"KeyWordResults,omitnil,omitempty" name:"KeyWordResults"`
}

type SentenceWords struct {
	// Word text.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// Start time offset in the sentence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OffsetStartMs *int64 `json:"OffsetStartMs,omitnil,omitempty" name:"OffsetStartMs"`

	// End time offset in the sentence.
	// Note: This field may return null, indicating that no valid values can be obtained.
	OffsetEndMs *int64 `json:"OffsetEndMs,omitnil,omitempty" name:"OffsetEndMs"`
}

type Task struct {
	// Task ID. This ID can be used to obtain the recognition status and results through polling. The data type of TaskId is ** uint64 **.
	// ** Note: The task ID is valid for 24 hours, and duplicate task IDs of different dates may exist. Do not use task ID as the unique ID in your business system. **
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type TaskStatus struct {
	// Task ID. Note: The data type of TaskId is uint64.
	TaskId *uint64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// Task status code. 0: waiting; 1: in process; 2: success; 3: failed.
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Task status. Valid values: waiting, in process, success, and failed.
	StatusStr *string `json:"StatusStr,omitnil,omitempty" name:"StatusStr"`

	// Recognition result.
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// Failure cause.
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// Recognition result details, including word time offsets for each sentence, which is generally used in subtitle generation scenarios. (This field is not left blank when ResTextFormat in the recording recognition request is set to 1.)
	// Note: This field may return null, indicating that no valid values can be obtained.
	ResultDetail []*SentenceDetail `json:"ResultDetail,omitnil,omitempty" name:"ResultDetail"`

	// Audio duration (seconds).
	// Note: This field may return null, indicating that no valid values can be obtained.
	AudioDuration *float64 `json:"AudioDuration,omitnil,omitempty" name:"AudioDuration"`
}
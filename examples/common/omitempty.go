package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/regions"
	trtc "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/trtc/v20190722"
	"os"
)

func main() {
	// The SDK uses the `omitnil` tag to serialize your request objects, as this allows it to
	// distinguish between `nil arrays` and `arrays with a length of 0`.

	// By default, the SDK will send `arrays with a length of 0` and ignore `nil arrays`.
	sendJsonRequest()

	// When you do not want to send an `array with a length of 0`, you can disable this
	// feature by setting json.OmitBehaviour.
	json.OmitBehaviour = json.OmitEmpty
	sendJsonRequest()
}

func sendJsonRequest() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	cpf := profile.NewClientProfile()
	cpf.Debug = true

	client, err := trtc.NewClient(credential, regions.Guangzhou, cpf)

	req := trtc.NewUpdatePublishCdnStreamRequest()

	// Due to the omitempty tag, the PublishCdnParams field will not be included in the request.
	// The request body sent will be `{}`.
	req.PublishCdnParams = []*trtc.McuPublishCdnParam{}

	response, err := client.UpdatePublishCdnStream(req)
	if err != nil {
		fmt.Printf("fail to invoke api: %v \n", err)
	}

	fmt.Println(response.ToJsonString())
}

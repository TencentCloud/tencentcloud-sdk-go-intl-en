package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"sort"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

const (
	SHA256 = "HmacSHA256"
	SHA1   = "HmacSHA1"
)

func Sign(s, secretKey, method string) string {
	var hashed hash.Hash
	switch method {
	case SHA256:
		hashed = hmac.New(sha256.New, []byte(secretKey))
	default:
		hashed = hmac.New(sha1.New, []byte(secretKey))
	}
	hashed.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func signRequest(request tchttp.Request, credential CredentialIface, method string) (err error) {
	if method != SHA256 {
		method = SHA1
	}
	params := request.GetParams()
	secId, secKey, token := credential.GetCredential()
	params["SecretId"] = secId
	if len(token) != 0 {
		params["Token"] = token
	}
	params["SignatureMethod"] = method
	delete(params, "Signature")
	s := getStringToSign(request)
	signature := Sign(s, secKey, method)
	request.GetParams()["Signature"] = signature
	return
}

func getStringToSign(request tchttp.Request) string {
	method := request.GetHttpMethod()
	domain := request.GetDomain()
	path := request.GetPath()

	var buf bytes.Buffer
	buf.WriteString(method)
	buf.WriteString(domain)
	buf.WriteString(path)
	buf.WriteString("?")

	params := request.GetParams()
	// sort params
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

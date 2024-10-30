package blockeditor

import "net/url"

func query(queryData map[string]string) string {
	if len(queryData) < 1 {
		return ""
	}

	urlValues := url.Values{}
	for key, value := range queryData {
		urlValues.Set(key, value)
	}

	return httpBuildQuery(urlValues)
}

func httpBuildQuery(queryData url.Values) string {
	return queryData.Encode()
}

package youtube

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

/*{
    "kind": "youtube#channelListResponse",
    "etag": "GnCS2l5ScPvua0fsDTz5E7boLZQ",
    "pageInfo": {
        "totalResults": 1,
        "resultsPerPage": 5
    },
    "items": [
        {
            "kind": "youtube#channel",
            "etag": "UiFpG12dLcDxdUOMxRRUpS7cT2U",
            "id": "UCV9WL7sW6_KjanYkUUaIDfQ",
            "statistics": {
                "viewCount": "23695712",
                "subscriberCount": "126000",
                "hiddenSubscriberCount": false,
                "videoCount": "139"
            }
        }
    ]
}*/

func TestRequest(t *testing.T) {
	resp, err := request("channels")
	fmt.Println("Request-Response:", resp)

	assert.Equal(t, err, nil)
}

func TestGetSubscriberCount(t *testing.T) {
	resp, err := request("channels")

	fmt.Println(resp.Items)
	fmt.Println(resp.Items[0].Stat.SubscriberCount)

	assert.Equal(t, err, nil)
}

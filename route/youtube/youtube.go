package youtube

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind string     `json:"kind"`
	Id   string     `json:"id"`
	Stat Statistics `json:"statistics"`
}

type Statistics struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

func createYoutubeReq(resource string) (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/"+resource, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("part", "statistics")
	q.Add("id", os.Getenv("YOUTUBE_CHANNEL_ID"))
	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func request(req *http.Request) (Response, error) {
	var response Response

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return response, err
	}
	//fmt.Println("response status: ", resp.Status)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("Body:", string(body))
	if err != nil {
		log.Fatalln(err)
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
		return response, err
	}
	return response, nil
}

func GetYoutubeSubscribers(c *gin.Context) {
	req, _ := createYoutubeReq("channels")
	resp, err := request(req)
	if err != nil {
		c.JSON(400, gin.H{
			"message":         "failure",
			"subscriberCount": nil,
		})
	}

	c.JSON(200, gin.H{
		"message":         "success",
		"subscriberCount": resp.Items[0].Stat.SubscriberCount,
	})
}

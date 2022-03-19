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
	Items []item `json:"items"`
}

type item struct {
	Kind string     `json:"kind"`
	Id   string     `json:"id"`
	Stat statistics `json:"statistics"`
}

type statistics struct {
	ViewCount             string `json:"viewCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

func request(resource string) (*Response, error) {
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

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer resp.Body.Close()
	//fmt.Println("response status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("Body:", string(body))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &response, nil
}

func GetYoutubeChannelStat(c *gin.Context) {
	resp, err := request("channels")
	if err != nil {
		c.JSON(400, gin.H{
			"message":         "failure",
			"subscriberCount": nil,
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
		"stat":    resp.Items[0].Stat,
	})
}

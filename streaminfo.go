package niconamanotify

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const getStreamInfoURL = "http://live.nicovideo.jp/api/getstreaminfo/lv%v"

func getStreamInfo(lv string, c *http.Client) (*StreamInfo, error) {
	if c == nil {
		c = http.DefaultClient
	}
	resp, err := c.Get(fmt.Sprintf(getStreamInfoURL, lv))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := &StreamInfo{}
	data, _ := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(data, info); err != nil {
		return nil, err
	}
	return info, nil
}

type StreamInfo struct {
	Status     string `xml:"status,attr"`
	RequestID  string `xml:"request_id"`
	StreamInfo struct {
		Title            string `xml:"title"`
		Description      string `xml:"description"`
		ProviderType     string `xml:"provider_type"`
		DefaultCommunity string `xml:"default_community"`
		MemberOnly       int    `xml:"member_only"`
	} `xml:"streaminfo"`
	CommunityInfo struct {
		Name      string `xml:"name"`
		Thumbnail string `xml:"thumbnail"`
	} `xml:"communityinfo"`
}

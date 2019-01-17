package niconamanotify

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const getAlertInfoURL = "https://live.nicovideo.jp/api/getalertinfo"

func getAlertInfo(c *http.Client) (*alertInfo, error) {
	if c == nil {
		c = http.DefaultClient
	}
	resp, err := c.Get(getAlertInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := &alertInfo{}
	data, _ := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(data, info); err != nil {
		return nil, err
	}
	return info, nil
}

type alertInfo struct {
	Status   string `xml:"status,attr"`
	Time     int    `xml:"time,attr"`
	UserID   string `xml:"user_id"`
	UserHash string `xml:"user_hash"`
	MS       struct {
		Addr   string `xml:"addr"`
		Port   string `xml:"port"`
		Thread string `xml:"thread"`
	} `xml:"ms"`
}

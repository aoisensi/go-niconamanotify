package niconamanotify

import (
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
)

var DefaultNotifyer = &Notifyer{}

var chatRegexp = regexp.MustCompile("<chat[^>]+>(.*)</chat>")

type Notifyer struct {
	info    *alertInfo
	Handler func(*StreamInfo)
	client  *http.Client
	conn    net.Conn
}

func (n *Notifyer) Init(client *http.Client) error {
	if client == nil {
		client = http.DefaultClient
	}
	info, err := getAlertInfo(client)
	if err != nil {
		return err
	}
	n.info = info
	return nil
}

func (n *Notifyer) Listen() error {
	ms := n.info.MS
	addr := fmt.Sprintf("%v:%v", ms.Addr, ms.Port)
	var err error
	n.conn, err = net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer n.conn.Close()
	thread := fmt.Sprintf("<thread thread=\"%v\" version=\"20061206\" res_from=\"-1\"/>\x00", ms.Thread)
	_, err = n.conn.Write([]byte(thread))

	if err != nil {
		return err
	}
	buf := make([]byte, 4096)
	for {
		cn, err := n.conn.Read(buf)
		if err != nil {
			return err
		}
		x := string(buf[:cn])
		if !strings.HasPrefix(x, "<chat ") {
			continue
		}
		chat := struct {
			ID string `xml:",chardata"`
		}{}
		if err = xml.Unmarshal([]byte(x), &chat); err != nil {
			return err
		}
		ids := strings.Split(chat.ID, ",")

		liveID := ids[0]
		//comID := ids[1]
		//userID := ids[2]
		go func() {
			info, err := getStreamInfo(liveID, n.client)
			if err != nil {
				return
			}
			if n.Listen == nil {
				return
			}
			n.Handler(info)
		}()
	}
}
func (n *Notifyer) Close() error {
	if n.conn == nil {
		return nil
	}
	return n.conn.Close()
}

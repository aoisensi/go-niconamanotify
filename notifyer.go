package niconamanotify

var DefaultNotifyer = &Notifyer{}

type Notifyer struct {
	info *alertInfo
}

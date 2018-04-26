package flexvolume

type Status string

const (
	StatusSuccess      Status = "Success"
	StatusFailure      Status = "Failure"
	StatusNotSupported Status = "Not Supported"
)

type FlexVolume interface {
	Init() Response
	GetVolumeName(map[string]string) Response
	Attach(map[string]string) Response
	WaitForAttach(string, map[string]string) Response
	IsAttached(map[string]string, string) Response
	Detach(string, string) Response
	MountDevice(string, string, map[string]string) Response
	UnmountDevice(string) Response
	Mount(string, map[string]string) Response
	Unmount(string) Response
}

type Response struct {
	Status     Status `json:"status"`
	Message    string `json:"message"`
	Device     string `json:"device"`
	VolumeName string `json:"volumeName"`
	Attached   bool   `json:"attached"`
}

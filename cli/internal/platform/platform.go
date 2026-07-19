package platform

import "runtime"

type Info struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

func Current() Info {
	return Info{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
}

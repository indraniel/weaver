// Code generated by go-bindata.
// sources:
// views/base.html
// Rscripts/install-packages.r
// Rscripts/knit.r
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _viewsBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\xed\x6e\xdb\x36\x17\xfe\x9d\x5c\x05\x5f\xbe\x43\x95\x60\x91\x28\x07\xed\x36\xd8\x96\x8b\x2d\x59\x97\x0d\x31\x5a\x2c\x2e\xb0\x2d\x0b\x0a\x4a\xa2\x25\xba\x14\x29\x90\xf4\x17\x82\xfc\xdf\x55\xec\xe2\x76\x25\x3b\x14\x6d\xd8\x8e\xac\xa1\x18\xb6\x31\x08\xcc\xaf\xc3\xf3\x3c\xe7\x53\xc3\xff\x5d\xbf\xbd\x9a\xfc\xfc\xee\x5b\x54\xda\x4a\x8c\x4e\x87\xee\x07\x09\x2a\x8b\x04\x33\x89\x47\xa7\x08\xc6\xb0\x64\x34\xf7\xd3\x66\x59\x31\x4b\x51\x56\x52\x6d\x98\x4d\xf0\xfb\xc9\x9b\xf0\x2b\x8c\xc8\xde\x05\xc1\xe5\x47\xa4\x99\x48\x30\x15\x96\x69\x49\x2d\xc3\xc8\xae\x6b\x06\x1b\x75\x2d\x78\x46\x2d\x57\x92\x68\x63\x3e\x5f\x55\x02\x8e\xb8\x15\x70\xf6\xe3\xdd\x1d\x46\xa5\x66\xd3\x04\x97\xd6\xd6\x7d\x42\x96\xcb\x65\x54\x51\x9d\xa9\x48\xe9\xc2\xdd\x3f\xd4\xd3\x00\x91\xb4\x02\xd9\x05\x67\xcb\x5a\x69\x8b\x51\xa6\xa4\x65\x12\x80\x2d\x79\x6e\xcb\x24\x67\x0b\x9e\xb1\xb0\x59\x5c\x20\x2e\xb9\xe5\x54\x84\x26\xa3\xa0\xb1\x17\xc5\x5d\xc0\x8d\x5d\x0b\x66\x4a\xc6\xec\x16\xb9\x65\x2b\x4b\x32\x07\x61\x07\xd1\x00\xc6\x2c\x97\x91\xa6\xcb\x82\xdb\x28\x53\x15\xe1\x32\xd7\x54\x72\x26\xc8\x92\xd1\x05\xd3\x64\x11\x47\xf0\x47\xa8\x01\x6b\x19\x62\x2c\x70\xcf\x08\xcf\x99\x43\x41\xa5\x89\xb2\xe7\xac\x1a\x10\x8d\x8e\x60\x63\x86\x29\x50\x32\x51\xa1\x54\x21\x18\xad\xb9\x69\x14\x81\xdc\xeb\x29\xad\xb8\x58\x27\xdf\x4b\x20\x6d\x94\xa0\x96\xf6\x5f\xc6\xf1\xc5\x97\x71\x1c\x34\x34\x82\x1d\x8d\xc0\xd3\x08\xb6\x34\x82\x7f\x8b\x76\x95\x87\x52\x59\x16\x5a\xa5\x84\x71\xec\x7b\xd1\xa5\x7b\x80\xc0\xd5\x72\x9e\x86\x25\x2f\x4a\x01\xff\xb6\x83\x79\x0b\xc7\x3f\x60\xee\x94\x17\x53\xa5\x6c\x98\xb3\x29\x9d\x0b\xaf\xf9\x3f\x51\xec\x9f\x6b\xe9\x6b\x02\x7e\xf4\xf8\x88\xa2\x89\x9b\xa1\xa7\xa7\x21\xf1\x7b\xa7\x3e\xe3\xc8\x2e\xe5\x86\xa9\xca\xd7\x7b\xb2\x39\x5f\x20\x9e\x27\xb8\xa2\xc6\xba\x5b\x29\xcd\x3e\x16\x5a\xcd\x65\x8e\x47\x2f\x64\x6a\xea\xc1\x90\xc0\x9d\x9d\xc4\x4e\xd4\xb0\xcc\x25\x5e\x23\x5e\x2b\x63\x0d\xde\xe8\xeb\x7c\x1c\x8f\x0e\xce\x7d\x35\xe8\xb5\x37\x9b\x03\xba\xb1\xd8\xff\xf1\x33\x6a\xb4\x43\xc0\xd4\xd4\x83\xb1\x5c\xae\x69\xaa\xe6\x2e\x7d\x05\x58\x30\xc1\x53\xfa\x41\x2a\x5d\x51\x71\x04\xc1\x76\x78\xb6\x2f\xd2\xb9\x10\x83\x2d\x73\xf7\xe4\x11\xcc\xe4\x18\xe8\xf6\xb5\xba\x41\x93\x33\x93\x69\x5e\x3b\x53\x75\x69\x07\x7e\x4a\xa3\xe8\x1d\xd5\xb4\x32\xd1\x5d\x09\x95\xe7\x7a\x27\x85\x30\x06\xde\x47\x50\xd4\x87\xcf\x6d\x1c\xf5\xcc\x07\x54\x43\xe4\xb8\x50\x68\xc9\x3b\x8f\x30\xdd\x61\xcc\xba\xdb\x50\x7f\xfc\xf6\x7b\x9b\xec\x76\x5c\x69\x06\xf5\x39\xef\xbb\xa8\xac\x18\xca\x61\xe1\x26\x49\xe0\x7c\x78\x0d\xab\xe8\x8d\x73\x84\x45\xf8\x32\x8e\xbf\x08\xe3\x5e\x18\x5f\x4e\x7a\xaf\xfa\xf1\xcb\x7e\xfc\xea\x17\x47\x34\x40\xf5\x3c\x75\x72\x10\x55\x7e\xe2\x03\xa0\x43\x18\x6f\xc2\xbd\x3a\x42\x71\x3b\xbc\x53\x3b\x8f\xdf\xd7\x79\x37\xe6\xcd\xe1\xdf\x86\xdd\x2d\xff\x29\xc8\x3b\x6d\xdd\xf2\xbe\xdf\xec\xf2\xa9\x43\x72\x33\x19\xdf\x7e\x03\xd9\xff\x3c\x98\x20\xa3\xba\x62\x64\x3b\x4e\x4f\x4e\x4e\x0e\xd6\x07\xf2\xae\x18\x32\x8d\x9a\xea\xe4\x4b\x7c\x48\xa1\x26\xcb\x3e\xca\xa0\x75\x32\x3d\xe8\x88\x7b\x08\x32\xe0\x77\x94\xc9\xe6\xf8\x3b\x26\x99\x76\xe6\x43\xe9\x1a\xf9\xca\x78\x24\xec\xbd\xfe\xbd\xe2\x03\x89\xeb\x6b\xd3\xfe\x9e\xcf\xa7\xfd\x3e\xb4\x0a\xc1\x23\xe5\x8c\xae\x42\x68\x78\x53\x5e\x1c\xc0\x1c\xc3\xd1\x0f\x74\x15\xdd\xcc\xd3\xe8\xaa\x39\x3e\x7b\x6c\xc1\x84\x67\x2e\x41\xbe\x8f\x1e\xb9\x84\xaa\xcf\x9c\x50\x1f\xdd\xdf\x07\x9f\x05\x17\xf0\xff\x70\x81\xee\x83\x5f\x61\x9c\xc1\xd2\xfd\x9e\x07\x0f\x0f\xfb\xc6\x7f\x3a\x1f\xec\x83\x6e\x10\x8e\xfe\x0a\xf2\x8c\x2e\xa8\xdf\xc5\x07\x60\x8c\xce\x0e\x7b\xcb\x86\x59\xf3\xa5\xb3\x99\x13\xe8\xe8\xcc\x58\xb2\x65\x36\x33\xaf\x3d\xef\x64\xc2\x7e\x0a\xbf\x1e\xdf\x85\xe3\xf1\xed\x07\x17\x25\x4a\xc3\x6c\xcf\x1a\xdd\xd0\x5a\x7a\x3f\xbd\x89\xcf\x0c\xd9\x35\xef\x1a\xda\x0e\x00\xc2\x1d\x5c\x47\x9d\x08\x46\xa5\x98\x99\xc8\x7d\x88\xdd\x6c\x1f\xe3\xb2\x78\x2b\x6f\x15\xcd\xcf\xce\x07\x87\x72\x43\xe2\xbb\x1f\xe4\x49\xf3\x7d\xfa\x67\x00\x00\x00\xff\xff\x8d\x96\x7f\x67\xb0\x0a\x00\x00")

func viewsBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsBaseHtml,
		"views/base.html",
	)
}

func viewsBaseHtml() (*asset, error) {
	bytes, err := viewsBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/base.html", size: 2736, mode: os.FileMode(420), modTime: time.Unix(1440098365, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _rscriptsInstallPackagesR = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\xdb\x6e\xdb\x38\x13\xbe\xf7\x53\x30\x2a\x02\x4b\x80\x4d\xc5\x49\x4f\x28\x7e\xfd\x40\x0f\x7b\xb1\x8b\xa2\x2d\x9c\x76\x6f\x82\xa0\xa0\x25\x5a\xe2\x9a\x22\xb5\x24\xe5\xac\xbb\xed\xbb\xef\x90\x12\xad\x43\xba\x4d\x36\x6a\x81\xa0\xa2\x38\xf3\xcd\xf0\x9b\xe1\x68\xc6\x8f\x4e\xe2\x5a\xab\x78\xc3\x44\x4c\xc5\x1e\xad\x75\xaa\x58\x65\x66\xb3\xd9\x23\xb4\x55\xb2\x44\x85\x31\x95\x7e\x11\xc7\x39\xd3\x06\xe7\xcc\x14\xf5\x06\xa7\xb2\x8c\xf7\x6c\xc7\x44\x1e\xaf\x9e\x9c\x5d\x3c\xbb\x78\x3a\x63\x42\x1b\xc2\x39\xae\x48\xba\x23\x39\xc5\x7b\xaa\x34\x93\x02\xfd\x6f\x89\xb6\xb5\x48\x0d\x3c\x87\x33\x84\xda\xed\x05\x3c\x7a\x89\x04\xbd\xfb\xf4\xf6\xad\x7d\xa3\x68\x25\x35\xac\x73\x6a\xde\x57\x4e\x63\xee\x5e\xcd\x23\xbb\x6b\x0e\x15\x1d\x6c\x06\xd5\x2e\xff\x08\x2f\x83\x28\x9a\xfd\x0d\x02\xa9\x14\x46\xb1\x4d\xad\xb8\xb5\xda\xae\x30\x2c\x43\x87\xb2\x70\x08\x11\x08\x92\x3d\x61\x9c\x6c\x38\xb5\x72\xc7\x85\x77\x5d\x87\x1d\x90\x95\x66\x5b\x14\xb6\x5b\xe8\x94\x89\x53\xa4\xe4\x0d\x16\xa4\x04\xc1\xa3\x6e\x14\x21\xeb\x01\xf8\x50\x2b\x45\x85\xe9\x1f\xff\x28\x74\xe5\x0f\x8f\xe6\xbf\x37\xdb\xf3\x6b\xa7\x64\x2d\x30\x8d\x45\xcd\x79\xd8\x2a\x46\xe8\xeb\xd7\x8e\xa1\x64\x8c\xeb\xcd\x81\xee\x90\x77\x1d\x1e\x8d\x78\x36\xdb\xc3\xf7\xd8\x49\x7a\x8b\x85\xe7\xd5\x93\x63\xff\x29\x6a\x6a\x25\xc2\x66\xf9\x6d\x66\xff\x1a\x7e\xed\x79\xf2\x2f\xf0\x10\x5a\x5a\x75\xa5\x98\x30\xdb\x30\x38\xd5\xb1\x56\x69\xdc\xa2\xc6\x2f\x55\x5a\xb0\x3d\xc5\x2a\xd3\x41\xeb\x47\xb4\x40\x81\xda\xd8\x48\x01\xfd\xcd\xb6\xc5\x52\x94\x64\xeb\x37\x97\x96\x71\xbb\x93\x72\xa9\x69\xb3\x68\x89\x3f\x19\x52\xdf\xd2\xde\x00\x44\x47\xd6\xb5\x91\x55\xe7\x4d\x2a\x6b\x9e\x89\xb9\x41\x5b\x26\x32\x9f\x71\x68\x7e\xaa\xe7\xe0\x4d\xbb\x74\x8e\xd8\x93\x31\xb1\x95\x2e\x4a\x0d\xe6\x95\x0f\xd2\xf5\xb5\x77\x61\x1c\x19\x6f\xf4\x11\xfa\xad\xd6\x06\xe5\x8a\x6c\x90\x29\x28\xe2\xc4\x50\x58\x4b\x41\x31\xfa\x58\x30\x8d\x6e\x18\xe7\xb0\xe4\x07\x54\x90\xaa\xa2\xc2\xa2\x59\xc1\xd6\x42\x0b\x52\x10\x8d\x36\x14\x76\x2b\x30\x42\xb3\xe6\xd2\xbd\x5e\xbf\x7c\x87\x9d\x80\xbf\x4f\x15\x31\x85\xf5\xd3\xfa\x7b\xc5\xa9\xc8\x4d\x11\xda\xe7\xc8\xa6\xd0\x37\x44\xb9\xa6\xad\x5f\x63\x8d\x8a\x68\x43\xbb\xb4\x08\xe2\x8e\x05\x58\x7d\x86\x55\x7b\x32\x58\x61\x43\x14\xce\xbf\xc0\x3b\x4d\xab\x24\x08\xa2\x63\x7e\x1e\x03\xd1\xe0\xba\x68\x38\xfb\x5d\x2a\x0e\xa3\xe0\x93\xd7\xf2\x8e\x80\x0d\x26\xf6\x84\x33\x38\x9f\x54\xe3\x98\x1c\xed\xf7\x83\xe3\x13\xaf\x3b\x50\x7b\xb1\xef\xc8\xba\xf8\xf4\x98\x74\x8b\x01\x17\x16\x94\xcb\x94\xf0\x23\x33\x5b\xc6\x9b\xad\xd0\xd0\xb2\xca\x98\x0a\x21\x4d\x37\x44\x53\x9b\x66\x83\xf3\x46\xbe\x10\x64\xf2\x46\x70\x49\x32\x6c\x75\xc3\x9e\x67\x8b\x1e\x76\x84\x4e\x12\x74\x36\x48\xce\x2e\x29\x3d\x82\xb3\xfe\x02\x75\xc1\xc0\x6d\xbd\x71\x77\xed\xd6\xb5\xee\xd0\x6f\xdd\xec\xc1\x05\xfe\x66\x0b\xf7\xa5\x21\x22\x23\x2a\xf3\x30\x33\x6d\x32\xdc\x3e\x0f\x2a\x72\xe3\xe3\x2d\x63\xc1\x4e\x30\xa3\x3c\x8d\x49\x60\xbf\x01\xf0\x09\x48\x15\x11\xf8\x06\x92\x9e\x63\x9a\xd5\x2e\x39\x6e\xab\xaa\x3f\xb4\x14\x0f\x53\xcd\xf3\x8a\x4b\x73\xfe\x30\x65\x59\x99\x8a\x28\x4d\xef\xd6\x76\x14\x7d\x60\x42\xd0\x8e\xa0\xca\x2d\xef\xcf\x91\x2f\xc4\x61\x50\xb2\x92\x76\x39\x9c\x04\x67\xf8\xe2\x01\x07\xe8\xf0\x34\x24\xb3\xc8\xd9\x10\xf2\xc9\xf2\xc9\x24\xd0\x92\xe4\x8a\x19\x17\xd2\x23\xea\x0a\x4f\xc3\xa4\x70\x9f\x6b\xa8\x79\x43\x4f\x9f\xe1\x87\xc4\xaf\x43\xcd\x18\x44\xd3\x0c\x31\x9f\xe2\xe7\x93\x30\xa1\xe4\x94\xc4\xac\x87\x87\x9f\xe6\x66\xc1\xf2\x42\x8d\x62\x34\x31\x42\x6a\x67\x4b\xc3\x98\xcd\x67\x3f\x21\x99\x46\x61\x3f\xc3\x67\x93\x40\x0f\xa4\xe4\x7d\xc4\x73\xbc\xc2\xab\x69\x49\xef\xeb\x4d\xcf\xcb\xd5\x6a\x12\xa2\x2f\x43\x3d\x32\xc1\xcf\x69\x31\x4a\x25\x97\x4a\xc3\x6b\x3a\x4a\xa5\xe5\xd3\x49\xb8\xeb\xb4\xaa\x86\xae\xae\xce\x27\xc6\x68\xfd\xda\xfa\xfa\x4a\xd1\x1b\x3a\xe6\x75\x39\xf5\x82\xa6\x05\xb4\x28\xc4\x0c\x53\xe0\x6c\x39\xcd\xe1\xb2\x16\x9a\x72\x3e\xa4\xe1\xf1\xc4\x5b\x0a\x6d\x37\xe5\x70\x01\x7e\x66\x7d\xae\xf8\x61\xc4\xe8\xf3\x89\x88\xb9\xb1\xd3\xc1\x28\x01\x26\x9e\x5c\x51\x0d\x4d\x27\x3d\x1f\x7a\xfa\x18\x4f\xbb\x54\x1a\x7a\x10\xaa\xc7\xb7\x6a\xda\xa5\xaa\x94\x34\x72\x14\xa0\xe5\x6a\x5a\x32\x75\x8d\xc4\xa0\xec\x4d\x3b\x3c\x0c\x9e\xd0\x63\x8c\x6e\xfe\xd4\x5a\xda\xeb\x5a\x7a\xb0\x17\xf7\x41\x75\x6d\x4c\x3b\x74\x64\x74\x4f\xb9\xac\x4a\x18\x11\x11\xcd\x98\xed\x5b\x34\x5a\x2e\x51\x49\x0e\x48\x48\x03\x83\x15\x25\x76\x10\xf1\xed\x0e\xc8\xff\x97\x7e\x70\x0e\xf2\x46\x4a\xae\xe7\xf7\x38\x2b\x67\x1b\x45\xd4\x21\xf4\x3a\xbd\xf3\x7f\x6e\x7e\x3f\x80\x4f\x08\x2b\x6a\x16\x77\x5d\xe6\x36\x81\x0f\x20\x8c\x29\x2a\xf8\x9e\x34\x24\xdd\x5f\xe9\x26\x76\x25\x3d\xee\xfa\xcb\x1f\x2b\x15\x24\xe3\xf4\x10\xf7\x7b\xca\x1f\x2b\x18\x45\xf7\x52\xf1\x2c\x1e\x74\x92\x7d\x1d\xc7\xb8\x1d\x7a\x6c\x3b\x61\x27\xd8\x8c\xe9\x8a\x93\xc3\xac\xfd\xdf\x53\xaa\xbf\xc3\x29\x4c\x17\xf6\xed\xe5\x41\x63\x48\x24\x2a\xf6\x61\x0a\x55\xfa\xf3\xdb\x5f\x5f\x5d\x36\x53\x71\xb5\xcb\x75\x33\xe0\x39\x0c\x68\x46\xbb\xde\x9f\x6d\x30\xf4\xff\x49\x80\x63\x65\xc5\x9c\xfb\x6e\x0a\x0a\xed\xf2\x6a\x01\x50\x1f\x1a\x61\xf0\x39\x68\x7f\x5c\xb0\x8f\x6f\x28\xcc\x9d\x30\x85\x47\xd7\xce\x7b\xa2\x1a\x1b\xa9\x2c\x4b\x98\x10\x5e\xc2\x32\x34\x8a\x30\x5b\x23\xdf\x43\x7e\x24\x1f\xd7\x9f\x7e\x81\xd9\xdb\x4e\x3a\x4e\x36\x49\x50\x00\x71\x0c\x9a\x33\x34\x26\x83\xe5\xf2\xff\xe8\x03\x55\x96\x06\xd0\x43\x04\xd9\xf4\x98\x7b\xc7\x9d\x73\xbd\xfc\x0a\xc1\x72\x33\x9e\x0e\x50\x9b\x86\xfb\x2e\xe0\x46\x6a\x88\x3d\x6c\xd5\xff\x05\x5e\xb7\x13\xd0\x5d\x06\xbc\xdc\xd0\x44\x6f\x5c\xea\xf0\x7b\x38\x6f\xa4\x9d\xe4\x76\x42\xde\xa0\x02\xfe\x8c\xf4\xda\x27\x4e\xfd\xcf\x9a\x99\x50\x93\x3d\x4d\x02\x61\x4b\x1b\x6c\x99\x5a\x27\x2b\x17\x83\x94\x40\xd4\xdc\x48\x1e\xd8\xc0\x8f\x53\x07\xec\xfd\x13\x00\x00\xff\xff\x9e\x90\xcf\x03\x92\x13\x00\x00")

func rscriptsInstallPackagesRBytes() ([]byte, error) {
	return bindataRead(
		_rscriptsInstallPackagesR,
		"Rscripts/install-packages.r",
	)
}

func rscriptsInstallPackagesR() (*asset, error) {
	bytes, err := rscriptsInstallPackagesRBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Rscripts/install-packages.r", size: 5010, mode: os.FileMode(493), modTime: time.Unix(1440092183, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _rscriptsKnitR = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x53\x41\x6b\xdc\x3c\x10\xbd\xeb\x57\xe8\xd3\x06\x56\x06\xef\x06\xbe\x63\xa8\x0f\x85\x12\xe8\x25\x29\xa5\xb7\x34\x18\xad\x3d\xb6\x45\x6c\xc9\x8c\xb4\xbb\xd9\x7f\xdf\x19\xd9\xde\x35\x4d\xe9\xa5\xbe\x18\x6b\xde\xcc\x7b\x7a\x6f\xbc\xf9\xef\xfe\x18\xf0\xfe\x60\xdd\x3d\xb8\x93\xfc\x1e\x2a\xb4\x63\x14\xa2\xb7\x07\x34\x78\xd1\xd1\xfb\x3e\x64\xd7\x4f\x3f\xc6\xd1\x60\x80\xdb\xc9\x9b\xb3\x11\x33\x21\xa8\x12\xf6\xbd\x0d\x51\x7e\xda\x49\x7e\x6b\x21\xe9\x19\xcc\x1b\x94\x54\xb3\xde\xe9\x4a\xab\x9d\x57\xb9\x54\xbb\x9d\x3f\xc6\xda\xa2\xca\x72\x69\x2a\xae\x15\x2a\x44\x8f\x40\xc5\xd4\xb5\x7e\xe2\x65\x84\x42\x55\x9d\x41\x82\x02\x12\xa4\x86\xc6\x1c\xfb\x58\xa8\xc6\x3a\xd3\xab\xfc\x43\x4b\x07\xfd\x58\xa8\x1f\x1d\x48\xe2\x19\x8f\x51\x12\x17\x54\x44\x70\x91\xd1\xcb\xb1\x37\x15\xc8\x48\xd5\xc6\xb6\x47\x84\x20\x5f\xe6\x89\x0f\x72\x9b\x46\x6e\x5f\x49\x99\xf8\xa3\x7e\x3b\xe9\xb7\x8e\xc6\xfe\xa3\xfc\xa7\xcf\x7f\x53\x9e\x18\x88\x1e\xdf\x6a\x7f\x76\x24\xb5\x87\x24\x1e\x7d\x05\x21\xa8\x4c\xcc\x9e\xb3\xdd\x29\x92\xd2\x60\x1b\xf4\x73\x92\xfa\x8d\x0f\x50\x4f\xba\x4b\x4e\xa3\xb8\xe6\x93\x51\x63\x65\xa2\x1e\x4d\x88\xa0\xd5\xd7\xc4\xf3\x48\xe3\x1f\x24\x89\x63\xd8\x5d\xe2\xa6\x7b\xfe\x74\x8a\xd0\x2b\xf0\xf3\x64\xe7\x63\xf2\x4d\x7e\x59\x5c\xbd\x75\x4e\xb9\x5e\x5b\x93\xc2\x92\x57\xe4\x2e\x40\xd4\x07\x13\x60\x4f\xf5\x62\x85\xcd\xe4\x46\xf2\x39\x67\x24\xcf\x1d\x60\xba\x66\x30\xa7\x6b\x3c\x42\x6c\xe4\x08\xd8\x78\x1c\x52\x6a\x08\xae\x06\xb4\xae\x15\xd6\x9d\x6c\xb0\x87\x1e\xd2\x16\xea\x9b\x76\xa6\xde\xc8\xc1\x9f\xa6\x9c\xe7\x2d\x58\x3c\xec\x3d\x5d\x89\x8c\x91\xbe\x99\xca\x68\x5b\x8e\x7d\xf6\xfc\xba\x2c\x82\x75\xb1\x2c\xb2\x98\x5e\xce\x0c\xb0\xe6\x10\xdc\xc7\xc6\x71\x9d\xa1\x1f\x00\xe8\x7d\x7c\xa2\x43\x06\x30\x79\x39\x9a\xd8\x95\xc1\xb8\x50\xc2\x3b\xe9\x9d\xfb\x33\x31\xd4\xcb\x9c\xc9\xe8\xa5\x31\x97\xdb\xfd\x50\x6f\x73\x19\x80\xd6\x82\x42\x17\x3c\x26\x2c\xbf\xd8\x3e\x7d\xe9\x16\xe2\xb9\xd6\x74\x67\x4b\xc6\xbc\x73\xb1\x45\x18\xf5\x34\x34\x4f\xcc\x21\x5b\x99\x95\xb6\x8e\x4f\xf7\x38\x69\xbe\xad\x61\x83\x7e\x28\x52\x89\xa5\x2e\x93\xe7\x19\x2f\x69\xfe\x6b\xb6\x5a\xdb\xe8\x57\xe8\xd9\xae\xdf\xd0\x09\x9c\xd6\x75\xc3\x29\xfd\xdf\xc5\xa1\xd7\x4d\xba\xf6\xaf\x00\x00\x00\xff\xff\xb6\xe2\xd5\xfd\x7d\x04\x00\x00")

func rscriptsKnitRBytes() ([]byte, error) {
	return bindataRead(
		_rscriptsKnitR,
		"Rscripts/knit.r",
	)
}

func rscriptsKnitR() (*asset, error) {
	bytes, err := rscriptsKnitRBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Rscripts/knit.r", size: 1149, mode: os.FileMode(493), modTime: time.Unix(1440092183, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"views/base.html": viewsBaseHtml,
	"Rscripts/install-packages.r": rscriptsInstallPackagesR,
	"Rscripts/knit.r": rscriptsKnitR,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"Rscripts": &bintree{nil, map[string]*bintree{
		"install-packages.r": &bintree{rscriptsInstallPackagesR, map[string]*bintree{
		}},
		"knit.r": &bintree{rscriptsKnitR, map[string]*bintree{
		}},
	}},
	"views": &bintree{nil, map[string]*bintree{
		"base.html": &bintree{viewsBaseHtml, map[string]*bintree{
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, filepath.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var _viewsBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\xed\x6e\xe4\x34\x14\xfd\xbd\x7d\x0a\x63\xd0\xa6\x15\x4d\x9c\x56\xbb\x80\xa6\x93\x59\x41\xcb\x52\x50\xab\x5d\xd1\x59\x09\x28\xd5\xca\x93\x78\x12\x77\x1d\x3b\xb2\x3d\x33\x1d\x55\xfd\xcf\x53\xf0\x70\x3c\x09\xd7\x71\xc2\x7c\x1a\x01\x02\x2c\x45\xf1\xd7\xf1\x3d\xe7\xfa\xfa\xda\xc3\x8f\x2e\xde\x9c\x8f\x7f\x7c\xfb\x35\xaa\x6c\x2d\x46\x07\x43\xf7\x43\x82\xca\x32\xc3\x4c\xe2\xd1\x01\x82\x32\xac\x18\x2d\x7c\xb5\x6d\xd6\xcc\x52\x94\x57\x54\x1b\x66\x33\xfc\x6e\xfc\x3a\xfe\x02\x23\xb2\x36\x41\x70\xf9\x01\x69\x26\x32\x4c\x85\x65\x5a\x52\xcb\x30\xb2\xcb\x86\x41\x47\xd3\x08\x9e\x53\xcb\x95\x24\xda\x98\x4f\x1f\x6a\x01\x43\xdc\x0a\x18\xfb\xfe\xe6\x06\xa3\x4a\xb3\x69\x86\x2b\x6b\x9b\x01\x21\x8b\xc5\x22\xa9\xa9\xce\x55\xa2\x74\xe9\xe6\x6f\xda\x69\x89\x48\x5a\x03\x76\xce\xd9\xa2\x51\xda\x62\x94\x2b\x69\x99\x04\x62\x0b\x5e\xd8\x2a\x2b\xd8\x9c\xe7\x2c\x6e\x1b\xc7\x88\x4b\x6e\x39\x15\xb1\xc9\x29\x58\x3c\x49\xd2\x10\x71\x63\x97\x82\x99\x8a\x31\xdb\x33\xb7\xec\xc1\x92\xdc\x51\x58\x51\x34\xc0\xb1\xe4\x36\x2f\x64\xe2\xa0\x44\xb3\x46\x11\x2e\x0b\x4d\x25\x67\x82\x2c\x18\x9d\x33\x4d\x6a\x6a\xc0\x0b\x84\x1a\xf0\x97\x21\xc6\x82\xfa\x9c\xf0\x82\x39\x1e\x54\x9a\x24\xdf\xd6\xd5\xd2\x68\xad\x44\x9d\x23\xa6\x20\xca\x24\xa5\x52\xa5\x60\xb4\xe1\x80\x51\xb5\x23\xf3\x6a\x4a\x6b\x2e\x96\xd9\xb7\x12\x64\x1b\x25\xa8\xa5\x83\x17\x69\x7a\xfc\x79\x9a\x46\xad\x90\x68\x25\x24\xf2\x42\xa2\x5e\x48\xf4\xdf\x09\xaf\x8b\x58\x2a\xcb\x62\xab\x94\x30\x64\x9e\x26\x27\xc9\xa9\x5b\xc2\x61\xaa\xd9\x24\xae\x78\x59\x09\xf8\x6c\x40\xfb\x0e\x93\x7f\xc5\xe5\x13\x5e\x4e\x95\xb2\x71\xc1\xa6\x74\x26\xbc\xed\xff\xc9\xb4\x5f\x70\xc7\x62\x1b\xf8\xa3\xc7\x47\x94\x8c\x5d\x0d\x3d\x3d\x0d\x89\xef\x3b\xf0\x27\x8f\xac\x8e\xde\x70\xa2\x8a\xe5\x1a\xb6\xe0\x73\xc4\x8b\x0c\x3b\x7b\x6e\xd6\x84\xe6\x1f\x4a\xad\x66\xb2\xc0\xa3\xe7\x72\x62\x9a\xb3\x21\x81\x39\x2b\xc4\x0a\x6a\x58\xee\x0e\x60\x0b\x6f\x94\xb1\x06\x77\xf6\x82\x8b\xe3\xd1\xc6\xb8\xcf\x0a\x27\xbb\x9d\xed\x00\xed\x7c\xf6\x31\xde\x92\x46\x03\x00\xd3\x50\x4f\xc6\x72\xb9\xa4\x13\x35\x73\xc7\x58\x80\x07\x33\x3c\xa5\xef\xa5\xd2\x35\x15\x7b\x18\xf4\xc5\xab\x7d\x3e\x99\x09\x71\xd6\x2b\x77\x4b\xee\xe1\x4c\xf6\x91\xde\x9d\xd6\xb4\x6c\x0a\x66\x72\xcd\x1b\xe7\xaa\x90\x75\x6f\x14\x39\x99\x6f\xa9\xa6\xb5\x49\x6e\x2a\xc8\x42\x17\x2b\x24\x08\xef\x66\xed\x61\xd3\x6c\x2e\xdb\x6d\xd8\xd6\x5e\x50\x0d\x11\xe4\x42\x62\x07\xef\x76\x86\xe9\x80\x53\x9b\xb0\xc3\x7e\xfb\xe5\xd7\x5d\xd1\x7d\x39\xd7\x0c\xf2\x75\x31\x70\xd1\x59\x33\x54\x40\xc3\x55\xb2\xc8\x89\xbc\x80\x56\xf2\xda\x6d\x88\x45\xf8\x34\x4d\x3f\x8b\xd3\x93\x38\x3d\x1d\x9f\xbc\x1c\xa4\x2f\x06\xe9\xcb\x9f\x30\xe8\x8d\x50\x33\x9b\x38\x1c\x44\x97\xaf\xf8\x40\x08\x80\x71\x17\xf6\xf5\x1e\x89\x7d\x09\x78\xb0\x2f\xef\x9a\x22\xcc\xb9\x1b\xfc\xc7\xb4\xc3\xf8\xbf\xc2\x3c\xe8\xeb\x9d\xdd\xf7\x9d\xa1\x3d\x75\x4c\x2e\xc7\xd7\x57\x5f\x41\x16\x00\xb3\x5b\x71\x13\x8c\x91\xbe\x1c\x3c\x7b\xf6\x6c\xa3\xbd\x81\x77\x69\x91\x69\xd4\x66\x29\x9f\xf0\x63\x0a\xf9\x59\x0e\x50\x0e\x57\x29\xd3\x67\x81\xf8\x87\x20\x03\x7d\x7b\x95\x74\xc3\x41\x42\xdf\x30\xc9\xb4\xf3\x2b\x9a\x2c\x83\x93\xfe\x48\x25\x6b\xe9\x17\xee\x8f\xf6\xf6\xdb\x4e\xbc\x78\xe4\xff\xe1\x34\xb3\x7b\xdc\xbc\xee\xb5\xe4\x07\x89\xc3\xe7\xc6\xf5\x3e\x7f\x96\xd7\x6f\xc3\x87\x18\x22\xa1\xba\xa7\x0f\x31\x5c\xbb\x53\x5e\x6e\xb8\xe7\x1a\x86\xbe\xa3\x0f\xc9\x25\x10\x3d\x6f\x87\x0f\x1f\x77\x08\xc1\x32\xa7\x80\x1f\xa0\x47\x2e\xe1\x32\x61\x0e\x34\x40\xb7\xb7\xd1\x27\xd1\x31\x7c\x77\xc7\xe8\x36\xfa\x19\xca\x21\x34\xdd\xff\x28\xba\xbb\x5b\xdf\xf4\xa7\xa3\xb3\x75\xd2\x2d\xc3\xd1\x9f\x51\xbe\xa7\x73\xea\x7b\xf1\x06\x19\xa3\xf3\x95\x7b\xdd\xd5\xd6\x29\x6b\x5f\x5c\x5d\x9d\xc0\xbb\x82\x19\x4b\x7a\x65\xf7\xe6\x95\xd7\x9d\x8d\xd9\x0f\xf1\x97\xd7\x37\xf1\xf5\xf5\xd5\x7b\x17\x9d\x4a\x43\x6d\xcd\x1b\x61\x6a\x1b\x76\xff\xee\x53\xe2\xde\x90\xd5\x13\xa2\x81\x8b\x0f\x28\xe1\x80\xda\x51\x90\xc3\xa8\x12\xf7\x26\x71\x4f\xc2\xcb\x7e\x31\x2e\xcb\x37\xf2\x4a\xd1\xe2\xf0\xe8\x6c\x13\x37\x24\xfe\xfe\x85\x13\xda\xbe\x94\x7f\x0f\x00\x00\xff\xff\x46\xc0\xa8\x5a\x3a\x0b\x00\x00")

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

	info := bindataFileInfo{name: "views/base.html", size: 2874, mode: os.FileMode(436), modTime: time.Unix(1468866017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rscriptsInstallPackagesR = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\x5b\x6f\xdc\x36\x13\x7d\xdf\x5f\xc1\x28\x30\x56\x02\xbc\x94\xd7\xce\x0d\x1f\x3e\x15\xc8\xed\xa1\x45\x9a\x04\xeb\xa4\x2f\x86\x11\x70\x25\xae\xc4\x2e\x45\xb2\x24\xb5\xce\xa6\xc9\x7f\xef\x90\xba\xab\x71\x92\x5a\x09\x60\x44\x14\x67\x0e\x0f\xcf\x0c\xa9\x99\xbd\x7f\x2f\xae\x8c\x8e\xb7\x4c\xc4\x54\x1c\xd0\xc6\xa4\x9a\x29\xbb\x58\x2c\xee\xa3\x9d\x96\x25\x2a\xac\x55\xe6\x7f\x71\x9c\x33\x63\x71\xce\x6c\x51\x6d\x71\x2a\xcb\xf8\xc0\xf6\x4c\xe4\xf1\xfa\xe1\xd9\xc5\xe3\x8b\x47\x0b\x26\x8c\x25\x9c\x63\x45\xd2\x3d\xc9\x29\x3e\x50\x6d\x98\x14\xe8\xff\x2b\xb4\xab\x44\x6a\xe1\x39\x5c\x20\xd4\x4c\x9f\xc2\x63\x6b\x91\xa0\xd7\xef\x5f\xbd\x72\x6f\x34\x55\xd2\xc0\x38\xa7\xf6\x8d\xf2\x1e\x4b\xff\x6a\x19\xb9\x59\x7b\x54\x74\x34\x19\xa8\x7d\xfe\x0e\x5e\x06\x51\xb4\xf8\x1b\x0c\x52\x29\xac\x66\xdb\x4a\x73\xb7\x6a\x33\xc2\x30\x0c\x3d\xca\xa9\x47\x88\xc0\x90\x1c\x08\xe3\x64\xcb\xa9\xb3\xeb\x06\x2d\x75\x13\xf6\x40\xce\x9a\xed\x50\xd8\x4c\xa1\x13\x26\x4e\x90\x96\x37\x58\x90\x12\x0c\x3b\xdf\x28\x42\x8e\x01\x70\xa8\xb4\xa6\xc2\x0e\xb7\xdf\x19\x5d\xb5\x9b\x47\xcb\x3f\xea\xe9\xe5\xb5\x77\x72\x2b\x30\x83\x45\xc5\x79\xd8\x38\x46\xe8\xf3\xe7\x5e\xa1\x64\x8a\xdb\x2e\x07\xbe\x63\xdd\x4d\xd8\x2d\xd2\xaa\xd9\x6c\x7e\xa0\x4e\x32\x18\x9c\xb6\xba\xb6\xe2\xb8\x7f\x9a\xda\x4a\x8b\xb0\x1e\x7e\x59\xb8\xbf\x5a\x5f\xb7\x9f\xfc\x13\x3c\x84\x4e\x56\xa3\x34\x13\x76\x17\x06\x27\x26\x36\x3a\x8d\x1b\xd4\xf8\x77\x6a\x49\x4c\x74\x5a\xb0\x03\xc5\x3a\x33\x41\x43\x26\x3a\x45\x81\xde\xba\x70\x41\x0c\xea\x69\x07\xa8\x29\xc9\x36\x2f\x2e\x9d\xec\x6e\x26\xe5\xd2\xd0\x7a\xd0\xa8\x7f\x6f\xac\x7f\xa3\x7d\x0d\x10\x75\xd2\x1b\x2b\x55\x4f\x29\x95\x15\xcf\xc4\xd2\xa2\x1d\x13\x59\x9b\x76\x68\x79\x62\x96\xc0\xa6\x19\x7a\x22\x6e\x7b\x4c\xec\xa4\x0f\x55\x8d\x79\xd5\x46\xea\xfa\x7a\x98\x2d\xd3\xdc\x36\x9e\xbc\xbc\xa9\xf9\x38\x8c\x8e\xf1\x34\x9a\x2d\xc7\xfb\xe8\xb7\xca\x58\x94\x6b\xb2\x45\xb6\xa0\x88\x13\x4b\x61\x2c\x05\xc5\xe8\x5d\xc1\x0c\xba\x61\x9c\xc3\x90\x1f\x51\x41\x94\xa2\xc2\xa1\x39\xc3\x66\xe9\x06\xa4\x20\x06\x6d\x29\xcc\x2a\x58\x84\x66\xf5\x41\x7d\xbe\x79\xfa\x1a\x7b\x83\x96\xa7\x22\xb6\x70\x1c\x1d\xb5\x2b\x4e\x45\x6e\x8b\x9a\xa6\xdb\xd6\x17\x44\xb9\xa1\x0d\xaf\xa9\x87\x22\xc6\xd2\x3e\x95\x82\xb8\x17\x0d\x46\x1f\x60\xd4\xec\x0c\x46\xd8\x12\x8d\xf3\x4f\xf0\xce\x50\x95\x04\x41\xd4\xe5\x74\x17\xb7\x1a\xd7\x07\xef\x76\x35\xa3\x3e\xa9\xc7\xa1\x6c\x8f\x81\x0b\x1e\x02\x8d\x98\x38\x10\xce\x60\xd7\x52\x4f\x03\xdb\xb1\x1a\x46\xb8\x4d\xe1\x7e\x9b\xcd\x15\x71\x5b\xfe\x3e\xad\xd3\x20\x3e\xe9\x32\xf7\x74\xa4\x90\xbf\x15\x6e\x39\x77\xd8\x9f\xa9\xf6\xec\xb9\x9b\x2d\x5a\x7c\x71\x97\xe9\xa5\x25\x22\x23\x3a\x6b\x3d\x17\xc6\x66\xb8\x79\x1e\xdd\x92\xb5\x0c\xff\xc2\x0f\xf6\x82\x59\xdd\x12\x4a\x02\x77\x2f\xc3\xb5\x9c\x6a\x22\xf0\x0d\x24\x15\xc7\x34\xab\x82\xaf\x52\x0b\xf4\x9f\x46\x8a\xbb\xb9\xe6\xb9\xe2\xd2\x9e\xdf\xc1\x19\x7c\x35\xcb\x5e\x7e\xb4\x9a\xdc\x6d\x69\xa9\xac\x22\xda\xd0\xef\x7b\x7b\x81\xdf\x32\x21\x68\x2f\xaf\xf2\xc3\x1f\x57\xb8\xcd\xc3\x30\x28\x59\x49\xfb\x5c\x4a\x82\x33\x7c\x71\x87\x0d\xf4\x78\x06\x92\x4a\xe4\x6c\x0c\xf9\x70\xf5\x70\x16\x68\x49\x40\x5d\xeb\x13\xa2\x43\x5d\xe3\x79\x98\x14\xce\x55\x05\x37\xd2\x98\xe9\x63\x7c\x97\xe8\xf7\xa8\x19\x83\x68\xda\x31\xe6\x23\xfc\x64\x16\x26\x1c\xfd\x92\xd8\xcd\x78\xf3\xf3\x68\x16\x2c\x2f\xf4\x24\x46\x33\x23\xa4\xf7\x19\x7c\x1d\xa6\x6a\x3e\xfe\x09\xc9\x34\x09\xfb\x19\x3e\x9b\x05\x7a\x24\x25\x1f\x22\x9e\xe3\x35\x5e\xcf\x4b\xfa\xf6\xb6\x1a\xb0\x5c\xaf\x67\x21\xb6\x97\xd8\x40\x4c\xe0\x39\x2f\x46\xa9\xe4\x52\x1b\x78\x4d\x27\xa9\xb4\x7a\x34\x0b\x77\x93\x2a\x35\xa6\xba\x3e\x9f\x19\xa3\xcd\x73\xc7\xf5\x99\xa6\x37\x74\xaa\xeb\x6a\xee\x01\x4d\x0b\x28\x20\x88\x1d\xa7\xc0\xd9\x6a\x1e\xe1\xb2\x12\x86\x72\x3e\x96\xe1\xc1\xcc\x53\x0a\xa5\x03\xe5\x70\x00\x7e\xe6\xfd\xac\xf8\x71\xa2\xe8\x93\x99\x88\xb9\x75\x15\xce\x24\x01\x66\xee\x5c\x53\x03\x25\x21\x3d\x1f\x33\x7d\x80\xe7\x1d\x2a\x93\x12\x4e\xcd\xf4\x54\xcd\x3b\x54\x4a\x4b\x2b\x27\x01\x5a\xad\xe7\x25\x53\x5f\x86\x8c\xae\xbd\x79\x9b\x1f\xd6\x27\xc3\xc4\x9f\x79\x52\xa1\x43\x85\xd2\x65\x72\xa1\xcc\x05\x1d\x14\x43\x03\xd8\x8b\x1f\x41\xf5\xd5\x51\xd3\x69\x64\xf4\x40\xb9\x54\x25\xf4\x92\x88\x66\xcc\xfa\x36\x66\xb5\x42\x25\x39\x22\x21\x2d\x34\x5f\x94\xb8\xee\xa3\xad\xa2\xc0\xfe\xbf\x14\xa9\x4b\xb0\xb7\x52\x72\xb3\xfc\x81\xbd\x72\xb6\xd5\x44\x1f\xc3\xd6\x67\xb0\xff\x0f\xf5\x0f\x0d\xf0\x65\x62\x45\xc5\xe2\xbe\xf4\xdd\x25\xf0\x5d\x85\xde\x44\x07\x5f\xb3\x86\x5c\xfe\x98\x6e\x63\xff\xa5\x88\xfb\xa2\xf7\xdb\x4e\x05\xc9\x38\x3d\xc6\xc3\x42\xf7\xdb\x0e\x5b\xa2\x2c\x83\xd9\xd8\xe5\x0f\xed\xea\xdb\x6f\x3b\x59\x4d\x0f\x52\xf3\x2c\x1e\x55\xb5\x43\x1f\x1f\x26\xd7\x9e\xb9\xd2\xc6\xb5\xc6\x19\x33\x8a\x93\xe3\xa2\xf9\xbf\x8d\x83\xf9\x4a\x20\x32\xa6\xdd\xdb\xcb\xa3\xc1\x90\x7d\x54\x1c\xc2\x14\xbe\x18\x1f\x5e\xfd\xfa\xec\xb2\x6e\xb7\xd5\x3e\x37\x75\x2b\xe8\x31\xa0\x30\xee\x62\x06\x61\xc0\x5c\xa6\x09\x60\x78\x4b\xd7\x13\x85\xce\xfe\xea\x14\x40\xde\xd6\x66\xc0\x36\x68\x7e\xb4\x70\x8f\x2f\x28\xf4\xa6\xd0\xd8\x47\xd7\x9e\x37\xd1\x35\x7a\x2a\xcb\x12\xba\x9c\xa7\x30\x0c\x41\x16\xe6\x6e\xea\x37\x90\x4e\xc9\xbb\xcd\xfb\x97\xd0\x1c\xbb\xa6\xd0\xdb\x26\x09\x0a\x20\xec\x41\xcd\xbe\x5e\x32\x58\xad\x7e\x41\x6f\xa9\x76\x02\x80\x1f\x22\xc8\x65\xd3\xb2\xa5\xec\x55\x1d\xa4\x63\x08\x2b\xd7\x2d\xec\x08\xb5\x2e\xfb\xbf\x07\x5c\x5b\x8d\xb1\xc7\x0d\xc3\x2d\xf0\xa6\xe9\xe2\xbe\xb7\x40\x6b\x37\x5e\x62\xd0\xf2\xf5\xf8\x03\x9c\x17\xd2\xfd\x62\xb1\x17\xf2\x06\x15\xf0\x67\x65\xeb\x7d\xcf\xbb\xff\x55\x31\x1b\x1a\x72\xa0\x49\x20\xdc\x05\x0b\x53\xb6\x32\xc9\xda\xc7\x20\x25\x10\x35\xdf\xb6\x07\x2e\xe4\xd3\xa4\x81\xf5\xfe\x09\x00\x00\xff\xff\x61\x7d\x0d\x0d\xea\x13\x00\x00")

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

	info := bindataFileInfo{name: "Rscripts/install-packages.r", size: 5098, mode: os.FileMode(509), modTime: time.Unix(1468863878, 0)}
	a := &asset{bytes: bytes, info: info}
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

	info := bindataFileInfo{name: "Rscripts/knit.r", size: 1149, mode: os.FileMode(509), modTime: time.Unix(1468863878, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"Rscripts": &bintree{nil, map[string]*bintree{
		"install-packages.r": &bintree{rscriptsInstallPackagesR, map[string]*bintree{}},
		"knit.r": &bintree{rscriptsKnitR, map[string]*bintree{}},
	}},
	"views": &bintree{nil, map[string]*bintree{
		"base.html": &bintree{viewsBaseHtml, map[string]*bintree{}},
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


// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sqls/1.sql.down.tmpl (306B)
// sqls/1.sql.up.tmpl (5.222kB)
// sqls/2.sql.down.tmpl (1B)
// sqls/2.sql.up.tmpl (5.611kB)

package postgres

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _sqls1SqlDownTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\x41\xaa\x42\x31\x0c\x45\xe7\x7f\x15\x1d\x7e\xd7\xf0\x16\x53\x6a\xbd\x96\x40\xfb\x12\x6e\x52\xdd\xbe\x20\x88\x22\x7d\x4e\xcf\xc9\x49\x72\xa1\x5a\xba\x09\xee\xc9\xb4\x4b\x15\x78\x36\x70\x88\xbb\xe8\xee\xdb\xdf\xc2\x53\x3b\x5e\x26\xca\xb9\x23\xa9\x95\x3c\x1d\x5c\xd3\x83\xe0\x98\x2e\x3e\x78\x0f\xfc\x72\xa5\xc6\x92\x13\xae\x93\xf5\xeb\x5c\x53\x5a\x1e\xd2\x58\x3e\xab\xeb\xdc\x9f\x5b\x52\x50\x5a\x03\xb3\x23\x72\xc8\x80\x47\x19\xf6\x7f\xda\x1e\x01\x00\x00\xff\xff\xfd\xeb\x29\xcd\x32\x01\x00\x00")

func sqls1SqlDownTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls1SqlDownTmpl,
		"sqls/1.sql.down.tmpl",
	)
}

func sqls1SqlDownTmpl() (*asset, error) {
	bytes, err := sqls1SqlDownTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/1.sql.down.tmpl", size: 306, mode: os.FileMode(0644), modTime: time.Unix(1595836250, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3f, 0x40, 0x4, 0xc0, 0xd9, 0x9a, 0x73, 0xb0, 0x49, 0xeb, 0x46, 0x2f, 0x53, 0xd1, 0xa1, 0xfa, 0x4, 0x82, 0x7c, 0x55, 0xec, 0x58, 0x5f, 0x7d, 0x33, 0x4f, 0xa0, 0x74, 0x4, 0xbe, 0x4, 0x13}}
	return a, nil
}

var _sqls1SqlUpTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x97\xdf\x6f\xe2\x38\x10\xc7\xdf\xf3\x57\xcc\x43\xa5\x26\x12\x42\x77\x27\xed\x4b\xd1\x3e\xa4\x60\x38\x4e\xbd\x84\x0b\xa0\xee\x3d\x45\xde\xc4\x65\xad\x86\x24\x6b\x1b\x5a\x6e\xb5\xff\xfb\xc9\xce\x2f\x27\x24\xa1\x05\x7a\xba\x3c\x81\x33\xf3\x9d\xf1\xf8\x33\xb6\x33\xf6\x90\xbd\x42\xe0\x7a\xe0\xa1\xc5\x83\x3d\x46\x30\x5d\x3b\xe3\xd5\xdc\x75\x40\x30\xba\xd9\x10\xe6\x73\x22\x7c\x41\xb7\x84\x0b\xbc\x4d\x4d\xcb\x00\x00\xf0\xd0\x6a\xed\x39\x4b\x58\x79\xf3\xd9\x0c\x79\x60\x2f\x8d\x9b\x1b\xe3\x1e\xcd\xe6\x8e\x7a\xef\xa0\xc7\xe1\x2e\x0d\xb1\x20\xa1\x8f\x05\x7c\x06\xc7\x7d\x34\xad\x91\xe6\x2b\x4d\x46\x06\x72\x26\x23\xe3\xe6\x06\x1e\x6c\x67\xb6\xb6\x67\x08\xd2\x28\xdd\xf0\xef\xd1\xc8\x30\x02\x46\xb0\x20\x20\xf0\xd7\x88\x00\x7d\x82\x38\x11\x40\x5e\x29\x17\x1c\x36\x09\x4b\xfd\x2d\xdd\x30\x2c\x68\x12\x73\xc3\x54\xba\x34\x84\xe2\x11\xe4\x55\x28\x87\x78\x17\x45\x46\x31\x1a\x24\x31\x17\x0c\xd3\x58\x34\x15\xfc\xf4\x99\x1c\x4a\x3b\xf9\xa4\x8c\x6e\x31\x3b\xc0\x33\x39\x0c\xd4\x0b\x9c\xa6\x11\xcd\x66\x53\x16\x03\x5e\xa8\xf8\xa6\xfe\xc2\x3f\x49\x4c\x0c\x6b\x64\x18\x38\x12\x84\xe5\x59\x37\xf3\x94\x3a\xc9\x4b\x2c\xdf\x27\xf0\xe3\xc7\x50\xfd\xfe\xf9\xb3\x7f\xb2\x49\x8a\x7d\x1c\x74\x4c\x14\xf6\x98\x05\xdf\x30\x33\x7f\xfb\xd5\x82\xde\x19\x6b\x32\x6f\x98\x6d\x8c\xb7\xe4\x28\xc4\xa7\x4f\x56\x19\x22\x33\x0b\x09\x0f\x18\x4d\xa5\x6a\xbb\x19\x84\xe4\x09\xef\x22\x01\xb7\xb7\x99\x47\x36\x51\x55\xc6\xaa\x8e\xe6\x2f\x2d\x1e\x71\xf2\x62\x5a\x99\x93\x46\xd2\x5b\x9c\x8e\xd6\x41\x2f\x61\xf7\x1a\x4c\x3c\x77\x51\x02\x3d\x9f\x02\xfa\x32\x5f\xae\x96\x50\xc3\x5f\xb9\xbb\x8e\xae\x38\x32\xf2\x1e\x2a\x5c\x8f\x1d\xee\xd1\xd4\xf5\x10\xac\x17\x13\x7b\x85\x5a\x24\xd4\xd0\xd4\xf5\x00\xd9\xe3\xdf\xc1\x73\x1f\x0d\xf4\x05\x8d\xd7\x2b\x04\x0b\xcf\x1d\xa3\xc9\xda\x43\x5d\xdd\x78\x9a\x9d\x94\xb0\x2d\xe5\xfc\x1a\xfc\x68\x52\x6f\xe9\x18\x35\x39\x5f\x45\x6b\x0d\x93\x99\x31\xc2\x93\x1d\x0b\x88\x34\xec\x31\xfb\x2f\xc0\xc9\x9c\x9e\xc9\xe1\xb8\x3c\x12\xea\xf2\xa9\xa0\xbe\xbb\x93\xaf\x71\x20\x59\xdb\x63\x76\xa0\xf1\xa6\xaa\x62\x1b\x86\xfa\x6a\x5c\x0f\x45\x4d\xf5\x12\x1c\x9b\xc9\x7d\x14\x92\xc5\x92\xb7\x00\xf9\x2e\x1e\x4b\x9d\xf7\xed\x68\xad\x3b\xd5\xe9\x18\x52\xc1\x6f\xc6\xd9\xc5\xf4\xfb\x8e\x1c\x21\x7a\x0e\xa1\xe7\xee\x6c\x55\x35\xaf\x07\x54\xa9\x79\x09\x4e\xf5\xc4\x3e\x0c\xa6\x24\x22\x2d\x9b\x9c\x1a\x2e\xc1\xd2\xb0\xca\xea\x5e\x39\x34\xf6\x9d\xf6\x12\x37\x63\x74\x57\xfa\x0d\xd9\x5e\x8c\xbd\xd4\x38\x1f\xf9\x7e\xd9\xff\x2d\xe5\xaa\x70\x57\x24\x5c\xea\x5d\x44\x77\x99\xd0\x47\x91\xbd\xe3\x84\xd5\x80\x51\x03\x1d\xd4\x34\x0e\xd5\x1a\xfd\x7d\x96\x1a\x04\x71\xe2\x87\xbb\x34\xa2\x01\x16\xc4\xcf\x63\xf9\x85\x52\x86\x00\x98\xf9\xf8\xa0\x08\xd1\xbe\x5a\x5a\xea\x67\xb7\x8a\xd4\x68\xbb\xb2\x9c\xee\x97\xde\xeb\x6c\xb3\x1d\x32\x0b\xb2\xc5\x34\xea\xb0\x68\x54\xac\x73\x15\x4a\x31\xca\xd5\xcd\x6e\x5f\xc4\xfc\x9a\x24\x95\xf6\x11\xf3\x4f\x38\xe2\x79\x6f\xa5\x2c\xd9\xd3\x90\xb0\xde\x2c\x8e\xee\xd4\x78\x8f\x05\x66\x27\x72\xef\xbd\x89\xcb\xe7\xac\x4b\xd5\xbb\x1c\xbb\x50\xa3\xa1\x1f\x26\x5b\x4c\x63\x5f\xad\x82\x5f\x56\x61\xed\xcc\xff\x5a\x23\x30\xd5\xf0\xa0\xac\x4e\x37\x72\xd7\xdc\x20\x94\xde\x25\x1b\x44\x95\xd0\xa5\x1b\x44\xc2\x80\x91\x34\xc2\x01\x81\x3d\x25\x2f\x90\x26\x11\x0d\xa8\x3c\x03\xaa\x23\xc9\xcc\xaf\xdb\x12\xf5\x41\x75\xa9\xce\xff\xca\x5e\x95\x3f\x2d\xc0\xdc\x58\xa2\x07\x34\x5e\xe9\x1f\x20\xc3\xbc\x41\xec\x25\xe8\x32\x45\x43\xd5\x0e\xf3\xcc\xd6\x5e\x36\x62\xd4\x6c\x65\xe3\x0f\xcb\xa6\x93\xb6\x45\x02\xc6\xd4\x73\xff\x04\xd3\x34\xcd\xce\x73\xf5\x0f\x77\xae\xed\xb0\xb2\x9a\x95\x31\x1f\xd2\x10\x3e\xb7\x9e\xc9\xc3\x62\x43\xb2\xac\xba\x8c\x66\x53\x89\xe9\x8e\xdd\x92\xb5\x5b\xc2\x91\x70\x5e\xbc\x4a\xb4\xa8\x66\x29\xa8\x6b\x95\x9f\x43\x85\x0e\xd4\x27\x5b\xd4\x57\x9b\x70\x59\xf2\x56\x41\xed\xc3\x49\x4a\xba\xde\x04\x79\x70\xff\x77\x63\x05\x1a\x9d\xd2\x86\xce\xe9\x2d\xba\x13\x40\x15\xa7\x71\x26\xc8\xa8\x5c\x07\xad\x66\x3d\x2c\x8c\x8b\x1a\x70\xc1\x68\xbc\xf1\xf1\x66\x63\x9a\x75\xcb\x8a\xda\xbb\x3b\x41\x5e\xc5\x00\x6e\x07\xb7\xd9\x4f\xab\x06\x15\xcf\xa9\xd2\xc0\x56\xcd\x37\x54\xbb\xb4\xbd\x84\x66\xcc\x56\x54\x6b\x98\x66\x16\x39\xac\x2d\x27\x5a\x63\xf5\x54\xb8\x6c\xe5\x4c\x3d\x7c\x9e\x39\x7c\x86\x86\x48\x51\x86\xdc\xa0\xc6\x44\x93\x8c\xaa\x0d\x6a\x7d\xd0\xad\x5d\x74\x42\x53\xbb\x86\x48\x99\x62\x7d\x7d\x8c\x99\xe7\xae\x17\xd2\xa8\x75\xd9\xba\x70\xea\x3f\xeb\xff\x0d\x00\x00\xff\xff\x2d\xa2\x7e\xa8\x66\x14\x00\x00")

func sqls1SqlUpTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls1SqlUpTmpl,
		"sqls/1.sql.up.tmpl",
	)
}

func sqls1SqlUpTmpl() (*asset, error) {
	bytes, err := sqls1SqlUpTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/1.sql.up.tmpl", size: 5222, mode: os.FileMode(0644), modTime: time.Unix(1595836129, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x34, 0xfd, 0xe1, 0xcd, 0x6d, 0x56, 0x62, 0x51, 0xb2, 0xcc, 0x48, 0x52, 0xe2, 0xfe, 0xb1, 0xd0, 0x48, 0x7, 0x50, 0xfb, 0x1, 0xbc, 0x2c, 0x66, 0x43, 0xfa, 0x31, 0xde, 0x19, 0x4d, 0xf5, 0x9}}
	return a, nil
}

var _sqls2SqlDownTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x06\x04\x00\x00\xff\xff\xa9\x06\x09\x63\x01\x00\x00\x00")

func sqls2SqlDownTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls2SqlDownTmpl,
		"sqls/2.sql.down.tmpl",
	)
}

func sqls2SqlDownTmpl() (*asset, error) {
	bytes, err := sqls2SqlDownTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/2.sql.down.tmpl", size: 1, mode: os.FileMode(0644), modTime: time.Unix(1594966792, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0xb8, 0x5, 0xea, 0x7a, 0xc0, 0x14, 0xe2, 0x35, 0x56, 0xe9, 0x8b, 0xb3, 0x74, 0x70, 0x2a, 0x8, 0x34, 0x42, 0x68, 0xf9, 0x24, 0x89, 0xa0, 0x2f, 0x8, 0x80, 0x84, 0x93, 0x94, 0xa1, 0xe4}}
	return a, nil
}

var _sqls2SqlUpTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x5b\x73\xa2\x48\x18\x7d\xdf\x5f\xc1\x1b\x49\x0d\x8c\xad\xe0\x25\xee\xcb\xaa\xa0\xf1\x12\xf0\x82\x97\x30\x33\x4b\xb5\xd2\x6a\x23\x08\xe1\xe6\x65\x6b\xe7\xb7\x6f\x81\x49\x26\xa9\xa5\x99\x4c\xa2\x96\x35\x2f\xd2\x36\x56\x7d\xe7\x9c\xaf\xbf\x73\xba\x6c\x4a\x03\xb1\xaf\x50\x4d\x49\x91\x29\x27\x98\x9a\x78\xf6\xd9\x76\xa0\x06\x67\x3e\xb6\xd7\x1e\x75\x85\x75\x86\x5a\x43\x0b\x31\x94\x8e\xbc\x99\x8b\x9d\x68\x9f\xa1\x66\x2e\x82\x3e\xd2\x35\xe8\x33\x54\xe0\xe8\x8f\xeb\x6b\x6a\x54\xe9\x0c\xc5\x01\x75\x45\xcf\x8c\xe1\xc0\xec\x09\x7d\x6d\x57\x2d\xb8\x43\xbd\x61\xb9\x46\xe5\x81\x66\x28\xba\x21\x2a\x34\x73\xf8\xa4\xe8\x1c\xc8\x01\x36\x0b\x58\x0e\x28\xe0\xa6\xcc\x97\xca\x3c\xf8\x04\x4a\x65\x00\xd2\x5f\x5e\xff\xf9\xc7\xa9\x70\x5b\x46\x75\x92\xeb\x8f\xbd\xba\x3f\x19\x07\x45\xb7\xe2\x6a\x9d\xad\x15\xa1\xe9\xca\x83\x08\xf8\xe1\x71\x89\xc8\x95\x46\x3d\x64\xc5\x65\xd5\xda\x36\xed\x0e\x36\x0c\xa5\x5f\xda\xa0\x18\xf9\x30\x06\x3e\xbc\x50\xdc\xbb\xfa\xf4\x61\xb8\xd2\xfb\x63\xa1\x58\x15\x44\x79\x51\x5a\x4b\xa6\x13\xa1\x11\xc4\x8e\xa8\x88\x34\xf3\xbc\xf8\x08\x7a\x12\x7c\x17\x79\x76\xe0\xce\xd0\x2b\x02\x3f\xc5\xdc\xbd\x1d\x6c\x72\xea\x2e\x7f\x9f\x15\x6e\x6d\x4d\xea\xdc\xb1\x4a\xbe\x1e\x61\xc8\x40\x07\x67\xae\xbe\xfc\xfd\xf5\x6b\xe6\xdb\xa7\xeb\x0c\x0c\xfc\xe5\x3e\xe3\xda\x26\xf2\x32\x5f\x20\xbb\xaf\xb0\x2a\x60\x6f\x34\xf6\xfb\xb7\x7f\x72\xd9\x7f\x4f\xd2\x8f\xf7\x11\x5a\x59\xa3\x85\xdf\x62\x57\xd0\xa8\x5b\xb9\x12\xcf\xcf\x55\x49\xe5\x93\x08\xe9\x7a\xc6\x0b\xac\x0b\x42\x5e\xd2\x24\x76\x51\x6a\x67\xef\xf7\x82\x9f\xc7\x82\x6a\xd7\x2a\x86\x90\x80\xdc\xc7\x33\x1f\xce\xa2\xc7\x05\x81\x97\xf7\x83\xfb\x5b\x21\x0b\xa7\x48\xf6\x47\x1e\xbf\x68\xe0\xd5\xc8\x4c\x01\x0f\x2f\x09\x3c\x37\xc0\xb2\xc5\xb2\xeb\x7d\xef\x6e\x3a\x6e\xc8\xfa\xc0\x2c\xc2\x61\xfa\x10\x9c\x66\x86\x1d\xe4\x5a\xd8\xf3\x7e\xd8\xd0\xc1\x93\xb4\x68\xf9\x44\x2d\xfe\x92\x4c\x89\xa1\x56\x68\xf7\x82\x97\xd3\x1a\xf6\xfc\x55\x2d\x6f\x86\xde\xb6\x2f\xda\x02\x36\x6b\xcb\xd8\x48\x89\x99\x46\x54\xe2\x9d\x5c\x0f\x31\x59\x4e\xd1\x91\xdc\xcb\x23\x8b\xd1\xde\x6a\x9c\xe8\x4c\xe5\x20\x3f\x9e\x3b\x6e\x2e\x17\x9a\xf5\xba\x98\x2a\x06\xd1\x1b\x4f\x25\x46\x92\xb3\x9e\x4d\x1f\xc1\xb5\x65\x71\x72\x2b\xe1\x96\x16\x2a\x9c\x13\xf4\x2c\xa5\x17\xb7\x9e\x78\x91\x20\x5a\xed\xfb\xf5\x89\x2e\x25\xff\x13\xe8\xd1\xa9\xcf\xa6\x04\xe0\x5b\x4d\xae\xd5\x69\x6f\xf3\x45\xb0\x44\xe2\x74\x50\xca\xb6\xe7\xa9\x4a\x10\xad\xfb\xb8\x4a\xbc\x70\xfe\xb3\x89\x61\x14\x97\xab\x8a\xa2\x86\x77\xb0\x3f\x0a\xf9\x51\x07\x1b\xa2\xbc\x4b\x1d\x1b\x62\x14\x1c\x75\x6c\x5e\x04\x49\xda\xa5\x28\x1a\xab\x5f\xca\x82\x06\xd8\xaf\xd5\xb1\x92\x87\x46\x53\x69\x2a\x37\xc5\xb0\x30\x9e\xe0\x08\x83\xbd\x59\x23\xf7\x34\xa1\xf5\xcb\x20\x87\xdd\x2d\x10\x45\xb1\xd5\xaf\x71\x77\x7b\x60\xfa\x1b\x47\xe0\x17\x51\x79\xa4\x63\xdf\xbe\x14\x94\x5a\xbf\xbb\x6d\x4f\xd4\x82\xe8\x71\x85\xef\x9a\xda\x0d\x56\x66\x10\x5f\xdd\x43\x8c\x36\x1f\xd5\x32\x0d\xe6\xeb\x19\x88\x77\x22\xd8\x3f\x76\x35\xac\xbf\xa5\xe5\xc9\xf9\x99\x2e\xd1\xb1\x6a\x27\xc7\xd5\x79\x6a\x27\x47\xc1\x79\x6a\x27\x9b\xef\x79\x6a\x27\x7b\xdd\x51\x6b\x13\x27\xf7\x0c\xbc\x89\xb5\xcf\xc0\x9b\xe8\x05\xc4\xb3\x46\x9c\xf0\xc0\x43\xee\x93\x1b\xc5\xeb\x38\xd9\x0e\x00\x5e\x54\x44\x55\xb4\x9e\xb7\x83\xda\xbe\x82\x2b\x35\x6d\x3a\x7f\x10\x95\xcd\x3d\xcd\x10\xba\x4f\x26\xfb\xc6\x72\x55\x98\x1d\xe9\x4b\xbb\xca\x15\x97\x6a\x47\xc6\xa3\xea\x54\x5b\xb6\x69\x86\x20\xfa\x87\xcb\x01\x7f\xb7\xc1\x85\x65\xe1\xc6\xd2\xbb\x6d\x49\xd2\xa4\x55\xb3\x58\xa5\x19\x82\xce\x29\x7e\x19\xd5\x78\x65\xeb\xc8\x82\xd8\x64\xa8\xe7\xda\xd8\x8b\xff\x19\x09\x11\x43\x39\xae\x1d\x62\x1d\xb9\x0c\x05\x43\xe8\x43\xf7\x0d\x19\x40\xe8\x02\x45\x3f\x47\xea\x5f\x68\x0b\x2d\xc7\x44\x9f\x67\xb6\xf5\xf8\xc2\x77\x03\xc4\x50\xf4\xc2\xb6\x17\x26\x7a\xfa\x71\x9c\x06\x80\x63\xb3\x80\x02\xc5\x32\xc8\x97\xb9\x6c\xc2\xf6\x4d\x99\xe7\x7f\xa2\xee\x09\xd9\x12\x0e\xc1\x81\xc0\x21\x9b\x7f\x27\xba\x84\x43\x78\x20\x70\x08\xf9\x93\xd0\xfd\x2f\x00\x00\xff\xff\xa6\xd9\x80\xc7\xeb\x15\x00\x00")

func sqls2SqlUpTmplBytes() ([]byte, error) {
	return bindataRead(
		_sqls2SqlUpTmpl,
		"sqls/2.sql.up.tmpl",
	)
}

func sqls2SqlUpTmpl() (*asset, error) {
	bytes, err := sqls2SqlUpTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqls/2.sql.up.tmpl", size: 5611, mode: os.FileMode(0644), modTime: time.Unix(1604022676, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x40, 0x3b, 0xad, 0xe0, 0x31, 0xb6, 0x12, 0x3e, 0xcb, 0x20, 0x84, 0x7e, 0x5b, 0x41, 0x1f, 0x31, 0xa, 0xf3, 0x5a, 0x4e, 0x53, 0x3c, 0x3, 0x1c, 0x16, 0xcf, 0xa5, 0xa5, 0xbb, 0x87, 0x72, 0xe0}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"sqls/1.sql.down.tmpl": sqls1SqlDownTmpl,
	"sqls/1.sql.up.tmpl":   sqls1SqlUpTmpl,
	"sqls/2.sql.down.tmpl": sqls2SqlDownTmpl,
	"sqls/2.sql.up.tmpl":   sqls2SqlUpTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"sqls": &bintree{nil, map[string]*bintree{
		"1.sql.down.tmpl": &bintree{sqls1SqlDownTmpl, map[string]*bintree{}},
		"1.sql.up.tmpl":   &bintree{sqls1SqlUpTmpl, map[string]*bintree{}},
		"2.sql.down.tmpl": &bintree{sqls2SqlDownTmpl, map[string]*bintree{}},
		"2.sql.up.tmpl":   &bintree{sqls2SqlUpTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}

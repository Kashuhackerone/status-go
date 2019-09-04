// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 0001_app.down.sql (230B)
// 0001_app.up.sql (2.246kB)
// doc.go (74B)

package migrations

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

var __0001_appDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\xe6\x42\x12\x4c\x4c\x4e\xce\x2f\xcd\x2b\x41\x15\x4c\x2a\xca\x2f\x2f\x4e\x2d\xc2\x2e\x18\x9f\x91\x59\x5c\x92\x5f\x54\x89\x22\x99\x92\x58\x50\x80\xaa\xbc\x20\xb5\x28\x37\xb3\xb8\x38\x33\x3f\x0f\x55\xbc\xa4\x28\x31\xaf\x38\x0d\xc3\xf0\x9c\xfc\xe4\x6c\xec\x2e\x8b\x2f\xc9\x8f\xc7\x22\x9d\x9b\x98\x99\x53\x9c\x5a\x54\x06\x36\x09\x10\x00\x00\xff\xff\xd2\xa8\x66\x23\xe6\x00\x00\x00")

func _0001_appDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_appDownSql,
		"0001_app.down.sql",
	)
}

func _0001_appDownSql() (*asset, error) {
	bytes, err := _0001_appDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_app.down.sql", size: 230, mode: os.FileMode(0644), modTime: time.Unix(1567452784, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc6, 0x64, 0x5c, 0xd5, 0x35, 0x13, 0xe3, 0x6e, 0x9c, 0x61, 0x16, 0x4, 0x5d, 0x61, 0x79, 0xb2, 0xfd, 0x44, 0xa3, 0x6a, 0x1d, 0x8e, 0xd9, 0x50, 0x92, 0x2e, 0x34, 0xb2, 0x6, 0x3, 0xb7, 0xf1}}
	return a, nil
}

var __0001_appUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xdd\x92\xa2\x3c\x10\xbd\xcf\x53\xf4\xa5\x56\xf1\x06\x73\x85\x1a\x1d\xea\xe3\x0b\xbb\xfc\xec\xcc\x5c\x51\x11\xa2\x52\xf2\xb7\x49\x1c\xc7\xb7\xdf\x0a\x21\x80\x23\xe2\x5a\x7b\x47\xd2\xdd\xc7\x73\x4e\xba\xdb\xa5\x8f\xed\x10\x43\x68\x2f\x5c\x0c\xce\x1a\x88\x17\x02\x7e\x77\x82\x30\x00\xc1\xa4\xcc\xca\xbd\x80\x19\x92\x97\x9a\xc1\x2f\xdb\x5f\xbe\xda\x3e\xfc\xf0\x9d\xff\x6d\xff\x03\xfe\xc3\x1f\x16\xfa\xa4\xf9\x89\xc1\xc2\xf5\x16\x68\x0e\x6f\x4e\xf8\xea\x45\x21\xf8\xde\x9b\xb3\x7a\x41\x68\x02\x9c\x26\x49\x75\x2a\xa5\x02\xa7\x69\xca\x99\x10\xe3\xf8\x67\x9a\xe7\x4c\xc2\xc2\xf3\x5c\x6c\x13\x0b\x25\x07\x3a\x38\x35\xbc\x42\xfc\x1e\x5a\x48\xc8\x8a\xd3\xbd\x39\xd5\xa7\xed\x91\x5d\x1a\x5e\x16\xaa\xa9\x3c\xb4\xf7\x25\x2d\x4c\x4a\x52\xe5\x15\x6f\xbe\xef\x33\x8f\x88\xf3\x33\xc2\xe0\x90\x15\x7e\x87\x53\x99\xfd\x3e\xb1\x58\x33\x8a\x0d\x6b\x8f\x0c\xb4\xe8\xd8\x1c\xde\x5e\xb1\x8f\xbb\xe3\xcb\x14\x9c\x12\x34\x0e\xa6\x22\x1d\x54\x73\x98\xb6\x74\xcb\xab\xb3\x60\x5c\x59\x9a\xa5\x8d\xb0\x6b\x2b\x3b\xed\x4d\x11\x89\x5c\xd7\x42\x32\x2b\x98\x90\xb4\xa8\x21\x0a\x36\xce\x86\xe0\x15\x2c\x9c\x8d\x43\x42\x0b\xa5\xb4\xae\x8d\xd3\xb0\xc2\x6b\x3b\x72\x43\xd8\xd1\x5c\x30\x0b\x1d\x32\x65\xf7\xc5\x29\x53\xf6\x05\x11\x09\x74\xa5\x43\x26\xac\x9c\x62\x1c\xb7\x78\x30\x43\xed\x55\x6c\x14\xf4\x54\x4d\x8e\x7e\xbd\xb5\xe7\x63\x67\x43\x94\xb2\x59\x5f\x33\x07\x1f\xaf\xb1\x8f\xc9\x12\xf7\xe8\x33\x75\xef\x29\x0d\x2e\x0e\x31\x2c\xed\x60\x69\xaf\x30\x7a\xe0\xa6\x92\xaf\xac\xec\x5d\x1b\x98\xf9\x9c\xcc\x9a\xf1\x22\x13\x22\xab\x4a\x05\xa8\x80\xe3\xb1\xb7\xe8\xd3\xbe\x47\x86\x62\xbb\xf2\x2b\xad\x0d\xdb\x99\xbe\x1e\x97\x3a\x45\x50\x72\x5a\x8a\x9d\x6e\x9d\x92\xc9\x73\xc5\x8f\xea\x01\xba\x87\xd5\x2d\x31\x7c\x0b\x2a\x0e\xdd\xbc\xf6\xd7\xdf\x27\xb9\x8f\x6c\xf3\x63\x7c\xa7\x48\x7e\xb5\x63\x2a\x58\x99\x32\x3e\x92\xc1\x59\xc2\xb2\x5a\xb6\x69\x79\xb5\x6f\xbf\xae\xb6\xd2\xb8\x5b\xbd\x1a\xcb\x50\xb8\xee\x91\xbc\x4a\x8e\x62\x98\xa6\x53\x6e\x3c\xb4\xd0\xd2\x23\x41\xe8\xdb\xca\x88\x76\x74\x8d\x6d\x71\xcd\xb8\x19\xe1\xe6\xbb\x85\x33\xf3\x3e\x53\x98\x56\x9b\x60\xf5\xbf\x35\x7f\xd4\x83\x9a\xdd\x3f\x3e\x4a\x79\x2a\xb6\x8c\xdf\xa6\x0f\x46\xff\x3e\x24\xa3\x69\xb3\x03\xba\x05\xb0\xb6\xdd\x60\xd4\x8c\x86\xeb\xa8\xfa\xef\xe6\xde\x2d\xd6\x4c\x1f\x61\xe8\xac\x87\xde\x99\x3d\x1a\xcb\x2a\x7e\xce\xc7\xe9\x2e\xbe\x67\xa7\xb8\x94\x09\x34\x8b\x73\xa2\xff\x5a\xee\xd3\x1d\x68\x92\xfe\xaa\x07\x0b\x5a\xd7\x59\xb9\x8f\x77\x15\x8f\x5b\xc9\x9d\xe2\x51\x27\x4d\x1b\xf6\x74\x9e\xe9\xc8\x82\x66\xb9\x60\xfc\x53\xef\x0a\x00\x80\x2c\x1d\xff\xe3\x56\xb1\x66\xcb\xdd\xda\xa8\x42\xf7\x4d\x56\xd1\x9a\x0a\x71\xae\x78\x07\xad\x6f\x77\x39\x63\xf2\xa6\x42\x51\xfe\x13\x00\x00\xff\xff\xe9\x5c\xa2\xa8\xc6\x08\x00\x00")

func _0001_appUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_appUpSql,
		"0001_app.up.sql",
	)
}

func _0001_appUpSql() (*asset, error) {
	bytes, err := _0001_appUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_app.up.sql", size: 2246, mode: os.FileMode(0644), modTime: time.Unix(1567452763, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6e, 0xc4, 0x4d, 0x62, 0x9b, 0x4, 0x4e, 0xb1, 0x57, 0x51, 0x5c, 0xf7, 0x9, 0x9, 0x30, 0x38, 0x10, 0x1f, 0xc1, 0xcc, 0xe4, 0x56, 0x67, 0xa, 0x17, 0x3b, 0x51, 0x1f, 0x84, 0x12, 0xf4, 0x15}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x0d\xc4\x20\x0c\x05\xd0\x9e\x29\xfe\x02\xd8\xfd\x6d\xe3\x4b\xac\x2f\x44\x82\x09\x78\x7f\xa5\x49\xfd\xa6\x1d\xdd\xe8\xd8\xcf\x55\x8a\x2a\xe3\x47\x1f\xbe\x2c\x1d\x8c\xfa\x6f\xe3\xb4\x34\xd4\xd9\x89\xbb\x71\x59\xb6\x18\x1b\x35\x20\xa2\x9f\x0a\x03\xa2\xe5\x0d\x00\x00\xff\xff\x60\xcd\x06\xbe\x4a\x00\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1566753968, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x7c, 0x28, 0xcd, 0x47, 0xf2, 0xfa, 0x7c, 0x51, 0x2d, 0xd8, 0x38, 0xb, 0xb0, 0x34, 0x9d, 0x4c, 0x62, 0xa, 0x9e, 0x28, 0xc3, 0x31, 0x23, 0xd9, 0xbb, 0x89, 0x9f, 0xa0, 0x89, 0x1f, 0xe8}}
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
	"0001_app.down.sql": _0001_appDownSql,

	"0001_app.up.sql": _0001_appUpSql,

	"doc.go": docGo,
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
	"0001_app.down.sql": &bintree{_0001_appDownSql, map[string]*bintree{}},
	"0001_app.up.sql":   &bintree{_0001_appUpSql, map[string]*bintree{}},
	"doc.go":            &bintree{docGo, map[string]*bintree{}},
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

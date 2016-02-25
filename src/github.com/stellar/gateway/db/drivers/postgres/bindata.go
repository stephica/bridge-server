// Code generated by go-bindata.
// sources:
// migrations/01_init.sql
// DO NOT EDIT!

package postgres

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

var _migrations01_initSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xcd\x4e\xeb\x30\x10\x85\xf7\x7e\x8a\x59\xb6\xba\xb7\x12\x20\x95\x4d\x57\x81\x1a\xa9\x22\xb4\x21\x24\x8b\xae\x22\xd7\x1e\x05\x8b\xc4\x8e\xec\x49\x28\x6f\x8f\xcb\x02\x12\xb7\xb0\x4c\xbe\x33\x3f\xe7\x78\x16\x0b\xf8\xd7\xea\xda\x09\x42\x28\x3b\x76\x9f\xf3\xa4\xe0\x50\x24\x77\x29\x87\x1c\x25\xea\x01\x55\x26\x3e\x5a\x34\x04\x33\x06\xa0\x15\x78\x74\x5a\x34\xff\xc3\x87\xed\x30\x14\x6a\x6b\xaa\xf0\x7b\x10\x4e\xbe\x0a\x37\xbb\x59\x2e\xe7\x50\x6e\x37\xcf\x25\x87\xed\xae\x80\x6d\x99\xa6\x27\x71\xe7\xac\x44\xef\x51\x55\x82\x80\x74\x8b\x9e\x44\xdb\x4d\x25\xa2\xd6\xa6\xae\xc8\xbe\xa1\x99\xf6\x1b\xab\x42\x1d\xf5\xfe\x77\x9e\xe5\x9b\xa7\x24\xdf\xc3\x23\xdf\xc3\x4c\xab\x39\x9b\xaf\xd8\xd4\xd8\x4b\x70\x53\x38\x61\xbc\x90\xa7\xed\xcf\x8d\x45\x23\xae\xaf\xa2\x0d\x6c\xef\x24\x7e\xe3\xe5\x6d\x84\xfb\x43\xab\x89\xfe\x72\xea\x7b\x29\x11\x55\x2c\x59\xf3\x87\xa4\x4c\x7f\x64\x0d\xaa\x1a\x1d\x1c\x74\xc8\x85\xce\x28\x9a\x01\x9b\xf0\x06\xd5\x51\x39\x20\x3c\xd2\x64\x84\x43\xdf\x37\xf4\xc5\x26\x51\xc5\x5d\x2e\xc6\x35\x3e\x8b\xb5\x7d\x37\x6c\x9d\xef\xb2\xcb\x67\xb1\x1a\xb3\x28\xd9\x15\xfb\x0c\x00\x00\xff\xff\x94\x78\xf0\x22\x60\x02\x00\x00")

func migrations01_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initSql,
		"migrations/01_init.sql",
	)
}

func migrations01_initSql() (*asset, error) {
	bytes, err := migrations01_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_init.sql", size: 608, mode: os.FileMode(420), modTime: time.Unix(1455035399, 0)}
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
	"migrations/01_init.sql": migrations01_initSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"01_init.sql": &bintree{migrations01_initSql, map[string]*bintree{}},
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
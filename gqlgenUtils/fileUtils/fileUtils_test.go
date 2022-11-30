package fileUtils_test

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
)

func TestFindDirectory(t *testing.T) {
	mockDirName := "server/gql/models"

	t.Run("Success", func(t *testing.T) {
		response := fileUtils.FindDirectory(mockDirName)
		if response != "" {
			assert.Equal(t, true, response != "")
		}
	})
}

func TestMakeDirectory(t *testing.T) {

	type args struct {
		path    string
		dirName string
	}

	cases := []struct {
		name string
		req  args
		err  bool
	}{
		{
			name: "Success",
			req: args{
				path:    "/go-works/cli-app",
				dirName: "dirName",
			},
			err: false,
		},
		{
			name: "Fail",
			req: args{
				path:    "/go-works/cli-app",
				dirName: "dirName",
			},
			err: true,
		},
	}
	for _, tt := range cases {

		patchMkdir := gomonkey.ApplyFunc(
			os.Mkdir,
			func(name string, perm fs.FileMode) error {
				if tt.err {
					return fmt.Errorf("Error")
				} else {
					return nil
				}
			},
		)
		defer patchMkdir.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := fileUtils.MakeDirectory(tt.req.path, tt.req.dirName)
			if err != nil {
				assert.Equal(t, tt.err, err != nil)
			} else {
				assert.Equal(t, tt.err, err != nil)
			}
		})
	}
}

func TestMakeFile(t *testing.T) {
	type args struct {
		path     string
		fileName string
	}
	cases := []struct {
		name string
		err  bool
		req  args
	}{
		{
			name: "Success",
			req: args{
				// path: "/Users/ijasmohamad/go-works/cli-app",
				path:     "path",
				fileName: "fileName",
			},
			err: false,
		},
		{
			name: "Fail",
			err:  true,
			req: args{
				path:     "path",
				fileName: "fileName",
			},
		},
	}
	for _, tt := range cases {
		patchCreate := gomonkey.ApplyFunc(
			os.Create,
			func(name string) (*os.File, error) {
				if tt.err {
					return nil, fmt.Errorf("no such file or directory")
				} else {
					return nil, nil
				}
			},
		)
		defer patchCreate.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := fileUtils.MakeFile(tt.req.path, tt.req.fileName)
			if err != nil {
				assert.Equal(t, true, strings.Contains(err.Error(), "no such file or directory"))
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	cases := []struct {
		name string
		req  string
		resp bool
	}{
		{
			name: "Success",
			req:  "server/gql/models",
			resp: true,
		},
		{
			name: "Failure",
			resp: false,
		},
	}

	for _, tt := range cases {
		patchStat := gomonkey.ApplyFunc(
			os.Stat,
			func(string) (fs.FileInfo, error) {
				if tt.resp {
					return nil, nil
				} else {
					return nil, fmt.Errorf("Error")
				}
			},
		)
		defer patchStat.Reset()

		t.Run(tt.name, func(t *testing.T) {
			response := fileUtils.DirExists(tt.req)
			if response {
				assert.Equal(t, true, response)
			} else {
				assert.Equal(t, false, response)
			}
		})
	}
}

func TestIsExists(t *testing.T) {
	patchStat := gomonkey.ApplyFunc(
		os.Stat,
		func(string) (fs.FileInfo, error) {
			return nil, nil
		},
	)
	defer patchStat.Reset()
	t.Run("Success", func(t *testing.T) {
		response := fileUtils.IsExists("path", "dirName")
		assert.Equal(t, true, response)
	})
}

func TestWriteToFile(t *testing.T) {
	cases := []struct {
		name           string
		openFileErr    bool
		writeStringErr bool
		syncErr        bool
	}{
		{
			name:           "Success",
			openFileErr:    false,
			writeStringErr: false,
			syncErr:        false,
		},
		{
			name:        "Fail in openFile",
			openFileErr: true,
		},
		{
			name:           "Fail in writeString",
			writeStringErr: true,
		},
		{
			name:    "Fail in sync",
			syncErr: true,
		},
	}
	for _, tt := range cases {
		patchOpenFile := gomonkey.ApplyFunc(
			os.OpenFile,
			func(name string, flag int, perm fs.FileMode) (*os.File, error) {
				if tt.openFileErr {
					return nil, fmt.Errorf("Error in openFile")
				} else {
					return &os.File{}, nil
				}
			},
		)
		defer patchOpenFile.Reset()

		var openFile *os.File
		patchWriteString := gomonkey.ApplyMethod(
			reflect.TypeOf(openFile),
			"WriteString",
			func(*os.File, string) (int, error) {
				if tt.writeStringErr {
					return 0, fmt.Errorf("Error in writeString")
				} else {
					return 0, nil
				}
			},
		)
		defer patchWriteString.Reset()

		patchSync := gomonkey.ApplyMethod(
			reflect.TypeOf(openFile),
			"Sync",
			func(*os.File) error {
				if tt.syncErr {
					return fmt.Errorf("Error in sync")
				} else {
					return nil
				}
			},
		)
		defer patchSync.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := fileUtils.WriteToFile("path", "file", "data")
			if err != nil {
				if tt.openFileErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in openFile"))
				} else if tt.writeStringErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in writeString"))
				} else if tt.syncErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in sync"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestCurrentDirectory(t *testing.T) {
	cases := []struct {
		name string
		err  bool
	}{
		{
			name: "Success",
			err:  false,
		},
		{
			name: "Fail",
			err:  true,
		},
	}

	for _, tt := range cases {

		patchAbs := gomonkey.ApplyFunc(
			filepath.Abs,
			func(string) (string, error) {
				if tt.err {
					return "", fmt.Errorf("Error in filepathAbs")
				} else {
					return "path", nil
				}
			},
		)
		defer patchAbs.Reset()

		t.Run("Success", func(t *testing.T) {
			response := fileUtils.CurrentDirectory()
			if response == "" {
				assert.Equal(t, true, tt.err)
			} else {
				assert.Equal(t, response, "path")
			}
		})
	}
}

package modelUtils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wednesday-solutions/negt/gqlgenUtils/fileUtils"
	"github.com/wednesday-solutions/negt/gqlgenUtils/modelUtils"
)

func TestCreateGqlModelFiles(t *testing.T) {

	type args struct {
		modelName string
		dirName   string
		files     []string
		testFiles []string
	}
	cases := []struct {
		name           string
		req            args
		makeDirErr     bool
		makeFileErr    bool
		secondMakeDir  bool
		secondMakeFile bool
	}{
		{
			name:       "Success",
			makeDirErr: false,
			req: args{
				modelName: "testModel",
				dirName:   "testDir/models",
				files:     []string{"file1"},
				testFiles: []string{"testFile1"},
			},
		},
		{
			name:       "Fail in makeDir",
			makeDirErr: true,
			req: args{
				modelName: "testModel",
				dirName:   "testDir/models",
				files:     []string{"file1"},
				testFiles: []string{"testFile1"},
			},
		},
		{
			name:       "Fail in makeDir",
			makeDirErr: true,
			req: args{
				modelName: "testModel",
				dirName:   "testDir/models",
				files:     []string{"file1"},
				testFiles: []string{"testFile1"},
			},
			secondMakeDir: true,
		},
		{
			name:        "Fail in makeFile",
			makeFileErr: true,
			req: args{
				modelName: "testModel",
				dirName:   "testDir/models",
				files:     []string{"file1"},
				testFiles: []string{"testFile1"},
			},
		},
		{
			name:        "Fail in makeFile",
			makeFileErr: true,
			req: args{
				modelName: "testModel",
				dirName:   "testDir/models",
				files:     []string{"file1"},
				testFiles: []string{"testFile1"},
			},
			secondMakeFile: true,
		},
	}

	for _, tt := range cases {

		patchFindDir := gomonkey.ApplyFunc(
			fileUtils.FindDirectory,
			func(string) string {
				return "path"
			},
		)
		defer patchFindDir.Reset()

		patchMakeDir := gomonkey.ApplyFunc(
			fileUtils.MakeDirectory,
			func(string, string) error {
				if tt.makeDirErr {
					if !tt.secondMakeDir {
						return fmt.Errorf("no such file or directory")
					}
				} else {
					return nil
				}
				tt.secondMakeDir = false
				return nil
			},
		)
		defer patchMakeDir.Reset()

		patchMakeFile := gomonkey.ApplyFunc(
			fileUtils.MakeFile,
			func(string, string) error {
				if tt.makeFileErr {
					if !tt.secondMakeFile {
						return fmt.Errorf("no such file or directory")
					}
				} else {
					return nil
				}
				tt.secondMakeFile = false
				return nil
			},
		)
		defer patchMakeFile.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.CreateGqlModelFiles(tt.req.modelName, tt.req.dirName, tt.req.files, tt.req.testFiles)
			if err != nil {
				assert.Equal(t, true, err != nil)
				if tt.makeDirErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "no such file or directory"))
				} else if tt.makeFileErr {
					fmt.Println("Error in makeFileErr: ", err)
					assert.Equal(t, true, strings.Contains(err.Error(), "no such file or directory"))
				}
			} else {
				assert.Equal(t, true, err == nil)
			}
		})
	}
}

func TestCreateCustomResolverFiles(t *testing.T) {
	type args struct {
		modelName         string
		dirName           string
		resolverFiles     []string
		resolverTestFiles []string
	}
	cases := []struct {
		name            string
		req             args
		err             bool
		resMakeFileErr  bool
		testMakeFileErr bool
		secondMakeFile  bool
	}{
		{
			name: "Success",
			err:  false,
			req: args{
				modelName:         "testModel",
				dirName:           "testDir",
				resolverFiles:     []string{"file1"},
				resolverTestFiles: []string{"testFile1"},
			},
		},
		{
			name: "Success-resMakeFile",
			err:  true,
			req: args{
				modelName:         "testModel",
				dirName:           "testDir",
				resolverFiles:     []string{"file1"},
				resolverTestFiles: []string{"testFile1"},
			},
			resMakeFileErr: true,
		},
		{
			name: "Success-testMakeFile",
			err:  true,
			req: args{
				modelName:         "testModel",
				dirName:           "testDir",
				resolverFiles:     []string{"file1"},
				resolverTestFiles: []string{"testFile1"},
			},
			resMakeFileErr:  false,
			testMakeFileErr: true,
			secondMakeFile:  true,
		},
	}

	for _, tt := range cases {

		patchFindDir := gomonkey.ApplyFunc(
			fileUtils.FindDirectory,
			func(string) string {
				return "path"
			},
		)
		defer patchFindDir.Reset()

		patchMakeFile := gomonkey.ApplyFunc(
			fileUtils.MakeFile,
			func(string, string) error {
				if tt.err {
					if tt.resMakeFileErr {
						return fmt.Errorf("Error in resMakeFile")
					} else if tt.testMakeFileErr {
						if !tt.secondMakeFile {
							return fmt.Errorf("Error in testMakeFile")
						}
					}
				}
				tt.secondMakeFile = false
				return nil
			},
		)
		defer patchMakeFile.Reset()

		t.Run(tt.name, func(t *testing.T) {
			err := modelUtils.CreateCustomResolverFiles(
				tt.req.modelName,
				tt.req.dirName,
				tt.req.resolverFiles,
				tt.req.resolverTestFiles,
			)
			if err != nil {
				assert.Equal(t, tt.err, err != nil)
				if tt.resMakeFileErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in resMakeFile"))
				} else if tt.testMakeFileErr {
					assert.Equal(t, true, strings.Contains(err.Error(), "Error in testMakeFile"))
				}
			} else {
				assert.Equal(t, tt.err, err != nil)
			}
		})
	}
}

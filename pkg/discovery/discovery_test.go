package discovery

import (
	"reflect"
	"testing"

	"os"
	"time"

	goplugin "github.com/hashicorp/go-plugin"
)

type fileInfoMock struct {
	name  string
	isDir bool
}

func (fi fileInfoMock) Name() string {
	return fi.name
}

func (fi fileInfoMock) IsDir() bool {
	return fi.isDir
}

func (fi fileInfoMock) Size() int64 {
	return 0
}

func (fi fileInfoMock) Mode() os.FileMode {
	return 0
}

func (fi fileInfoMock) ModTime() time.Time {
	return time.Time{}
}

func (fi fileInfoMock) Sys() interface{} {
	return nil
}

func Test_pluginMap(t *testing.T) {
	type args struct {
		prefix      string
		pluginPaths []string
		listDir     listDir
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]goplugin.Plugin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := pluginMap(tt.args.prefix, tt.args.pluginPaths, tt.args.listDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("pluginMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pluginMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandName(t *testing.T) {
	type args struct {
		name        string
		prefix      string
		pluginPaths []string
		listDir     listDir
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandName(tt.args.name, tt.args.prefix, tt.args.pluginPaths, tt.args.listDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("expandName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPlugins(t *testing.T) {
	type args struct {
		prefix      string
		pluginPaths []string
		listDir     listDir
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPlugins(tt.args.prefix, tt.args.pluginPaths, tt.args.listDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPlugins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPlugins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPluginsInPathAux(t *testing.T) {
	type args struct {
		prefix  string
		path    string
		listDir listDir
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "SingleDirPath",
			args: args{
				prefix: "plugin-",
				path:   "testdir",
				listDir: func(dir string) ([]os.FileInfo, error) {
					return []os.FileInfo{
						fileInfoMock{
							name:  "plugin-one",
							isDir: false,
						},
						fileInfoMock{
							name:  "plugin-two",
							isDir: false,
						},
					}, nil
				},
			},
			want:    []string{"plugin-one", "plugin-two"},
			wantErr: false,
		},
		{
			name: "MultiDirPath",
			args: args{
				prefix: "plugin-",
				path:   "testdir1:testdir2",
				listDir: func(dir string) ([]os.FileInfo, error) {
					if dir == "testdir1" {
						return []os.FileInfo{
							fileInfoMock{
								name:  "plugin-one",
								isDir: false,
							},
							fileInfoMock{
								name:  "plugin-two",
								isDir: false,
							},
						}, nil
					}
					return []os.FileInfo{
						fileInfoMock{
							name:  "plugin-three",
							isDir: false,
						},
						fileInfoMock{
							name:  "plugin-four",
							isDir: false,
						},
					}, nil
				},
			},
			want:    []string{"plugin-one", "plugin-two", "plugin-three", "plugin-four"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPluginsInPathAux(tt.args.prefix, tt.args.path, tt.args.listDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginsInPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPluginsInPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPluginsInDirectory(t *testing.T) {
	type args struct {
		prefix  string
		dir     string
		listDir listDir
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "MultiplePlugins",
			args: args{
				prefix: "plugin-",
				dir:    "testdir",
				listDir: func(string) ([]os.FileInfo, error) {
					return []os.FileInfo{
						fileInfoMock{
							name:  "plugin-one",
							isDir: false,
						},
						fileInfoMock{
							name:  "plugin-two",
							isDir: false,
						},
					}, nil
				},
			},
			want:    []string{"plugin-one", "plugin-two"},
			wantErr: false,
		},
		{
			name: "OnePlugin",
			args: args{
				prefix: "plugin-",
				dir:    "testdir",
				listDir: func(string) ([]os.FileInfo, error) {
					return []os.FileInfo{
						fileInfoMock{
							name:  "plugin-one",
							isDir: false,
						},
						fileInfoMock{
							name:  "not-plugin",
							isDir: false,
						},
					}, nil
				},
			},
			want:    []string{"plugin-one"},
			wantErr: false,
		},
		{
			name: "NoPlugins",
			args: args{
				prefix: "plugin-",
				dir:    "testdir",
				listDir: func(string) ([]os.FileInfo, error) {
					return []os.FileInfo{
						fileInfoMock{
							name:  "not-plugin1",
							isDir: false,
						},
						fileInfoMock{
							name:  "not-plugin2",
							isDir: false,
						},
					}, nil
				},
			},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "Recursive",
			args: args{
				prefix: "plugin-",
				dir:    "testdir",
				listDir: func(dir string) ([]os.FileInfo, error) {
					if dir == "dir" {
						return []os.FileInfo{
							fileInfoMock{
								name:  "plugin-dir-one",
								isDir: false,
							},
							fileInfoMock{
								name:  "plugin-dir-two",
								isDir: false,
							},
						}, nil
					}

					return []os.FileInfo{
						fileInfoMock{
							name:  "plugin-one",
							isDir: false,
						},
						fileInfoMock{
							name:  "dir",
							isDir: true,
						},
					}, nil
				},
			},
			want: []string{
				"plugin-one", "plugin-dir-one", "plugin-dir-two",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPluginsInDirectory(tt.args.prefix, tt.args.dir, tt.args.listDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginsInDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPluginsInDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

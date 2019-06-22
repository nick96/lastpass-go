package discovery

import (
	"reflect"
	"testing"

	goplugin "github.com/hashicorp/go-plugin"
)

func TestMap(t *testing.T) {
	type args struct {
		prefix      string
		pluginPaths []string
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := Map(tt.args.prefix, tt.args.pluginPaths)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpandName(t *testing.T) {
	type args struct {
		name        string
		prefix      string
		pluginPaths []string
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandName(tt.args.name, tt.args.prefix, tt.args.pluginPaths)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExpandName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPlugins(t *testing.T) {
	type args struct {
		prefix      string
		pluginPaths []string
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPlugins(tt.args.prefix, tt.args.pluginPaths)
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

func Test_findPluginsInPath(t *testing.T) {
	type args struct {
		prefix string
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPluginsInPath(tt.args.prefix)
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
		prefix string
		dir    string
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := findPluginsInDirectory(tt.args.prefix, tt.args.dir)
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

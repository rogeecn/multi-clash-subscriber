package source

import (
	"multi-clash-subscriber/config"
	"reflect"
	"testing"
)

func TestSource_download(t *testing.T) {
	type fields struct {
		source *config.Source
	}
	tests := []struct {
		name    string
		fields  fields
		want    *config.Clash
		wantErr bool
	}{
		{
			name: "test download",
			fields: fields{
				source: &config.Source{
					Name:         "test",
					SubscribeURL: "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0&flag=clash",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test download",
			fields: fields{
				source: &config.Source{
					Name:         "test",
					SubscribeURL: "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Source{
				source: tt.fields.source,
			}
			_, err := c.download()
			if (err != nil) != tt.wantErr {
				t.Errorf("Source.download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSource_filterProxies(t *testing.T) {
	type fields struct {
		source *config.Source
	}
	type args struct {
		proxies []config.Proxy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []config.Proxy
	}{
		{
			name: "test filterProxies",
			fields: fields{
				source: &config.Source{
					Name:         "test",
					SubscribeURL: "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0&flag=clash",
					FilterChars:  []string{"{emoji}", "{space}"},
					IgnoreNodes:  []string{"abc"},
				},
			},
			args: args{
				proxies: []config.Proxy{
					{Name: "abc🎯"},
					{Name: "a b c "},
					{Name: "abc"},
				},
			},
			want: []config.Proxy{},
		},
		{
			name: "test filterProxies",
			fields: fields{
				source: &config.Source{
					Name:         "test",
					SubscribeURL: "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0&flag=clash",
					FilterChars:  []string{"{emoji}", "{space}"},
					IgnoreNodes:  []string{"abc"},
				},
			},
			args: args{
				proxies: []config.Proxy{
					{Name: "abc1"},
				},
			},
			want: []config.Proxy{
				{
					Name:     "abc1",
					Type:     "",
					Server:   "",
					Port:     0,
					Cipher:   "",
					Password: "",
					UDP:      false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Source{
				source: tt.fields.source,
			}
			if got := c.filterProxies(tt.args.proxies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Source.filterProxies() = %v, want %v", got, tt.want)
			}
		})
	}
}

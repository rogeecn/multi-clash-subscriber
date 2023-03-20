package source

import (
	"multi-clash-subscriber/config"
	"reflect"
	"testing"
)

func TestSource_download(t *testing.T) {
	type fields struct {
		subscribe *config.Subscribe
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
				subscribe: &config.Subscribe{
					Name: "test",
					URL:  "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0&flag=clash",
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test download",
			fields: fields{
				subscribe: &config.Subscribe{
					Name: "test",
					URL:  "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Subscribe{
				subscribe: tt.fields.subscribe,
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
		subscribe *config.Subscribe
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
				subscribe: &config.Subscribe{
					Name:        "test",
					URL:         "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0&flag=clash",
					FilterChars: []string{"{emoji}", "{space}"},
					IgnoreChars: []string{"abc"},
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
				subscribe: &config.Subscribe{
					Name:        "test",
					URL:         "https://sub.wl-sub1.com/api/v1/client/subscribe?token=f6bf7db5bfe5e2bb3ca809c68ed1d3a0&flag=clash",
					FilterChars: []string{"{emoji}", "{space}"},
					IgnoreChars: []string{"abc"},
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
			c := Subscribe{
				subscribe: tt.fields.subscribe,
			}
			if got := c.filterProxies(tt.args.proxies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Source.filterProxies() = %v, want %v", got, tt.want)
			}
		})
	}
}

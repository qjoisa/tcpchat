package client

import (
	"net"
	"reflect"
	"testing"
)

func TestClient_Start(t *testing.T) {
	type fields struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				conn: tt.fields.conn,
			}
			c.Start()
		})
	}
}

func TestClient_clientReader(t *testing.T) {
	type fields struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				conn: tt.fields.conn,
			}
			c.clientReader()
		})
	}
}

func TestClient_serverWriter(t *testing.T) {
	type fields struct {
		conn net.Conn
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				conn: tt.fields.conn,
			}
			c.serverWriter(tt.args.msg)
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		network string
		address string
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.network, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
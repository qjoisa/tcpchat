package server

import (
	"net"
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		network string
		address string
	}
	tests := []struct {
		name string
		args args
		want Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.network, tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	type fields struct {
		Listener net.Listener
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Listener: tt.fields.Listener,
			}
			if err := s.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_broadcaster(t *testing.T) {
	type fields struct {
		Listener net.Listener
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Listener: tt.fields.Listener,
			}
			s.broadcaster()
		})
	}
}

func TestServer_clientReader(t *testing.T) {
	type fields struct {
		Listener net.Listener
	}
	type args struct {
		conn net.Conn
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
			s := Server{
				Listener: tt.fields.Listener,
			}
			s.clientReader(tt.args.conn)
		})
	}
}

func TestServer_clientWriter(t *testing.T) {
	type fields struct {
		Listener net.Listener
	}
	type args struct {
		conn net.Conn
		ch   <-chan string
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
			s := Server{
				Listener: tt.fields.Listener,
			}
			s.clientWriter(tt.args.conn, tt.args.ch)
		})
	}
}

func TestServer_handleConn(t *testing.T) {
	type fields struct {
		Listener net.Listener
	}
	type args struct {
		conn net.Conn
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
			s := Server{
				Listener: tt.fields.Listener,
			}
			s.handleConn(tt.args.conn)
		})
	}
}

func TestServer_writeMessages(t *testing.T) {
	type fields struct {
		Listener net.Listener
	}
	type args struct {
		conn net.Conn
		msg  []byte
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
			s := Server{
				Listener: tt.fields.Listener,
			}
			s.writeMessages(tt.args.conn, tt.args.msg)
		})
	}
}
package uuid

import (
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试uuid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(UUID())
		})
	}
}

func TestUUID32(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试uuid32",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UUID32()
			t.Log(got)
		})
	}
}

func Test_uuid32(t *testing.T) {
	type args struct {
		u *uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试uuid32-private",
			args: args{u: &uuid.UUID{'a', 'b', 'c'}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uuid32(tt.args.u)
			t.Log(got)
		})
	}
}

func TestUUID22(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "测试UUID22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UUID22()
			t.Log(got)
		})
	}
}

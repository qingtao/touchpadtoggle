package touchpad

import (
	"context"
	"testing"
)

func TestFindTouchpad(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantId  string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				ctx: context.TODO(),
			},
			wantId:  "15",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := FindTouchpad(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTouchpad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("FindTouchpad() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestGetStateOfTouchpad(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name        string
		args        args
		wantEnabled bool
		wantErr     bool
	}{
		{
			name: "enabled",
			args: args{
				ctx: context.TODO(),
				id:  "15",
			},
			wantEnabled: true,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEnabled, err := GetStateOfTouchpad(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStateOfTouchpad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotEnabled != tt.wantEnabled {
				t.Errorf("GetStateOfTouchpad() = %v, want %v", gotEnabled, tt.wantEnabled)
			}
		})
	}
}

func TestToggleStateOfTouchpad(t *testing.T) {
	SetDebug(true)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "toggle",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := FindTouchpad(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToggleStateOfTouchpad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 切换状态
			if err := ToggleStateOfTouchpad(tt.args.ctx, id); (err != nil) != tt.wantErr {
				t.Errorf("ToggleStateOfTouchpad() error = %v, wantErr %v", err, tt.wantErr)
			}

			// 还原状态
			if err := ToggleStateOfTouchpad(tt.args.ctx, id); (err != nil) != tt.wantErr {
				t.Errorf("ToggleStateOfTouchpad() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

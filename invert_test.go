package invert

import "testing"

func TestInvert(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test with correct three letter hex color",
			args: args{
				color: "#FFF",
			},
			want:    "#000000",
			wantErr: false,
		},
		{
			name: "Test with correct six length hex color",
			args: args{
				color: "#FFFFFF",
			},
			want:    "#000000",
			wantErr: false,
		},
		{
			name: "Test with correct six length hex color without hash prefix",
			args: args{
				color: "fff",
			},
			want:    "000000",
			wantErr: false,
		},
		{
			name: "Test with invalid length hex color",
			args: args{
				color: "ff",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test with valid length but invalid hex color",
			args: args{
				color: "ff$",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Invert(tt.args.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("Invert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Invert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

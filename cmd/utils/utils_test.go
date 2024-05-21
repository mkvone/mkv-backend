package utils

import "testing"

func TestEncodeValidatorAddress(t *testing.T) {
	type args struct {
		pubKeyBase64 string
		valaddr      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{

		{
			name: "Test EncodeValidatorAddress",
			args: args{
				pubKeyBase64: "58mtQQ+B1UbLVFcpk2sTgjz6CZrUmLczjPl0LC2i34w=",
				valaddr:      "odinvaloper1fupf80rhes3vu25vgjsk6sp8t68urpldecak8y",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeValidatorAddress(tt.args.pubKeyBase64, tt.args.valaddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeValidatorAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("EncodeValidatorAddress() = %v, want %v", got, tt.want)
			// }
			// if success then print the result
			t.Logf("EncodeValidatorAddress() = %v", got)

		})
	}
}

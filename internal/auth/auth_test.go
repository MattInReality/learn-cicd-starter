package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIHeader(t *testing.T) {
	type test struct {
		Name       string
		HeaderName string
		TokenValue string
		Want       string
		Err        bool
		ErrMsg     string
	}

	tests := []test{
		{
			Name:       "Test correct passing case",
			HeaderName: "Authorization",
			TokenValue: "ApiKe 123456897W89798123423",
			Want:       "123456897W89798123423",
			Err:        false,
			ErrMsg:     "",
		},
		{
			Name:       "Test incorrect token key",
			HeaderName: "Authorization",
			TokenValue: "Bearer 123456897W89798123423",
			Want:       "123456897W89798123423",
			Err:        true,
			ErrMsg:     "malformed authorization header",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var header http.Header = make(map[string][]string)
			header.Add(test.HeaderName, test.TokenValue)
			token, err := GetAPIKey(header)
			if err != nil {
				if test.Err {
					return
				}
				t.Errorf("could not parse properly formated token: %s.\nerror: %s\n", test.TokenValue, err.Error())
				return
			}
			if token != test.Want {
				t.Errorf("Output value: %s does not equal result: %s\n", token, test.Want)
				return
			}
		})
	}

}

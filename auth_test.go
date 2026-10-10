package twitter

import (
	"testing"
	errors2 "twitter/pkg/errors"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Validate(t *testing.T) {
	testCase := []struct {
		name string
		in   *RegisterInput
		err  error
	}{
		{
			name: "valid",
			in: &RegisterInput{
				Username:          "username",
				Password:          "password",
				Email:             "jafar@yahoo.com",
				ConfirmedPassword: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			in: &RegisterInput{
				Username:          "username",
				Password:          "password",
				Email:             "jafar@yahoo.com",
				ConfirmedPassword: "password",
			},
			err: errors2.ErrValidation,
		},
		{
			name: "username too short",
			in: &RegisterInput{
				Username:          "u",
				Password:          "password",
				Email:             "abbassmortazavi@gmail.com",
				ConfirmedPassword: "password",
			},
			err: errors2.ErrValidation,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.Validate()
			if tt.err == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.err)
			}
		})
	}
}

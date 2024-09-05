package command_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/v1/capability/command"
)

func TestTop(t *testing.T) {
	require.Equal(t, "/", command.Top().String())
}

func TestIsValidCommand(t *testing.T) {
	t.Parallel()

	t.Run("succeeds when", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range validTestcases(t) {
			testcase := testcase

			t.Run(testcase.name, func(t *testing.T) {
				t.Parallel()

				require.True(t, command.IsValid(testcase.inp))
			})
		}
	})

	t.Run("fails when", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range invalidTestcases(t) {
			testcase := testcase

			t.Run(testcase.name, func(t *testing.T) {
				t.Parallel()

				require.False(t, command.IsValid(testcase.inp))
			})
		}
	})
}

func TestParseCommand(t *testing.T) {
	t.Parallel()

	t.Run("succeeds when", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range validTestcases(t) {
			testcase := testcase

			t.Run(testcase.name, func(t *testing.T) {
				t.Parallel()

				cmd, err := command.Parse("/elem0/elem1/elem2")
				require.NoError(t, err)
				require.NotNil(t, cmd)
			})
		}
	})

	t.Run("fails when", func(t *testing.T) {
		t.Parallel()

		for _, testcase := range invalidTestcases(t) {
			testcase := testcase

			t.Run(testcase.name, func(t *testing.T) {
				t.Parallel()

				cmd, err := command.Parse(testcase.inp)
				require.ErrorIs(t, err, testcase.err)
				require.Nil(t, cmd)
			})
		}
	})
}

type testcase struct {
	name string
	inp  string
}

func validTestcases(t *testing.T) []testcase {
	t.Helper()

	cmds := []string{
		"/",
		"/crud",
		"/crud/create",
		"/stack/pop",
		"/crypto/sign",
		"/foo/bar/baz/qux/quux",
		"/ほげ/ふが",
	}

	testcases := make([]testcase, len(cmds))

	for i, inp := range cmds {
		testcases[i] = testcase{
			name: "Command is " + inp,
			inp:  inp,
		}
	}

	return testcases
}

type errorTestcase struct {
	testcase
	err error
}

func invalidTestcases(t *testing.T) []errorTestcase {
	t.Helper()

	return []errorTestcase{
		{
			testcase: testcase{
				name: "leading slash is missing",
				inp:  "elem0/elem1/elem2",
			},
			err: command.ErrRequiresLeadingSlash,
		},
		{
			testcase: testcase{
				name: "trailing slash is present",
				inp:  "/elem0/elem1/elem2/",
			},
			err: command.ErrDisallowsTrailingSlash,
		},
		{
			testcase: testcase{
				name: "uppercase character are present",
				inp:  "/elem0/Elem1/elem2",
			},
			err: command.ErrRequiresLowercase,
		},
	}
}

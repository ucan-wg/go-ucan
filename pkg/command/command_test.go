package command_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ucan-wg/go-ucan/pkg/command"
)

func TestTop(t *testing.T) {
	require.Equal(t, "/", command.Top().String())
}

func TestIsValidCommand(t *testing.T) {
	t.Run("succeeds when", func(t *testing.T) {
		for _, testcase := range validTestcases(t) {
			t.Run(testcase.name, func(t *testing.T) {
				require.True(t, command.IsValid(testcase.inp))
			})
		}
	})

	t.Run("fails when", func(t *testing.T) {
		for _, testcase := range invalidTestcases(t) {
			t.Run(testcase.name, func(t *testing.T) {
				require.False(t, command.IsValid(testcase.inp))
			})
		}
	})
}

func TestNew(t *testing.T) {
	require.Equal(t, command.Top(), command.New())
	require.Equal(t, "/foo", command.New("foo").String())
	require.Equal(t, "/foo/bar", command.New("foo", "bar").String())
	require.Equal(t, "/foo/bar/baz", command.New("foo", "bar/baz").String())
}

func TestParseCommand(t *testing.T) {
	t.Run("succeeds when", func(t *testing.T) {
		for _, testcase := range validTestcases(t) {
			t.Run(testcase.name, func(t *testing.T) {
				cmd, err := command.Parse("/elem0/elem1/elem2")
				require.NoError(t, err)
				require.NotEmpty(t, cmd)
			})
		}
	})

	t.Run("fails when", func(t *testing.T) {
		for _, testcase := range invalidTestcases(t) {
			t.Run(testcase.name, func(t *testing.T) {
				cmd, err := command.Parse(testcase.inp)
				require.ErrorIs(t, err, testcase.err)
				require.Zero(t, cmd)
			})
		}
	})
}

func TestEquality(t *testing.T) {
	require.True(t, command.MustParse("/foo/bar/baz") == command.MustParse("/foo/bar/baz"))
	require.False(t, command.MustParse("/foo/bar/baz") == command.MustParse("/foo/bar/bazz"))
	require.False(t, command.MustParse("/foo/bar") == command.MustParse("/foo/bar/baz"))
}

func TestJoin(t *testing.T) {
	require.Equal(t, "/foo", command.Top().Join("foo").String())
	require.Equal(t, "/foo/bar", command.Top().Join("foo/bar").String())
	require.Equal(t, "/foo/bar", command.Top().Join("foo", "bar").String())
	require.Equal(t, "/faz/boz/foo/bar", command.MustParse("/faz/boz").Join("foo/bar").String())
	require.Equal(t, "/faz/boz/foo/bar", command.MustParse("/faz/boz").Join("foo", "bar").String())
}

func TestSegments(t *testing.T) {
	require.Empty(t, command.Top().Segments())
	require.Equal(t, []string{"foo", "bar", "baz"}, command.MustParse("/foo/bar/baz").Segments())
}

func TestCovers(t *testing.T) {
	require.True(t, command.MustParse("/foo/bar/baz").Covers(command.MustParse("/foo/bar/baz")))
	require.True(t, command.MustParse("/foo/bar").Covers(command.MustParse("/foo/bar/baz")))
	require.False(t, command.MustParse("/foo/bar/baz").Covers(command.MustParse("/foo/bar")))
	require.True(t, command.MustParse("/").Covers(command.MustParse("/foo")))
	require.True(t, command.MustParse("/").Covers(command.MustParse("/foo/bar/baz")))
	require.False(t, command.MustParse("/foo").Covers(command.MustParse("/foo00")))
	require.False(t, command.MustParse("/foo/bar").Covers(command.MustParse("/foo/bar00")))
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

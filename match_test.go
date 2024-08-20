package policy

import (
	"testing"

	"github.com/ipfs/go-cid"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/storacha-network/go-ucanto/core/policy/literal"
	"github.com/storacha-network/go-ucanto/core/policy/selector"
	"github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	t.Run("equality string", func(t *testing.T) {
		np := basicnode.Prototype.String
		nb := np.NewBuilder()
		nb.AssignString("test")
		nd := nb.Build()

		pol := Policy{Equal(selector.MustParse("."), literal.String("test"))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.String("test2"))}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.Int(138))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("equality int", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{Equal(selector.MustParse("."), literal.Int(138))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.Int(1138))}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.String("138"))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("equality float", func(t *testing.T) {
		np := basicnode.Prototype.Float
		nb := np.NewBuilder()
		nb.AssignFloat(1.138)
		nd := nb.Build()

		pol := Policy{Equal(selector.MustParse("."), literal.Float(1.138))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.Float(11.38))}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.String("138"))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("equality IPLD Link", func(t *testing.T) {
		l0 := cidlink.Link{Cid: cid.MustParse("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq")}
		l1 := cidlink.Link{Cid: cid.MustParse("bafkreifau35r7vi37tvbvfy3hdwvgb4tlflqf7zcdzeujqcjk3rsphiwte")}

		np := basicnode.Prototype.Link
		nb := np.NewBuilder()
		nb.AssignLink(l0)
		nd := nb.Build()

		pol := Policy{Equal(selector.MustParse("."), literal.Link(l0))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.Link(l1))}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Equal(selector.MustParse("."), literal.String("bafybeif4owy5gno5lwnixqm52rwqfodklf76hsetxdhffuxnplvijskzqq"))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("equality string in map", func(t *testing.T) {
		np := basicnode.Prototype.Map
		nb := np.NewBuilder()
		ma, _ := nb.BeginMap(1)
		ma.AssembleKey().AssignString("foo")
		ma.AssembleValue().AssignString("bar")
		ma.Finish()
		nd := nb.Build()

		pol := Policy{Equal(selector.MustParse(".foo"), literal.String("bar"))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse(".[\"foo\"]"), literal.String("bar"))}
		ok = Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse(".foo"), literal.String("baz"))}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Equal(selector.MustParse(".foobar"), literal.String("bar"))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("equality string in list", func(t *testing.T) {
		np := basicnode.Prototype.List
		nb := np.NewBuilder()
		la, _ := nb.BeginList(1)
		la.AssembleValue().AssignString("foo")
		la.Finish()
		nd := nb.Build()

		pol := Policy{Equal(selector.MustParse(".[0]"), literal.String("foo"))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Equal(selector.MustParse(".[1]"), literal.String("foo"))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("inequality gt int", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{GreaterThan(selector.MustParse("."), literal.Int(1))}
		ok := Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("inequality gte int", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Int(1))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Int(138))}
		ok = Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("inequality gt float", func(t *testing.T) {
		np := basicnode.Prototype.Float
		nb := np.NewBuilder()
		nb.AssignFloat(1.38)
		nd := nb.Build()

		pol := Policy{GreaterThan(selector.MustParse("."), literal.Float(1))}
		ok := Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("inequality gte float", func(t *testing.T) {
		np := basicnode.Prototype.Float
		nb := np.NewBuilder()
		nb.AssignFloat(1.38)
		nd := nb.Build()

		pol := Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Float(1))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{GreaterThanOrEqual(selector.MustParse("."), literal.Float(1.38))}
		ok = Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("inequality lt int", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{LessThan(selector.MustParse("."), literal.Int(1138))}
		ok := Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("inequality lte int", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{LessThanOrEqual(selector.MustParse("."), literal.Int(1138))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{LessThanOrEqual(selector.MustParse("."), literal.Int(138))}
		ok = Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("negation", func(t *testing.T) {
		np := basicnode.Prototype.Bool
		nb := np.NewBuilder()
		nb.AssignBool(false)
		nd := nb.Build()

		pol := Policy{Not(Equal(selector.MustParse("."), literal.Bool(true)))}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{Not(Equal(selector.MustParse("."), literal.Bool(false)))}
		ok = Match(pol, nd)
		require.False(t, ok)
	})

	t.Run("conjunction", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{
			And(
				GreaterThan(selector.MustParse("."), literal.Int(1)),
				LessThan(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{
			And(
				GreaterThan(selector.MustParse("."), literal.Int(1)),
				Equal(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{And()}
		ok = Match(pol, nd)
		require.True(t, ok)
	})

	t.Run("disjunction", func(t *testing.T) {
		np := basicnode.Prototype.Int
		nb := np.NewBuilder()
		nb.AssignInt(138)
		nd := nb.Build()

		pol := Policy{
			Or(
				GreaterThan(selector.MustParse("."), literal.Int(138)),
				LessThan(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok := Match(pol, nd)
		require.True(t, ok)

		pol = Policy{
			Or(
				GreaterThan(selector.MustParse("."), literal.Int(138)),
				Equal(selector.MustParse("."), literal.Int(1138)),
			),
		}
		ok = Match(pol, nd)
		require.False(t, ok)

		pol = Policy{Or()}
		ok = Match(pol, nd)
		require.True(t, ok)
	})
}

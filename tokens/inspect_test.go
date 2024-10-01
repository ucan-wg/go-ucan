package tokens_test

// func TestInspect(t *testing.T) {
// 	t.Parallel()
//
// 	for _, filename := range []string{
// 		"example.dagcbor",
// 		"example.dagjson",
// 	} {
// 		t.Run(filename, func(t *testing.T) {
// 			t.Parallel()
//
// 			data := golden.Get(t, filename)
// 			expSig, err := base64.RawStdEncoding.DecodeString("fPqfwL3iFpbw9SvBiq0DIbUurv9o6c36R08tC/yslGrJcwV51ghzWahxdetpEf6T5LCszXX9I/K8khvnmAxjAg")
// 			require.NoError(t, err)
//
// 			info, err := tokens.Inspect(data)
// 			require.NoError(t, err)
// 			assert.Equal(t, expSig, info.Signature)
// 			assert.Equal(t, "ucan/example@v1.0.0-rc.1", info.Tag)
// 			assert.Equal(t, []byte{0x34, 0xed, 0x1, 0x71}, info.VarsigHeader)
// 		})
// 	}
// }

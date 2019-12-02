package crypto

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestECDSA_MarshalUnmarshal(t *testing.T) {
	priv, pub := generateECDSA(rand.Reader)
	require.NotNil(t, priv)
	require.NotNil(t, pub)

	data, err := priv.MarshalBinary()
	require.NoError(t, err)

	priv1 := new(ECDSAPriv)
	require.NoError(t, priv1.UnmarshalBinary(data))
	require.Equal(t, priv, priv1)

	data, err = pub.MarshalBinary()
	require.NoError(t, err)

	pub1 := new(ECDSAPub)
	require.NoError(t, pub1.UnmarshalBinary(data))
	require.Equal(t, pub, pub1)

	require.Error(t, pub1.UnmarshalBinary([]byte{0, 1, 2, 3}))
}
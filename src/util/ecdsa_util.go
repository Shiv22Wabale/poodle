package util

import (
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 32 byte array
func SumSHA256(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func SumSHA256d(data []byte) []byte {
	return SumSHA256(SumSHA256(data))
}

// 16 byte array
func SumMD5(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}

func AESPad(input []byte) []byte {
	if len(input) >= aes.BlockSize {
		return input[len(input)-aes.BlockSize:]
	}

	padded := make([]byte, aes.BlockSize)
	copy(padded[:], input)
	return padded
}

func ECDSAGenerateKey() *ecdsa.PrivateKey {
	curve := elliptic.P256()
	priv_key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return priv_key
	}
}

func ECDSASign(priv_key *ecdsa.PrivateKey, hash []byte) (r, s *big.Int, err error) {
	return ecdsa.Sign(rand.Reader, priv_key, hash)
}

func ECDSAVerify(pub_key *ecdsa.PublicKey, hash []byte, r, s *big.Int) bool {
	return ecdsa.Verify(pub_key, hash, r, s)
}

func ECDSAGetPrivateKey(priv_bytes []byte) *ecdsa.PrivateKey {
	priv_key := new(ecdsa.PrivateKey)
	priv_key.PublicKey.Curve = elliptic.P256()
	priv_key.D = new(big.Int)
	priv_key.D.SetBytes(priv_bytes)
	priv_key.PublicKey.X, priv_key.PublicKey.Y = priv_key.PublicKey.Curve.ScalarBaseMult(priv_bytes)
	return priv_key
}

func ECDSAGetPublicKey(pub_bytes []byte) *ecdsa.PublicKey {
	pub_key := new(ecdsa.PublicKey)
	pub_key.Curve = elliptic.P256()
	pub_key.X = new(big.Int)
	pub_key.X.SetBytes(pub_bytes)
	pub_key.Y = ECDSACalculateY(pub_key.Curve, pub_key.X)
	return pub_key
}

func ECDSACalculateY(curve elliptic.Curve, x *big.Int) *big.Int {
	// y² = x³ - 3x + b
	//y2 := new(big.Int).Mul(y, y)
	//y2.Mod(y2, curve.Params().P)

	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	threeX := new(big.Int).Lsh(x, 1)
	threeX.Add(threeX, x)

	x3.Sub(x3, threeX)
	x3.Add(x3, curve.Params().B)
	x3.Mod(x3, curve.Params().P)

	y := new(big.Int)
	y = y.ModSqrt(x3, curve.Params().P)

	//return x3.Cmp(y2) == 0
	return y
}

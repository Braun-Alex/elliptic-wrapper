package ec

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
)

type ElCPoint struct {
	X *big.Int
	Y *big.Int
}

// Returning base point G on the elliptic curve secp521r1

func BasePointGGet() ElCPoint {
	ellipticSecp521r1Params := elliptic.P521().Params()
	return ElCPointGen(ellipticSecp521r1Params.Gx, ellipticSecp521r1Params.Gy)
}

// Returning ElCPoint structure wrapped in coordinates

func ElCPointGen(x, y *big.Int) ElCPoint {
	return ElCPoint{x, y}
}

// Checking that point is on curve secp521r1

func IsOnCurveCheck(a ElCPoint) bool {
	return elliptic.P521().IsOnCurve(a.X, a.Y)
}

func Eq(a, b ElCPoint) bool {
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}

// Adding two different elliptic curve points

func AddElCPoints(a, b ElCPoint) ElCPoint {
	return ElCPointGen(elliptic.P521().Add(a.X, a.Y, b.X, b.Y))
}

// Double multiplying of elliptic curve point

func DoubleElCPoints(a ElCPoint) ElCPoint {
	return ElCPointGen(elliptic.P521().Double(a.X, a.Y))
}

// Scalar multiplying of elliptic curve point

func ScalarMult(k big.Int, a ElCPoint) ElCPoint {
	return ElCPointGen(elliptic.P521().ScalarMult(a.X, a.Y, k.Bytes()))
}

func ElCPointToString(point ElCPoint) string {
	return string(elliptic.MarshalCompressed(elliptic.P521(), point.X, point.Y))
}

func StringToElCPoint(s string) ElCPoint {
	return ElCPointGen(elliptic.UnmarshalCompressed(elliptic.P521(), []byte(s)))
}

func PrintElCPoint(point ElCPoint) {
	const HexEncoding = 16
	fmt.Printf("Elliptic curve point has such coordinates: \n%s\n%s\n",
		point.X.Text(HexEncoding), point.Y.Text(HexEncoding))
}

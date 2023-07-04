package test

import (
	"crypto/rand"
	"elliptic-wrapper/pkg/ec"
	"fmt"
	"math/big"
	"testing"
)

func SetRandom(bits int) *big.Int {
	randomNumber, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2),
		big.NewInt(int64(bits)), nil))
	if err != nil {
		panic("Could not be generated random " + string(rune(bits)) + "-bit number")
	}
	return randomNumber
}

func TestGroupOperation(t *testing.T) {
	g := ec.BasePointGGet()
	randomFirstScalar := SetRandom(521)
	randomSecondScalar := SetRandom(521)
	randomFirstPoint := ec.ScalarMult(*randomFirstScalar, g)
	randomSecondPoint := ec.ScalarMult(*randomSecondScalar, g)
	resultPoint := ec.AddElCPoints(randomFirstPoint, randomSecondPoint)
	if !ec.IsOnCurveCheck(resultPoint) {
		t.Error("Group operation result does not belong to the elliptic curve")
	}
}

func TestGroupAssociativity(t *testing.T) {
	g := ec.BasePointGGet()
	k := SetRandom(521)
	d := SetRandom(521)
	h1 := ec.ScalarMult(*d, g)
	h2 := ec.ScalarMult(*k, h1)
	h3 := ec.ScalarMult(*k, g)
	h4 := ec.ScalarMult(*d, h3)
	if !ec.Eq(h2, h4) {
		t.Error("Group of elliptic curve points has no group associativity")
	}
}

func TestGroupNeutralElement(t *testing.T) {
	g := ec.BasePointGGet()
	randomScalar := SetRandom(521)
	randomPoint := ec.ScalarMult(*randomScalar, g)
	infinitePoint := ec.ElCPointGen(big.NewInt(0), big.NewInt(0))
	leftPoint := ec.AddElCPoints(randomPoint, infinitePoint)
	rightPoint := ec.AddElCPoints(infinitePoint, randomPoint)
	if !ec.Eq(leftPoint, rightPoint) || !ec.Eq(randomPoint, leftPoint) ||
		!ec.Eq(randomPoint, rightPoint) {
		t.Error("Group of elliptic curve points has no neutral element that a*0 = 0*a = a")
	}
}

func TestGroupCommutativity(t *testing.T) {
	g := ec.BasePointGGet()
	randomFirstScalar := SetRandom(521)
	randomSecondScalar := SetRandom(521)
	randomFirstPoint := ec.ScalarMult(*randomFirstScalar, g)
	randomSecondPoint := ec.ScalarMult(*randomSecondScalar, g)
	if !ec.Eq(ec.AddElCPoints(randomFirstPoint, randomSecondPoint),
		ec.AddElCPoints(randomSecondPoint, randomFirstPoint)) {
		t.Error("Group of elliptic curve points has no group commutativity")
	}
}

func TestMainElCOperations(t *testing.T) {
	privateKey := new(big.Int)
	privateKey.SetString("1e5f8e1c393587e9a5c1e69d48df6875aa2bc95944413ff835f94671f43cdf9"+
		"38b4464aab3ae565b6de59c21b14326d3bfddc6b4bdcdfda9a061b206d159c40d959", ec.HexEncoding)
	actualPublicKey := ec.ScalarMult(*privateKey, ec.BasePointGGet())
	expectedPublicKey := ec.StringToElCPoint("0301c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9f1" +
		"317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e69510a25caed69758156f1")
	if !ec.Eq(actualPublicKey, expectedPublicKey) {
		t.Error("Elliptic curve operations have not been properly implemented")
	}
}

func TestIsElCPointOnTheCurve(t *testing.T) {
	ellipticCurvePoint := ec.StringToElCPoint("0301c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9f1" +
		"317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e69510a25caed69758156f1")
	if !ec.IsOnCurveCheck(ellipticCurvePoint) {
		t.Error("Elliptic curve point is not on the elliptic curve")
	}
	fmt.Println(ellipticCurvePoint.X.Text(16), ellipticCurvePoint.Y.Text(16))
}

func TestEncodingOfElCPoint(t *testing.T) {
	x, _ := new(big.Int).SetString("1c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9"+
		"f1317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e6951"+
		"0a25caed69758156f1", ec.HexEncoding)
	y, _ := new(big.Int).SetString("1385b7eaaf883dbc5b71f8af3c60c576b48e70e1a5f6d1862b9dafac1"+
		"4a75038bca9b5baf41ad2d96c22e55a192f3c8516187166ee0b15"+
		"4d08b66ed8b6f654a645", ec.HexEncoding)
	point := ec.ElCPointGen(x, y)
	encodedPoint := ec.ElCPointToString(point)
	if encodedPoint != "0301c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9f1"+
		"317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e6951"+
		"0a25caed69758156f1" {
		t.Error("Encoded elliptic curve point has not been properly encoded")
	}
}

func TestDecodingOfElCPoint(t *testing.T) {
	decodedPoint := ec.StringToElCPoint("0301c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9f1" +
		"317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e69510a25caed69758156f1")
	x, _ := new(big.Int).SetString("1c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9"+
		"f1317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e6951"+
		"0a25caed69758156f1", ec.HexEncoding)
	y, _ := new(big.Int).SetString("1385b7eaaf883dbc5b71f8af3c60c576b48e70e1a5f6d1862b9dafac1"+
		"4a75038bca9b5baf41ad2d96c22e55a192f3c8516187166ee0b15"+
		"4d08b66ed8b6f654a645", ec.HexEncoding)
	actualPoint := ec.ElCPointGen(x, y)
	if !ec.Eq(decodedPoint, actualPoint) {
		t.Error("Encoded elliptic curve point has not been properly decoded")
	}
}

func TestConvertingOfElCPoint(t *testing.T) {
	x, _ := new(big.Int).SetString("1c877c1a9aca747b44817d61d9a307a3a50243f9920dbdac9"+
		"f1317557c75cfd8625b2d549797688a1498c611f9f6a0a4fa828e667263e6951"+
		"0a25caed69758156f1", ec.HexEncoding)
	y, _ := new(big.Int).SetString("1385b7eaaf883dbc5b71f8af3c60c576b48e70e1a5f6d1862b9dafac1"+
		"4a75038bca9b5baf41ad2d96c22e55a192f3c8516187166ee0b15"+
		"4d08b66ed8b6f654a645", ec.HexEncoding)
	actualPoint := ec.ElCPointGen(x, y)
	encodedPoint := ec.ElCPointToString(actualPoint)
	decodedPoint := ec.StringToElCPoint(encodedPoint)
	if !ec.Eq(decodedPoint, actualPoint) {
		t.Error("Converting of elliptic curve point has not been properly implemented")
	}
}

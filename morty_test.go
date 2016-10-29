package main

import (
	"bytes"
	"net/url"
	"testing"
)

type AttrTestCase struct {
	AttrName       []byte
	AttrValue      []byte
	ExpectedOutput []byte
}

var attrTestData []*AttrTestCase = []*AttrTestCase{
	&AttrTestCase{
		[]byte("href"),
		[]byte("./x"),
		[]byte(` href="./?mortyurl=http%3A%2F%2F127.0.0.1%2Fx"`),
	},
	&AttrTestCase{
		[]byte("src"),
		[]byte("http://x.com/y"),
		[]byte(` src="./?mortyurl=http%3A%2F%2Fx.com%2Fy"`),
	},
	&AttrTestCase{
		[]byte("action"),
		[]byte("/z"),
		[]byte(` action="./?mortyurl=http%3A%2F%2F127.0.0.1%2Fz"`),
	},
	&AttrTestCase{
		[]byte("onclick"),
		[]byte("console.log(document.cookies)"),
		nil,
	},
}

func TestAttrSanitizer(t *testing.T) {
	u, _ := url.Parse("http://127.0.0.1/")
	rc := &RequestConfig{nil, u}
	for _, testCase := range attrTestData {
		out := bytes.NewBuffer(nil)
		sanitizeAttr(rc, out, testCase.AttrName, testCase.AttrValue)
		res, _ := out.ReadBytes(byte(0))
		if !bytes.Equal(res, testCase.ExpectedOutput) {
			t.Errorf(
				`Attribute parse error. Name: "%s", Value: "%s", Expected: %s, Got: %s`,
				testCase.AttrName,
				testCase.AttrValue,
				testCase.ExpectedOutput,
				res,
			)
		}
	}
}
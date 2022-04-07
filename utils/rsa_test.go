package utils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAo5Hi2Lyef0KRQfmVT/VYeqiiW6u6CCmjtl/kp/0eQnnxbyj1
bgJsG6zXnIYsg5T6vpk34F2LSIkipNZ3H1qSKppItpsx05VnKqiMjPO0wmeA0U9M
VXzIReDQB5J4BZWIn0W8bAbkTKTDPjKKPVBJpS4D1TZ+CigSrMZzzy2dkPkGa9Qi
lWiYZh0uiGpFQm9b5vxJsZWtPDkVoRq8p6I33XDfapxhnc6/NvcaZ8SUiLvKHaoH
x4IdFAcCkxYji/q3Murxd7suLe6j9wvP8cNnWOn+tC4YvAQhfdVMahGcwa7S2MmG
+I0novJHrsDZJbXGpIhQVRtLCNOfLYw9eJJSIQIDAQABAoIBAQChiIPfCL757HJp
FiKneSLs1zankzq/rud3RtIYODc27RhKb8eY4ZmD5zhy4mp5c5oAum1dDORtOMjW
wtmXRCmPGxhuBWu6iRP+I7/RvycmyXKEwmwfDID3yWGH6NTQpvDqYcMHZ/Lz4c7S
0LSGgpCIKJ4uZIcqXpK6v61Ez8aIkGBRh7BuacrX0P7jSxxyGix+P0Bi7mpK33+P
KyGLw21w0Xea5HvLFHSLiXbINGtRxgxUhyhDCMFx9zTx1PGAk6uZNRHOi3najS+u
AHqdTJV6znIUY0FYz8lTx17qq20OBxdzD1KZtNZRyp+PJn7WK8iEYeBRTN5wUZVI
yv18D5gBAoGBANfONfm3YGjfpFAFRslnoE7oNcuBRyLlb2UzEbM8rZvA3bSY7lhF
KyxrIXknb9DhruVDwxiS3lneBhxSO1yc0XgB0YR/elWIk4B17Hyd5hBTHZZn2cwH
Ck2VvPJO4cYpS+rA5nwBvqf+YeBeozv6Y5urYh60Go+8+/teSIM+lfFBAoGBAMIJ
CSByHmkF6cAtRiBndDAhNV6KXnXMGmHkVfeeZWb+OKAl2OchJKYYR/8HNbZFQS84
FUf3BZQjbBeZ62+it7WPDd8sNvzcoA2ra9RElHtK4GYOW6xytbr5JvTn925K3/7Z
LaHD3eOdNAbfUHICP8rIwb0e6WM9UWDFd0IAL0jhAoGACDZiu3QNCTCpkhoBTZzR
7u7LTUlIKb3NdW5yF/4RJUI7ram8OEO5luEkzvV/PdW40CS6Ae9k4sqeNOm0NJwL
JKP9Eyu+vg6KfnL+YuenX6uqDpU5JpiF/dna7OWvldjweKiKnOF8HeyCUQWI9Llc
nQnym/SCpkGs3S4DWGtDTYECgYBmYwoz2sP9OyeLRtiYRL7knsFgr7cEhbzuJuTO
EwEDBeavawk3fR+vnWRwSEepdIXe+sMQBGMnNkHVelznvX8fCOuS9evWJj02uVve
2IlZPbuHXtA1ARqwHMTXuoev5UapXcRcddRiRROtBK2YW2N6B9FjkCorObH1VH2H
kD1jQQKBgD3Mb1Xwsc5S3pU5pqcAv1Cc7gIErgq0Jkehx4Cxhz2LSrc37mhexLlc
I43NJeZscWPcAPi9A9SWF3r+ZipOvUSBTSB+USoK9+I3kzLNs2HxURCN2qPklrtK
DbL1fd5aEiJMZUPHptUZQM6c18ZbwAUnCA84HZFQLTh2YJnK3dcv
-----END RSA PRIVATE KEY-----
`

var testPublicKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo5Hi2Lyef0KRQfmVT/VY\neqiiW6u6CCmjtl/kp/0eQnnxbyj1bgJsG6zXnIYsg5T6vpk34F2LSIkipNZ3H1qS\nKppItpsx05VnKqiMjPO0wmeA0U9MVXzIReDQB5J4BZWIn0W8bAbkTKTDPjKKPVBJ\npS4D1TZ+CigSrMZzzy2dkPkGa9QilWiYZh0uiGpFQm9b5vxJsZWtPDkVoRq8p6I3\n3XDfapxhnc6/NvcaZ8SUiLvKHaoHx4IdFAcCkxYji/q3Murxd7suLe6j9wvP8cNn\nWOn+tC4YvAQhfdVMahGcwa7S2MmG+I0novJHrsDZJbXGpIhQVRtLCNOfLYw9eJJS\nIQIDAQAB\n-----END PUBLIC KEY-----\n"

func TestRsaEncrypt(t *testing.T) {
	Convey("TestRsaEncrypt ", t, func() {
		Convey("normal", func() {
			plaintext := "this is a message"
			key := testPublicKey
			fmt.Println("plaintext", plaintext)
			cyphertext, err := RsaEncrypt(plaintext, key)
			So(err, ShouldBeNil)
			fmt.Println("cyphertext", cyphertext)
			res, err := RsaDecrypt(cyphertext, testPrivateKey)
			So(err, ShouldBeNil)
			So(plaintext, ShouldEqual, res)
			fmt.Println("plaintext", res)
		})
		Convey("unicode string", func() {
			plaintext := "\uFFFF\u0FFF\uF112\uf0f0\u7766\u1fff\uf1f1\uf1ff\u11f1\uf1f1\uf322"
			key := testPublicKey
			fmt.Println("plaintext", plaintext)
			cyphertext, err := RsaEncrypt(plaintext, key)
			So(err, ShouldBeNil)
			fmt.Println("cyphertext", cyphertext)
			res, err := RsaDecrypt(cyphertext, testPrivateKey)
			So(err, ShouldBeNil)
			So(plaintext, ShouldEqual, res)
			fmt.Println("plaintext", res)
		})
		Convey("empty plaintext", func() {
			plaintext := ""
			key := testPublicKey
			fmt.Println("plaintext", plaintext)
			cyphertext, err := RsaEncrypt(plaintext, key)
			So(err, ShouldBeNil)
			fmt.Println("cyphertext", cyphertext)
			res, err := RsaDecrypt(cyphertext, testPrivateKey)
			So(err, ShouldBeNil)
			So(plaintext, ShouldEqual, res)
			fmt.Println("plaintext", res)
		})
		Convey("empty public key", func() {
			plaintext := "this is a message"
			key := ""
			fmt.Println("plaintext", plaintext)
			_, err := RsaEncrypt(plaintext, key)
			So(err, ShouldNotBeNil)
		})
		Convey("empty private key", func() {
			plaintext := "this is a message"
			key := testPublicKey
			fmt.Println("plaintext", plaintext)
			cyphertext, err := RsaEncrypt(plaintext, key)
			So(err, ShouldBeNil)
			fmt.Println("cyphertext", cyphertext)
			_, err = RsaDecrypt(cyphertext, "")
			So(err, ShouldNotBeNil)
		})
		Convey("empty cypher text", func() {
			cyphertext := ""
			plaintext, err := RsaDecrypt(cyphertext, testPrivateKey)
			So(err, ShouldBeNil)
			So(plaintext, ShouldEqual, cyphertext)

		})
	})
}

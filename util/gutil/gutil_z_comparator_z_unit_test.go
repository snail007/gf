package gutil_test

import (
	"testing"

	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/util/gutil"
)

func Test_ComparatorString(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorString(1, 1), 0)
		gtest.Assert(gutil.ComparatorString(1, 2), -1)
		gtest.Assert(gutil.ComparatorString(2, 1), 1)
	})
}

func Test_ComparatorInt(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorInt(1, 1), 0)
		gtest.Assert(gutil.ComparatorInt(1, 2), -1)
		gtest.Assert(gutil.ComparatorInt(2, 1), 1)
	})
}

func Test_ComparatorInt8(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorInt8(1, 1), 0)
		gtest.Assert(gutil.ComparatorInt8(1, 2), -1)
		gtest.Assert(gutil.ComparatorInt8(2, 1), 1)
	})
}

func Test_ComparatorInt16(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorInt16(1, 1), 0)
		gtest.Assert(gutil.ComparatorInt16(1, 2), -1)
		gtest.Assert(gutil.ComparatorInt16(2, 1), 1)
	})
}

func Test_ComparatorInt32(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorInt32(1, 1), 0)
		gtest.Assert(gutil.ComparatorInt32(1, 2), -1)
		gtest.Assert(gutil.ComparatorInt32(2, 1), 1)
	})
}

func Test_ComparatorInt64(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorInt64(1, 1), 0)
		gtest.Assert(gutil.ComparatorInt64(1, 2), -1)
		gtest.Assert(gutil.ComparatorInt64(2, 1), 1)
	})
}

func Test_ComparatorUint(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorUint(1, 1), 0)
		gtest.Assert(gutil.ComparatorUint(1, 2), -1)
		gtest.Assert(gutil.ComparatorUint(2, 1), 1)
	})
}

func Test_ComparatorUint8(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorUint8(1, 1), 0)
		gtest.Assert(gutil.ComparatorUint8(2, 6), 252)
		gtest.Assert(gutil.ComparatorUint8(2, 1), 1)
	})
}

func Test_ComparatorUint16(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorUint16(1, 1), 0)
		gtest.Assert(gutil.ComparatorUint16(1, 2), 65535)
		gtest.Assert(gutil.ComparatorUint16(2, 1), 1)
	})
}

func Test_ComparatorUint32(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorUint32(1, 1), 0)
		gtest.Assert(gutil.ComparatorUint32(-1000, 2147483640), 2147482656)
		gtest.Assert(gutil.ComparatorUint32(2, 1), 1)
	})
}

func Test_ComparatorUint64(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorUint64(1, 1), 0)
		gtest.Assert(gutil.ComparatorUint64(1, 2), -1)
		gtest.Assert(gutil.ComparatorUint64(2, 1), 1)
	})
}

func Test_ComparatorFloat32(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorFloat32(1, 1), 0)
		gtest.Assert(gutil.ComparatorFloat32(1, 2), -1)
		gtest.Assert(gutil.ComparatorFloat32(2, 1), 1)
	})
}

func Test_ComparatorFloat64(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorFloat64(1, 1), 0)
		gtest.Assert(gutil.ComparatorFloat64(1, 2), -1)
		gtest.Assert(gutil.ComparatorFloat64(2, 1), 1)
	})
}

func Test_ComparatorByte(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorByte(1, 1), 0)
		gtest.Assert(gutil.ComparatorByte(1, 2), 255)
		gtest.Assert(gutil.ComparatorByte(2, 1), 1)
	})
}

func Test_ComparatorRune(t *testing.T) {

	gtest.Case(t, func() {
		gtest.Assert(gutil.ComparatorRune(1, 1), 0)
		gtest.Assert(gutil.ComparatorRune(1, 2), -1)
		gtest.Assert(gutil.ComparatorRune(2, 1), 1)
	})
}

func Test_ComparatorTime(t *testing.T) {
	gtest.Case(t, func() {
		j := gutil.ComparatorTime("2019-06-14", "2019-06-14")
		gtest.Assert(j, 0)

		k := gutil.ComparatorTime("2019-06-15", "2019-06-14")
		gtest.Assert(k, 1)

		l := gutil.ComparatorTime("2019-06-13", "2019-06-14")
		gtest.Assert(l, -1)
	})
}

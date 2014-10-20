package pathreflect

import (
	"testing"

	"github.com/getlantern/testify/assert"
)

type A struct {
	B      *B
	MapB   map[string]*B
	SliceB []*B
	E      map[string]interface{}
}

type B struct {
	S string
	I int
}

func TestSetOnEmptyRoot(t *testing.T) {
	var d *B

	err := Parse("B/E").Set(d, 50)
	assert.Error(t, err)
}

func TestSetOnEmptyParent(t *testing.T) {
	var d *B

	err := Parse("B/E/dude").Set(d, 50)
	assert.Error(t, err)
}

func TestNestedPrimitiveInStruct(t *testing.T) {
	d := makeData()

	err := Parse("B/S").Set(d, "10")
	assert.NoError(t, err)
	err = Parse("B/I").Set(d, 10)
	assert.NoError(t, err)

	assert.Equal(t, "10", d.B.S)
	assert.Equal(t, 10, d.B.I)
}

func TestNestedPrimitiveInMap(t *testing.T) {
	d := makeData()

	err := Parse("MapB/3/S").Set(d, "10")
	assert.NoError(t, err)
	err = Parse("MapB/3/I").Set(d, 10)
	assert.NoError(t, err)

	assert.Equal(t, "10", d.MapB["3"].S)
	assert.Equal(t, 10, d.MapB["3"].I)
}

func TestNestedPrimitiveInSlice(t *testing.T) {
	d := makeData()

	err := Parse("SliceB/1/S").Set(d, "10")
	assert.NoError(t, err)
	err = Parse("SliceB/1/I").Set(d, 10)
	assert.NoError(t, err)

	assert.Equal(t, "10", d.SliceB[1].S)
	assert.Equal(t, 10, d.SliceB[1].I)
}

func TestNestedField(t *testing.T) {
	d := makeData()
	orig := d.B

	err := Parse("B").Set(d, &B{
		S: "10",
		I: 10,
	})

	assert.NoError(t, err)
	assert.Equal(t, "10", d.B.S)
	assert.Equal(t, 10, d.B.I)
	assert.NotEqual(t, d.B, orig)
}

func TestNestedMapEntry(t *testing.T) {
	d := makeData()
	orig := d.MapB["3"]

	err := Parse("MapB/3").Set(d, &B{
		S: "10",
		I: 10,
	})

	assert.NoError(t, err)
	assert.Equal(t, "10", d.MapB["3"].S)
	assert.Equal(t, 10, d.MapB["3"].I)
	assert.NotEqual(t, d.B, orig)
}

func TestNestedSliceEntry(t *testing.T) {
	d := makeData()
	orig := d.SliceB[1]

	err := Parse("SliceB/1").Set(d, &B{
		S: "10",
		I: 10,
	})

	assert.NoError(t, err)
	assert.Equal(t, "10", d.SliceB[1].S)
	assert.Equal(t, 10, d.SliceB[1].I)
	assert.NotEqual(t, d.B, orig)
}

func makeData() *A {
	return &A{
		B: &B{
			S: "5",
			I: 5,
		},
		MapB: map[string]*B{
			"4": &B{
				S: "4",
				I: 4,
			},
			"3": &B{
				S: "3",
				I: 3,
			},
		},
		SliceB: []*B{
			&B{
				S: "0",
				I: 0,
			},
			&B{
				S: "1",
				I: 1,
			},
		},
	}
}

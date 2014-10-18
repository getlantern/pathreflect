package pathreflect

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MainSuite struct{}

var _ = Suite(&MainSuite{})

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

func (s *MainSuite) TestSetOnEmptyRoot(c *C) {
	var d *B

	err := Parse("B/E").Set(d, 50)
	c.Assert(err, Not(IsNil))
}

func (s *MainSuite) TestSetOnEmptyParent(c *C) {
	var d *B

	err := Parse("B/E/dude").Set(d, 50)
	c.Assert(err, Not(IsNil))
}

// func (s *MainSuite) TestSetOnEmptyField(c *C) {
// 	d := makeData()

// 	err := Parse("B/E").Set(d, map[string]interface{}{"dude": 50})
// 	c.Assert(err, IsNil)
// }

func (s *MainSuite) TestNestedPrimitiveInStruct(c *C) {
	d := makeData()

	err := Parse("B/S").Set(d, "10")
	c.Assert(err, IsNil)
	err = Parse("B/I").Set(d, 10)
	c.Assert(err, IsNil)

	c.Assert(d.B.S, Equals, "10")
	c.Assert(d.B.I, Equals, 10)
}

func (s *MainSuite) TestNestedPrimitiveInMap(c *C) {
	d := makeData()

	err := Parse("MapB/3/S").Set(d, "10")
	c.Assert(err, IsNil)
	err = Parse("MapB/3/I").Set(d, 10)
	c.Assert(err, IsNil)

	c.Assert(d.MapB["3"].S, Equals, "10")
	c.Assert(d.MapB["3"].I, Equals, 10)
}

func (s *MainSuite) TestNestedPrimitiveInSlice(c *C) {
	d := makeData()

	err := Parse("SliceB/1/S").Set(d, "10")
	c.Assert(err, IsNil)
	err = Parse("SliceB/1/I").Set(d, 10)
	c.Assert(err, IsNil)

	c.Assert(d.SliceB[1].S, Equals, "10")
	c.Assert(d.SliceB[1].I, Equals, 10)
}

func (s *MainSuite) TestNestedField(c *C) {
	d := makeData()
	orig := d.B

	err := Parse("B").Set(d, &B{
		S: "10",
		I: 10,
	})

	c.Assert(err, IsNil)
	c.Assert(d.B.S, DeepEquals, "10")
	c.Assert(d.B.I, DeepEquals, 10)
	c.Assert(orig, Not(Equals), d.B)
}

func (s *MainSuite) TestNestedMapEntry(c *C) {
	d := makeData()
	orig := d.MapB["3"]

	err := Parse("MapB/3").Set(d, &B{
		S: "10",
		I: 10,
	})

	c.Assert(err, IsNil)
	c.Assert(d.MapB["3"].S, DeepEquals, "10")
	c.Assert(d.MapB["3"].I, DeepEquals, 10)
	c.Assert(orig, Not(Equals), d.B)
}

func (s *MainSuite) TestNestedSliceEntry(c *C) {
	d := makeData()
	orig := d.SliceB[1]

	err := Parse("SliceB/1").Set(d, &B{
		S: "10",
		I: 10,
	})

	c.Assert(err, IsNil)
	c.Assert(d.SliceB[1].S, DeepEquals, "10")
	c.Assert(d.SliceB[1].I, DeepEquals, 10)
	c.Assert(orig, Not(Equals), d.B)
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

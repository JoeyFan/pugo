package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTree(t *testing.T) {
	tree := NewTree("")
	tree.Add("/abc.html", "abc", TreePage, 0)
	tree.Add("/abc/xyz.html", "abc-xyz", TreePage, 1)
	tree.Add("/abc/123.html", "abc-123", TreePost, 2)
	tree.Add("/abc/123/xyz.html", "abc-123-xyz", TreePost, 2)
	tree.Add("/abc/123/456.html", "abc-123-456", TreePost, 2)

	tree.Add("/666/abc.html", "666-abc", TreePage, 1)
	tree.Add("/666/def.html", "666-def", TreePage, 2)
	tree.Add("/666/hij.html", "666-hij", TreePage, 3)
	tree.Add("/666/klmn.html", "666-klmn", TreePage, 4)
	tree.Add("/666/opq.html", "666-opq", TreePage, 5)

	tree.Add("/c9c/rst.html", "c9c-rst", TreePage, 5)
	tree.Add("/c9c/uvw.html", "c9c-uvw", TreePage, 5)
	tree.Add("/c9c/", "c9c", TreePageNode, 0)

	tree.Add("/clc/", "clc", TreePageNode, 0)
	tree.Add("/clc/rst.html", "clc-rst", TreePage, 7)
	tree.Add("/clc/uvw.html", "clc-uvw", TreePage, 9)

	tree.Print("")

	Convey("Tree", t, func() {
		children := tree.Children("abc")
		So(children, ShouldHaveLength, 3)
		So(children[0].Title, ShouldEqual, "abc-xyz")
		So(children[1].Children(), ShouldHaveLength, 0)
		So(children[1].Type, ShouldEqual, TreePost)

		So(children[2].Link, ShouldEqual, "123")
		children2 := children[2].Children("xyz.html")
		So(children2, ShouldHaveLength, 1)

		dirs := tree.Dirs("")
		So(dirs, ShouldHaveLength, 2)
	})
}

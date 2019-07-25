package util_test

import (
	"path/filepath"
	"sort"
	"testing"

	. "gopkg.in/check.v1"
	"github.com/cloudperson/go-billy-ex/memfs"
	"github.com/cloudperson/go-billy-ex/util"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&UtilSuite{})

type UtilSuite struct{}

func (s *UtilSuite) TestCreate(c *C) {
	fs := memfs.New()
	util.WriteFile(fs, "foo/qux", nil, 0644)
	util.WriteFile(fs, "foo/bar", nil, 0644)
	util.WriteFile(fs, "foo/baz/foo", nil, 0644)

	names, err := util.Glob(fs, "*/b*")
	c.Assert(err, IsNil)
	c.Assert(names, HasLen, 2)
	sort.Strings(names)
	c.Assert(names, DeepEquals, []string{
		filepath.Join("foo", "bar"),
		filepath.Join("foo", "baz"),
	})

}

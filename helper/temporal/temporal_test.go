package temporal

import (
	"strings"
	"testing"

	"github.com/cloudperson/go-billy-ex/memfs"
	"github.com/cloudperson/go-billy-ex/test"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&TemporalSuite{})

type TemporalSuite struct {
	test.FilesystemSuite
}

func (s *TemporalSuite) SetUpTest(c *C) {
	fs := New(memfs.New(), "foo")
	s.FilesystemSuite = test.NewFilesystemSuite(fs)
}

func (s *TemporalSuite) TestTempFileDefaultPath(c *C) {
	fs := New(memfs.New(), "foo")
	f, err := fs.TempFile("", "bar")
	c.Assert(err, IsNil)
	c.Assert(f.Close(), IsNil)

	c.Assert(strings.HasPrefix(f.Name(), fs.Join("foo", "bar")), Equals, true)
}

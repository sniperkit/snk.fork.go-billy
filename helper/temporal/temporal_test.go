/*
Sniperkit-Bot
- Date: 2018-08-12 12:10:57.421562389 +0200 CEST m=+0.041271547
- Status: analyzed
*/

package temporal

import (
	"strings"
	"testing"

	. "gopkg.in/check.v1"

	"github.com/sniperkit/snk.fork.go-billy.v4/memfs"
	"github.com/sniperkit/snk.fork.go-billy.v4/test"
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

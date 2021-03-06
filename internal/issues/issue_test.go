package issues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubIssues(t *testing.T) {
	uu := map[string]struct {
		i Issue
		e bool
	}{
		"root":  {New(Root, WarnLevel, "blah"), false},
		"rootf": {Newf(Root, WarnLevel, "blah %s", "blee"), false},
		"sub":   {New("sub1", WarnLevel, "blah"), true},
		"subf":  {Newf("sub1", WarnLevel, "blah %s", "blee"), true},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, u.i.IsSubIssue())
		})
	}
}

func TestBlank(t *testing.T) {
	uu := map[string]struct {
		i Issue
		e bool
	}{
		"blank":    {Issue{}, true},
		"notBlank": {New(Root, WarnLevel, "blah"), false},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, u.i.Blank())
		})
	}
}

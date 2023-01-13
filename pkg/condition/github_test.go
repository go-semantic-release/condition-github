package condition

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubValid(t *testing.T) {
	gha := GitHubActions{}
	os.Setenv("GITHUB_REF", "")
	err := gha.RunCondition(map[string]string{"defaultBranch": ""})
	assert.EqualError(t, err, "this test run is not running on a branch build")
}

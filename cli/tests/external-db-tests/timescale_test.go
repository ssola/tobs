package external_db_tests

import (
	"fmt"
	test_utils "github.com/timescale/tobs/cli/tests/test-utils"
	"testing"
)

func TestTimescale(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TimescaleDB tests")
	}

	fmt.Println("Performing TimescaleDB tests for external db setup...")
	releaseInfo := test_utils.ReleaseInfo{
		Release:   RELEASE_NAME,
		Namespace: NAMESPACE,
	}

	releaseInfo.TestTimescaleGetPassword(t)
	releaseInfo.TestTimescaleChangePassword(t, "battery")
	releaseInfo.VerifyTimescalePassword(t, "battery")

	releaseInfo.TestTimescaleGetPassword(t)
	releaseInfo.TestTimescaleChangePassword(t, "chips")
	releaseInfo.VerifyTimescalePassword(t, "chips")

	releaseInfo.TestTimescaleSuperUserConnect(t, true)
	releaseInfo.TestTimescaleSuperUserConnect(t, false)
	releaseInfo.TestTimescaleSuperUserConnect(t, false)
	releaseInfo.TestTimescaleSuperUserConnect(t, false)
}

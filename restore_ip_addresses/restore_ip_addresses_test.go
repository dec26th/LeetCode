package restore_ip_addresses

import (
	"fmt"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	test := "010010"
	fmt.Println(restoreIpAddresses(test))
}

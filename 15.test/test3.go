package main

import (
	"fmt"
	"regexp"
)

var IllegalCommands = []string{
	"rm",
	"mv",
	"cp",
	"wget",
	"curl",
	"scp",
	"chown",
	"chmod",
	"tar",
	"zip",
	"unzip",
	"find",
	"perl",
	"python",
	"shutdown",
	"reboot",
	"dd",
	"iptables",
	"nc",
	"telnet",
}

func CheckScript(script string) error {
	for _, cmd := range IllegalCommands {
		pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(cmd))
		re := regexp.MustCompile(pattern)
		if re.MatchString(script) {
			return fmt.Errorf("illegal command or keyword found: %s", cmd)
		}
		pattern_sda := `>\s+/dev/sd[abcd]`
		re = regexp.MustCompile(pattern_sda)
		if re.MatchString(script) {
			return fmt.Errorf("illegal command or keyword found: %s", pattern_sda)
		}
	}
	return nil
}

func main() {
	fmt.Println(CheckScript(">     /dev/sda"))
}

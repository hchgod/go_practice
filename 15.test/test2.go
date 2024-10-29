package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// SafeCommandCheck 检查输入字符串是否安全
func SafeCommandCheck(command []string) error {

	commandStr := strings.Join(command, " ")


	// 定义不安全的关键字
	unsafeKeywords := []string{"rm", "> /dev/sda"}

	// 检查是否包含不安全的关键字
	for _, keyword := range unsafeKeywords {
		if strings.Contains(commandStr, keyword) {
			return errors.New("命令包含不安全的部分: " + keyword)
		}
	}

	// 定义十六进制转义序列的正则表达式
	hexPattern := `\\x[0-9a-fA-F]{2}`

	// 编译正则表达式
	re := regexp.MustCompile(hexPattern)

	// 检查是否存在十六进制命令
	if re.MatchString(commandStr) {
		return errors.New("命令包含不安全的十六进制序列")
	}

	// 命令是安全的
	return nil
}

var illegalCommands = []string{
    "rm",
    "mv",
    "cp",
    "wget",
    "curl",
    "ssh",
    "scp",
    "sudo",
    "chown",
    "chmod",
    "tar",
    "zip",
    "unzip",
    "find",
    "grep",
    "sed",
    "awk",
    "perl",
    "python",
    "bash",
    "sh",
    "source",
}

func checkScript(script string) error {
    for _, cmd := range illegalCommands {
        pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(cmd))
        re := regexp.MustCompile(pattern)
        if re.MatchString(script) {
            return fmt.Errorf("illegal command or keyword found: %s", cmd)
        }
    }
    return nil
}

func main() {
	// 测试样例
	testCommands := []string{
		"cp -p /bin/sh /tmp/.beyond; chmod 4755 /tmp/.beyond;",
		"rm -rf /; echo hello;",
		"char esp[] attribute ((section(\".text\"))) = \"\\xeb\\x3e\"",
		"ls -l",
		"echo 'Hello World'",
		"cat > /dev/sda",
		"gcc -o myprog myprog.c; ./myprog",
	}

	for _, cmd := range testCommands {
		if err := checkScript(cmd); err != nil {
			fmt.Printf("不安全的命令: %s, 错误: %s\n", cmd, err)
		} else {
			fmt.Printf("安全的命令: %s\n", cmd)
		}
	}
}

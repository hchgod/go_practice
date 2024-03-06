package main

import "fmt"

// "fmt"

func reverseWords(s string) (res string) {
	left,right := len(s)-1,len(s)-1
	for fast:=len(s)-1; fast>=0; fast-- {
		if fast >0{
			if s[fast] != ' ' && s[fast-1] == ' '{
				left = fast		
			}else if s[fast] == ' ' && s[fast-1] != ' ' {
				right = fast
			}else if s[fast] == ' ' && s[fast-1] == ' '{
				continue
			}
			if s[left] != ' ' && s[right-1] != ' ' && right>=left {
				res = res + s[left:right] + " "
			}	
		}else{
			if s[fast] != ' '{
				left = fast
				res = res + s[left:right] + " "
			}
		}	
	}
	fmt.Println(res)
	return res[:len(res)-1]
}


func main() {
	s := " a good   example  "
	reverseWords(s)	
}

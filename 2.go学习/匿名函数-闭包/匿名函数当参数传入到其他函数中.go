package main
import "fmt"
func main() {
 // 项目经理的一天
 work(func() {
  fmt.Println("组织部门开会")
  fmt.Println("给部门员工分配今天工作任务")
  fmt.Println("检查部门员工昨天提交的代码")
  fmt.Println("... ...")
 })
 // 程序员的一天
 work(func() {               //匿名函数当参数传入到其他函数中
  fmt.Println("参加部门会议")
  fmt.Println("修改测试提交的BUG")
  fmt.Println("完成老大今天安排的任务")
  fmt.Println("... ...")
 })
}
// 假设我们需要编写一个函数,用于描述一个人每天上班都需要干嘛
// 职场中的人每天上班前,上班后要做的事几乎都是相同的, 但是每天上班过程中要做的事确实不确定的
// 所以此时我们可以使用匿名函数来解决, 让上班的人自己觉得自己每天上班需要干什么
func work(custom func())  {
 // 上班前
 fmt.Println("起床")
 fmt.Println("刷牙")
 fmt.Println("洗脸")
 fmt.Println("出门")
 fmt.Println("上班打卡")
 fmt.Println("开电脑")

 // 上班中
 custom()

 // 上班后
 fmt.Println("关电脑")
 fmt.Println("下班打卡")
 fmt.Println("出门")
 fmt.Println("到家")
 fmt.Println("吃饭")
 fmt.Println("睡觉")

}
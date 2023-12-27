- 这是关于taurus_go的代码示例

- 如果遇到错误`This file is within module "taurus_go_demo", which is not included in your workspace.
To fix this problem, you can add this module to your go.work file`
请在go.work添加taurus_go_demo
``` go.work
use /path/to/taurus_go_demo
```
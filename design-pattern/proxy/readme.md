代理模式：当我们不能直接的去做一些事情的时候需要通过第三方进行代理，以完成某些事情
代码中的场景：
1. 对于不能直接访问的对象
2. 不想直接访问的对象
3. 访问存在困难的对象

优点：职责清晰； 高扩展性；智能化
缺点：
	1. 由于交互双方之间多了一层代理，因此会代理性能上的问题
	2. 由于代理模式需要额外的工作因此可能会使得实现变得复杂
常见的用法：
- 远程代理
- 虚拟代理
- 安全代理
- 智能指针
- 同步代理

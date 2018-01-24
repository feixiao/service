## shipping

shipping演示了一个更真实的由多个服务组成的应用程序。 程序详细情况参考书籍[《领域驱动设计》](https://book.douban.com/subject/26819666/)。


### 结构

shipping由booking、handling和tracking三个服务组成。每个服务都是跟之前的例子一样独立的Gokit服务。

+ booking 

  **航运公司**用来预订和规划货船。

+ handling

   **工作人员**记录货船什么时候到达、装货等等。

+ tracking

   **航运公司顾客**使用沿着路线跟踪货物。

有一些纯粹的领域包(domain pakcages),它们包含了一些复杂的业务逻辑。它们提供领域对象和领域服务，

这些领域对象和服务为每个应用程序服务使用从而为用户提供有趣的用例。

+ inmem 包括了 in-memory实现, 可以在领域包中找到。
+ route 包提供了领域服务，用于查询外部应用程序以获得可能的路由。

### 包说明

+ location	模拟位置信息

+ voyage   模拟航线信息

  ​

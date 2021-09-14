## solid 原则

### 单一职责

### 开闭原则

### 里斯替换原则

### 接口隔离

### 依赖反转
#### 不使用注入的问题
* 业务(也就是领域模型)不应该依赖具体的基础设施(db, redis, mq等)
* 领域依赖基础设施,那么就是业务依赖具体技术细节; 技术细节的改变对业务逻辑产生影响,使其不得不改变,这样不合理
* 复用困难, 核心领域的复用价值很高,如果依赖基础设施,复用困难

#### 使用依赖反转
**接口定义一般放在消费者,不放在实现接口的包**
* 一般来说某个包实现了某个功能,但是还未抽象出接口,也就是说这个具体实现包并不知道自己将会被怎么使用,实现了某个接口
* 消费者(一般是业务层), 需要调用具体实现包来完成业务, 这时候为了抽象出底层,可以抽象出接口,该接口应该属于消费者(因为消费者知道自己需要什么)
* 为了解耦消费者与底层具体实现包的强依赖,抽象接口解耦: 消费者依赖抽象接口,不依赖具体实现包, 由外部注入到消费者
* 接口注入能更好的mock test
* 如果接口定义在实现包,那么又成了预先定义接口的方式, 不满足go 隐式实现接口的设计
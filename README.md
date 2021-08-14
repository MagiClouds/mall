# mall
golang-micro kratos go-zero


商城服务流程

用户登录 -》 商品分类、商品列表 -》 商品详情 -》 加入购物车 -》 下单 -》 支付 -》 已购买订单


下单（限流（令牌 漏斗 分布式限流）、秒杀、分布式锁， 分布式id）：
支付（分布式事物）： 扣钱，清空购物车，更改订单状态，增加用户商品

商品搜索（mqsql-dump + cdc + es）

logging + tracing + metrics



服务拆分:

业务层(bff)

商城 shop
后管 admin


应用层

用户服务        user
商品服务        product

订单服务        order
购物车服务      cart
支付服务        payment

## GOQUERY学习
阅读 https://www.flysnow.org/2018/01/20/golang-goquery-examples-selector.html

# 考核
结合http包学习编写https://cn.idgcapital.com/news  爬虫  爬取所有新闻详情页

### 思路
Pipeline
1. 爬取全量分页  
2. 协议分页获得 当前LIST下所有详情页面URL
3. 爬取详情页URL

可以采用Pipeline方式  分为三层 通过channel链接


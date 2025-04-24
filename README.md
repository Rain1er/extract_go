# extract_go is a tool to extract backup urls from recon files.

## 目标
`httpx -t 150 -l output.txt -cl -oa output.txt`

```shell
https://zzzx.suse.edu.cn/1.rar [66]
https://zzzx.suse.edu.cn/2021.zip [66]
https://zzzx.suse.edu.cn/2022.zip [66]
https://zzzx.suse.edu.cn/2012.rar [66]
```

对如上标准输出进行数据处理，对于输出3次以上`content_length`相同的url进行去重，并筛选掉1MB以下的文件，打印余下的结果


## 设计思路

1. 从output.txt中读入所有行
2. 扫描所有行，放到urlMap中（url，content_length），对`content_length`进行统计，放到countMap中（content_length，count）
3. 遍历countMap,从countMap中取出`count`小于3的`content_length`
4. 遍历urlMap,从urlMap中取出`content_length`在countMap中存在的url
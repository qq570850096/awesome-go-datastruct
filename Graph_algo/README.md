# golang图论算法
本项目的初衷是通过golang实现图论常用算法，熟悉golang语言的一个开源项目。


[TOC]

## 图的表示方法

在Adj文件夹中保存了三种图的表示方法

1. Matrix：邻接矩阵
2. Table：邻接表
3. Hash：也是邻接表，不过使用了哈希表的存储方式加快速度，本项目中后续算法实现都是基于此结构。

## 文件结构
* Adj(这里存放了关于图的三种数据结构)
* BFS 存放的是图的广度优先搜索算法
* DFS 存放了图的深度优先算法，带连通分量(
Connected Component)统计的版本
* search 这里存放的是二分图查找，图中是否包含环，图中的单源路径查找问题.

g.txt 一个普通的图

g2.txt 一个包含一个环的图

g2_noCycle.txt 在g2基础上删除了环

notBip.txt 一个非二分图

package block7game

// BlockNums - 3
const BlockNums = 3

// DefaultMaxBlockLevel - 7
const DefaultMaxBlockLevel = 7

// DefaultMaxBlockNums - 7
const DefaultMaxBlockNums = 7

// DefaultVersion - 5
const DefaultVersion = 5

const (
	GameStateRunning = 1

	GameStateSucess = 2

	GameStateFail = 3
)

// FuncHasBlock - has block
type FuncHasBlock func(x, y, z int) bool

// FuncHasBlockEx - has block
type FuncHasBlockEx func(x, y, z int, w, h int) bool

// FuncIsLess - is less
type FuncIsLess func(i, j int) bool

// FuncCountBlockValue - count value
type FuncCountBlockValue func(x, y, z int) int

// 001 冰块 覆盖层1 需要点击周围的一个其他的块，冰块才能消失
// 002 蛋糕 特殊方块 场景中增加一个蛋糕和三个铲子 点击铲子则铲子消失 蛋糕减少一块 点击三个铲子 蛋糕消失
// 003 问号 覆盖层0 点亮的问号 同时只有一个问号方块可以看到其中的内容
// 004 窗帘 覆盖层2 每次其他方块点击 点亮的开关改变一次状态 关闭状态下 方块不可点击
// 005 瓢虫 覆盖层3 初始生命值3 点击周围的方块生命减少一点 然后随机飞到其他方块上
// 006 杂草 覆盖层1 点击周围块消除杂草 如果一次点击没有消除杂草 则往周围生长 最后一个覆盖块可以点击消除杂草
// 007 冰淇淋 覆盖层2+特殊方块 遮挡一个3*2的区域 运用三个勺子可以清除整个覆盖区域（相当于更大的蛋糕）
// 008 炸弹 特殊方块 点击后自身消失 同时消除场景中的其他三个相同的方块
// 009 炸弹1 特殊方块 点击后自身消失 同时消除场景中的其他六个相同的方块
// 010 炸弹2 特殊方块 点击后自身消失 同时消除场景中的其他九个相同的方块
// 011 炸弹3 特殊方块 点击后自身消失 同时消除场景中的其他十二个相同的方块
// 012 传送门 特殊方块 点击后自身消失 把其他位置的一个方块传送到当前位置
// 013 彩虹 特殊方块 点击后自身消失 改变三个相同方块为另外一种图案
// 014 冰块1 覆盖层2 生命值为2的冰块
// 015 冰块2 覆盖层2 生命值为3的冰块
// 016 冰块3 覆盖层2 生命值为4的冰块
// 017 逆转 覆盖层1 当方块没有被点亮的时候不可见
// 其他：
// 锁 覆盖层2+特殊方块 锁的数据记录在地图中 最多有5中不同的锁 每种锁对应三把钥匙 三把钥匙放到消除栏消除后 打开所有同颜色的锁
// 扑克牌方块 暂时没用到

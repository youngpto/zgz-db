package model

// PropertyOffset 玩家存储的属性偏移量
type PropertyOffset struct {
	Id              int64 `xorm:"pk autoincr"`        // 主键
	TotalOffset     int   `xorm:"'total_offset'"`     // 总偏移量
	LifeOffset      int   `xorm:"'life_offset'"`      // 生命偏移
	ReasonOffset    int   `xorm:"'reason_offset'"`    // 理智偏移
	PowerOffset     int   `xorm:"'power_offset'"`     // 力量偏移
	AgileOffset     int   `xorm:"'agile_offset'"`     // 敏捷偏移
	KnowledgeOffset int   `xorm:"'knowledge_offset'"` // 知识偏移
	WillOffset      int   `xorm:"'will_offset'"`      // 意志偏移
}

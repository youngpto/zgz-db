package model

import "github.com/youngpto/zgz-db/enum"

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

func ConvEnum2Str(property enum.HeroProperty) string {
	switch property {
	case enum.BaseHp:
		return "life_offset"
	case enum.BaseSan:
		return "reason_offset"
	case enum.BaseStrength:
		return "power_offset"
	case enum.BaseAgility:
		return "agile_offset"
	case enum.BaseKnowledge:
		return "knowledge_offset"
	case enum.BaseWillPower:
		return "will_offset"
	}
	return ""
}

// IncrOrDecrByType 计算某项属性自增或自减后的值
func (p *PropertyOffset) IncrOrDecrByType(property enum.HeroProperty, isIncr bool) (int, bool) {
	if isIncr {
		if p.LifeOffset+p.ReasonOffset+p.PowerOffset+p.AgileOffset+p.KnowledgeOffset+p.WillOffset >= p.TotalOffset {
			return 0, false
		}
	}

	changeLayer := 0
	if isIncr {
		changeLayer++
	} else {
		changeLayer--
	}

	var number int
	switch property {
	case enum.BaseHp:
		number = p.LifeOffset + changeLayer
	case enum.BaseSan:
		number = p.ReasonOffset + changeLayer
	case enum.BaseStrength:
		number = p.PowerOffset + changeLayer
	case enum.BaseAgility:
		number = p.AgileOffset + changeLayer
	case enum.BaseKnowledge:
		number = p.KnowledgeOffset + changeLayer
	case enum.BaseWillPower:
		number = p.WillOffset + changeLayer
	}
	if number < 0 {
		return 0, false
	}
	return number, true
}

// ResetPropertyByType 设置某项属性最新值
func (p *PropertyOffset) ResetPropertyByType(property enum.HeroProperty, newNumber int) {
	switch property {
	case enum.BaseHp:
		p.LifeOffset = newNumber
	case enum.BaseSan:
		p.ReasonOffset = newNumber
	case enum.BaseStrength:
		p.PowerOffset = newNumber
	case enum.BaseAgility:
		p.AgileOffset = newNumber
	case enum.BaseKnowledge:
		p.KnowledgeOffset = newNumber
	case enum.BaseWillPower:
		p.WillOffset = newNumber
	}
}

package dto

// GainExperienceRes 玩家获取经验后的返回结果
type GainExperienceRes struct {
	OriginLevel   int     // 原始等级
	OriginExp     float64 // 原始经验
	OriginMaxExp  float64 // 原始最高经验值
	CurrentLevel  int     // 当前等级
	CurrentExp    float64 // 当前经验值
	CurrentMaxExp float64 // 当前等级最高经验值
	AddExpNumber  float64 // 增加的经验值
}

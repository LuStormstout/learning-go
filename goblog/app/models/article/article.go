package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

// Article 文章模型
type Article struct {
	ID    uint64
	Title string
	Body  string
}

// Get 通过 ID 获取文章
func Get(idStr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idStr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

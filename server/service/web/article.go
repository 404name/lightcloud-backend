package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
)

type ArticleService struct {
}

// CreateArticle 创建Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article web.Article) (err error) {
	err = global.GVA_DB.Create(&article).Error
	return err
}

// DeleteArticle 删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticle(article web.Article) (err error) {
	err = global.GVA_DB.Delete(&article).Error
	return err
}

// DeleteArticleByIds 批量删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.Article{},"id in ?",ids.Ids).Error
	return err
}

// UpdateArticle 更新Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)UpdateArticle(article web.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据id获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticle(id uint) (article web.Article, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

// GetArticleInfoList 分页获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticleInfoList(info webReq.ArticleSearch) (list []web.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&web.Article{})
    var articles []web.Article
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CategoryId != nil {
        db = db.Where("category_id = ?",info.CategoryId)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Uid != nil {
        db = db.Where("uid = ?",info.Uid)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&articles).Error
	return  articles, total, err
}

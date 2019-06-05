package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 定义tag表的模型和模型操作方法

type Tag struct {
	Model	// 引入公共部分

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

// 根据分页数目 每页多少 及其他 获取tags  返回的是数组
func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
	// gorm语法   这个查手册或者莫放这里也可以
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	// 这里会返回 : return tags
	return
}

// 获取tag的总数
func GetTagTotal(maps interface {}) (count int){
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// 更具名字获取tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// 添加tag
func AddTag(name string, state int, createdBy string) bool{
	db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface {}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

// 自动执行触发器 =======================================================================
/*
gorm所支持的回调方法：
创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
删除：BeforeDelete、AfterDelete
查询：AfterFind
*/
// 添加前自动触发
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// 修改前自动触发
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

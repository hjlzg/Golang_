package models

type Tag struct{
	Model

	Name string `json:"name"`
	CreateBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int,pageSize int,mpas interface{})(tags []Tag){
	//db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

package core

const (
	// DirTypeProductFile 产品文件
	DirTypeProductFile DirType = "product"
	// DirTypeProductRelease 产品发布文件
	DirTypeProductRelease DirType = "product-release"
	// DirTypePublicDisk 公共云盘
	DirTypePublicDisk DirType = "public"
	// DirTypePrivateDisk 私有云盘
	DirTypePrivateDisk DirType = "private"
	// DirTypeFileResource 普通文件
	DirTypeFileResource DirType = "resource"
	// DirTypeSystemFile 系统文件文件
	DirTypeSystemFile DirType = "system"
	// DirTypeCourse 课程资源
	DirTypeCourse DirType = "course"
	// DirTypeCourseTimeRecord 课程资源
	DirTypeCourseTimeRecord DirType = "record"
	// DirTypeOrgRelease 版本发布文件
	DirTypeOrgRelease DirType = "org-release"
)

// DirType 文件类型
type DirType string

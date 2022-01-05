package cloudfile

// 文件类型
const (
	// TypeProductFile 产品文件
	TypeProductFile Type = "product"
	// TypeProductRelease 产品发布文件
	TypeProductRelease Type = "product-release"
	// TypePublicDisk 公共云盘
	TypePublicDisk Type = "public"
	// TypePrivateDisk 私有云盘
	TypePrivateDisk Type = "private"
	// TypeFileResource 普通文件
	TypeFileResource Type = "resource"
	// TypeSystemFile 系统文件文件
	TypeSystemFile Type = "system"
	// TypeCourse 课程资源
	TypeCourse Type = "course"
	// TypeCourseTimeRecord 课程录课资源
	TypeCourseTimeRecord Type = "record"
	// TypeOrgRelease 版本发布文件
	TypeOrgRelease Type = "org-release"
)

type Type string

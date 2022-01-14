package cloudfile

// 文件类型
const (
	// TypeProductFile 产品文件
	TypeProductFile DirType = "product"
	// TypeProductRelease 产品发布文件
	TypeProductRelease DirType = "product-release"
	// TypePublicDisk 公共云盘
	TypePublicDisk DirType = "public"
	// TypePrivateDisk 私有云盘
	TypePrivateDisk DirType = "private"
	// TypeFileResource 普通文件
	TypeFileResource DirType = "resource"
	// TypeSystemFile 系统文件文件
	TypeSystemFile DirType = "system"
	// TypeCourse 课程资源
	TypeCourse DirType = "course"
	// TypeCourseTimeRecord 课程录课资源
	TypeCourseTimeRecord DirType = "record"
	// TypeOrgRelease 版本发布文件
	TypeOrgRelease DirType = "org-release"
)

type DirType string

package lesson

import (
	`fmt`
)

func (u *User) LessonCacheKey(lessonId, yunkeId int64) string {
	return fmt.Sprintf("%d:%d:%d", yunkeId, lessonId, u.Id)
}

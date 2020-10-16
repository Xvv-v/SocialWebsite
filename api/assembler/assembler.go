package assembler

import (
	"SocialWebsite/api/model"
	"SocialWebsite/domain/user/models"
)

//将 interfaceModel ---> domainModel

//VO（或DOT） ---> DO

//UserToDO 转为DO对象
func UserToDO(r *model.UserVO) *models.User {

	return &models.User{}
}

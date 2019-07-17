package services
import	"../domains"
import "../utils"


func	GetUserFromApi(userID int64)(*domains.User,*utils.ApiError){
	user:= &domains.User{
		ID:userID,
	}
	if err:=user.Get(); err!= nil {
		return nil, err
	}
	return user,nil
}

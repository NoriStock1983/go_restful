package controllers

import (
	"GOTEST/cmd/models"
	"GOTEST/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

/***********************************************************************/
/* 関数名：Allusermst
/* 引数　：
/* 戻り値：[]models.AllUsemst（JSON形式）
/* 内容　：Usermstテーブルに格納されているデータを全て取得し、JSON形式で返す。
/* 作成者：Norifumi Sato
/* 作成日：2023/5/17
/* 備考　：
/***********************************************************************/
func Allusermst(c *gin.Context) {
	usermst := []models.AllUsemst{
		{Usercd: "a0006802", User_f_name: "Noriufmi", User_l_name: "Sato", Belonged_company_cd: "0001", Belonged_dept_cd: "32P2", Authcd: "01", Activeflg: true, Createdby: "a0006802", Created_date: "2022-04-01", Updatedby: "a0006802", Updated_date: "2022-04-01", UpdateCounter: 1},
		{Usercd: "a0006801", User_f_name: "Noriufmi", User_l_name: "Sato", Belonged_company_cd: "0001", Belonged_dept_cd: "32P2", Authcd: "01", Activeflg: true, Createdby: "a0006802", Created_date: "2022-04-01", Updatedby: "a0006802", Updated_date: "2022-04-01", UpdateCounter: 1},
	}
	c.JSON(http.StatusOK, usermst)
}

/***********************************************************************/
/* 関数名：GetByUsercd
/* 引数　：
/* 戻り値：models.Usemst（JSON形式）
/* 内容　：引数で指定されたusercdに紐づくユーザ情報をusermstテーブルから取得し、
/* 　　　　JSON形式で返す。
/* 作成者：Norifumi Sato
/* 作成日：2023/5/17
/* 備考　：
/***********************************************************************/
func GetByUsercd(c *gin.Context) {

	usercd := c.Param("usercd")

	countbyuser := services.CountByUsercd(usercd)

	if countbyuser == 0 {

	} else {
		usermst := models.AllUsemst{Usercd: "a0006802", User_f_name: "Noriufmi", User_l_name: "Sato", Belonged_company_cd: "0001", Belonged_dept_cd: "32P2", Authcd: "01", Activeflg: true, Createdby: "a0006802", Created_date: "2022-04-01", Updatedby: "a0006802", Updated_date: "2022-04-01", UpdateCounter: 1}
		c.JSON(http.StatusOK, usermst)
	}
}

/***********************************************************************/
/* 関数名：UpdateByUsercd
/* 引数　：models.AllUsemst（JSON形式）
/* 戻り値：
/* 内容　：引数で指定されたusercdに紐づくユーザ情報をusermstテーブルから取得し、
/* 　　　　JSON形式で返す。
/* 作成者：Norifumi Sato
/* 作成日：2023/5/17
/* 備考　：
/***********************************************************************/
func UpdateByUsercd(c *gin.Context) {

}

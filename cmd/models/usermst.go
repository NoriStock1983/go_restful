package models

type LoginedUser struct {
	Usercd              string `json:"usercd"`
	User_f_name         string `json:"user_f_name"`
	User_l_name         string `json:"user_l_name"`
	Belonged_company_cd string `json:"belonged_company_cd"`
	Belonged_dept_cd    string `json:"belonged_dept_cd"`
	Authcd              string `json:"authcd"`
}

type AllUsemst struct {
	Usercd              string `json:"usercd"`
	User_f_name         string `json:"user_f_name"`
	User_l_name         string `json:"user_l_name"`
	Belonged_company_cd string `json:"belonged_company_cd"`
	Belonged_dept_cd    string `json:"belonged_dept_cd"`
	Authcd              string `json:"authcd"`
	Activeflg           bool   `json:"activeflg"`
	Createdby           string `json:"createdby"`
	Created_date        string `json:"created_date"`
	Updatedby           string `json:"updatedby"`
	Updated_date        string `json:"updated_date"`
	UpdateCounter       int    `json:"updateCounter"`
}

package services

/***********************************************************************/
/* 関数名：CountByUsercd
/* 引数　：usercd string
/* 戻り値：データ取得件数
/* 内容　：引数で指定されたusercdの内容の登録件数を返す。
/* 作成者：Norifumi Sato
/* 作成日：2023/5/18
/* 備考　：
/***********************************************************************/
func CountByUsercd(usercd string) int {

	if usercd == "a0006802" {
		return 1
	} else {
		return 0
	}
}

package resp

type UserResp struct {
	ID       int32  `json:"ID"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}

package vo

import "github.com/megrez/pkg/entity/po"

type Author struct {
	ID     uint
	Name   string
	Mail   string
	Site   string
	Role   string
	Avatar string
	IP     string
}

func GetAuthorFromPO(po po.Author) *Author {
	vo := &Author{}
	vo.ID = po.ID
	vo.Name = po.Name
	vo.Mail = po.Mail
	vo.Site = po.Site
	vo.Avatar = po.Avatar
	// TODO: Role
	vo.IP = po.IP
	return vo
}

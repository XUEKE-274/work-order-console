package roleEnum

import "work-order-console/domain/enum"

const (

	//SYSTEM    系统管理员 （不能被删除， 所有权限）
	SYSTEM     enum.RoleEnum  =  "SYSTEM"

	//ADMIN     超级管理员  （所有权限）
	ADMIN      enum.RoleEnum  =  "ADMIN"

	//WriterOrder    写操作 + 读操作  工单
	WriterOrder     enum.RoleEnum  =  "WRITER_ORDER"

	//ReaderOrder    读操作       工单
	ReaderOrder     enum.RoleEnum  =  "READER_ORDER"


	//WriterKf     写操作 + 读操作 客服
	WriterKf     enum.RoleEnum  =  "WRITER_KF"

	//ReaderKf     读操作  客服
	ReaderKf     enum.RoleEnum  =  "READER_KF"


	//WriterQt    写操作 + 读操作
	WriterQt     enum.RoleEnum  =  "WRITER_QT"

	//ReaderQt    读操作
	ReaderQt    enum.RoleEnum  =  "READER_QT"


)

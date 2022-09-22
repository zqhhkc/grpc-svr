package pb

import (
	context "context"
	"log"
)

type Server struct {
}

func (svr *Server) InsertAccount(ctx context.Context, req *AccountInsertReq) (*AccountInsertRes, error) {

	res := AccountInsertRes{}
	res.Status = IMCoreStatus_StatusSucc
	log.Println("InsertAccount")
	return &res, nil
}

func (svr *Server) UpdateAccount(ctx context.Context, req *AccountUpdateReq) (*AccountUpdateRes, error) {
	res := AccountUpdateRes{}
	res.Status = IMCoreStatus_AccountNotFound
	log.Println("UpdateAccount")
	return &res, nil
}

func (svr *Server) DeleteAccount(ctx context.Context, req *AccountDeleteReq) (*AccountDeleteRes, error) {
	res := AccountDeleteRes{}
	res.Status = IMCoreStatus_StatusSucc
	log.Println("DeleteAccount")
	return &res, nil
}

func (svr *Server) QueryAccount(ctx context.Context, req *AccountQueryReq) (*AccountQueryRes, error) {
	res := AccountQueryRes{}
	res.Status = IMCoreStatus_StatusSucc
	log.Println("QueryAccount")
	return &res, nil
}

func (svr *Server) ListAllAccount(ctx context.Context, req *ListAllAccountReq) (*ListAllAccountRes, error) {
	res := ListAllAccountRes{}
	res.Status = IMCoreStatus_StatusSucc
	log.Println("ListAllAccount")
	return &res, nil
}

func (svr *Server) mustEmbedUnimplementedIMMainServiceServer() {

}

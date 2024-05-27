package user

//func TestLogin(t *testing.T) {
//	ctx := context.Background()
//	clientConn, _ := grpc.DialContext(ctx, "localhost:8082", grpc.WithInsecure())
//	defer clientConn.Close()
//
//	userClient := user.NewUserClient(clientConn)
//	resp, err := userClient.Login(ctx, &user.LoginReq{
//		Type:     0,
//		Account:  "admin",
//		Password: "adminadmin",
//		Captcha:  "zxcf",
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(resp)
//}
//
//func TestRegister(t *testing.T) {
//	ctx := context.Background()
//	clientConn, _ := grpc.DialContext(ctx, "localhost:8082", grpc.WithInsecure())
//	defer clientConn.Close()
//
//	userClient := user.NewUserClient(clientConn)
//	resp, err := userClient.Register(ctx, &user.RegisterReq{
//		Account:  "admin",
//		Password: "adminadmin",
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(resp)
//}

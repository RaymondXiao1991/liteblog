package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetAbout",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetComment",
			Router: `/comment/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetDetail",
			Router: `/details/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetMessage",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetReg",
			Router: `/reg`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:IndexController"] = append(beego.GlobalControllerRouter["controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:MessageController"] = append(beego.GlobalControllerRouter["controllers:MessageController"],
		beego.ControllerComments{
			Method: "Count",
			Router: `/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:MessageController"] = append(beego.GlobalControllerRouter["controllers:MessageController"],
		beego.ControllerComments{
			Method: "NewMessage",
			Router: `/new/?:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:MessageController"] = append(beego.GlobalControllerRouter["controllers:MessageController"],
		beego.ControllerComments{
			Method: "Query",
			Router: `/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:NoteController"] = append(beego.GlobalControllerRouter["controllers:NoteController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:NoteController"] = append(beego.GlobalControllerRouter["controllers:NoteController"],
		beego.ControllerComments{
			Method: "EditPage",
			Router: `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:NoteController"] = append(beego.GlobalControllerRouter["controllers:NoteController"],
		beego.ControllerComments{
			Method: "NewPage",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:NoteController"] = append(beego.GlobalControllerRouter["controllers:NoteController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/save/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:PraiseController"] = append(beego.GlobalControllerRouter["controllers:PraiseController"],
		beego.ControllerComments{
			Method: "Parse",
			Router: `/:type/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:UploadController"] = append(beego.GlobalControllerRouter["controllers:UploadController"],
		beego.ControllerComments{
			Method: "UploadFile",
			Router: `/uploadfile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:UploadController"] = append(beego.GlobalControllerRouter["controllers:UploadController"],
		beego.ControllerComments{
			Method: "UploadImg",
			Router: `/uploadimg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:UserController"] = append(beego.GlobalControllerRouter["controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:UserController"] = append(beego.GlobalControllerRouter["controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["controllers:UserController"] = append(beego.GlobalControllerRouter["controllers:UserController"],
		beego.ControllerComments{
			Method: "Reg",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}

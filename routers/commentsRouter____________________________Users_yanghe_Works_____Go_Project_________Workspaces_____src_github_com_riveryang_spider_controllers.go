package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/riveryang/spider/controllers:CodecController"] = append(beego.GlobalControllerRouter["github.com/riveryang/spider/controllers:CodecController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:codec/:page`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/riveryang/spider/controllers:CodecController"] = append(beego.GlobalControllerRouter["github.com/riveryang/spider/controllers:CodecController"],
		beego.ControllerComments{
			Method: "Auto",
			Router: `/:codec/auto/:page`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/riveryang/spider/controllers:CodecController"] = append(beego.GlobalControllerRouter["github.com/riveryang/spider/controllers:CodecController"],
		beego.ControllerComments{
			Method: "Stop",
			Router: `/:codec/stop`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}

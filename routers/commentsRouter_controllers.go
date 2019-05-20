package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["nim_163/controllers:FriendController"] = append(beego.GlobalControllerRouter["nim_163/controllers:FriendController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:FriendController"] = append(beego.GlobalControllerRouter["nim_163/controllers:FriendController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:FriendController"] = append(beego.GlobalControllerRouter["nim_163/controllers:FriendController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/get`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:FriendController"] = append(beego.GlobalControllerRouter["nim_163/controllers:FriendController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:MsgController"] = append(beego.GlobalControllerRouter["nim_163/controllers:MsgController"],
        beego.ControllerComments{
            Method: "BroadcastMsg",
            Router: `/broadcastMsg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:MsgController"] = append(beego.GlobalControllerRouter["nim_163/controllers:MsgController"],
        beego.ControllerComments{
            Method: "Recall",
            Router: `/recall`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:MsgController"] = append(beego.GlobalControllerRouter["nim_163/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendAttachMsg",
            Router: `/sendBatchAttachMsg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:MsgController"] = append(beego.GlobalControllerRouter["nim_163/controllers:MsgController"],
        beego.ControllerComments{
            Method: "SendMsg",
            Router: `/sendMsg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:MsgController"] = append(beego.GlobalControllerRouter["nim_163/controllers:MsgController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "ListBlackAndMuteList",
            Router: `/listBlackAndMuteList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "Mute",
            Router: `/mute`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "MuteAv",
            Router: `/muteAv`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "SetDesktopPush",
            Router: `/setDonnop`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "SetSpacialRelation",
            Router: `/setSpecialRelation`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["nim_163/controllers:UserController"] = append(beego.GlobalControllerRouter["nim_163/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

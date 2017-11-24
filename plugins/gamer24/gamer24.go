////////////////////////////////////////////////////////////////////////////
// Package: gamer24
// Purpose: Play the 24 game (http://rosettacode.org/wiki/24_game) in wechat
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gamer24 // 以插件名命令包名

import (
	"bytes"
	"strings"

	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb" // 导入协议包

	"github.com/suntong/game24"
)

var (
	g24 *game24.Game
	buf bytes.Buffer
)

// Register plugin
// 必须有的插件注册函数
// 指定session, 可以对不同用户注册不同插件
func Register(session *wxweb.Session) {
	// 将插件注册到session
	// 第一个参数: 指定消息类型, 所有该类型的消息都会被转发到此插件
	// 第二个参数: 指定消息处理函数, 消息会进入此函数
	// 第三个参数: 自定义插件名，不能重名，switcher插件会用到此名称
	session.HandlerRegister.Add(wxweb.MSG_TEXT, wxweb.Handler(gamer24), "gamer24")

	if err := session.HandlerRegister.EnableByName("gamer24"); err != nil {
		logs.Error(err)
	}

	g24 = game24.NewGame(30, &buf)
}

// 消息处理函数
func gamer24(session *wxweb.Session, msg *wxweb.ReceivedMessage) {

	// 取收到的内容
	if strings.Index(msg.Content, "3824") == 0 {
		logs.Info("command '3824' received")
		g24.Play()
		logs.Info("game generated: " + buf.String())
		session.SendText(buf.String(),
			session.Bot.UserName, wxweb.RealTargetUserName(session, msg))
		logs.Info("wechat message sent")
		buf.Reset() // resets to empty for future writes
		logs.Info("gamer24 exit gracefully")
	}

	// debug
	//who := session.Cm.GetContactByUserName(msg.FromUserName)
	//logs.Debug("who send", who)
	//if msg.IsGroup {
	//	mm, err := wxweb.CreateMemberManagerFromGroupContact(session, who)
	//	if err != nil {
	//		logs.Error(err)
	//		return
	//	}
	//	info := mm.GetContactByUserName(msg.Who)
	//	logs.Debug(info)
	//}
}

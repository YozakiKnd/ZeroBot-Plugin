/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
为 Golang 移植版，语料、素材均来自上述项目
本项目遵守 AGPL v3 协议进行开源
*/
package atri

import (
	"encoding/base64"
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
)

type datagetter func(string, bool) ([]byte, error)

func (dgtr datagetter) randImage(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.ImageBytes(data)
}

func (dgtr datagetter) randRecord(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.Record("base64://" + base64.StdEncoding.EncodeToString(data))
}

func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}

// isAtriSleeping 凌晨0点到6点，ATRI 在睡觉，不回应任何请求
func isAtriSleeping(ctx *zero.Ctx) bool {
	if now := time.Now().Hour(); now >= 1 && now < 6 {
		return false
	}
	return true
}

func init() { // 插件主体
	engine := control.Register("atri", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "atri人格文本回复",
		Help: "本插件基于 ATRI ，为 Golang 移植版\n" +
			"- ATRI醒醒\n- ATRI睡吧\n- 萝卜子\n- 喜欢 | 爱你 | 爱 | suki | daisuki | すき | 好き | 贴贴 | 老婆 | 亲一个 | mua\n" +
			"- 草你妈 | 操你妈 | 脑瘫 | 废柴 | fw | 废物 | 战斗 | 爬 | 爪巴 | sb | SB | 傻B\n- 早安 | 早哇 | 早上好 | ohayo | 哦哈哟 | お早う | 早好 | 早 | 早早早\n" +
			"- 中午好 | 午安 | 午好\n- 晚安 | oyasuminasai | おやすみなさい | 晚好 | 晚上好\n- 高性能 | 太棒了 | すごい | sugoi | 斯国一 | よかった\n" +
			"- 没事 | 没关系 | 大丈夫 | 还好 | 不要紧 | 没出大问题 | 没伤到哪\n- 好吗 | 是吗 | 行不行 | 能不能 | 可不可以\n- 啊这\n- 我好了\n- ？ | ? | ¿\n" +
			"- 离谱\n- 答应我",
		PublicDataFolder: "Atri",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("嗯呜呜……夏生先生……？"))
		},
		OnDisable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("Zzz……Zzz……"))
		},
	})
	engine.UsePreHandler(isAtriSleeping)
	var dgtr datagetter = engine.GetLazyData
	engine.OnFullMatch("萝卜子").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(2) {
			case 0:
				ctx.SendChain(randText("萝卜子是对机器人的蔑称！", "是亚托莉......萝卜子可是对机器人的蔑称"))
			case 1:
				ctx.SendChain(dgtr.randRecord("RocketPunch.amr"))
			}
		})
	
	engine.OnKeyword("ke的攻略").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("ke.png"))
		})
	engine.OnKeyword("rm的攻略").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("rm.png"))
		})
	engine.OnKeyword("im的攻略").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("im.png"))
		})
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("buff：幻象之剑(证实，常驻buff必须维持)，克己(证实，减伤霸体，超越后免疫6Sdebuff)，无尽追击者(致命，原力，增伤移速)\n道中：最终一击 (厚重，施法较慢)，狂怒切割，无尽追击者(致命，原力），剑刃之轮 (原力输出吃X轴不如非原力稳定)\n对王：真烈焰斩(厚重)，剑狱降临(对巨，弱人型，原力吃站位)，索命残杀(致命/强化，自由指向不吃站位)，狂怒切割（原力，施法快可抢伤害，输出略高于非原力），幻想之舞(厚重，范围小需配合幻象之剑，脱手快伤害慢)，无限打击 (厚重，大体积，4把剑以上用原版，以下用原力但优先级低)\n特殊被动：剑的记忆（第三下被动10%的单技伤加成，尽量做到高伤技能吃到就行，不要因为第三下被动而特意去等，属于单技伤加成。推荐：巨型:剑雨>无限打击  人形:剑雨>幻舞>索命）\n特殊主动：剑气护盾(轻便，施法快，CD短。补充红剑道用）输出要点：追击者持续时间内尽量丢高伤技能，第三下被动尽量给高伤技能吃，但不要特意停下浪费输出时间。 \n连招思路：续好幻想之剑，无尽追击者起手，在无尽追击者持续时间内丢2~3个技能（CD转不过来需要更多），这个时候无尽追击者就结束了，接着再丢无尽追击者来继续2-3个及以上技能循环。 如果无尽追击者CD还没好，就优先丢高伤技能，根据剩余的技能CD来判断是否丢无限打击\n输出技能优先级参考：\n人形：剑狱降临>幻像之舞>索命残杀>真烈焰斩(无限打击)>狂怒切割，无限打击比较特殊，因为施法时间太长。要在CD实在是转不过来的时候才能丢，不然就按照上面的优先级，否则排在真爷爷之前或者之后。\n中型：剑狱降临>索命残杀、真烈焰斩、幻像之舞（贴脸的时候）>狂怒切割，无限打击同上\n巨型：剑狱降临>无限打击>真烈焰斩>索命残杀(远)、幻象之物(近)>狂怒切割ps：技能分支紫霸状态多的情况是 致命>厚重>强化，基本没紫霸的时候 厚重>强化>致命（本攻略为夜秋叶的IM笔记，感谢提供者和笔记制作者，由于原攻略过长，已尽量缩减"))
		})
		
	
	engine.OnFullMatchGroup([]string{"喜欢", "爱你", "爱", "suki", "daisuki", "すき", "好き", "贴贴", "老婆", "亲一个", "mua"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(4) {
			case 0:
				ctx.SendChain(randText("mua~"))
			case 1:
				ctx.SendChain(randText("那种事...也是可以的哦"))
			case 2:
				ctx.SendChain(randText("二刺螈,kimo"))	
			case 3:
				ctx.SendChain(randText("爬"))	
			}
		})
	engine.OnKeywordGroup([]string{"草你妈", "操你妈", "脑瘫", "废柴", "fw", "five", "废物", "战斗", "爬", "爪巴", "sb", "SB", "傻B"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(dgtr.randImage("FN.jpg", "WQ.jpg", "WQ1.jpg"))
		})
	engine.OnFullMatchGroup([]string{"早安", "早哇", "早上好", "ohayo", "哦哈哟", "お早う", "早好", "早", "早早早"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 9:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"啊......早上好...(哈欠)",
					"唔......吧唧...早上...哈啊啊~~~\n早上好......",
					"早上好......",
					"早上好呜......呼啊啊~~~~",
					"啊......早上好。\n昨晚也很激情呢！",
					"吧唧吧唧......怎么了...已经早上了么...",
					"早上好！",
					"......看起来像是傍晚，其实已经早上了吗？",
					"早上好......欸~~~脸好近呢",
				))
			case now >= 9 && now < 18:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"哼！这个点还早啥，昨晚干啥去了！？",
					"熬夜了对吧熬夜了对吧熬夜了对吧？？？！",
					"是不是熬夜是不是熬夜是不是熬夜？！",
				))
			case now >= 18 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"早个啥？哼唧！我都准备洗洗睡了！",
					"不是...你看看几点了，哼！",
					"晚上好哇",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"中午好", "午安", "午好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			if now > 11 && now < 15 { // 中午
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，ATRI会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"晚安", "oyasuminasai", "おやすみなさい", "晚好", "晚上好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"zzz...好涩哦..zzz....",
					"别...不要..zzz..那..zzz..",
					"嘻嘻..zzz..呐~..zzzz..",
					"...zzz....哧溜哧溜....",
				))
			case now >= 6 && now < 11:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"你可猝死算了吧！",
					"？啊这",
					"亲，这边建议赶快去睡觉呢~~~",
					"不可忍不可忍不可忍！！为何这还不猝死！！",
				))
			case now >= 11 && now < 15:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安w",
					"午觉要好好睡哦，ATRI会陪伴在你身旁的w",
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"睡你午觉去！哼唧！！",
				))
			case now >= 15 && now < 19:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"难不成？？晚上不想睡觉？？现在休息",
					"就......挺离谱的...现在睡觉",
					"现在还是白天哦，睡觉还太早了",
				))
			case now >= 19 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"嗯哼哼~睡吧，就像平常一样安眠吧~o(≧▽≦)o",
					"......(打瞌睡)",
					"呼...呼...已经睡着了哦~...呼......",
					"......我、我会在这守着你的，请务必好好睡着",
				))
			}
		})
	engine.OnKeywordGroup([]string{"高性能", "太棒了", "すごい", "sugoi", "斯国一", "よかった"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"当然，我是高性能的嘛~！",
				"小事一桩，我是高性能的嘛",
				"怎么样？还是我比较高性能吧？",
				"哼哼！我果然是高性能的呢！",
				"因为我是高性能的嘛！嗯哼！",
				"因为我是高性能的呢！",
				"哎呀~，我可真是太高性能了",
				"正是，因为我是高性能的",
				"是的。我是高性能的嘛♪",
				"毕竟我可是高性能的！",
				"嘿嘿，我的高性能发挥出来啦♪",
				"我果然是很高性能的机器人吧！",
				"是吧！谁叫我这么高性能呢！哼哼！",
				"交给我吧，有高性能的我陪着呢",
				"呣......我的高性能，毫无遗憾地施展出来了......",
			))
		})
	engine.OnKeywordGroup([]string{"没事", "没关系", "大丈夫", "还好", "不要紧", "没出大问题", "没伤到哪"}, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"当然，我是高性能的嘛~！",
				"没事没事，因为我是高性能的嘛！嗯哼！",
				"没事的，因为我是高性能的呢！",
				"正是，因为我是高性能的",
				"是的。我是高性能的嘛♪",
				"毕竟我可是高性能的！",
				"那种程度的事不算什么的。\n别看我这样，我可是高性能的",
				"没问题的，我可是高性能的",
			))
		})

	engine.OnKeywordGroup([]string{"好吗", "是吗", "行不行", "能不能", "可不可以"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(dgtr.randImage("YES.png", "NO.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"啊这"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(dgtr.randImage("AZ.jpg", "AZ1.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"我好了"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText("不许好！", "憋回去！"))
		})
	engine.OnFullMatchGroup([]string{"？", "?", "¿"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(dgtr.randImage("WH.jpg", "WH1.jpg", "WH2.jpg", "WH3.jpg"))
			}
		})
	engine.OnKeyword("离谱").SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(5) {
			case 0:
				ctx.SendChain(randText("?", "？", "嗯？", "(。´・ω・)ん?", "ん？"))
			case 1, 2:
				ctx.SendChain(dgtr.randImage("WH.jpg"))
			}
		})
	engine.OnKeyword("答应我", zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText("我无法回应你的请求"))
		})
}

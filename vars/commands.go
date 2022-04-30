package vars

const (
	CmdStartCode int8 = iota
	CmdSettingsCode
	CmdHelpCode
	CmdMeCode
)

const CmdStart = "/start"
const CmdStartText = "⭐ Главное меню"

const CmdSettings = "/settings"
const CmdSettingsText = "⚙ Настройки"

const CmdHelp = "/help"
const CmdHelpText = "🆘 Справка"

const CmdMe = "/me"
const CmdMeText = "👤 Информация о пользователе"

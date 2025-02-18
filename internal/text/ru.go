package text

var Ru = [...]string{
	Civilian:  "Мирный житель",
	Mafia:     "Мафия",
	Detective: "Шериф",
	Doctor:    "Доктор",

	GMafia:  "Мафия",
	GBunker: "Бункер",

	Default: "%s",

	Join:   "Присоединиться",
	Create: "Создать",

	GameStarted:     "Игра началась!",
	GameNotSelected: "",
	GameMafia:       "Мафия",

	IsMafiaF:      "%s - мафия!",
	IsNotMafiaF:   "%s - не мафия!",
	HealedF:       "Вы вылечили %s!",
	ChosenToKillF: "Вы выбрали убить %s!",
	NightFalls:    "Ночь наступает. Мирные жители засыпают",
	DayIsComing:   "День наступает. Мирные жители просыпаются",
	IsWakingUpF:   "%s просыпается!",
	FallAsleepF:   "%s сделала выбор и уснула!",
	MakeChoiceF:   "[%s] %s, сделай свой выбор:",
	MafiaSuccessF: "Плохие новости!\nМафия убила %s!",
	MafiaFailed:   "Хорошие новости!\nЭтой ночью доктор спас жизнь жертвы, и никто не умер!",

	MafiaWon:     "Мафия победила!",
	CiviliansWon: "Мирные жители победили!",

	Voting:         "Проголосуйте, кого вы хотите выгнать:",
	VotedF:         "Вы проголосовали за %s",
	VotingResultsF: "Результаты голосования:\n%s",
	VotesAreEqual:  "Голоса равны! Проголосуйте еще раз",
	VoteKickF:      "%s выгнан! Он был %s",

	GameChosenF:   "Вы выбрали %s.\nЧтобы начать игру, отправьте /game_start",
	CantStartGame: "Вы не можете начать игру, пока не выберете тип игры!",

	AtLeastMembersF: "В лобби должно быть как минимум %d участников, чтобы начать эту игру!",
	MaximumMembersF: "В лобби может быть максимум %d участников, чтобы начать эту игру!",
	RoleF:           "Вы - %s!",

	KeyCreatedF:        "Ваш ключ к лобби: %s\nТеперь выберите, в какую игру хотите сыграть!",
	SendKey:            "Напишите ключ к лобби, чтобы присоединиться, например: O1F5",
	IncorrectKey:       "Некорректный ключ!\nУбедитесь, что вы используете правильный ключ",
	LobbyNotExists:     "Лобби с этим ключом не существует или ваш ключ некорректен!",
	LobbyGameIsStarted: "В этом лобби игра уже началась! Вы не можете присоединиться.",

	PlayerJoinedLobbyF: "%s присоединился к лобби!\nТекущие участники в лобби:\n%s",
	AlreadyInLobbyF:    "Вы уже находитесь в лобби с ключом: %s!\nПокиньте это лобби, а затем сможете присоединиться.",

	UnknownCommand: "Я не знаю этой команды",
	StartCommandF:  "Привет, %s, я GFF бот!\nВыберите, чтобы присоединиться к лобби или создать одно, чтобы начать игру!",

	CreatingLobbyError: "Что-то пошло не так при создании лобби! Пожалуйста, попробуйте снова позже.",
	CantFindUser:       "Ошибка!\nНе могу вас найти в данных! Пожалуйста, перезапустите GFFBot",
	SomethingWentWrong: "Что-то пошло не так! Пожалуйста, перезапустите бота",

	You: "Вы",

	Yes: "Да",
	No:  "Нет",

	SugestToKillF: "%s предлагает выбрать %s как жертву",
	AcceptedF:     "%s принял ваше решение",
	DeclinedF:     "%s отверг ваше решение",

	GameStoped: "Что-то пошло не так. Игра остановлена, начните новую.",
}

const (
	Profession       = "Профессия"
	BoilogicalParams = "Биологические параметры"
	HealthStatus     = "Состояние здоровья"
	Hobby            = "Хобби"
	Phopia           = "Фобия"
	Character        = "Характер"
	Skill            = "Навык"
	Knowledge        = "Знание"
	Baggage          = "Багаж"
	ActionCard       = "Карта действия"
	ConditionCard    = "Карта условия"

	CantVotForYourSelf = "Нельзя голосовать за самого себя!\nПроголосуйте ещё раз за другого"
	VotedForF          = "Вы проголосовали за %s"

	ToOpen			= "Выберите что хотите раскрыть."
	ToKick			= "Выберите кого хотите выгнать"
	KickedF         = "%s выгнан из бункера"
	VotesAreEqualRu = "Количество голосов равно! Проголосуйте ещё раз"
	GameEnd 		= "Игра закончилась"
)

const ApocalypseExamplesLen = len(ApocalypseExamples)

var ApocalypseExamples = [...]string{
    "Ядерная война\nБункер глубоко под землей, с системой фильтрации воздуха, запасом еды на 50 лет и противорадиационными стенами. Есть лаборатория для исследований, но нет окон.",
    "Зомби-апокалипсис\nУкрепленный бункер с бронированными дверями, ловушками на подступах и оранжереей. Запас оружия и еды на 20 лет, но вентиляция уязвима для проникновения.",
    "Климатический коллапс\nПодземный комплекс с геотермальным генератором, запасами пресной воды и теплицами. Стены защищены от наводнений, но нет связи с внешним миром.",
    "Инопланетное вторжение\nБункер, замаскированный под обычный дом, с электромагнитной защитой и клонирующей капсулой. Есть сканеры, но энергия ограничена.",
    "Пандемия смертельного вируса\nСтерильный бункер с автономной системой жизнеобеспечения, лабораторией и карантинными зонами. Запас вакцин на 10 лет, но нет солнечного света.",
    "Падение астероида\nКуполообразный бункер под горой, с запасами еды на 30 лет и роботами-ремонтниками. Есть телескоп, но возможны подземные толчки.",
    "Искусственный интеллект восстал\nБункер с электромагнитными щитами, ручным управлением и библиотекой бумажных книг. Нет цифровых систем, но есть печатный станок.",
    "Биоинженерная катастрофа\nПодводный бункер с генетическими модификаторами растений и системой очистки воды. Есть подлодка, но опасность мутировавших существ.",
    "Солнечная супервспышка\nБункер с экранированными стенами, солнечными батареями и подземными фермами. Электроника защищена, но нет доступа к интернету.",
    "Магический апокалипсис\nБункер, защищенный древними рунами, с артефактами и гримуарами. Есть магический кристалл энергии, но возможны порталы в другие миры.",
    "Роботизированная чума\nБункер с электромагнитными импульсными пушками и мастерской для создания дронов. Запас запчастей на 15 лет, но роботы могут взломать защиту.",
    "Восстание животных\nКрепость на дереве с самонаводящимися ловушками и запасами семян. Есть дрон для разведки, но хищники умнее людей.",
}

const (
	ProfessionsFemale = 0
	ProfessionsMale	  = 19
	ProfessionsLen	  = len(ProfessionsRu)
)

var ProfessionsRu = [...]string{
	"Кальянщик",
	"Безработный",
	"Врач",
	"Фармацевт",
	"Инженер",
	"Военный",
	"Ученый",
	"Фермер",
	"Повар",
	"Няня",
	"Клоун",
	"Механик",
	"Программист",
	"Строитель",
	"Химик",
	"Сексолог",
	"Темщик",
	"Уролог",
	"Гинеколог",
	"Таролог",
	"Сутенёр",
	"Веб-камщица",
	"Проститутка",
}

const BiologicalParamsLen = len(BiologicalParamsRu)

var BiologicalParamsRu = [...]string{
	"гетеросексуал",
	"бисексуал",
	"асексуал",
	"пансексуал",
	"гомосексуал",
}

const (
	HealthStatusFemale = 1
	HealthStatusMale   = 21
	HealthStatusLen    = len(HealthStatusRu)
)

var HealthStatusRu = [...]string{
	"Евнух",
	"Здоров",
	"Рак в ремиссии",
	"ВИЧ-позитивный",
	"Травма позвоночника",
	"Слепота на один глаз",
	"Глухота",
	"Протез ноги",
	"COVID-19",
	"Слепота",
	"Глисты",
	"Аллергия на растения",
	"Аллергия на фрукты",
	"Шизофрения",
	"Деменция",
	"Синдром Туретта",
	"Немота",
	"Цинга",
	"Биполярное расстройство",
	"Гонорея",
	"Диарея",
}

const HobbiesLen = len(HobbiesRu)

var HobbiesRu = [...]string{
	"Садоводство",
	"Готовка",
	"Робототехника",
	"Чтение",
	"Охота",
	"Рыбалка",
	"Боевые искусства",
	"Пьянство",
}

const PhopiaLen = len(PhobiasRu)

var PhobiasRu = [...]string{
	"Нет фобии",
	"Клаустрофобия",
	"Арахнофобия",
	"Акрофобия (страх высоты)",
	"Агорафобия (страх открытых пространств)",
	"Гемофобия (страх крови)",
}

const CharacterLen = len(CharacterRu)

var CharacterRu = [...]string{
	"Лидер",
	"Оптимист",
	"Параноик",
	"Альтруист",
	"Эгоист",
	"Агрессор",
	"Паникёр",
}

const SkillsLen = len(SkillsRu)

var SkillsRu = [...]string{
	"Медицинские навыки",
	"Ремонт электроники",
	"Строительство укрытий",
	"Взлом замков",
	"Радиосвязь",
}

const BaggageLen = len(BaggageRu)

var BaggageRu = [...]string{
	"Аптечка первой помощи",
	"Набор инструментов",
	"Мешок картошки",
	"Ноутбук",
	"Семена пшеницы",
	"Огнетушитель",
	"Lego",
}
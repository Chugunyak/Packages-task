package animals

type Animal interface {
	Sound() string
	Move() string
	Feed() string
	Sleep() string
	Special() string
}

type Swimmer interface {
	Swim() string
	IsSwimming() bool
}

type Lion struct{}

func (l Lion) Sound() string {
	return "Рычит"
}

func (l Lion) Move() string {
	return "Бежит"
}

func (l Lion) Feed() string {
	return "Лев ест мясо"
}

func (l Lion) Sleep() string {
	return "Лев спит на солнце"
}

func (l Lion) Special() string {
	return "У льва есть грива"
}

type Dolphin struct{}

func (d Dolphin) Sound() string {
	return "Поёт"
}

func (d Dolphin) Move() string {
	return "Плывет"
}

func (d Dolphin) Feed() string {
	return "Дельфин ест рыбу"
}

func (d Dolphin) Sleep() string {
	return "Дельфин спит, выплывая на поверхность"
}

func (d Dolphin) Swim() string {
	return "Дельфин плавает"
}

func (d Dolphin) Special() string {
	return "Дельфины очень умные"
}

func (d Dolphin) IsSwimming() bool {
	return true
}

type Elephant struct{}

func (e Elephant) Sound() string {
	return "Топает"
}

func (e Elephant) Move() string {
	return "Идет"
}

func (e Elephant) Feed() string {
	return "Слон ест траву"
}

func (e Elephant) Sleep() string {
	return "Слон спит, сидя"
}

func (e Elephant) SprayWater() string {
	return "Слон распыляет воду"
}

func (e Elephant) Special() string {
	return "У слона есть хобот"
}

type Frog struct{}

func (f Frog) Sound() string {
	return "Квакает"
}

func (f Frog) Move() string {
	return "Прыгает"
}

func (f Frog) Feed() string {
	return "Жаба ест насекомых"
}

func (f Frog) Sleep() string {
	return "Жаба спит под листом"
}

func (f Frog) Jump() string {
	return "Жаба делает большой прыжок"
}

func (f Frog) Special() string {
	return "У лягушки длинный язык"
}

type Penguin struct{}

func (p Penguin) Sound() string {
	return "Пищит"
}

func (p Penguin) Move() string {
	return "Скользит по льду"
}

func (p Penguin) Feed() string {
	return "Пингвин ест рыбу"
}

func (p Penguin) Sleep() string {
	return "Пингвин спит на льду"
}

func (p Penguin) Swim() string {
	return "Пингвин отлично плавает"
}

func (p Penguin) Special() string {
	return "Пингвины умеют передвигаться на животе"
}

func (p Penguin) IsSwimming() bool {
	return true
}

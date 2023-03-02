package accounts

// Account Struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account => 매개변수 생성자 대체 Go 패턴함수 예시
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // account 메모리 주소를 가져와 반환 => 메모리 낭비 배제
}

// Deposit x amount on your account
func (acc Account) Deposit(amount int) {
	acc.balance += amount
}

// Balance of your account
func (acc Account) Balance() int {
	return acc.balance
}

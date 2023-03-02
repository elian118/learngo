package accounts

import "errors"

// Account Struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw.")

// NewAccount creates Account => 매개변수 생성자 대체 Go 패턴함수 예시
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // account 메모리 주소를 가져와 반환 => 메모리 낭비 배제
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// acc *Account -> Account 복사본 생성 없이 지정된 메모리가 가진 정보을 가져온다.
	a.balance += amount // 재할당
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// Go 는 try-catch-error 핸들러가 따로 없으므로, 직접 정의해야 함
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount // 재할당
	return nil          // null
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

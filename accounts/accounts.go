package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw.")

// NewAccount creates Account => 매개변수 생성자 대체 Go 패턴함수 예시
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // account 메모리 주소를 가져와 *참조값을 반환 => 메모리 낭비 배제
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

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// String은 자바의 Object.toString()처럼 struct를 문자열로 바꿔 반환하는 슈퍼 메서드이고, 오버라이드가 가능하다.
// 클래스를 문자열로 변환하는 슈퍼 메서드 String() 오버라이드
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}

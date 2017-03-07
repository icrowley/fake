package fake

import (
	"sync"
)

var mutex = sync.Mutex{}

func lock() func() {
	mutex.Lock()
	return mutex.Unlock
}

func CurrencyCodeLock() string {
	defer lock()()
	return CurrencyCode()
}

func EmailAddressLock() string {
	defer lock()()
	return EmailAddress()
}

func FirstNameLock() string {
	defer lock()()
	return FirstName()
}

func FullNameLock() string {
	defer lock()()
	return FullName()
}

func IPv4Lock() string {
	defer lock()()
	return IPv4()
}

func LastNameLock() string {
	defer lock()()
	return LastName()
}

func PhoneLock() string {
	defer lock()()
	return Phone()
}

func StreetLock() string {
	defer lock()()
	return Street()
}

func UserNameLock() string {
	defer lock()()
	return UserName()
}

func ZipLock() string {
	defer lock()()
	return Zip()
}

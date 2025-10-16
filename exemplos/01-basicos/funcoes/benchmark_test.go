package funcoes

import (
	"testing"
)

func BenchmarkBadProcessUserData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessUserData(
			"John",
			"Doe",
			"john@example.com",
			"123456789",
			"Street 1",
			"City",
			"State",
			"Country",
			"12345",
			30,
			true,
			true,
			true,
		)
	}
}

func BenchmarkGoodProcessUserData(b *testing.B) {
	userData := UserData{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "123456789",
		Address: Address{
			Street:  "Street 1",
			City:    "City",
			State:   "State",
			Country: "Country",
			ZipCode: "12345",
		},
		Flags: UserFlags{
			IsActive:  true,
			IsAdmin:   true,
			IsPremium: true,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessUser(userData, 30)
	}
}

func BenchmarkBadRecursion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BadRecursion(10)
	}
}

func BenchmarkGoodRecursion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumRecursive(10)
	}
}

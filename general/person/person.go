package person

import "fmt"

type Addable interface {
	int | float32
}

type Person struct {
	Name          string
	Age           int
	Address       string
	CreditScoring float32
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) GetAddress() string {
	return p.Address
}

func (p *Person) GetCreditScoring() float32 {
	return p.CreditScoring
}

func (p *Person) SayHello() (string, error) {
	if p.Name == "" {
		return "", fmt.Errorf("Name is empty")
	}

	return fmt.Sprintf("Hello %s", p.Name), nil
}


// Writing a generic function:
func SumIntOrFloat[V Addable](a, b V) V {
	return a + b
}
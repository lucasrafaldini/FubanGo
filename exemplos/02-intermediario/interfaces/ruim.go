package interfaces

import (
	"fmt"
)

// Interface grande e não coesa
type BigInterface interface {
	ReadData() []byte
	WriteData([]byte)
	ProcessData()
	ValidateData()
	TransformData()
	CompressData()
	EncryptData()
	SendData()
	ReceiveData()
	ParseData()
	CacheData()
	LogData()
	BackupData()
	RestoreData()
}

// Implementação que força muitos métodos desnecessários
type BigImplementation struct{}

func (b *BigImplementation) ReadData() []byte { return nil }
func (b *BigImplementation) WriteData([]byte) {}
func (b *BigImplementation) ProcessData()     {}
func (b *BigImplementation) ValidateData()    {}
func (b *BigImplementation) TransformData()   {}
func (b *BigImplementation) CompressData()    {}
func (b *BigImplementation) EncryptData()     {}
func (b *BigImplementation) SendData()        {}
func (b *BigImplementation) ReceiveData()     {}
func (b *BigImplementation) ParseData()       {}
func (b *BigImplementation) CacheData()       {}
func (b *BigImplementation) LogData()         {}
func (b *BigImplementation) BackupData()      {}
func (b *BigImplementation) RestoreData()     {}

// Interface que expõe detalhes de implementação
type BadDatabase interface {
	ConnectToMySQL()
	ExecuteSQLQuery(string)
	CloseMySQL()
}

// Interface que depende de tipos concretos
type BadProcessor interface {
	Process(*BigImplementation)
	Handle(*CustomError)
}

// Interface que viola ISP (Interface Segregation Principle)
type BadWorker interface {
	BadDoWork()
	SendEmail()
	GenerateReport()
	UpdateDatabase()
	NotifyAdmin()
}

// Implementação que precisa de todos os métodos mesmo só usando um
type BadSimpleWorker struct{}

func (s *SimpleWorker) BadDoWork()      { fmt.Println("working...") }
func (s *SimpleWorker) SendEmail()      { /* não usa, mas precisa implementar */ }
func (s *SimpleWorker) GenerateReport() { /* não usa, mas precisa implementar */ }
func (s *SimpleWorker) UpdateDatabase() { /* não usa, mas precisa implementar */ }
func (s *SimpleWorker) NotifyAdmin()    { /* não usa, mas precisa implementar */ }

// Interface com método Accept que aceita interface{}
type BadAcceptor interface {
	Accept(interface{})
}

// Implementação que usa type assertions excessivas
type BadTypeAssert struct{}

func (b *BadTypeAssert) ProcessAnything(data interface{}) {
	// Type assertions em cadeia
	if str, ok := data.(string); ok {
		fmt.Println("string:", str)
	} else if num, ok := data.(int); ok {
		fmt.Println("int:", num)
	} else if fl, ok := data.(float64); ok {
		fmt.Println("float64:", fl)
	} else if bl, ok := data.(bool); ok {
		fmt.Println("bool:", bl)
	} else if sl, ok := data.([]string); ok {
		fmt.Println("[]string:", sl)
	}
}

// Erro personalizado mal implementado
type CustomError struct {
	message string
}

// Não implementa Error() corretamente
func (e CustomError) String() string {
	return e.message
}

// Interface vazia usada desnecessariamente
type BadContainer struct {
	data interface{}
}

func (c *BadContainer) Store(value interface{}) {
	c.data = value
}

func (c *BadContainer) Retrieve() interface{} {
	return c.data
}

// Uso incorreto de embedding de interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Interface grande criada por embedding desnecessário
type BadReadWriter interface {
	Reader
	Writer
	Close()
	Flush()
	Reset()
	String() string
	Size() int64
	IsEmpty() bool
}

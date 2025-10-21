package interfaces

import (
	"fmt"
	"io"
)

// DataReader é uma interface focada e coesa
type DataReader interface {
	Read() ([]byte, error)
}

// DataWriter é uma interface focada e coesa
type DataWriter interface {
	Write([]byte) error
}

// DataProcessor é uma interface focada e coesa
type DataProcessor interface {
	Process([]byte) ([]byte, error)
}

// Implementação simples que combina interfaces quando necessário
type FileHandler struct{}

func (f *FileHandler) Read() ([]byte, error) {
	return []byte("dados"), nil
}

func (f *FileHandler) Write(data []byte) error {
	return nil
}

func (f *FileHandler) Process(data []byte) ([]byte, error) {
	return data, nil
}

// Database abstrai detalhes de implementação
type Database interface {
	Connect() error
	Execute(query string) error
	Close() error
}

// Processor depende de interfaces, não implementações
type GoodProcessor interface {
	Process(DataReader) error
	HandleError(error) error
}

// Worker segregado em interfaces menores
type Worker interface {
	DoWork() error
}

type EmailSender interface {
	SendEmail(string) error
}

type ReportGenerator interface {
	GenerateReport() ([]byte, error)
}

// Implementação que só precisa do que usa
type SimpleWorker struct{}

func (s *SimpleWorker) DoWork() error {
	return nil
}

// Generic Container com type safety
type Container[T any] struct {
	data T
}

func NewContainer[T any](value T) *Container[T] {
	return &Container[T]{data: value}
}

func (c *Container[T]) Store(value T) {
	c.data = value
}

func (c *Container[T]) Retrieve() T {
	return c.data
}

// Erro personalizado bem implementado
type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v (code: %d)", e.Message, e.Err, e.Code)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Composição sensata de interfaces
type ReadCloser interface {
	io.Reader
	io.Closer
}

// Implementação limpa de multiplas interfaces
type DataHandler struct{}

func (d *DataHandler) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (d *DataHandler) Close() error {
	return nil
}

// Factory function que retorna interface
func NewDataHandler() ReadCloser {
	return &DataHandler{}
}

// Interface que aceita genéricos ao invés de interface{}
type Processor[T any] interface {
	Process(input T) (T, error)
}

// Implementação com tipo específico
type StringProcessor struct{}

func (sp *StringProcessor) Process(input string) (string, error) {
	return input + " processado", nil
}

// Uso de composição ao invés de herança
type EnhancedReader struct {
	reader DataReader
}

func NewEnhancedReader(r DataReader) *EnhancedReader {
	return &EnhancedReader{reader: r}
}

func (er *EnhancedReader) Read() ([]byte, error) {
	data, err := er.reader.Read()
	if err != nil {
		return nil, err
	}
	// Adiciona funcionalidade sem modificar a interface original
	return append(data, []byte(" melhorado")...), nil
}

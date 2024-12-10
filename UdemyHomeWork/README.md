In-Memory Cache

Описание проекта
Этот пакет предоставляет in-memory кэш для хранения пар ключ-значение в оперативной памяти. Кэш поддерживает три основных операции:

Set(key string, value interface{}) — добавить значение по ключу.
Get(key string) — получить значение по ключу.
Delete(key string) — удалить значение по ключу.
Пакет написан на Go и может быть легко интегрирован в любой проект.

Установка
Для установки пакета используйте команду:


go get -u github.com/yourusername/cache

Замените yourusername на имя своего пользователя GitHub или название модуля.


Пример использования
Вот пример кода, демонстрирующий, как использовать пакет cache:
package main

import (
	"fmt"
	"github.com/yourusername/cache" // Замените на путь к вашему модулю
)

func main() {
	// Создаем новый экземпляр кэша
	c := cache.New()

	// Добавляем значение по ключу
	c.Set("userId", 42)

	// Получаем значение по ключу
	userId := c.Get("userId")
	fmt.Println("userId:", userId) // Вывод: userId: 42

	// Удаляем значение по ключу
	c.Delete("userId")

	// Проверяем, что значение удалено
	userId = c.Get("userId")
	fmt.Println("userId after delete:", userId) // Вывод: userId after delete: <nil>
}


Методы
New() *Cache
Создаёт и возвращает новый экземпляр кэша.

Set(key string, value interface{})

Добавляет значение value в кэш по ключу key.
c.Set("username", "JohnDoe")


Get(key string) interface{}
Возвращает значение по ключу key. Если ключ не найден, возвращает nil.

Пример:
value := c.Get("username")
fmt.Println(value) // Вывод: JohnDoe

Delete(key string)
Удаляет значение из кэша по ключу key.
c.Delete("username")
fmt.Println(c.Get("username")) // Вывод: <nil>

cache/
│── cache.go        # Реализация кэша
│── go.mod          # Модуль Go
└── README.md       # Документация

Лицензия
Этот проект распространяется под лицензией MIT.

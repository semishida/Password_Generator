# Password Generator (аналог pwgen)

Этот проект реализует генератор паролей на языке Go.

## Требования

- Go 1.20 или новее

## Установка

1. Клонируйте репозиторий:

    ```sh
    git clone https://github.com/yourusername/password_generator.git
    ```

2. Перейдите в директорию проекта:

    ```sh
    cd password_generator
    ```

3. Инициализируйте модуль Go:

    ```sh
    go mod init password_generator
    ```

## Использование

Для запуска генератора паролей выполните следующую команду:

```sh
go run main.go
 ```

## Тестирование

Для запуска тестов выполните следующие команды:

Запуск тестов:
```sh
go test
  ```
Запуск тестов с генерацией отчета о покрытии:
```sh
go test -coverprofile coverage.out
  ```
или 
```sh
go test -coverprofile=coverage.out
  ```
Генерация HTML-отчета о покрытии:
```sh
go tool cover -html coverage.out -o coverage.html
  ```
или 
```sh
go tool cover -html=coverage.out -o coverage.html
  ```

Этот README.md файл содержит краткое описание проекта, инструкцию по установке и использованию, а также информацию о тестировании и вкладе в проект.

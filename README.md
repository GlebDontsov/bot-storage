# bot-storage
Этот бот на Go предназначен для хранения ссылок и предоставления пользователю доступа к ним с помощью нескольких команд.

# Команды
- /rnd - получить случайную ссылку из хранилища;
- /last - получить последнюю добавленную ссылку;
- /first - получить первую добавленную ссылку;
- /all - получить список всех ссылок;
- /count - получить количество ссылок в хранилище;
- /help - получить справку о боте.

# Запуск бота
1. Склонируйте репозиторий и перейдите в директорию проекта
```
git clone https://github.com/GlebDontsov/bot-storage.git
```
2. Установите все необходимые зависимости:
```
go mod download
```
3.  Скомпилируйте свой код и создайте исполняемый файл:
```
go build
```
4. Запустите бота, передав токен в качестве аргумента командной строки:
```
./bot-storage -tg-bot-token=<YOUR_BOT_TOKEN>
```
5. Добавьте бота в список контактов Telegram и напишите ему /start, чтобы начать использование.



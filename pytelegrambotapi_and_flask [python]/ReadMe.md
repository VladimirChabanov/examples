# Запуск pyTelegramBotAPI и Flask одновременно

- `requirements.txt` - зависимости;
- `main.py` - тут происходит запуск;
- `flask_app.py` - логика работы Flask;
- `bot_app.py` - логика работы pyTelegramBotAPI;



## Подготовка и запуск

Чтобы не захламлять глобальное окружение нашими пакетами и уменьшить вероятность конфликта версий, будем работать в виртуальном окружении.

- Клонируйте репозиторий и перейдите в папку с исходниками;

- Создайте виртуальное окружение в текущей папке:

  ```
  python -m venv venv
  ```

- Активируйте виртуальное окружение (терминал git-a):

  ```
  source venv/Scripts/activate
  ```

  или терминал cmd:

  ```
  venv\Scripts\activate.bat
  ```

  или терминал PowerShell:

  ```
  .\venv\Scripts\Activate.ps1
  ```

  **Внимание:** для powershell по умолчанию запрещено выполнение скриптов, поэтому его нужно [разрешить](https://winitpro.ru/index.php/2020/06/03/powershell-execution-policy-zapusk-scriptov/).

- Установите зависимости:

  Нужно установить два пакета [pyTelegramBotAPI](https://pypi.org/project/pyTelegramBotAPI/) и [flask](https://flask.palletsprojects.com/en/3.0.x/):

  ```
  pip install pyTelegramBotAPI
  pip install Flask
  ```

  Приложение тестировалось на Windows 10 x64 с Python 3.9 и зависимостями указанными в requirements.txt. Если у вас такая же система, то можно установить зависимости командой:

  ```
  pip install -r requirements.txt
  ```

- Запустите скрипт (виртуальное окружение должно быть активировано):

  ```
  python main.py
  ```




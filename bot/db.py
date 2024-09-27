import os
import sqlite3
from datetime import datetime


class Database():
    def __init__(self):
        self.con = None
        self.bd_type = os.getenv('DB_TYPE')
        self.bd = "../database.db"

    def __del__(self):
        self.close()

    def open(self):
        if self.con is None:
            self.con = sqlite3.connect(self.bd, isolation_level=None)
            self.con.row_factory = sqlite3.Row
            self.cur = self.con.cursor()

    def close(self):
        self.con.close()
        self.cur = None
        self.con = None

    def execute(self, sql, args=()):
        cursor = self.con.execute(sql, args)
        return cursor

    def first(self, sql, args=()):
        cursor = self.execute(sql, args)
        if cursor is None:
            return None
        return cursor.fetchone()

    def get(self, sql, args):
        row = self.first(sql, args)
        if row is None:
            return None
        return row[0]


db = Database()
time_now = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
db.open()

# Создает базу данных при первом запуске
def create_db():
    # Проверим по tg_id, есть ли пользователь в базе users
    db.open()  # Подключимся к БД
    # запрос в БД
    db.cur.execute(f'''CREATE TABLE users(
        "id" INT PRIMARY KEY AUTOINCREMENT,
        "name" VARCHAR(60),
        "email" VARCHAR(60) UNIQUE,
        "password" VARCHAR(255),
        "telegram_id" VARCHAR UNIQUE,
        "role_id" INT DEFAULT 0,
        "licensed" BOOLEAN DEFAULT 0,
        "license_end" TIMESTAMP,
        "created_at" TIMESTAMP,
        "updated_at" TIMESTAMP,
    )''')
    res = False
    if db.cur.fetchone():  # Извлечем данные
        res = True
    db.close()  # Закроем соединение
    return res

def check_user(telegram_id):
    # Проверим по tg_id, есть ли пользователь в базе users
    db.open()  # Подключимся к БД
    # запрос в БД
    db.cur.execute(f"SELECT * FROM users WHERE telegram_id = '{telegram_id}'")
    res = False
    if db.cur.fetchone():  # Извлечем данные
        res = True
    db.close()  # Закроем соединение
    return res


def registration(telegram_id, name):
    user_exist = check_user(telegram_id)
    
    if user_exist != True:
        
        timestamp = datetime.now()
        db.open()
        user_id = db.cur.execute(f"INSERT INTO users (telegram_id, name, created_at, updated_at)\
                        VALUES ({telegram_id}, '{name}', '{timestamp}', '{timestamp}')")
        db.con.commit()
        db.close()
        return user_by_id(user_id)


def finish(telegram_id, finish):
    # Сделаем запись об успешном прохождении опроса
    db.open()
    db.cur.execute(f"UPDATE users SET finish='{finish}',\
                     result_time='{time_now}'\
                     WHERE tg_id={tg_id}")
    db.con.commit()
    db.close()
    return True


def update_one_var(tg_id, row, var):
    # Запишем / обновим одну переменную для tg_id в базе users
    db.open()
    db.cur.execute(f"UPDATE users SET {row}='{var}' WHERE tg_id={tg_id}")
    db.con.commit()
    db.close()
    return True


def targets(tg_id):
    # Выведем список
    db.open()
    db.cur.execute(f"SELECT name, question1, question2, question3, question4, question5, question6, question7, question8, question9, question10, question11, question12, question13, question14, question15, question16, question17, question18 FROM users WHERE (active='True' AND telegram_id='{tg_id}')")
    data = db.cur.fetchall()
    res = []
    
    db.close()
    return res


def user_by_id(id):
    db.open()
    sql = "SELECT * FROM users WHERE id=?"
    user_data = db.first(sql, (id,))
    if user_data is None:
        return None
    db.close()
    return dict(user_data)
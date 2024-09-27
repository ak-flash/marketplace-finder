import db
from aiogram import types, F, Router
from aiogram.types import Message
from aiogram.filters import Command


router = Router()


@router.message(Command("start"))
async def start_handler(msg: Message):
    tg_id = msg.chat.id
    name = msg.chat.first_name
    
    await msg.answer(f"Привет, {name}! Я помогу тебе отслеживать и находить лучшие цены на товары со СберМегаМаркета")

    # Запишем в базу данные юзера  
    db.registration(tg_id, name)



@router.message(F.text.lower() == "get-targets")
async def get_targets(msg: Message):
    await msg.answer(f"Твой ID: {msg.from_user.id}")


@router.message()
async def message_handler(msg: Message):
    await msg.answer("Используйте меню для выполнения команд")
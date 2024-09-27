import asyncio
import logging
import os

from aiogram.client.bot import DefaultBotProperties
from aiogram import Bot, Dispatcher
from aiogram.enums.parse_mode import ParseMode
from aiogram.fsm.storage.memory import MemoryStorage

from dotenv import load_dotenv
from handlers import router


async def main():
    tg_token = str(os.getenv('BOT_TOKEN'))
    bot = Bot(token=tg_token, default=DefaultBotProperties(parse_mode=ParseMode.HTML))
    dp = Dispatcher(storage=MemoryStorage())
    dp.include_router(router)
    await bot.delete_webhook(drop_pending_updates=True)
    await dp.start_polling(bot, allowed_updates=dp.resolve_used_update_types())


if __name__ == "__main__":
    load_dotenv("../.env")
    logging.basicConfig(level=logging.INFO)
    asyncio.run(main())
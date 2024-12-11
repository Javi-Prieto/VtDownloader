import asyncio
from queue import Queue

from fastapi import APIRouter, WebSocket
from starlette.responses import StreamingResponse


router = APIRouter(prefix='/vt-sse')

message_queue = Queue()

@router.get("/download")
async def sse_download():
    return StreamingResponse(send_download_message(), media_type="text/event-stream")

async def send_download_message():
    while True:
        message = await get_message()
        yield f"data: {message}\n\n"
        await asyncio.sleep(0.1)

async def get_message():
    while message_queue.empty():
        await asyncio.sleep(1)
    return message_queue.get()

def receive_download_message(download_id: int, status: str):
    msg = f'{{"download_id": {download_id}, "status": "{status}"}}'
    message_queue.put(msg)


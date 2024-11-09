import uvicorn
from fastapi import FastAPI
from MOD_D_Downloads.endpoint import downloads_controller
from common.MOD_S_SSE import sse

app = FastAPI()

app.include_router(router=downloads_controller.router)
app.include_router(router=sse.router)

if __name__ == '__main__':
    uvicorn.run(app='main:app', reload=True)

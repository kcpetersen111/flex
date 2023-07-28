import ffmpeg
import asyncio
import websockets

async def stream(websocket):
    
    out, _ = (
        ffmpeg
        .input('in.mp4')
        .output('pipe:', format='rawvideo', pix_fmt='rgb24')
        .run(capture_stdout=True)
    )

   
    await websocket.send(out)

async def main():
    async with websockets.serve(stream, "", 3410):
        await asyncio.Future()

if __name__ == "__main__":
    asyncio.run(main())
print(out)

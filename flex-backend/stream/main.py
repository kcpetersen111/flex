import ffmpeg
import asyncio
import websockets
import numpy as np
import json

input_vid = 'in.mp4'
frame_num = 0

def extract_frame(input_vid, frame_num):
    # at some point should stop ignoring the error
    out, _ = (
        ffmpeg
        .input(input_vid)
        .filter_('select', 'gte(n,{})'.format(frame_num))
        .output('pipe:', format='rawvideo', pix_fmt='rgb24', vframes=1)
        .run(capture_stdout=True, capture_stderr=True)
    )
    return out
    # return np.frombuffer(out, np.uint16).reshape(-1, height, width, 2)


async def stream(websocket):

    print(ffmpeg.probe(input_vid))
    stats = ffmpeg.probe(input_vid)
    
    frameRate = stats['avg_frame_rate'].split('/')
    frame = []
    for i in range(frameRate[0]):
        frame_num+=i
        frame.append(extract_frame(input_vid,frame_num))
   # print(video_stream)
   
    await websocket.send(frame)

async def main():
    async with websockets.serve(stream, "", 3410):
        await asyncio.Future()

if __name__ == "__main__":
    asyncio.run(main())
print(out)

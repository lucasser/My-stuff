import asyncio
import websockets

connected = set()

async def server(websocket, path):
    connected.add(websocket)
    print(connected)
    try:
        async for message in websocket:
            for conn in connected:
                await conn.send(websocket, message)
    finally:
        connected.remove(websocket)
    

start_server = websockets.serve(server, "localhost", 5000)

asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
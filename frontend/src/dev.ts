let ws: WebSocket | undefined;
let reconnectInterval: number | undefined = undefined;
let wasDisconnected = false;

export const WebSocketConnect = () => {
  try {
    ws = new WebSocket("ws://localhost:8080/ws");
    ws.onopen = () => {
      if (wasDisconnected) {
        window.location.reload();
      }
      clearInterval(reconnectInterval);
      if (ws) {
        ws.onclose = () => {
          wasDisconnected = true;
          reconnectInterval = setInterval(() => {
            console.log("try to connect to server");
            ws = undefined;
            WebSocketConnect();
          }, 1000);
        };
      }
    };
    ws.onerror = () => {
      ws?.close();
    };
  } catch (error) {
    console.error(error);
  }
};

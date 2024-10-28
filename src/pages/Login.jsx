import { useEffect, useRef, useState } from "react";
import "../index.css";

export default function Login() {
  const wsRef = useRef(null);
  const [message, setMessage] = useState();
  useEffect(() => {
    wsRef.current = new WebSocket(`ws://localhost:5000/ws`);
  }, []);
  const handleSendingMessage = (e) => {
    wsRef.current.send(message);
  };
  return (
    <div className="h-screen w-screen flex items-center justify-center bg-slate-500 text-white">
      <input onChange={(e) => setMessage(e.target.value)} type="text" />
      <button onClick={handleSendingMessage}>Send Message</button>
    </div>
  );
}

import { useState } from "react";

export function PingButton() {
  const [msg, setMsg] = useState<string | null>(null);

  async function pingServer() {
    const res = await fetch("api/ping");
    const data = await res.json();
    setMsg(data.message);
  }

  return (
    <>
      <button onClick={pingServer}>Ping Go</button>
      {msg && <p>Response: {msg}</p>}
    </>
  );
}

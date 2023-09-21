'use client'
import { useEffect, useState } from 'react'

export default function Home() {
  const [response, setResponse] = useState('');

  const testAPI = async () => {
    const res = await fetch('http://localhost:8080/api/info');
    const data = await res.json();
    console.log("Fetched data:", data);
    if (data && data.message) {
      setResponse(data.message);
    }
  };

  useEffect(() => {
    testAPI();
  }, []); 

  return (
    <main>
      {response}
    </main>
  );
}
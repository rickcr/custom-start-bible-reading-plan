"use client";

import { signIn, signOut, useSession } from "next-auth/react";
import { useState } from "react";

export default function Home() {
  const { data: session, status } = useSession();
  const [backendData, setBackendData] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const fetchBackendData = async () => {
    if (!session?.accessToken) return;

    setLoading(true);
    setError("");
    try {
      // The Go backend will expect the token in the Authorization header
      const res = await fetch("http://localhost:8080/api/user", {
        headers: {
          Authorization: `Bearer ${session.accessToken}`,
        },
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status}`);
      }

      const data = await res.json();
      setBackendData(data);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  if (status === "loading") {
    return <div className="flex min-h-screen items-center justify-center">Loading...</div>;
  }

  if (session) {
    return (
      <div className="flex min-h-screen flex-col items-center justify-center p-24 bg-gray-50 dark:bg-zinc-900">
        <div className="max-w-md w-full items-center justify-between font-mono text-sm flex flex-col gap-6 bg-white dark:bg-zinc-800 p-8 rounded-xl shadow-lg border border-gray-200 dark:border-zinc-700">
          <div className="flex flex-col items-center gap-4">
            {session.user?.image && (
              <img
                src={session.user.image}
                alt="Profile"
                className="w-20 h-20 rounded-full border-4 border-indigo-500 shadow-sm"
              />
            )}
            <h1 className="text-2xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-indigo-500 to-purple-600">
              Welcome, {session.user?.name}
            </h1>
            <p className="text-gray-500 dark:text-gray-400">{session.user?.email}</p>
          </div>

          <button
            onClick={fetchBackendData}
            disabled={loading}
            className="w-full px-6 py-3 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg font-medium transition-all transform active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed shadow-md flex justify-center items-center"
          >
            {loading ? (
              <svg className="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            ) : "Call Go Backend"}
          </button>

          {error && (
            <div className="w-full p-4 bg-red-50 dark:bg-red-900/30 text-red-600 dark:text-red-400 rounded-lg text-center text-sm border border-red-200 dark:border-red-800">
              {error}
            </div>
          )}

          {backendData && (
            <div className="w-full p-4 bg-green-50 dark:bg-green-900/30 text-green-700 dark:text-green-400 rounded-lg border border-green-200 dark:border-green-800 break-all">
              <p className="font-semibold mb-2 flex items-center gap-2">
                <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" /></svg>
                Go Backend Response:
              </p>
              <pre className="text-xs bg-white/50 dark:bg-black/20 p-2 rounded">
                {JSON.stringify(backendData, null, 2)}
              </pre>
            </div>
          )}

          <button
            onClick={() => signOut()}
            className="mt-4 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 text-sm transition-colors"
          >
            Sign out
          </button>
        </div>
      </div>
    );
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24 bg-white dark:bg-black bg-[radial-gradient(ellipse_at_top,_var(--tw-gradient-stops))] from-indigo-100/50 via-white to-white dark:from-zinc-900 dark:via-black dark:to-black">
      <div className="z-10 max-w-sm w-full items-center justify-between font-mono text-sm flex flex-col gap-8 text-center">
        <div className="space-y-4">
          <h1 className="text-5xl font-extrabold tracking-tight bg-clip-text text-transparent bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 pb-2">
            Bible Plan App
          </h1>
          <p className="text-gray-500 dark:text-gray-400 text-lg">
            Sign in to sync your custom reading plan.
          </p>
        </div>

        <button
          onClick={() => signIn("google")}
          className="group relative flex w-full justify-center rounded-xl border border-transparent bg-white dark:bg-white/10 px-6 py-4 text-sm font-semibold text-gray-900 dark:text-white shadow-xl hover:bg-gray-50 dark:hover:bg-white/20 hover:shadow-2xl transition-all duration-200 overflow-hidden"
        >
          <div className="absolute inset-0 w-full h-full bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 opacity-0 group-hover:opacity-10 transition-opacity duration-300"></div>
          <div className="flex items-center gap-3 z-10">
            <svg className="w-5 h-5" viewBox="0 0 24 24">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4" />
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853" />
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05" />
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335" />
              <path d="M1 1h22v22H1z" fill="none" />
            </svg>
            Sign in with Google
          </div>
        </button>
      </div>
    </main>
  );
}

import type { MetaFunction } from "@remix-run/node";
import { useState } from "react";
import { Link } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "Hello - ArtWork" },
    {
      name: "description",
      content: "バックエンドAPIとの連携テストページ",
    },
  ];
};

export default function Hello() {
  const [response, setResponse] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  const handleHelloRequest = async () => {
    setLoading(true);
    setError("");
    setResponse("");

    try {
      // バックエンドのhelloエンドポイントにアクセス
      const res = await fetch("http://localhost:8080/api/v1/hello");

      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }

      const data = await res.text();
      setResponse(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "エラーが発生しました");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-white">
      {/* ヘッダー */}
      <header className="bg-white border-b border-amber-100 sticky top-0 z-50 shadow-sm">
        <div className="max-w-5xl mx-auto px-6 py-5">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl md:text-3xl font-light text-gray-800 mb-1 tracking-wider">
                <span className="text-amber-600 font-normal">Art</span>Work
              </h1>
              <p className="text-gray-500 text-sm md:text-base hidden sm:block font-light tracking-wide">
                上質なアートと美の世界
              </p>
            </div>

            {/* ナビゲーション */}
            <nav className="hidden md:flex space-x-8">
              <Link
                to="/"
                className="text-gray-600 hover:text-amber-600 transition-colors font-light tracking-wide relative group"
              >
                ホーム
                <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-amber-400 transition-all duration-300 group-hover:w-full"></span>
              </Link>
              <Link
                to="/hello"
                className="text-amber-600 font-light tracking-wide relative group"
              >
                Hello
                <span className="absolute -bottom-1 left-0 w-full h-0.5 bg-amber-400"></span>
              </Link>
            </nav>
          </div>
        </div>
      </header>

      {/* メインコンテンツ */}
      <main className="max-w-4xl mx-auto px-6 py-16">
        <div className="text-center mb-12">
          <h2 className="text-3xl md:text-4xl font-light text-gray-800 mb-4 tracking-wide">
            Hello API テスト
          </h2>
          <p className="text-gray-600 text-lg font-light tracking-wide">
            バックエンドのhelloエンドポイントとの連携をテストします
          </p>
        </div>

        {/* APIテストセクション */}
        <div className="bg-gradient-to-br from-amber-50 to-orange-50 rounded-xl p-8 md:p-12 shadow-sm border border-amber-100">
          <div className="text-center mb-8">
            <h3 className="text-xl md:text-2xl font-light text-gray-800 mb-4 tracking-wide">
              バックエンドAPIにアクセス
            </h3>
            <p className="text-gray-600 font-light tracking-wide mb-6">
              下のボタンをクリックして、バックエンドのhelloエンドポイントからレスポンスを取得します
            </p>

            <button
              onClick={handleHelloRequest}
              disabled={loading}
              className={`px-8 py-4 rounded-lg font-light tracking-wide transition-all duration-300 ${
                loading
                  ? "bg-gray-300 text-gray-500 cursor-not-allowed"
                  : "bg-gradient-to-r from-amber-400 to-yellow-500 text-white hover:from-amber-500 hover:to-yellow-600 shadow-lg hover:shadow-xl transform hover:scale-105"
              }`}
            >
              {loading ? (
                <div className="flex items-center">
                  <svg
                    className="animate-spin -ml-1 mr-3 h-5 w-5 text-gray-500"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                  >
                    <circle
                      className="opacity-25"
                      cx="12"
                      cy="12"
                      r="10"
                      stroke="currentColor"
                      strokeWidth="4"
                    ></circle>
                    <path
                      className="opacity-75"
                      fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                    ></path>
                  </svg>
                  リクエスト中...
                </div>
              ) : (
                "Hello API を呼び出す"
              )}
            </button>
          </div>

          {/* レスポンス表示エリア */}
          {(response || error) && (
            <div className="mt-8">
              <h4 className="text-lg font-light text-gray-800 mb-4 tracking-wide">
                レスポンス結果
              </h4>

              {error ? (
                <div className="bg-red-50 border border-red-200 rounded-lg p-6">
                  <div className="flex items-center mb-2">
                    <svg
                      className="w-5 h-5 text-red-500 mr-2"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        strokeLinecap="round"
                        stroke-linejoin="round"
                        strokeWidth={2}
                        d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                      />
                    </svg>
                    <span className="text-red-800 font-medium">エラー</span>
                  </div>
                  <p className="text-red-700 font-light">{error}</p>
                </div>
              ) : (
                <div className="bg-green-50 border border-green-200 rounded-lg p-6">
                  <div className="flex items-center mb-2">
                    <svg
                      className="w-5 h-5 text-green-500 mr-2"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth={2}
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                      />
                    </svg>
                    <span className="text-green-800 font-medium">成功</span>
                  </div>
                  <div className="bg-white rounded-lg p-4 mt-3 border border-green-100">
                    <code className="text-green-800 font-mono text-lg">
                      {response}
                    </code>
                  </div>
                </div>
              )}
            </div>
          )}

          {/* API情報 */}
          <div className="mt-8 pt-8 border-t border-amber-200">
            <h4 className="text-lg font-light text-gray-800 mb-4 tracking-wide">
              API情報
            </h4>
            <div className="bg-white rounded-lg p-6 border border-amber-100">
              <div className="space-y-3 text-sm font-mono">
                <div className="flex items-center">
                  <span className="text-gray-500 w-20">URL:</span>
                  <span className="text-gray-800">
                    http://localhost:8080/api/v1/hello
                  </span>
                </div>
                <div className="flex items-center">
                  <span className="text-gray-500 w-20">Method:</span>
                  <span className="text-gray-800">GET</span>
                </div>
                <div className="flex items-center">
                  <span className="text-gray-500 w-20">Response:</span>
                  <span className="text-gray-800">&quot;helloworld&quot;</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* ホームに戻るリンク */}
        <div className="text-center mt-12">
          <Link
            to="/"
            className="inline-flex items-center text-amber-600 hover:text-amber-700 font-light transition-all duration-300 text-base tracking-wide group"
          >
            <svg
              className="mr-2 w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={1.5}
                d="M15 19l-7-7 7-7"
              />
            </svg>
            ホームに戻る
          </Link>
        </div>
      </main>

      {/* フッター */}
      <footer className="bg-gradient-to-b from-white to-amber-50 border-t border-amber-100 mt-20">
        <div className="max-w-5xl mx-auto px-6 py-10 md:py-12">
          <div className="text-center">
            <p className="text-sm text-gray-500 font-light tracking-wide">
              &copy; 2025 ArtWork. All rights reserved.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}

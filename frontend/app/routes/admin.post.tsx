import type {
  MetaFunction,
  LoaderFunctionArgs,
  ActionFunctionArgs,
} from "@remix-run/node";
import {
  Form,
  useLoaderData,
  useActionData,
  useNavigation,
} from "@remix-run/react";
import { requireAdmin } from "~/utils/auth.server";
import { redirect } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "記事投稿 - ArtWork" },
    {
      name: "description",
      content: "管理者専用の記事投稿ページ",
    },
  ];
};

// 管理者権限が必要なページとして設定
export async function loader({ request }: LoaderFunctionArgs) {
  const token = await requireAdmin(request);
  return { token };
}

// 記事投稿処理
export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();
  const title = formData.get("title") as string;
  const content = formData.get("content") as string;
  const category = formData.get("category") as string;

  console.log("📝 Article submission:", {
    title,
    contentLength: content?.length,
    category,
  });

  // バリデーション
  if (!title || !content || !category) {
    return {
      error: "タイトル、内容、カテゴリをすべて入力してください",
      success: false,
    };
  }

  try {
    // 実際の実装では、バックエンドのAPIに記事を送信
    console.log("✅ Article created successfully");

    return {
      success: true,
      message: "記事が正常に投稿されました",
    };
  } catch (error) {
    console.log("❌ Article creation failed:", error);
    return {
      error: "記事の投稿に失敗しました",
      success: false,
    };
  }
}

export default function AdminPost() {
  const { token } = useLoaderData<typeof loader>();
  const actionData = useActionData<typeof action>();
  const navigation = useNavigation();
  const isSubmitting = navigation.state === "submitting";

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
                管理者専用 - 記事投稿
              </p>
            </div>

            {/* 管理者権限表示とログアウトボタン */}
            <div className="flex items-center space-x-4">
              <div className="bg-red-50 border border-red-200 rounded-lg px-4 py-2">
                <div className="flex items-center">
                  <svg
                    className="w-5 h-5 text-red-500 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
                    />
                  </svg>
                  <span className="text-red-800 text-sm font-medium">
                    管理者権限
                  </span>
                </div>
              </div>

              {/* ログアウトボタン */}
              <Form method="post" action="/logout">
                <button
                  type="submit"
                  className="px-4 py-2 bg-gray-100 hover:bg-gray-200 text-gray-700 rounded-lg text-sm font-medium transition-colors"
                >
                  ログアウト
                </button>
              </Form>
            </div>
          </div>
        </div>
      </header>

      {/* メインコンテンツ */}
      <main className="max-w-4xl mx-auto px-6 py-16">
        <div className="text-center mb-12">
          <h2 className="text-3xl md:text-4xl font-light text-gray-800 mb-4 tracking-wide">
            記事投稿
          </h2>
          <p className="text-gray-600 text-lg font-light tracking-wide">
            新しい記事を作成して投稿します
          </p>
        </div>

        {/* 記事投稿フォーム */}
        <div className="bg-gradient-to-br from-amber-50 to-orange-50 rounded-xl p-8 md:p-12 shadow-sm border border-amber-100">
          <Form method="post" className="space-y-6">
            {/* 成功メッセージ */}
            {actionData?.success && (
              <div className="bg-green-50 border border-green-200 rounded-lg p-4">
                <div className="flex items-center">
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
                  <span className="text-green-800 text-sm font-medium">
                    {actionData.message}
                  </span>
                </div>
              </div>
            )}

            {/* エラーメッセージ */}
            {actionData?.error && (
              <div className="bg-red-50 border border-red-200 rounded-lg p-4">
                <div className="flex items-center">
                  <svg
                    className="w-5 h-5 text-red-500 mr-2"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                  <span className="text-red-800 text-sm font-medium">
                    {actionData.error}
                  </span>
                </div>
              </div>
            )}

            {/* タイトルフィールド */}
            <div>
              <label
                htmlFor="title"
                className="block text-sm font-medium text-gray-700 mb-2"
              >
                タイトル
              </label>
              <input
                id="title"
                name="title"
                type="text"
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors font-light tracking-wide"
                placeholder="記事のタイトルを入力"
                disabled={isSubmitting}
              />
            </div>

            {/* カテゴリフィールド */}
            <div>
              <label
                htmlFor="category"
                className="block text-sm font-medium text-gray-700 mb-2"
              >
                カテゴリ
              </label>
              <select
                id="category"
                name="category"
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors font-light tracking-wide"
                disabled={isSubmitting}
              >
                <option value="">カテゴリを選択</option>
                <option value="絵画">絵画</option>
                <option value="工芸">工芸</option>
                <option value="写真">写真</option>
                <option value="展覧会">展覧会</option>
                <option value="その他">その他</option>
              </select>
            </div>

            {/* 内容フィールド */}
            <div>
              <label
                htmlFor="content"
                className="block text-sm font-medium text-gray-700 mb-2"
              >
                内容
              </label>
              <textarea
                id="content"
                name="content"
                rows={10}
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors font-light tracking-wide resize-vertical"
                placeholder="記事の内容を入力してください"
                disabled={isSubmitting}
              />
            </div>

            {/* 投稿ボタン */}
            <div className="flex justify-end space-x-4">
              <a
                href="/"
                className="px-6 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors font-light tracking-wide"
              >
                キャンセル
              </a>
              <button
                type="submit"
                disabled={isSubmitting}
                className={`px-8 py-3 rounded-lg font-light tracking-wide transition-all duration-300 ${
                  isSubmitting
                    ? "bg-gray-300 text-gray-500 cursor-not-allowed"
                    : "bg-gradient-to-r from-amber-400 to-yellow-500 text-white hover:from-amber-500 hover:to-yellow-600 shadow-lg hover:shadow-xl transform hover:scale-105"
                }`}
              >
                {isSubmitting ? (
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
                    投稿中...
                  </div>
                ) : (
                  "記事を投稿"
                )}
              </button>
            </div>
          </Form>
        </div>

        {/* 管理者情報 */}
        <div className="mt-12 text-center">
          <div className="bg-white rounded-lg p-6 border border-amber-100 shadow-sm">
            <h3 className="text-lg font-light text-gray-800 mb-4 tracking-wide">
              管理者情報
            </h3>
            <div className="space-y-2 text-sm text-gray-600">
              <p>認証トークン: {token?.substring(0, 20)}...</p>
              <p>権限: 管理者</p>
              <p>アクセス日時: {new Date().toLocaleString("ja-JP")}</p>
            </div>
          </div>
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

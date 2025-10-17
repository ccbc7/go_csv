import type {
  MetaFunction,
  ActionFunctionArgs,
  LoaderFunctionArgs,
} from "@remix-run/node";
import { Form, useActionData, useNavigation, redirect } from "@remix-run/react";
import { getTokenFromCookie, authCookie } from "~/utils/auth.server";

/**
 * 管理者ログインページのメタデータ
 * @returns 管理者ログインページのメタデータ
 */
export const meta: MetaFunction = () => [
  { title: "管理者ログイン - ArtWork" },
  {
    name: "description",
    content: "管理者認証ページ",
  },
];

/**
 * サーバーサイドで認証状態をチェック
 * @param request
 * @returns 認証が成功した場合はリダイレクト, 失敗した場合はnull
 */
export async function loader({ request }: LoaderFunctionArgs) {
  // 既にログインしている場合は管理者ダッシュボードにリダイレクト
  const token = await getTokenFromCookie(request);

  if (token) {
    // JWTトークンが有効かチェック
    try {
      const response = await fetch("http://backend:8080/api/v1/auth/verify", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        return redirect("/admin/dashboard");
      }
    } catch (error) {
      console.error("Token verification failed:", error);
    }
  }

  return null;
}

/**
 * ログインフォーム送信時の処理
 * @param request
 * @returns ログインが成功した場合はリダイレクト, 失敗した場合はエラー
 */
export async function action({ request }: ActionFunctionArgs) {
  const formData = await request.formData();
  const email = formData.get("email") as string;
  const password = formData.get("password") as string;

  console.log("🔍 Login attempt:", { email, passwordLength: password?.length });

  // バリデーション
  if (!email || !password) {
    console.log("❌ Validation failed: missing fields");
    return {
      error: "メールアドレスとパスワードを入力してください",
      success: false,
    };
  }

  try {
    console.log("📡 Sending request to backend...");
    // バックエンドの認証APIを呼び出し
    const response = await fetch("http://backend:8080/api/v1/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email,
        password,
      }),
    });

    console.log("📊 Response status:", response.status);
    console.log(
      "📊 Response headers:",
      Object.fromEntries(response.headers.entries())
    );

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      console.log("❌ Login failed:", errorData);
      return {
        error: errorData.error || `ログインに失敗しました (${response.status})`,
        success: false,
      };
    }

    const data = await response.json();
    console.log("✅ Login successful:", { hasToken: !!data.token });

    // JWTトークンを取得
    if (data.token) {
      // クッキーにトークンを設定して管理者ダッシュボードにリダイレクト
      return redirect("/admin/dashboard", {
        headers: {
          "Set-Cookie": await authCookie.serialize(data.token),
        },
      });
    }

    return {
      error: "トークンの取得に失敗しました",
      success: false,
    };
  } catch (error) {
    console.log("💥 Network error:", error);
    return {
      error: `サーバーエラーが発生しました: ${
        error instanceof Error ? error.message : "Unknown error"
      }`,
      success: false,
    };
  }
}

/**
 * 管理者ログインページ
 * @returns 管理者ログインページ
 * @description 管理者ログインページは、管理者ログインフォームを表示する
 */
export default function ManagementAuth() {
  const actionData = useActionData<typeof action>();
  const navigation = useNavigation();
  const isSubmitting = navigation.state === "submitting";

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-gray-100 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        {/* ヘッダー */}
        <div className="text-center">
          <div className="mx-auto h-16 w-16 bg-gradient-to-r from-amber-400 to-yellow-500 rounded-full flex items-center justify-center shadow-lg">
            <svg
              className="h-8 w-8 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
              />
            </svg>
          </div>
          <h2 className="mt-6 text-3xl font-light text-gray-900 tracking-wide">
            管理者ログイン
          </h2>
          <p className="mt-2 text-sm text-gray-600 font-light">
            管理システムにアクセスするには認証が必要です
          </p>
        </div>

        {/* ログインフォーム */}
        <div className="bg-white py-8 px-6 shadow-xl rounded-lg border border-gray-200">
          <Form method="post" className="space-y-6">
            {/* エラーメッセージ */}
            {actionData?.error && (
              <div className="bg-red-50 border border-red-200 rounded-lg p-4">
                <div className="flex items-start">
                  <svg
                    className="w-5 h-5 text-red-500 mr-2 mt-0.5 flex-shrink-0"
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
                  <div>
                    <span className="text-red-800 text-sm font-medium block mb-1">
                      ログインエラー
                    </span>
                    <span className="text-red-700 text-sm">
                      {actionData.error}
                    </span>
                    <div className="mt-2 text-xs text-red-600">
                      <p>💡 ヒント:</p>
                      <ul className="list-disc list-inside ml-2 space-y-1">
                        <li>メールアドレス: test1@example.com</li>
                        <li>パスワード: password123</li>
                        <li>
                          ブラウザの開発者ツール（F12）でコンソールログを確認してください
                        </li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
            )}

            {/* メールアドレスフィールド */}
            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700 mb-2"
              >
                メールアドレス
              </label>
              <input
                id="email"
                name="email"
                type="email"
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors font-light tracking-wide"
                placeholder="test1@example.com"
                disabled={isSubmitting}
              />
            </div>

            {/* パスワードフィールド */}
            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-gray-700 mb-2"
              >
                パスワード
              </label>
              <input
                id="password"
                name="password"
                type="password"
                required
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-amber-500 focus:border-amber-500 transition-colors font-light tracking-wide"
                placeholder="パスワードを入力"
                disabled={isSubmitting}
              />
            </div>

            {/* ログインボタン */}
            <button
              type="submit"
              disabled={isSubmitting}
              className={`w-full py-3 px-4 rounded-lg font-medium tracking-wide transition-all duration-300 ${
                isSubmitting
                  ? "bg-gray-300 text-gray-500 cursor-not-allowed"
                  : "bg-gradient-to-r from-amber-400 to-yellow-500 text-white hover:from-amber-500 hover:to-yellow-600 shadow-lg hover:shadow-xl transform hover:scale-105"
              }`}
            >
              {isSubmitting ? (
                <div className="flex items-center justify-center">
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
                  認証中...
                </div>
              ) : (
                "ログイン"
              )}
            </button>
          </Form>
        </div>

        {/* フッター情報 */}
        <div className="text-center">
          <p className="text-xs text-gray-500 font-light">
            &copy; 2025 ArtWork Management System
          </p>
        </div>
      </div>
    </div>
  );
}

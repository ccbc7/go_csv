import type { MetaFunction, LoaderFunctionArgs } from "@remix-run/node";
import { Link, useLoaderData, Form } from "@remix-run/react";
import { requireAdmin } from "~/utils/auth.server";

export const meta: MetaFunction = () => {
  return [
    { title: "管理者ダッシュボード - ArtWork" },
    {
      name: "description",
      content: "管理者専用のダッシュボードページ",
    },
  ];
};

// 管理者権限が必要なページとして設定
export async function loader({ request }: LoaderFunctionArgs) {
  const token = await requireAdmin(request);
  return { token };
}

export default function AdminDashboard() {
  const { token } = useLoaderData<typeof loader>();

  // 管理機能のメニュー
  const adminMenus = [
    {
      title: "記事投稿",
      description: "新しい記事を作成・投稿",
      href: "/admin/post",
      icon: "📝",
      color: "from-blue-400 to-blue-600",
    },
    {
      title: "記事管理",
      description: "投稿済み記事の編集・削除",
      href: "/admin/articles",
      icon: "📰",
      color: "from-green-400 to-green-600",
    },
    {
      title: "ユーザー管理",
      description: "ユーザーアカウントの管理",
      href: "/admin/users",
      icon: "👥",
      color: "from-purple-400 to-purple-600",
    },
    {
      title: "設定",
      description: "システム設定・カテゴリ管理",
      href: "/admin/settings",
      icon: "⚙️",
      color: "from-gray-400 to-gray-600",
    },
  ];

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-gray-100">
      {/* ヘッダー */}
      <header className="bg-white border-b border-gray-200 sticky top-0 z-50 shadow-sm">
        <div className="max-w-7xl mx-auto px-6 py-5">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl md:text-3xl font-light text-gray-800 mb-1 tracking-wider">
                <span className="text-amber-600 font-normal">Art</span>Work
              </h1>
              <p className="text-gray-500 text-sm md:text-base hidden sm:block font-light tracking-wide">
                管理者ダッシュボード
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
              <Form method="post" action="/auth/logout">
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
      <main className="max-w-7xl mx-auto px-6 py-16">
        {/* ウェルカムセクション */}
        <div className="text-center mb-16">
          <h2 className="text-4xl md:text-5xl font-light text-gray-800 mb-4 tracking-wide">
            管理者ダッシュボード
          </h2>
          <p className="text-gray-600 text-xl font-light tracking-wide max-w-2xl mx-auto">
            ArtWorkの管理機能にアクセスして、サイトの運営を行います
          </p>
        </div>

        {/* 管理機能メニュー */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
          {adminMenus.map((menu, index) => (
            <Link
              key={index}
              to={`${menu.href}?token=${token}`}
              className="group bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105 border border-gray-100 overflow-hidden"
            >
              <div className={`h-2 bg-gradient-to-r ${menu.color}`}></div>
              <div className="p-8">
                <div className="text-4xl mb-4">{menu.icon}</div>
                <h3 className="text-xl font-medium text-gray-800 mb-2 group-hover:text-amber-600 transition-colors">
                  {menu.title}
                </h3>
                <p className="text-gray-600 text-sm font-light leading-relaxed">
                  {menu.description}
                </p>
                <div className="mt-4 flex items-center text-amber-600 text-sm font-medium group-hover:text-amber-700">
                  <span>アクセス</span>
                  <svg
                    className="w-4 h-4 ml-1 transform group-hover:translate-x-1 transition-transform"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 5l7 7-7 7"
                    />
                  </svg>
                </div>
              </div>
            </Link>
          ))}
        </div>

        {/* 統計情報セクション */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 mb-16">
          <div className="bg-white rounded-xl shadow-lg p-8 border border-gray-100">
            <div className="flex items-center">
              <div className="bg-blue-100 rounded-lg p-3 mr-4">
                <svg
                  className="w-8 h-8 text-blue-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                  />
                </svg>
              </div>
              <div>
                <h4 className="text-lg font-medium text-gray-800">記事数</h4>
                <p className="text-3xl font-light text-blue-600">12</p>
                <p className="text-sm text-gray-500">今月 +3</p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-xl shadow-lg p-8 border border-gray-100">
            <div className="flex items-center">
              <div className="bg-green-100 rounded-lg p-3 mr-4">
                <svg
                  className="w-8 h-8 text-green-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"
                  />
                </svg>
              </div>
              <div>
                <h4 className="text-lg font-medium text-gray-800">
                  ユーザー数
                </h4>
                <p className="text-3xl font-light text-green-600">156</p>
                <p className="text-sm text-gray-500">今月 +8</p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-xl shadow-lg p-8 border border-gray-100">
            <div className="flex items-center">
              <div className="bg-purple-100 rounded-lg p-3 mr-4">
                <svg
                  className="w-8 h-8 text-purple-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                  />
                </svg>
              </div>
              <div>
                <h4 className="text-lg font-medium text-gray-800">閲覧数</h4>
                <p className="text-3xl font-light text-purple-600">2,847</p>
                <p className="text-sm text-gray-500">今月 +156</p>
              </div>
            </div>
          </div>
        </div>

        {/* 最近のアクティビティ */}
        <div className="bg-white rounded-xl shadow-lg p-8 border border-gray-100">
          <h3 className="text-2xl font-light text-gray-800 mb-6 tracking-wide">
            最近のアクティビティ
          </h3>
          <div className="space-y-4">
            <div className="flex items-center py-3 border-b border-gray-100">
              <div className="bg-blue-100 rounded-full p-2 mr-4">
                <svg
                  className="w-5 h-5 text-blue-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                  />
                </svg>
              </div>
              <div className="flex-1">
                <p className="text-gray-800 font-medium">
                  新しい記事「水彩画の技法」を投稿
                </p>
                <p className="text-gray-500 text-sm">2時間前</p>
              </div>
            </div>
            <div className="flex items-center py-3 border-b border-gray-100">
              <div className="bg-green-100 rounded-full p-2 mr-4">
                <svg
                  className="w-5 h-5 text-green-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
              </div>
              <div className="flex-1">
                <p className="text-gray-800 font-medium">
                  新規ユーザー「田中さん」が登録
                </p>
                <p className="text-gray-500 text-sm">5時間前</p>
              </div>
            </div>
            <div className="flex items-center py-3">
              <div className="bg-purple-100 rounded-full p-2 mr-4">
                <svg
                  className="w-5 h-5 text-purple-600"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
                  />
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                </svg>
              </div>
              <div className="flex-1">
                <p className="text-gray-800 font-medium">システム設定を更新</p>
                <p className="text-gray-500 text-sm">1日前</p>
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* フッター */}
      <footer className="bg-white border-t border-gray-200 mt-20">
        <div className="max-w-7xl mx-auto px-6 py-10 md:py-12">
          <div className="text-center">
            <p className="text-sm text-gray-500 font-light tracking-wide">
              &copy; 2025 ArtWork Management System. All rights reserved.
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}

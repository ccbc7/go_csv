import type { ActionFunctionArgs } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import { authCookie } from "~/utils/auth.server";

// ログアウト処理
export async function action({ request }: ActionFunctionArgs) {
  // クッキーを削除してログインページにリダイレクト
  return redirect("/auth/login", {
    headers: {
      "Set-Cookie": await authCookie.serialize("", {
        maxAge: 0, // 即座に期限切れにする
      }),
    },
  });
}

// GETリクエストでもログアウトできるようにする
export async function loader() {
  return redirect("/auth/login");
}

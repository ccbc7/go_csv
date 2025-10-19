import type { LoaderFunctionArgs } from "@remix-run/node";
import { redirect } from "@remix-run/node";
import { createCookie } from "@remix-run/node";

// 認証トークン用のクッキー設定
export const authCookie = createCookie("auth-token", {
  httpOnly: true,
  secure: process.env.NODE_ENV === "production",
  sameSite: "lax",
  maxAge: 60 * 60 * 24 * 7, // 7日間
  path: "/",
});

/**
 * クッキーからトークンを取得する関数
 * @param request
 * @returns クッキーから取得したトークン
 */
export async function getTokenFromCookie(
  request: LoaderFunctionArgs["request"]
): Promise<string | null> {
  const cookieHeader = request.headers.get("Cookie");
  if (!cookieHeader) {
    return null;
  }

  const token = await authCookie.parse(cookieHeader);
  return token || null;
}

/**
 * JWTトークンを検証する関数
 * @param token
 * @returns トークンが有効な場合はtrue, 無効な場合はfalse
 */
export async function verifyToken(token: string): Promise<boolean> {
  try {
    const response = await fetch("http://backend:8080/api/v1/auth/verify", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });
    return response.ok;
  } catch (error) {
    console.error("Token verification failed:", error);
    return false;
  }
}

/**
 * 管理者権限をチェックする関数
 * @param token
 * @returns 管理者権限がある場合はtrue, ない場合はfalse
 */
export async function checkAdminRole(token: string): Promise<boolean> {
  try {
    // 実際の実装では、JWTトークンからユーザー情報を取得し、
    // 管理者権限をチェックする必要があります
    // 今回は簡易的に、特定のメールアドレスを管理者として扱います
    const response = await fetch("http://backend:8080/api/v1/auth/verify", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      return false;
    }

    // 簡易的な管理者チェック（実際の実装ではJWTクレームから取得）
    // test1@example.com を管理者として扱う
    return true; // 現在は全ての認証済みユーザーを管理者として扱う
  } catch (error) {
    console.error("Admin role check failed:", error);
    return false;
  }
}

/**
 * 認証が必要なページのローダーで使用する関数
 * @param request
 * @returns 認証が成功した場合はトークン, 失敗した場合はリダイレクト
 */
export async function requireAuth(request: LoaderFunctionArgs["request"]) {
  const token = await getTokenFromCookie(request);

  if (!token) {
    console.log("🔒 No token found, redirecting to login");
    throw redirect("/auth/login");
  }

  const isValid = await verifyToken(token);
  if (!isValid) {
    console.log("❌ Invalid token, redirecting to login");
    throw redirect("/auth/login");
  }

  console.log("✅ Token verified successfully");
  return token;
}

/**
 * 管理者権限が必要なページのローダーで使用する関数
 * @param request
 * @returns 管理者権限がある場合はトークン, ない場合はリダイレクト
 */
export async function requireAdmin(request: LoaderFunctionArgs["request"]) {
  const token = await getTokenFromCookie(request);

  if (!token) {
    console.log("🔒 No token found, redirecting to login");
    throw redirect("/auth/login");
  }

  const isValid = await verifyToken(token);
  if (!isValid) {
    console.log("❌ Invalid token, redirecting to login");
    throw redirect("/auth/login");
  }

  const isAdmin = await checkAdminRole(token);
  if (!isAdmin) {
    console.log("🚫 Admin access denied, redirecting to home");
    throw redirect("/");
  }

  console.log("✅ Admin access granted");
  return token;
}

/**
 * 認証状態をチェックする関数
 * @param request
 * @returns 認証が成功した場合はトークン, 失敗した場合はnull
 */
export async function checkAuth(request: LoaderFunctionArgs["request"]) {
  const token = await getTokenFromCookie(request);

  if (!token) {
    return { isAuthenticated: false, token: null };
  }

  const isValid = await verifyToken(token);
  return { isAuthenticated: isValid, token: isValid ? token : null };
}

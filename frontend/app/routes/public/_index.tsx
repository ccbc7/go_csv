import type { MetaFunction } from "@remix-run/node";
import { useState, useEffect } from "react";
import { Link } from "@remix-run/react";

export const meta: MetaFunction = () => {
  return [
    { title: "ArtWork - 上質なアートの世界" },
    {
      name: "description",
      content: "洗練されたアートと美の世界をお届けする、上質なギャラリーサイト",
    },
  ];
};

// スライドショー用の画像データ
const slideImages = [
  {
    id: 1,
    url: "/sample1.jpg",
    title: "美しき水彩の世界",
    subtitle: "透明感あふれる色彩の調べ",
  },
  {
    id: 2,
    url: "/sample2.jpg",
    title: "モダンアートの精神",
    subtitle: "現代に息づく芸術の心",
  },
  {
    id: 3,
    url: "/sample3.jpg",
    title: "伝統工芸の美学",
    subtitle: "受け継がれる技と心",
  },
  {
    id: 4,
    url: "/sample4.jpg",
    title: "創造の瞬間",
    subtitle: "アーティストの眼差し",
  },
];

// スライドショーコンポーネント
function Slideshow() {
  const [currentSlide, setCurrentSlide] = useState(0);
  const [isAutoPlay, setIsAutoPlay] = useState(true);

  // 自動切り替え
  useEffect(() => {
    if (!isAutoPlay) return;

    const interval = setInterval(() => {
      setCurrentSlide(prev => (prev + 1) % slideImages.length);
    }, 4000); // 4秒間隔

    return () => clearInterval(interval);
  }, [isAutoPlay]);

  // ドットクリック時の処理
  const handleDotClick = (index: number) => {
    setCurrentSlide(index);
    setIsAutoPlay(false); // 自動再生を停止

    // 5秒後に自動再生を再開
    setTimeout(() => {
      setIsAutoPlay(true);
    }, 5000);
  };

  return (
    <div
      className="relative w-full h-64 md:h-80 lg:h-96 overflow-hidden bg-gradient-to-br from-amber-50 to-orange-50"
      onMouseEnter={() => setIsAutoPlay(false)}
      onMouseLeave={() => setIsAutoPlay(true)}
    >
      {/* 画像スライド */}
      <div className="relative w-full h-full">
        {slideImages.map((slide, index) => (
          <div
            key={slide.id}
            className={`absolute inset-0 transition-opacity duration-1000 ease-in-out ${
              index === currentSlide ? "opacity-100" : "opacity-0"
            }`}
          >
            <img
              src={slide.url}
              alt={slide.title}
              className="w-full h-full object-cover"
            />
            {/* 上品なオーバーレイとテキスト */}
            <div className="absolute inset-0 bg-gradient-to-t from-black/60 via-black/20 to-transparent flex items-center justify-center">
              <div className="text-center text-white px-6">
                <h2 className="text-2xl md:text-4xl font-light mb-2 md:mb-4 tracking-wide">
                  {slide.title}
                </h2>
                <p className="text-sm md:text-lg opacity-90 font-light tracking-wider">
                  {slide.subtitle}
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* エレガントなドットナビゲーション */}
      <div className="absolute bottom-6 left-1/2 transform -translate-x-1/2 flex space-x-3">
        {slideImages.map((_, index) => (
          <button
            key={index}
            onClick={() => handleDotClick(index)}
            className={`w-3 h-3 rounded-full transition-all duration-300 border ${
              index === currentSlide
                ? "bg-amber-400 border-amber-300 scale-110 shadow-lg shadow-amber-400/50"
                : "bg-white/70 border-white/50 hover:bg-white hover:border-white"
            }`}
            aria-label={`スライド ${index + 1} に移動`}
          />
        ))}
      </div>

      {/* ゴールドアクセントのプログレスバー */}
      <div className="absolute bottom-0 left-0 w-full h-0.5 bg-white/20">
        <div
          className="h-full bg-gradient-to-r from-amber-400 to-yellow-300 transition-all duration-100 ease-linear shadow-sm"
          style={{
            width: isAutoPlay
              ? `${((currentSlide + 1) / slideImages.length) * 100}%`
              : "0%",
          }}
        />
      </div>
    </div>
  );
}

// モックデータ
const blogPosts = [
  {
    id: 1,
    title: "水彩画で織りなす詩的な表現",
    excerpt:
      "透明感ある水彩の技法で、季節の移ろいや心の機微を繊細に表現する方法を、実際の作品制作とともにご紹介いたします。",
    date: "2024年1月15日",
    category: "絵画",
    readTime: "5分",
    image: "/placeholder-image-1.jpg",
  },
  {
    id: 2,
    title: "現代陶芸に宿る美意識",
    excerpt:
      "伝統的な陶芸の技法に現代的な感性を融合させた、新しい美の表現について、注目の作家の作品を通して探求いたします。",
    date: "2024年1月12日",
    category: "工芸",
    readTime: "8分",
    image: "/placeholder-image-2.jpg",
  },
  {
    id: 3,
    title: "日常に潜む美的瞬間",
    excerpt:
      "街角や自然の中に隠れている美しい瞬間を捉える眼差しと、それを芸術として昇華させるアーティストの哲学をお伝えします。",
    date: "2024年1月10日",
    category: "写真",
    readTime: "6分",
    image: "/placeholder-image-3.jpg",
  },
  {
    id: 4,
    title: "上質なアート空間との出会い",
    excerpt:
      "都内の選りすぐりのギャラリーと、作品との深い対話を楽しむための鑑賞のエッセンスを、専門家の視点からご案内いたします。",
    date: "2024年1月8日",
    category: "展覧会",
    readTime: "7分",
    image: "/placeholder-image-4.jpg",
  },
];

const categories = ["すべて", "絵画", "工芸", "一言", "展覧会"];

export default function Index() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  return (
    <div className="min-h-screen bg-white">
      {/* エレガントなヘッダー */}
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

            {/* ハンバーガーメニューボタン（モバイル用） */}
            <button
              onClick={() => setIsMenuOpen(!isMenuOpen)}
              className="md:hidden p-2 rounded-lg hover:bg-amber-50 transition-colors"
              aria-label="メニューを開く"
            >
              <svg
                className="w-6 h-6 text-gray-600"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                {isMenuOpen ? (
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={1.5}
                    d="M6 18L18 6M6 6l12 12"
                  />
                ) : (
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={1.5}
                    d="M4 6h16M4 12h16M4 18h16"
                  />
                )}
              </svg>
            </button>

            {/* デスクトップ用ナビゲーション */}
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
                className="text-gray-600 hover:text-amber-600 transition-colors font-light tracking-wide relative group"
              >
                Hello
                <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-amber-400 transition-all duration-300 group-hover:w-full"></span>
              </Link>
              <Link
                to="/management/auth"
                className="text-gray-600 hover:text-amber-600 transition-colors font-light tracking-wide relative group"
              >
                管理者
                <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-amber-400 transition-all duration-300 group-hover:w-full"></span>
              </Link>
              <a
                href="/gallery"
                className="text-gray-600 hover:text-amber-600 transition-colors font-light tracking-wide relative group"
              >
                ギャラリー
                <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-amber-400 transition-all duration-300 group-hover:w-full"></span>
              </a>
              <a
                href="/artists"
                className="text-gray-600 hover:text-amber-600 transition-colors font-light tracking-wide relative group"
              >
                アーティスト
                <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-amber-400 transition-all duration-300 group-hover:w-full"></span>
              </a>
              <a
                href="/contact"
                className="text-gray-600 hover:text-amber-600 transition-colors font-light tracking-wide relative group"
              >
                お問い合わせ
                <span className="absolute -bottom-1 left-0 w-0 h-0.5 bg-amber-400 transition-all duration-300 group-hover:w-full"></span>
              </a>
            </nav>
          </div>

          {/* モバイル用メニュー */}
          {isMenuOpen && (
            <nav className="md:hidden mt-6 pt-6 border-t border-amber-100">
              <div className="space-y-4">
                <Link
                  to="/"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  ホーム
                </Link>
                <Link
                  to="/hello"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  Hello
                </Link>
                <Link
                  to="/management/auth"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  管理者
                </Link>
                <a
                  href="/gallery"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  ギャラリー
                </a>
                <a
                  href="/artists"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  アーティスト
                </a>
                <a
                  href="/exhibitions"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  展覧会情報
                </a>
                <a
                  href="/contact"
                  className="block text-gray-600 hover:text-amber-600 transition-colors py-2 font-light tracking-wide"
                  onClick={() => setIsMenuOpen(false)}
                >
                  お問い合わせ
                </a>
              </div>
            </nav>
          )}
        </div>
      </header>

      {/* スライドショー */}
      <Slideshow />

      <main className="max-w-5xl mx-auto px-6 py-10 md:py-16">
        {/* 洗練されたカテゴリフィルター */}
        <div className="mb-10 md:mb-16">
          <div className="flex flex-wrap gap-3 md:gap-4 justify-center">
            {categories.map(category => (
              <button
                key={category}
                className={`px-5 py-2.5 md:px-6 md:py-3 rounded-full text-sm md:text-base font-light transition-all duration-300 tracking-wide ${
                  category === "すべて"
                    ? "bg-gradient-to-r from-amber-300 to-yellow-400 text-white shadow-lg shadow-amber-200 transform scale-105"
                    : "bg-white text-gray-600 hover:text-amber-600 border border-gray-200 hover:border-amber-200 hover:bg-amber-50 hover:shadow-md"
                }`}
              >
                {category}
              </button>
            ))}
          </div>
        </div>

        {/* 記事一覧 */}
        <div className="space-y-4 md:space-y-12">
          {blogPosts.map(post => (
            <article
              key={post.id}
              className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden hover:shadow-lg hover:border-amber-100 transition-all duration-300 group"
            >
              <div className="p-4 md:p-10">
                <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between mb-6 gap-3">
                  <div className="flex items-center text-sm text-gray-500 space-x-4 font-light">
                    <time>{post.date}</time>
                    <span className="text-amber-300">•</span>
                    <span className="inline-block px-4 py-2 text-sm font-light text-amber-700 bg-gradient-to-r from-amber-50 to-yellow-50 rounded-full w-fit border border-amber-100">
                      {post.category}
                    </span>
                  </div>
                </div>

                <h2 className="text-xl md:text-2xl font-light text-gray-800 mb-4 hover:text-amber-700 transition-colors cursor-pointer tracking-wide leading-relaxed">
                  {post.title}
                </h2>

                <p className="text-gray-600 leading-relaxed mb-8 text-sm md:text-base font-light tracking-wide">
                  {post.excerpt}
                </p>

                <button className="inline-flex items-center text-amber-600 hover:text-amber-700 font-light transition-all duration-300 text-sm md:text-base tracking-wide group-hover:translate-x-1">
                  詳しく見る
                  <svg
                    className="ml-2 w-4 h-4 transition-transform duration-300 group-hover:translate-x-1"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={1.5}
                      d="M9 5l7 7-7 7"
                    />
                  </svg>
                </button>
              </div>
            </article>
          ))}
        </div>
      </main>

      {/* 上品なフッター */}
      <footer className="bg-gradient-to-b from-white to-amber-50 border-t border-amber-100 mt-20 md:mt-28">
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

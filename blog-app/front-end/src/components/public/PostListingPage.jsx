import React from "react";

function PostListingPage() {
  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <header className="border-b border-gray-200 py-4">
        <div className="container mx-auto px-4 max-w-7xl flex justify-between items-center">
          <div className="flex items-center">
            <span className="text-2xl font-bold text-black">Medium</span>
          </div>

          <div className="hidden md:flex items-center space-x-6">
            <a href="#" className="text-gray-600 hover:text-black">
              Our story
            </a>
            <a href="#" className="text-gray-600 hover:text-black">
              Membership
            </a>
            <a href="#" className="text-gray-600 hover:text-black">
              Write
            </a>
            <a href="#" className="text-gray-600 hover:text-black">
              Sign In
            </a>
            <button className="bg-black text-white px-4 py-2 rounded-full text-sm font-medium hover:bg-gray-800 transition-colors">
              Get Started
            </button>
          </div>

          <button className="md:hidden text-gray-600">
            <svg
              className="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16m-7 6h7"
              ></path>
            </svg>
          </button>
        </div>
      </header>

      <div className="container mx-auto px-4 max-w-7xl py-8">
        <div className="flex flex-col lg:flex-row gap-10">
          {/* Main Content */}
          <div className="lg:w-2/3">
            {/* Navigation Tabs */}
            <div className="flex border-b border-gray-200 mb-8 overflow-x-auto">
              <button className="px-4 py-2 border-b-2 border-black text-black font-medium">
                Following
              </button>
              <button className="px-4 py-2 text-gray-600 hover:text-black">
                Recommended
              </button>
            </div>

            {/* Article Cards */}
            <div className="space-y-10">
              {/* Article 1 */}
              <div className="flex flex-col sm:flex-row gap-6">
                <div className="sm:w-2/3">
                  <div className="flex items-center space-x-2 mb-3">
                    <div className="w-6 h-6 bg-gray-300 rounded-full"></div>
                    <span className="text-sm font-medium text-gray-900">
                      Pied Heveline John
                    </span>
                    <span className="text-xs text-gray-500">in</span>
                    <span className="text-sm font-medium text-gray-900">
                      Thought Thinking
                    </span>
                  </div>
                  <h2 className="text-2xl font-bold text-gray-900 mb-3">
                    How to slice a pie
                  </h2>
                  <p className="text-gray-600 mb-4">
                    The art of perfect pie slicing techniques for any occasion
                  </p>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2 text-sm text-gray-500">
                      <span>May 7</span>
                      <span>•</span>
                      <span>5 min read</span>
                    </div>
                    <button className="p-1 text-gray-400 hover:text-black">
                      <svg
                        className="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"
                        ></path>
                      </svg>
                    </button>
                  </div>
                </div>
                <div className="sm:w-1/3">
                  <div className="w-full h-32 bg-gray-100 rounded"></div>
                </div>
              </div>

              {/* Article 2 */}
              <div className="flex flex-col sm:flex-row gap-6">
                <div className="sm:w-2/3">
                  <div className="flex items-center space-x-2 mb-3">
                    <div className="w-6 h-6 bg-gray-400 rounded-full"></div>
                    <span className="text-sm font-medium text-gray-900">
                      Jason McBride
                    </span>
                  </div>
                  <h2 className="text-2xl font-bold text-gray-900 mb-3">
                    Walking in the Underworld
                  </h2>
                  <p className="text-gray-600 mb-4">
                    A journey through mythological realms and personal
                    transformation
                  </p>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2 text-sm text-gray-500">
                      <span>4 days ago</span>
                      <span>•</span>
                      <span>8 min read</span>
                    </div>
                    <button className="p-1 text-gray-400 hover:text-black">
                      <svg
                        className="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"
                        ></path>
                      </svg>
                    </button>
                  </div>
                </div>
                <div className="sm:w-1/3">
                  <div className="w-full h-32 bg-gray-200 rounded"></div>
                </div>
              </div>

              {/* Article 3 */}
              <div className="flex flex-col sm:flex-row gap-6">
                <div className="sm:w-2/3">
                  <div className="flex items-center space-x-2 mb-3">
                    <div className="w-6 h-6 bg-gray-500 rounded-full"></div>
                    <span className="text-sm font-medium text-gray-900">
                      Dev Welton
                    </span>
                    <span className="text-xs text-gray-500">in</span>
                    <span className="text-sm font-medium text-gray-900">
                      Condemnation
                    </span>
                  </div>
                  <h2 className="text-2xl font-bold text-gray-900 mb-3">
                    Go 1.25: The Top 5 Changes Backend Developers Should Care
                    About
                  </h2>
                  <p className="text-gray-600 mb-4">
                    The first Release Candidate (RC) of Go 1.25 is out. Every
                    time I read the official Release Notes, it feels like
                    panning for gold
                  </p>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2 text-sm text-gray-500">
                      <span>Jan 21</span>
                      <span>•</span>
                      <span>6 min read</span>
                    </div>
                    <button className="p-1 text-gray-400 hover:text-black">
                      <svg
                        className="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"
                        ></path>
                      </svg>
                    </button>
                  </div>
                </div>
                <div className="sm:w-1/3">
                  <div className="w-full h-32 bg-gray-300 rounded"></div>
                </div>
              </div>

              {/* Article 4 */}
              <div className="flex flex-col sm:flex-row gap-6">
                <div className="sm:w-2/3">
                  <div className="flex items-center space-x-2 mb-3">
                    <div className="w-6 h-6 bg-gray-600 rounded-full"></div>
                    <span className="text-sm font-medium text-gray-900">
                      TwoDemonstration
                    </span>
                  </div>
                  <h2 className="text-2xl font-bold text-gray-900 mb-3">
                    Go Just Stole Rust's Secret Weapon—Without the Pain
                  </h2>
                  <p className="text-gray-600 mb-4">
                    The day Go learned Rust's memory tricks—without lifetimes,
                    borrow checkers, or sleepless nights
                  </p>
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-2 text-sm text-gray-500">
                      <span>Aug 18</span>
                      <span>•</span>
                      <span>7 min read</span>
                    </div>
                    <button className="p-1 text-gray-400 hover:text-black">
                      <svg
                        className="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M5 5a2 2 0 012-2h10a2 2 0 012 2v16l-7-3.5L5 21V5z"
                        ></path>
                      </svg>
                    </button>
                  </div>
                </div>
                <div className="sm:w-1/3">
                  <div className="w-full h-32 bg-gray-400 rounded"></div>
                </div>
              </div>
            </div>
          </div>

          {/* Sidebar */}
          <div className="lg:w-1/3">
            {/* Membership Card */}
            <div className="border border-gray-200 rounded-lg p-6 mb-8">
              <h3 className="text-xl font-bold text-gray-900 mb-3">
                Get unlimited access
              </h3>
              <p className="text-gray-600 mb-4">
                Enjoy the best of Medium for less than $5/month
              </p>
              <button className="w-full bg-black text-white py-3 px-4 rounded-full font-medium hover:bg-gray-800 transition-colors">
                Become a Member
              </button>
            </div>

            {/* Recommended Topics */}
            <div className="mb-8">
              <h3 className="text-lg font-bold text-gray-900 mb-4">
                Recommended topics
              </h3>
              <div className="flex flex-wrap gap-2">
                {[
                  "React",
                  "Productivity",
                  "Data Science",
                  "Relationships",
                  "Humor",
                  "Life",
                  "Technology",
                  "Programming",
                ].map((topic) => (
                  <span
                    key={topic}
                    className="bg-gray-100 text-gray-700 px-3 py-1.5 rounded-full text-sm hover:bg-gray-200 cursor-pointer"
                  >
                    {topic}
                  </span>
                ))}
              </div>
            </div>

            {/* Who to follow */}
            <div className="mb-8">
              <h3 className="text-lg font-bold text-gray-900 mb-4">
                Who to follow
              </h3>
              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-3">
                    <div className="w-10 h-10 bg-gray-300 rounded-full"></div>
                    <div>
                      <h4 className="font-medium text-gray-900">Adit Mishra</h4>
                      <p className="text-sm text-gray-500">Software Engineer</p>
                    </div>
                  </div>
                  <button className="text-gray-700 text-sm font-medium hover:text-black border border-gray-300 px-3 py-1 rounded-full">
                    Follow
                  </button>
                </div>
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-3">
                    <div className="w-10 h-10 bg-gray-400 rounded-full"></div>
                    <div>
                      <h4 className="font-medium text-gray-900">Gollet</h4>
                      <p className="text-sm text-gray-500">Senior Developer</p>
                    </div>
                  </div>
                  <button className="text-gray-700 text-sm font-medium hover:text-black border border-gray-300 px-3 py-1 rounded-full">
                    Follow
                  </button>
                </div>
              </div>
            </div>

            {/* Recently Saved */}
            <div>
              <h3 className="text-lg font-bold text-gray-900 mb-4">
                Recently saved
              </h3>
              <div className="space-y-4">
                <div className="flex items-center space-x-3">
                  <div className="w-7 h-7 bg-gray-500 rounded-full flex-shrink-0"></div>
                  <p className="text-sm font-medium text-gray-900">
                    How to Change Your Life Today
                  </p>
                </div>
                <div className="text-sm text-gray-500 pl-10">
                  Sep 5 · Saved from Mindful Living
                </div>
                <button className="text-sm text-gray-700 hover:text-black font-medium pl-10">
                  See all (3)
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default PostListingPage;

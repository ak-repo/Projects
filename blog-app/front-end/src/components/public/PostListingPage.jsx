import axios from "axios";
import React, { useEffect, useState } from "react";
import { postApi } from "../../api";

function PostListingPage() {
  const [postList, setPostList] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  // fetch posts
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        setIsLoading(true);
        const response = await axios.get(postApi);
        setPostList(response.data.posts);
        setError(null);
      } catch (error) {
        console.error("Error fetching posts:", error);
        setError("Failed to load posts. Please try again later.");
      } finally {
        setIsLoading(false);
      }
    };

    fetchPosts();
  }, []);

  useEffect(() => {
    console.log(postList);
  }, [postList]);

  // // Format date function
  // const formatDate = (dateString) => {
  //   const options = { year: 'numeric', month: 'long', day: 'numeric' };
  //   return new Date(dateString).toLocaleDateString(undefined, options);
  // };

  // // Truncate content for preview
  // const truncateContent = (content, maxLength = 150) => {
  //   if (content.length <= maxLength) return content;
  //   return content.substring(0, maxLength) + '...';
  // };

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="container mx-auto px-4 max-w-4xl">
        <header className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-800 mb-2">Blog Posts</h1>
          <p className="text-lg text-gray-600">
            Discover the latest articles and insights
          </p>
        </header>

        {/* Loading State */}
        {isLoading && (
          <div className="flex justify-center items-center h-64">
            <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
          </div>
        )}

        {/* Error State */}
        {error && (
          <div className="bg-red-50 border-l-4 border-red-500 p-4 mb-6 rounded">
            <div className="flex">
              <div className="flex-shrink-0">
                <svg
                  className="h-5 w-5 text-red-400"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fillRule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                    clipRule="evenodd"
                  />
                </svg>
              </div>
              <div className="ml-3">
                <p className="text-sm text-red-700">{error}</p>
              </div>
            </div>
          </div>
        )}

        {/* Posts Grid */}
        {!isLoading && !error && (
          <>
            {postList.length > 0 ? (
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {postList.map((post) => (
                  <article
                    key={post?.PostID}
                    className="bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-300 hover:shadow-lg hover:-translate-y-1"
                  >
                    <div className="p-6">
                      <h2 className="text-xl font-semibold text-gray-800 mb-2 line-clamp-2">
                        {post?.title}
                      </h2>

                      {post?.date && (
                        <p className="text-sm text-gray-500 mb-3">
                          {post.date}
                        </p>
                      )}

                      <p className="text-gray-600 mb-4">{post?.content}</p>

                      <div className="flex justify-between items-center">
                        <button className="text-blue-600 hover:text-blue-800 font-medium flex items-center">
                          Read more
                          <svg
                            className="w-4 h-4 ml-1"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                            xmlns="http://www.w3.org/2000/svg"
                          >
                            <path
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth="2"
                              d="M9 5l7 7-7 7"
                            ></path>
                          </svg>
                        </button>

                        {post?.author && (
                          <span className="text-sm text-gray-500">
                            By {post.author}
                          </span>
                        )}
                      </div>
                    </div>
                  </article>
                ))}
              </div>
            ) : (
              <div className="text-center py-12">
                <svg
                  className="mx-auto h-12 w-12 text-gray-400"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
                <h3 className="mt-2 text-lg font-medium text-gray-900">
                  No posts found
                </h3>
                <p className="mt-1 text-gray-500">
                  Get started by creating a new post.
                </p>
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}

export default PostListingPage;

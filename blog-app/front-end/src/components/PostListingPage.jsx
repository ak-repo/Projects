import axios from "axios";
import React, { useEffect, useState } from "react";

export const postApi = "http://localhost:8080/posts";

function PostListingPage() {
  const [postList, setPostList] = useState([]);

  // fetch posts
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axios.get(postApi);
        setPostList(response.data); // store the data, not the promise
        console.log(response);
      } catch (error) {
        console.error("Error fetching posts:", error);
      }
    };

    fetchPosts();
  }, []);

  return (
    <div>
      <h1>Blogs</h1>
      {postList.length > 0 ? (
        <ul>
          {postList.map((post) => (
            <li key={post.PostID}>{post.title}</li> // adjust keys/fields as per your API
          ))}
        </ul>
      ) : (
        <p>No posts found</p>
      )}
    </div>
  );
}

export default PostListingPage;

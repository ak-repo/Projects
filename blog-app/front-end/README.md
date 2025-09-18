# React + Gin/go Blog app

- user - registartion, login, logout - post create, update,delete

- public - home, posts




src/
 ├── app/
 │   ├── store.js          # Redux store setup
 │   └── rootReducer.js    # Combine slices
 ├── features/
 │   ├── auth/
 │   │   ├── authSlice.js
 │   │   └── AuthPage.jsx
 │   ├── posts/
 │   │   ├── postsSlice.js
 │   │   ├── PostList.jsx
 │   │   ├── PostDetail.jsx
 │   │   └── PostForm.jsx
 │   └── comments/
 │       ├── commentsSlice.js
 │       └── CommentList.jsx
 ├── components/           # Reusable UI components
 │   ├── Navbar.jsx
 │   ├── Footer.jsx
 │   └── Loader.jsx
 ├── pages/                # Page-level components
 │   ├── Home.jsx
 │   ├── Login.jsx
 │   └── Register.jsx
 ├── utils/                # Helpers (API client, etc.)
 │   └── api.js
 ├── App.jsx
 └── index.js
